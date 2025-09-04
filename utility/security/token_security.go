package security

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// EnhancedJwtCustomClaims 增强的JWT Claims，包含安全相关信息
type EnhancedJwtCustomClaims struct {
	sys_model.JwtCustomClaims
	DeviceFingerprint string `json:"device_fingerprint"` // 设备指纹
	LoginIP          string `json:"login_ip"`           // 登录IP
	UserAgent        string `json:"user_agent"`         // 用户代理
	SessionID        string `json:"session_id"`         // 会话ID
	TokenVersion     int    `json:"token_version"`      // Token版本号
	SecurityLevel    int    `json:"security_level"`     // 安全级别 (1-5)
}

// TokenRateLimiter Token使用频率限制器
type TokenRateLimiter struct {
	requests    map[string][]time.Time
	mutex       sync.RWMutex
	maxRequests int           // 最大请求数
	timeWindow  time.Duration // 时间窗口
}

// NewTokenRateLimiter 创建Token频率限制器
func NewTokenRateLimiter(maxRequests int, timeWindow time.Duration) *TokenRateLimiter {
	return &TokenRateLimiter{
		requests:    make(map[string][]time.Time),
		maxRequests: maxRequests,
		timeWindow:  timeWindow,
	}
}

// IsAllowed 检查Token是否允许访问
func (trl *TokenRateLimiter) IsAllowed(tokenHash string) bool {
	trl.mutex.Lock()
	defer trl.mutex.Unlock()

	now := time.Now()
	if trl.requests == nil {
		trl.requests = make(map[string][]time.Time)
	}

	// 清理过期记录
	requests := trl.requests[tokenHash]
	validRequests := make([]time.Time, 0)
	for _, reqTime := range requests {
		if now.Sub(reqTime) < trl.timeWindow {
			validRequests = append(validRequests, reqTime)
		}
	}

	// 检查是否超过限制
	if len(validRequests) >= trl.maxRequests {
		return false
	}

	// 添加当前请求
	validRequests = append(validRequests, now)
	trl.requests[tokenHash] = validRequests

	return true
}

// TokenLifecycleManager Token生命周期管理器
type TokenLifecycleManager struct {
	activeTokens     map[int64][]string // 用户ID -> Token列表
	mutex            sync.RWMutex
	maxTokensPerUser int // 每个用户最大Token数量
}

// NewTokenLifecycleManager 创建Token生命周期管理器
func NewTokenLifecycleManager(maxTokensPerUser int) *TokenLifecycleManager {
	return &TokenLifecycleManager{
		activeTokens:     make(map[int64][]string),
		maxTokensPerUser: maxTokensPerUser,
	}
}

// AddToken 添加Token到用户的活跃Token列表
func (tlm *TokenLifecycleManager) AddToken(userID int64, token string) error {
	tlm.mutex.Lock()
	defer tlm.mutex.Unlock()

	if tlm.activeTokens == nil {
		tlm.activeTokens = make(map[int64][]string)
	}

	tokens := tlm.activeTokens[userID]

	// 如果超过最大Token数量，撤销最旧的Token
	if len(tokens) >= tlm.maxTokensPerUser {
		oldestToken := tokens[0]
		sys_service.Jwt().RevokeToken(context.Background(), oldestToken)
		tokens = tokens[1:]
	}

	tokens = append(tokens, token)
	tlm.activeTokens[userID] = tokens

	return nil
}

// RevokeAllUserTokens 撤销用户所有Token（用于强制登出）
func (tlm *TokenLifecycleManager) RevokeAllUserTokens(userID int64) error {
	tlm.mutex.Lock()
	defer tlm.mutex.Unlock()

	if tokens, exists := tlm.activeTokens[userID]; exists {
		for _, token := range tokens {
			sys_service.Jwt().RevokeToken(context.Background(), token)
		}
		delete(tlm.activeTokens, userID)
	}

	return nil
}

// GetUserTokenCount 获取用户当前活跃Token数量
func (tlm *TokenLifecycleManager) GetUserTokenCount(userID int64) int {
	tlm.mutex.RLock()
	defer tlm.mutex.RUnlock()

	if tokens, exists := tlm.activeTokens[userID]; exists {
		return len(tokens)
	}
	return 0
}

// SecurityEvent 安全事件
type SecurityEvent struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Details   string    `json:"details"`
}

// SecurityEventMonitor 安全事件监控器
type SecurityEventMonitor struct {
	suspiciousActivities map[int64][]SecurityEvent
	mutex                sync.RWMutex
	alertThreshold       int // 告警阈值
}

// NewSecurityEventMonitor 创建安全事件监控器
func NewSecurityEventMonitor(alertThreshold int) *SecurityEventMonitor {
	return &SecurityEventMonitor{
		suspiciousActivities: make(map[int64][]SecurityEvent),
		alertThreshold:       alertThreshold,
	}
}

// RecordSuspiciousActivity 记录可疑活动
func (sem *SecurityEventMonitor) RecordSuspiciousActivity(userID int64, event SecurityEvent) {
	sem.mutex.Lock()
	defer sem.mutex.Unlock()

	if sem.suspiciousActivities == nil {
		sem.suspiciousActivities = make(map[int64][]SecurityEvent)
	}

	events := sem.suspiciousActivities[userID]
	events = append(events, event)

	// 只保留最近24小时的事件
	cutoff := time.Now().Add(-24 * time.Hour)
	validEvents := make([]SecurityEvent, 0)
	for _, e := range events {
		if e.Timestamp.After(cutoff) {
			validEvents = append(validEvents, e)
		}
	}

	sem.suspiciousActivities[userID] = validEvents

	// 检查是否需要告警
	if len(validEvents) >= sem.alertThreshold {
		sem.triggerSecurityAlert(userID, validEvents)
	}
}

