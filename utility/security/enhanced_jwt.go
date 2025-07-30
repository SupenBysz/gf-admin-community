package security

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
)

// EnhancedJwtService 增强的JWT服务
type EnhancedJwtService struct {
	securityManager *TokenSecurityManager
}

// NewEnhancedJwtService 创建增强JWT服务
func NewEnhancedJwtService() *EnhancedJwtService {
	return &EnhancedJwtService{
		securityManager: NewTokenSecurityManager(),
	}
}

// RefreshTokenWithRotation Token轮换机制 - 自动刷新和撤销旧Token
func (s *EnhancedJwtService) RefreshTokenWithRotation(ctx context.Context, oldToken string) (*sys_model.TokenInfo, error) {
	// 验证旧token
	claims := sys_service.Jwt().MakeSession(ctx, oldToken)
	if claims == nil {
		return nil, gerror.New("invalid_token")
	}

	// 检查是否需要刷新（距离过期时间小于1小时才刷新）
	if time.Until(claims.ExpiresAt.Time) > time.Hour {
		g.Log().Info(ctx, "Token refresh skipped - not near expiration", g.Map{
			"user_id":    claims.SysUser.Id,
			"expires_at": claims.ExpiresAt.Time,
		})
		return &sys_model.TokenInfo{Token: oldToken}, nil
	}

	// 立即撤销旧token
	sys_service.Jwt().RevokeToken(ctx, oldToken)

	// 生成新token
	newTokenInfo, err := sys_service.Jwt().GenerateToken(ctx, &claims.SysUser)
	if err != nil {
		return nil, gerror.Wrap(err, "failed_to_generate_new_token")
	}

	// 添加到生命周期管理
	s.securityManager.lifecycleManager.AddToken(claims.SysUser.Id, newTokenInfo.Token)

	g.Log().Info(ctx, "Token rotated successfully", g.Map{
		"user_id":   claims.SysUser.Id,
		"old_token": s.hashToken(oldToken),
		"new_token": s.hashToken(newTokenInfo.Token),
	})

	return newTokenInfo, nil
}

