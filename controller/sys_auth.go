package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/service"
)

// Auth 鉴权
var Auth = cAuth{}

type cAuth struct{}

// Login 登陆
func (a *cAuth) Login(ctx context.Context, req *sysapi.LoginReq) (res *sysapi.LoginRes, err error) {
	result, err := service.SysAuth().Login(ctx, req.LoginInfo)
	if err != nil {
		return nil, err
	}

	return (*sysapi.LoginRes)(result), nil
}

// Register 注册
func (a *cAuth) Register(ctx context.Context, req *sysapi.RegisterReq) (res api_v1.BoolRes, err error) {
	_, err = service.SysAuth().Register(ctx, req.SysUserRegister)

	if err != nil {
		return false, err
	}
	return true, nil
}

// ForgotPassword 忘记密码
func (a *cAuth) ForgotPassword(ctx context.Context, req *sysapi.ForgotPasswordReq) (res *sysapi.ForgotPasswordRes, err error) {
	result, err := service.SysAuth().ForgotPassword(ctx, req.ForgotPassword)

	if err != nil {
		return nil, err
	}
	res = &sysapi.ForgotPasswordRes{IdKey: result}
	return res, nil
}

// ResetPassword 重置密码
func (a *cAuth) ResetPassword(ctx context.Context, req *sysapi.ResetPasswordReq) (res api_v1.BoolRes, err error) {
	_, err = service.SysAuth().ResetPassword(ctx, req.Password, req.Password, req.IdKey)
	if err != nil {
		return false, err
	}
	return true, nil
}
