// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IJwt interface {
		// InstallHook 安装Hook
		InstallHook(userType sys_enum.UserType, hookFunc sys_hook.JwtHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除所有Hook
		CleanAllHook()
		// GenerateToken 创建一个token
		GenerateToken(ctx context.Context, user *sys_model.SysUser) (response *sys_model.TokenInfo, err error)
		// CreateToken 创建一个token
		CreateToken(claims *sys_model.JwtCustomClaims) (string, error)
		// RevokeToken 注销token
		RevokeToken(ctx context.Context, tokenString string) error
		// IsTokenRevoked 检查token是否已被注销
		IsTokenRevoked(ctx context.Context, tokenString string) bool
		// RefreshToken 刷新Token,并发安全
		RefreshToken(oldToken string, claims *sys_model.JwtCustomClaims) (string, error)
		// Middleware 鉴权中间件函数
		Middleware(r *ghttp.Request)
		// MakeSession 构建会话
		MakeSession(ctx context.Context, tokenString string) *sys_model.JwtCustomClaims
		
		// 增强安全功能
		// GenerateEnhancedToken 生成增强安全token
		GenerateEnhancedToken(ctx context.Context, user *sys_model.SysUser, deviceFingerprint, userAgent, ip string) (response *sys_model.TokenInfo, err error)
		// RefreshTokenWithRotation 带轮换的token刷新
		RefreshTokenWithRotation(ctx context.Context, oldToken, deviceFingerprint, userAgent, ip string) (response *sys_model.TokenInfo, err error)
		// RevokeAllUserTokens 撤销用户所有token
		RevokeAllUserTokens(ctx context.Context, userId int64) error
		// ValidateTokenSecurity 验证token安全性
		ValidateTokenSecurity(ctx context.Context, tokenString, deviceFingerprint, userAgent, ip string) (bool, error)
		// GetActiveTokenCount 获取用户活跃token数量
		GetActiveTokenCount(ctx context.Context, userId int64) (int, error)
	}
)

var (
	localJwt IJwt
)

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}
