package sys_jwt

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/response"
	"github.com/yitter/idgenerator-go/idgen"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
)

type hookInfo sys_model.KeyValueT[int64, sys_model.JwtHookInfo]

type sJwt struct {
	SigningKey []byte
	hookArr    []hookInfo
}

var (
	ConcurrencyControl = &singleflight.Group{}
)

func init() {
	sys_service.RegisterJwt(New())
}

// New MiddlewareJwt 权限控制
func New() *sJwt {
	return &sJwt{
		SigningKey: []byte(g.Cfg().MustGet(gctx.New(), "service.tokenSignKey").String()),
	}
}

// InstallHook 安装Hook
func (s *sJwt) InstallHook(userType sys_enum.UserType, hookFunc sys_model.JwtHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_model.JwtHookInfo{Key: userType, Value: hookFunc}}

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
func (s *sJwt) GenerateToken(user *sys_entity.SysUser) (*sys_model.TokenInfo, error) {
	customClaims := sys_model.JwtCustomClaims{
		Id:        user.Id,
		Username:  user.Username,
		State:     user.State,
		Type:      user.Type,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "免啦街",
			Subject:   "筷满客",
		},
	}

	token, err := s.CreateToken(&customClaims)

	if err != nil {
		return nil, gerror.New("创建登录令牌失败")
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

// RefreshToken 刷新Token,并发安全
func (s *sJwt) RefreshToken(oldToken string, claims *sys_model.JwtCustomClaims) (string, error) {
	v, err, _ := ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return s.CreateToken(claims)
	})
	return v.(string), err
}

// CustomMiddleware 自定义调用JWT网关认证
func (s *sJwt) CustomMiddleware(r *ghttp.Request) {
	Authorization := r.Header.Get("Authorization")
	claimsUser, err := s.ParseToken(Authorization)
	if err != nil {
		response.JsonExit(r, 401, err.Error())
		return
	}
	var userUnionMainId int64
	g.Try(r.Context(), func(ctx context.Context) {

		userInfo := sys_entity.SysUser{
			Id:       claimsUser.Id,
			Username: claimsUser.Username,
			State:    claimsUser.State,
			Type:     claimsUser.Type,
		}
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&claimsUser.Type == claimsUser.Type {
				userUnionMainId, err = hook.Value.Value(ctx, userInfo)
				if err != nil {
					break
				}
			}
		}
	})
	claimsUser.UnionMainId = userUnionMainId
	sys_service.BizCtx().SetUser(r.Context(), claimsUser)
}

func (s *sJwt) Middleware(r *ghttp.Request) {
	s.CustomMiddleware(r)
	r.Middleware.Next()
}

func (s *sJwt) ParseToken(tokenString string) (*sys_model.JwtCustomClaims, error) {
	if gstr.HasPrefix(tokenString, "Bearer ") {
		tokenString = gstr.SubStr(tokenString, 7)
	}

	token, err := jwt.ParseWithClaims(tokenString, &sys_model.JwtCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return s.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, gerror.NewSkip(401, "无效TOKEN")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, gerror.NewSkip(401, "TOKEN 已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, gerror.NewSkip(401, "TOKEN 未激活")
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, gerror.NewSkip(401, "TOKEN 签名无效")
			} else {
				return nil, gerror.NewSkip(401, "解析TOKEN失败")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*sys_model.JwtCustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, gerror.NewSkip(401, "解析TOKEN失败")

	} else {
		return nil, gerror.NewSkip(401, "解析TOKEN失败")
	}
}
