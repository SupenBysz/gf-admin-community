package sys_jwt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-community/utility/response"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
)

type hookInfo sys_model.KeyValueT[int64, sys_hook.JwtHookInfo]

type sJwt struct {
	SigningKey    []byte
	hookArr       []hookInfo
	revokedTokens map[string]time.Time // token黑名单，存储被注销的token及其过期时间
	mutex         sync.RWMutex         // 保护revokedTokens的读写锁
	tokenCache    *gredis.Redis        // Redis缓存实例，用于缓存token验证结果
}

var (
	ConcurrencyControl = &singleflight.Group{}
)

func init() {
	sys_service.RegisterJwt(New())
}

// New MiddlewareJwt 权限控制
func New() sys_service.IJwt {
	return &sJwt{
		SigningKey:    []byte(g.Cfg().MustGet(gctx.New(), "service.tokenSignKey").String()),
		hookArr:       make([]hookInfo, 0),
		revokedTokens: make(map[string]time.Time),
		tokenCache:    g.Redis(),
	}
}

// InstallHook 安装Hook
func (s *sJwt) InstallHook(userType sys_enum.UserType, hookFunc sys_hook.JwtHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_hook.JwtHookInfo{Key: userType, Value: hookFunc}}

	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sJwt) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除所有Hook
func (s *sJwt) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// GenerateToken 创建一个token
func (s *sJwt) GenerateToken(ctx context.Context, user *sys_model.SysUser) (response *sys_model.TokenInfo, err error) {
	user.Password = ""

	customClaims := &sys_model.JwtCustomClaims{
		SysUser:      *user,
		IsSuperAdmin: user.Type == sys_enum.User.Type.SuperAdmin.Code(),
		IsAdmin:      user.Type == sys_enum.User.Type.Admin.Code(),

		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "免啦街",
			Subject:   "筷满客",
		},
	}

	_ = g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&user.Type == user.Type || (user.Type == 64 && hook.Value.Key.Code() == 32) {
				customClaims, err = hook.Value.Value(ctx, customClaims)
				if err != nil {
					break
				}
			}
		}
	})

	token, err := s.CreateToken(customClaims)

	if err != nil {
		return nil, gerror.New("error_token_creation_failed")
	}

	return &sys_model.TokenInfo{
		Token:    token,
		ExpireAt: customClaims.RegisteredClaims.ExpiresAt.Time,
	}, nil
}

// CreateToken 创建一个token
func (s *sJwt) CreateToken(claims *sys_model.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.SigningKey)
}

// RevokeToken 注销token
func (s *sJwt) RevokeToken(ctx context.Context, tokenString string) error {
	if gstr.HasPrefix(tokenString, "Bearer ") {
		tokenString = gstr.SubStr(tokenString, 7)
	}

	// 解析token获取过期时间
	token, err := jwt.ParseWithClaims(tokenString, &sys_model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.SigningKey, nil
	})

	if err != nil {
		return gerror.New("invalid_token")
	}

	if claims, ok := token.Claims.(*sys_model.JwtCustomClaims); ok {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		// 将token添加到黑名单，存储到过期时间
		s.revokedTokens[tokenString] = claims.RegisteredClaims.ExpiresAt.Time

		// 清理已过期的token（可选的优化）
		now := time.Now()
		for token, expireTime := range s.revokedTokens {
			if now.After(expireTime) {
				delete(s.revokedTokens, token)
			}
		}

		return nil
	}

	return gerror.New("invalid_token_claims")
}

// IsTokenRevoked 检查token是否已被注销
func (s *sJwt) IsTokenRevoked(ctx context.Context, tokenString string) bool {
	if gstr.HasPrefix(tokenString, "Bearer ") {
		tokenString = gstr.SubStr(tokenString, 7)
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	expireTime, exists := s.revokedTokens[tokenString]
	if !exists {
		return false
	}

	// 如果token已过期，从黑名单中移除
	if time.Now().After(expireTime) {
		s.mutex.RUnlock()
		s.mutex.Lock()
		delete(s.revokedTokens, tokenString)
		s.mutex.Unlock()
		s.mutex.RLock()
		return false
	}

	return true
}

// RefreshToken 刷新Token,并发安全
func (s *sJwt) RefreshToken(oldToken string, claims *sys_model.JwtCustomClaims) (string, error) {
	v, err, _ := ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return s.CreateToken(claims)
	})
	return v.(string), err
}

