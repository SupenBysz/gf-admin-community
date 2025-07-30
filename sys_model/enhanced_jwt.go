package sys_model

import (
	"time"
)

// EnhancedJwtCustomClaims 增强的JWT自定义声明
type EnhancedJwtCustomClaims struct {
	JwtCustomClaims
	// 设备指纹，用于验证请求是否来自同一设备
	DeviceFingerprint string `json:"device_fingerprint,omitempty"`
	// 登录IP，用于验证请求是否来自同一IP或IP段
	LoginIP string `json:"login_ip,omitempty"`
	// 用户代理，用于验证请求是否来自同一浏览器或应用
	UserAgent string `json:"user_agent,omitempty"`
	// 会话ID，用于唯一标识一个会话
	SessionID string `json:"session_id,omitempty"`
	// Token版本，用于在需要时强制使所有token失效
	TokenVersion int `json:"token_version,omitempty"`
	// 安全级别，用于控制不同安全级别的操作
	SecurityLevel int `json:"security_level,omitempty"`
}

// EnhancedTokenInfo 增强的Token信息
type EnhancedTokenInfo struct {
	TokenInfo
	// 会话ID
	SessionID string `json:"session_id,omitempty"`
	// 安全级别
	SecurityLevel int `json:"security_level,omitempty"`
}

// SecurityEvent 安全事件记录
type SecurityEvent struct {
	// 用户ID
	UserID int64 `json:"user_id"`
	// 事件类型
	EventType string `json:"event_type"`
	// 事件描述
	Description string `json:"description"`
	// 事件时间
	EventTime time.Time `json:"event_time"`
	// 事件IP
	IP string `json:"ip"`
	// 设备指纹
	DeviceFingerprint string `json:"device_fingerprint"`
	// 用户代理
	UserAgent string `json:"user_agent"`
	// 严重程度 1-5
	Severity int `json:"severity"`
	// 相关Token
	TokenID string `json:"token_id,omitempty"`
}

// TokenSecurityConfig Token安全配置
type TokenSecurityConfig struct {
	// 是否启用设备指纹验证
	EnableDeviceValidation bool `json:"enable_device_validation"`
	// 是否启用IP验证
	EnableIPValidation bool `json:"enable_ip_validation"`
	// 是否启用Token轮换
	EnableTokenRotation bool `json:"enable_token_rotation"`
	// 是否启用Token频率限制
	EnableRateLimiting bool `json:"enable_rate_limiting"`
	// 是否启用安全事件监控
	EnableSecurityMonitoring bool `json:"enable_security_monitoring"`
	// 每个用户最大活跃Token数
	MaxActiveTokensPerUser int `json:"max_active_tokens_per_user"`
	// IP变化阈值（公里）
	IPChangeThreshold int `json:"ip_change_threshold"`
	// 安全事件警报阈值
	SecurityAlertThreshold int `json:"security_alert_threshold"`
	// Token轮换间隔（小时）
	TokenRotationInterval int `json:"token_rotation_interval"`
}