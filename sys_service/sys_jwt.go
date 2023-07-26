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
		// RefreshToken 刷新Token,并发安全
		RefreshToken(oldToken string, claims *sys_model.JwtCustomClaims) (string, error)
		Middleware(r *ghttp.Request)
		MakeSession(ctx context.Context, tokenString string) *sys_model.JwtCustomClaims
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
