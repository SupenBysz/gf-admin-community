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
)

type (
	ISysAuth interface {
		// InstallHook 安装Hook
		InstallHook(actionType sys_enum.AuthActionType, userType sys_enum.UserType, hookFunc sys_hook.AuthHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除所有Hook
		CleanAllHook()
		// Login 登陆
		Login(ctx context.Context, req sys_model.LoginInfo, needCaptcha ...bool) (*sys_model.TokenInfo, error)
		// InnerLogin 内部登录，无需校验验证码和密码
		InnerLogin(ctx context.Context, user *sys_model.SysUser) (*sys_model.TokenInfo, error)
		// LoginByMobile 手机号 + 验证码登陆
		LoginByMobile(ctx context.Context, info sys_model.LoginByMobileInfo) (*sys_model.LoginByMobileRes, error)
		// LoginByMail 邮箱 + 密码登陆 (如果指定用户名，代表明确知道要登陆的是哪一个账号)
		LoginByMail(ctx context.Context, info sys_model.LoginByMailInfo) (*sys_model.LoginByMailRes, error)
		// Register 注册账号 (用户名+密码+图形验证码)
		Register(ctx context.Context, info sys_model.SysUserRegister) (*sys_model.SysUser, error)
		// RegisterByMobileOrMail 注册账号 (用户名+密码+ 手机号+验证码 或者 用户名+密码+ 邮箱+验证码)
		RegisterByMobileOrMail(ctx context.Context, info sys_model.SysUserRegisterByMobileOrMail) (res *sys_model.SysUser, err error)
		// ForgotPassword 忘记密码
		ForgotPassword(ctx context.Context, info sys_model.ForgotPassword) (int64, error)
		// ResetPassword 重置密码
		ResetPassword(ctx context.Context, password string, confirmPassword string, idKey string) (bool, error)
	}
)

var (
	localSysAuth ISysAuth
)

func SysAuth() ISysAuth {
	if localSysAuth == nil {
		panic("implement not found for interface ISysAuth, forgot register?")
	}
	return localSysAuth
}

func RegisterSysAuth(i ISysAuth) {
	localSysAuth = i
}
