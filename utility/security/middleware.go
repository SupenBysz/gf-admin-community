package security

import (
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// SecurityMiddleware 安全中间件
type SecurityMiddleware struct {
	tokenSecurityManager *TokenSecurityManager
	enhancedJwtService   *EnhancedJwtService
}

// NewSecurityMiddleware 创建安全中间件
func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{
		tokenSecurityManager: NewTokenSecurityManager(),
		enhancedJwtService:   NewEnhancedJwtService(),
	}
}

// TokenSecurityHandler Token安全处理中间件
func (sm *SecurityMiddleware) TokenSecurityHandler(r *ghttp.Request) {
	// 跳过不需要认证的路径
	if sm.shouldSkipAuth(r) {
		r.Middleware.Next()
		return
	}

	tokenString := sm.extractToken(r)
	if tokenString == "" {
		response.JsonExit(r, 401, "missing_token")
		return
	}

	// 基础Token验证
	claims := sys_service.Jwt().MakeSession(r.Context(), tokenString)
	if claims == nil {
		response.JsonExit(r, 401, "invalid_token")
		return
	}

	// 设备指纹验证
	deviceFingerprint := r.Header.Get("X-Device-Fingerprint")
	if deviceFingerprint != "" {
		if err := sm.tokenSecurityManager.ValidateDeviceFingerprint(r.Context(), claims, deviceFingerprint); err != nil {
			g.Log().Warning(r.Context(), "Device fingerprint validation failed", g.Map{
				"user_id": claims.SysUser.Id,
				"error":   err.Error(),
			})
			response.JsonExit(r, 403, "device_verification_failed")
			return
		}
	}

	// IP地址验证
	clientIP := r.GetClientIp()
	if err := sm.tokenSecurityManager.ValidateIPAddress(r.Context(), claims, clientIP); err != nil {
		g.Log().Warning(r.Context(), "IP address validation failed", g.Map{
			"user_id": claims.SysUser.Id,
			"error":   err.Error(),
		})
		response.JsonExit(r, 403, "ip_verification_failed")
		return
	}

	// Token使用频率限制
	tokenHash := sm.tokenSecurityManager.hashToken(tokenString)
	if !sm.tokenSecurityManager.rateLimiter.IsAllowed(tokenHash) {
		// 记录频率限制事件
		sm.tokenSecurityManager.eventMonitor.RecordSuspiciousActivity(claims.SysUser.Id, SecurityEvent{
			Type:      "rate_limit_exceeded",
			Timestamp: time.Now(),
			IP:        clientIP,
			UserAgent: r.Header.Get("User-Agent"),
			Details:   "Token usage rate limit exceeded",
		})

		response.JsonExit(r, 429, "too_many_requests")
		return
	}

	// 记录正常访问
	g.Log().Debug(r.Context(), "Token security validation passed", g.Map{
		"user_id":   claims.SysUser.Id,
		"client_ip": clientIP,
		"path":      r.URL.Path,
	})

	r.Middleware.Next()
}

// EnhancedTokenHandler 增强Token处理中间件（用于支持增强Token的场景）
func (sm *SecurityMiddleware) EnhancedTokenHandler(r *ghttp.Request) {
	// 跳过不需要认证的路径
	if sm.shouldSkipAuth(r) {
		r.Middleware.Next()
		return
	}

	tokenString := sm.extractToken(r)
	if tokenString == "" {
		response.JsonExit(r, 401, "missing_token")
		return
	}

	// 使用增强JWT服务验证Token
	claims, err := sm.enhancedJwtService.ValidateEnhancedToken(r.Context(), tokenString, r)
	if err != nil {
		g.Log().Warning(r.Context(), "Enhanced token validation failed", g.Map{
			"error": err.Error(),
			"ip":    r.GetClientIp(),
		})
		response.JsonExit(r, 401, "token_validation_failed")
		return
	}

	// 将claims信息存储到上下文中
	r.SetCtxVar("enhanced_claims", claims)

	g.Log().Debug(r.Context(), "Enhanced token validation passed", g.Map{
		"user_id":            claims.SysUser.Id,
		"device_fingerprint": claims.DeviceFingerprint,
		"security_level":     claims.SecurityLevel,
	})

	r.Middleware.Next()
}

// extractToken 提取Token
func (sm *SecurityMiddleware) extractToken(r *ghttp.Request) string {
	tokenString := gstr.Trim(r.Header.Get("Authorization"))
	
	// 从Header中获取
	if tokenString != "" {
		if gstr.HasPrefix(tokenString, "Bearer ") {
			return gstr.SubStr(tokenString, 7)
		}
		return tokenString
	}

	// 从GET参数中获取
	if gstr.ToUpper(r.Method) == "GET" {
		return r.Get("token", "").String()
	}

	return ""
}

// shouldSkipAuth 判断是否跳过认证
func (sm *SecurityMiddleware) shouldSkipAuth(r *ghttp.Request) bool {
	// 定义不需要认证的路径
	skipPaths := []string{
		"/api/v1/login",
		"/api/v1/register",
		"/api/v1/captcha",
		"/api/v1/forgot-password",
		"/api/v1/health",
		"/api/v1/ping",
	}

	path := r.URL.Path
	for _, skipPath := range skipPaths {
		if gstr.HasPrefix(path, skipPath) {
			return true
		}
	}

	return false
}

// GetTokenSecurityManager 获取Token安全管理器
func (sm *SecurityMiddleware) GetTokenSecurityManager() *TokenSecurityManager {
	return sm.tokenSecurityManager
}

// GetEnhancedJwtService 获取增强JWT服务
func (sm *SecurityMiddleware) GetEnhancedJwtService() *EnhancedJwtService {
	return sm.enhancedJwtService
}

// SecurityConfig 安全配置
type SecurityConfig struct {
	EnableDeviceFingerprint bool     `json:"enable_device_fingerprint"` // 启用设备指纹验证
	EnableIPValidation      bool     `json:"enable_ip_validation"`      // 启用IP验证
	EnableRateLimit         bool     `json:"enable_rate_limit"`         // 启用频率限制
	MaxTokensPerUser        int      `json:"max_tokens_per_user"`       // 每用户最大Token数量
	RateLimitRequests       int      `json:"rate_limit_requests"`       // 频率限制请求数
	RateLimitWindow         int      `json:"rate_limit_window"`         // 频率限制时间窗口（秒）
	AlertThreshold          int      `json:"alert_threshold"`           // 告警阈值
	SkipAuthPaths           []string `json:"skip_auth_paths"`           // 跳过认证的路径
}

// DefaultSecurityConfig 默认安全配置
func DefaultSecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		EnableDeviceFingerprint: false, // 默认关闭，需要前端支持
		EnableIPValidation:      true,  // 启用IP验证
		EnableRateLimit:         true,  // 启用频率限制
		MaxTokensPerUser:        5,     // 每用户最多5个Token
		RateLimitRequests:       100,   // 每分钟最多100次请求
		RateLimitWindow:         60,    // 60秒时间窗口
		AlertThreshold:          5,     // 5次可疑活动触发告警
		SkipAuthPaths: []string{
			"/api/v1/login",
			"/api/v1/register",
			"/api/v1/captcha",
			"/api/v1/forgot-password",
			"/api/v1/health",
			"/api/v1/ping",
		},
	}
}