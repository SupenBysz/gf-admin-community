package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	sys_api "github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// Auth 鉴权
var Auth = cAuth{}

type cAuth struct{}

// Login 登陆
func (a *cAuth) Login(ctx context.Context, req *sys_api.LoginReq) (res *sys_api.LoginRes, err error) {
	result, err := sys_service.SysAuth().Login(ctx, req.LoginInfo)
	if err != nil {
		return nil, err
	}

	return (*sys_api.LoginRes)(result), nil
}

// LoginByMobile 通过手机号码+验证码登陆
func (c *cAuth) LoginByMobile(ctx context.Context, req *sys_api.LoginByMobileReq) (res *sys_api.LoginByMobileRes, err error) {
	// 获取
	result, err := sys_service.SysAuth().LoginByMobile(ctx, req.LoginByMobileInfo)
	if err != nil {
		return nil, err
	}

	return (*sys_api.LoginByMobileRes)(result), nil
}

// Register 注册
func (a *cAuth) Register(ctx context.Context, req *sys_api.RegisterReq) (res api_v1.BoolRes, err error) {
	_, err = sys_service.SysAuth().Register(ctx, req.SysUserRegister)

	if err != nil {
		return false, err
	}
	return true, nil
}

// ForgotPassword 忘记密码
func (a *cAuth) ForgotPassword(ctx context.Context, req *sys_api.ForgotPasswordReq) (res *sys_api.ForgotPasswordRes, err error) {
	result, err := sys_service.SysAuth().ForgotPassword(ctx, req.ForgotPassword)

	if err != nil {
		return nil, err
	}
	res = &sys_api.ForgotPasswordRes{IdKey: result}
	return res, nil
}

// ResetPassword 重置密码
func (a *cAuth) ResetPassword(ctx context.Context, req *sys_api.ResetPasswordReq) (res api_v1.BoolRes, err error) {
	_, err = sys_service.SysAuth().ResetPassword(ctx, req.Password, req.ConfirmPassword, req.IdKey)
	if err != nil {
		return false, err
	}
	return true, nil
}