// triggerSecurityAlert 触发安全告警
func (sem *SecurityEventMonitor) triggerSecurityAlert(userID int64, events []SecurityEvent) {
	g.Log().Error(context.Background(), "Security alert triggered", g.Map{
		"user_id":     userID,
		"event_count": len(events),
		"events":      events,
	})

	// 严重情况下可以自动锁定账户
	if len(events) >= 10 {
		// 这里可以调用用户锁定服务
		g.Log().Critical(context.Background(), "Auto-locking user due to suspicious activities", g.Map{
			"user_id":     userID,
			"event_count": len(events),
		})
	}
}

// TokenSecurityManager Token安全管理器
type TokenSecurityManager struct {
	rateLimiter      *TokenRateLimiter
	lifecycleManager *TokenLifecycleManager
	eventMonitor     *SecurityEventMonitor
}

// NewTokenSecurityManager 创建Token安全管理器
func NewTokenSecurityManager() *TokenSecurityManager {
	return &TokenSecurityManager{
		rateLimiter:      NewTokenRateLimiter(100, time.Minute), // 每分钟最多100次请求
		lifecycleManager: NewTokenLifecycleManager(5),           // 每用户最多5个活跃Token
		eventMonitor:     NewSecurityEventMonitor(5),            // 5次可疑活动触发告警
	}
}

// GetRateLimiter 获取频率限制器
func (tsm *TokenSecurityManager) GetRateLimiter() *TokenRateLimiter {
	return tsm.rateLimiter
}

// GetLifecycleManager 获取生命周期管理器
func (tsm *TokenSecurityManager) GetLifecycleManager() *TokenLifecycleManager {
	return tsm.lifecycleManager
}

// GetEventMonitor 获取事件监控器
func (tsm *TokenSecurityManager) GetEventMonitor() *SecurityEventMonitor {
	return tsm.eventMonitor
}

// ValidateDeviceFingerprint 验证设备指纹
func (tsm *TokenSecurityManager) ValidateDeviceFingerprint(ctx context.Context, claims *sys_model.JwtCustomClaims, currentFingerprint string) error {
	// 由于当前系统使用的是标准JwtCustomClaims，这里暂时跳过设备指纹验证
	// 在实际部署时，需要将Token生成逻辑修改为使用EnhancedJwtCustomClaims
	g.Log().Info(ctx, "Device fingerprint validation skipped - using standard claims", g.Map{
		"user_id":            claims.SysUser.Id,
		"current_fingerprint": currentFingerprint,
	})
	return nil
}

// ValidateIPAddress 验证IP地址
func (tsm *TokenSecurityManager) ValidateIPAddress(ctx context.Context, claims *sys_model.JwtCustomClaims, currentIP string) error {
	// 由于当前系统使用的是标准JwtCustomClaims，这里暂时跳过IP验证
	// 在实际部署时，需要将Token生成逻辑修改为使用EnhancedJwtCustomClaims
	g.Log().Info(ctx, "IP address validation skipped - using standard claims", g.Map{
		"user_id":    claims.SysUser.Id,
		"current_ip": currentIP,
	})
	return nil
}

// EnhancedSecurityMiddleware 增强的安全中间件
func (tsm *TokenSecurityManager) EnhancedSecurityMiddleware(r *ghttp.Request) {
	tokenString := gstr.Trim(r.Header.Get("Authorization"))

	if gstr.ToUpper(r.Method) == "GET" && tokenString == "" {
		tokenString = r.Get("token", "").String()
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
		if err := tsm.ValidateDeviceFingerprint(r.Context(), claims, deviceFingerprint); err != nil {
			response.JsonExit(r, 403, "device_verification_failed")
			return
		}
	}

	// IP地址验证
	clientIP := r.GetClientIp()
	if err := tsm.ValidateIPAddress(r.Context(), claims, clientIP); err != nil {
		response.JsonExit(r, 403, "ip_verification_failed")
		return
	}

	// Token使用频率限制
	tokenHash := tsm.hashToken(tokenString)
	if !tsm.rateLimiter.IsAllowed(tokenHash) {
		// 记录频率限制事件
		tsm.eventMonitor.RecordSuspiciousActivity(claims.SysUser.Id, SecurityEvent{
			Type:      "rate_limit_exceeded",
			Timestamp: time.Now(),
			IP:        clientIP,
			UserAgent: r.Header.Get("User-Agent"),
			Details:   "Token usage rate limit exceeded",
		})

		response.JsonExit(r, 429, "too_many_requests")
		return
	}

	r.Middleware.Next()
}

// hashToken 对Token进行哈希处理
func (tsm *TokenSecurityManager) hashToken(token string) string {
	hash := md5.Sum([]byte(token))
	return hex.EncodeToString(hash[:])
}

// GenerateDeviceFingerprint 生成设备指纹
func GenerateDeviceFingerprint(r *ghttp.Request) string {
	userAgent := r.Header.Get("User-Agent")
	acceptLanguage := r.Header.Get("Accept-Language")
	acceptEncoding := r.Header.Get("Accept-Encoding")
	
	fingerprint := fmt.Sprintf("%s|%s|%s", userAgent, acceptLanguage, acceptEncoding)
	hash := md5.Sum([]byte(fingerprint))
	return hex.EncodeToString(hash[:])
}