// GenerateEnhancedToken 生成增强Token（包含设备指纹等安全信息）
func (s *EnhancedJwtService) GenerateEnhancedToken(ctx context.Context, user *sys_model.SysUser, r *ghttp.Request) (*sys_model.TokenInfo, error) {
	user.Password = ""

	// 生成设备指纹
	deviceFingerprint := GenerateDeviceFingerprint(r)
	
	// 创建增强的Claims
	customClaims := &EnhancedJwtCustomClaims{
		JwtCustomClaims: sys_model.JwtCustomClaims{
			SysUser:      *user,
			IsSuperAdmin: user.Type == 1, // 假设1为超级管理员类型
			IsAdmin:      user.Type == 2, // 假设2为管理员类型
			RegisteredClaims: jwt.RegisteredClaims{
				NotBefore: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
		DeviceFingerprint: deviceFingerprint,
		LoginIP:          r.GetClientIp(),
		UserAgent:        r.Header.Get("User-Agent"),
		SessionID:        s.generateSessionID(user.Id),
		TokenVersion:     1,
		SecurityLevel:    s.calculateSecurityLevel(user),
	}

	// 使用JWT库生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	
	// 获取签名密钥（这里需要从配置或服务中获取）
	signingKey := []byte("your-secret-key") // 实际应用中应从配置获取
	
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return nil, gerror.Wrap(err, "failed_to_sign_token")
	}

	tokenInfo := &sys_model.TokenInfo{
		Token:    tokenString,
		ExpireAt: customClaims.ExpiresAt.Time,
	}

	// 添加到生命周期管理
	s.securityManager.lifecycleManager.AddToken(user.Id, tokenString)

	g.Log().Info(ctx, "Enhanced token generated", g.Map{
		"user_id":            user.Id,
		"device_fingerprint": deviceFingerprint,
		"login_ip":           customClaims.LoginIP,
		"security_level":     customClaims.SecurityLevel,
	})

	return tokenInfo, nil
}

// ValidateEnhancedToken 验证增强Token
func (s *EnhancedJwtService) ValidateEnhancedToken(ctx context.Context, tokenString string, r *ghttp.Request) (*EnhancedJwtCustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &EnhancedJwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // 实际应用中应从配置获取
	})

	if err != nil {
		return nil, gerror.Wrap(err, "failed_to_parse_token")
	}

	claims, ok := token.Claims.(*EnhancedJwtCustomClaims)
	if !ok || !token.Valid {
		return nil, gerror.New("invalid_token_claims")
	}

	// 验证设备指纹
	currentFingerprint := GenerateDeviceFingerprint(r)
	if claims.DeviceFingerprint != currentFingerprint {
		s.securityManager.eventMonitor.RecordSuspiciousActivity(claims.SysUser.Id, SecurityEvent{
			Type:      "device_fingerprint_mismatch",
			Timestamp: time.Now(),
			IP:        r.GetClientIp(),
			UserAgent: r.Header.Get("User-Agent"),
			Details:   fmt.Sprintf("Expected: %s, Got: %s", claims.DeviceFingerprint, currentFingerprint),
		})
		return nil, gerror.New("device_fingerprint_mismatch")
	}

	// 验证IP地址（高安全级别用户）
	currentIP := r.GetClientIp()
	if claims.SecurityLevel >= 3 && claims.LoginIP != currentIP {
		s.securityManager.eventMonitor.RecordSuspiciousActivity(claims.SysUser.Id, SecurityEvent{
			Type:      "ip_address_changed",
			Timestamp: time.Now(),
			IP:        currentIP,
			UserAgent: r.Header.Get("User-Agent"),
			Details:   fmt.Sprintf("Original IP: %s, New IP: %s", claims.LoginIP, currentIP),
		})
		return nil, gerror.New("ip_address_changed_high_security")
	}

	// Token使用频率限制
	tokenHash := s.hashToken(tokenString)
	if !s.securityManager.rateLimiter.IsAllowed(tokenHash) {
		s.securityManager.eventMonitor.RecordSuspiciousActivity(claims.SysUser.Id, SecurityEvent{
			Type:      "rate_limit_exceeded",
			Timestamp: time.Now(),
			IP:        currentIP,
			UserAgent: r.Header.Get("User-Agent"),
			Details:   "Token usage rate limit exceeded",
		})
		return nil, gerror.New("too_many_requests")
	}

	return claims, nil
}

// generateSessionID 生成会话ID
func (s *EnhancedJwtService) generateSessionID(userID int64) string {
	data := fmt.Sprintf("%d-%d", userID, time.Now().UnixNano())
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// calculateSecurityLevel 计算用户安全级别
func (s *EnhancedJwtService) calculateSecurityLevel(user *sys_model.SysUser) int {
	// 根据用户类型和其他因素计算安全级别
	if user.Type == 1 { // 超级管理员
		return 5
	} else if user.Type == 2 { // 管理员
		return 4
	} else if user.Type == 3 { // 普通用户
		return 2
	}
	return 1 // 默认安全级别
}

// hashToken 对Token进行哈希处理
func (s *EnhancedJwtService) hashToken(token string) string {
	hash := md5.Sum([]byte(token))
	return hex.EncodeToString(hash[:])
}

// RevokeAllUserTokens 撤销用户所有Token
func (s *EnhancedJwtService) RevokeAllUserTokens(ctx context.Context, userID int64) error {
	return s.securityManager.lifecycleManager.RevokeAllUserTokens(userID)
}

// GetUserTokenCount 获取用户当前活跃Token数量
func (s *EnhancedJwtService) GetUserTokenCount(userID int64) int {
	return s.securityManager.lifecycleManager.GetUserTokenCount(userID)
}