package sysapi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/SupenBysz/gf-admin-community/model"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"登录" tags:"鉴权"`
	model.LoginInfo
}

type LoginRes model.TokenInfo

type RegisterReq struct {
	g.Meta `path:"/register" method:"post" summary:"注册" tags:"鉴权"`
	model.SysUserRegister
}

type ForgotPasswordReq struct {
	g.Meta `path:"/forgotPassword" method:"post" summary:"忘记密码" tags:"鉴权"`
	model.ForgotPassword
}
type ForgotPasswordRes struct {
	IdKey int64 `json:"id" dc:"ResetPassword 接口需要的key"`
}

type ResetPasswordReq struct {
	g.Meta   `path:"/resetPasswordReq" method:"post" summary:"重置密码" tags:"鉴权"`
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"登录密码"`
	IdKey    string `json:"key" v:"required#请输入KEY" dc:"KEY，通过ForgotPassword结构获取"`
}