// Middleware 鉴权中间件函数
func (s *sJwt) Middleware(r *ghttp.Request) {
	tokenString := gstr.Trim(r.Header.Get("Authorization"))

	if gstr.ToUpper(r.Method) == "GET" && tokenString == "" {
		tokenString = r.Get("token", "").String()
	}

	s.MakeSession(r.Context(), tokenString)
}

// MakeSession 构建会话
func (s *sJwt) MakeSession(ctx context.Context, tokenString string) *sys_model.JwtCustomClaims {
	if gstr.HasPrefix(tokenString, "Bearer ") {
		tokenString = gstr.SubStr(tokenString, 7)
	}

	// 检查token是否已被注销
	if s.IsTokenRevoked(ctx, tokenString) {
		isCustomSession := sys_service.SysSession().HasCustom(ctx)
		if !isCustomSession {
			response.JsonExit(g.RequestFromCtx(ctx), 401, "token_revoked")
		}
		return nil
	}

	// 尝试从缓存中获取token验证结果
	cacheKey := fmt.Sprintf("jwt:token:%s", tokenString)
	if s.tokenCache != nil {
		cachedData, err := s.tokenCache.Get(ctx, cacheKey)
		if err == nil && !cachedData.IsNil() {
			var claims sys_model.JwtCustomClaims
			if err := json.Unmarshal(cachedData.Bytes(), &claims); err == nil {
				// 验证缓存的token是否仍然有效
				if time.Now().Before(claims.RegisteredClaims.ExpiresAt.Time) {
					isCustomSession := sys_service.SysSession().HasCustom(ctx)
					if !isCustomSession {
						sys_service.SysSession().SetUser(ctx, &claims)
						g.RequestFromCtx(ctx).Middleware.Next()
					}
					return &claims
				} else {
					// 缓存的token已过期，删除缓存
					s.tokenCache.Del(ctx, cacheKey)
				}
			}
		}
	}

	token, err := jwt.ParseWithClaims(tokenString, &sys_model.JwtCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return s.SigningKey, nil
	})

	isCustomSession := sys_service.SysSession().HasCustom(ctx)

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = gerror.Wrap(err, "error_invalid_token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				err = gerror.Wrap(err, "error_token_expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = gerror.Wrap(err, "error_token_not_active")
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				err = gerror.Wrap(err, "error_token_signature_invalid")
			} else {
				err = gerror.Wrap(err, "error_token_parsing_failed")
			}
		}
		if !isCustomSession {
			response.JsonExit(g.RequestFromCtx(ctx), 401, "error_token_parsing_failed")
		}
		return nil
	}

	if token != nil {
		if claims, ok := token.Claims.(*sys_model.JwtCustomClaims); ok && token.Valid {
			// 将验证成功的token缓存到Redis，设置过期时间为token的剩余有效时间
			if s.tokenCache != nil {
				claimsData, err := json.Marshal(claims)
				if err == nil {
					ttl := time.Until(claims.RegisteredClaims.ExpiresAt.Time)
					if ttl > 0 {
						s.tokenCache.SetEX(ctx, cacheKey, claimsData, int64(ttl.Seconds()))
					}
				}
			}

			if !isCustomSession {
				sys_service.SysSession().SetUser(ctx, claims)
				g.RequestFromCtx(ctx).Middleware.Next()
			}
			return claims
		}
	}

	if !isCustomSession {
		response.JsonExit(g.RequestFromCtx(ctx), 401, "error_token_parsing_failed")
	}
	return nil
}
