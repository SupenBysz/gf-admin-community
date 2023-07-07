package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"登录" tags:"鉴权"`
	sys_model.LoginInfo
}

type LoginByMobileReq struct {
	g.Meta `path:"/loginByMobile" method:"post" summary:"手机号登录" tags:"鉴权"`
	sys_model.LoginByMobileInfo
}

type LoginByMailReq struct {
	g.Meta `path:"/loginByMail" method:"post" summary:"邮箱登录" tags:"鉴权"`
	sys_model.LoginByMailInfo
}

//type LoginByMobileRes sys_model.TokenInfo

type LoginRes sys_model.TokenInfo

type RegisterReq struct {
	g.Meta `path:"/register" method:"post" summary:"注册" tags:"鉴权"`
	sys_model.SysUserRegister
}

type ForgotPasswordReq struct {
	g.Meta `path:"/forgotPassword" method:"post" summary:"忘记密码" tags:"鉴权"`
	sys_model.ForgotPassword
}
type ForgotPasswordRes struct {
	IdKey int64 `json:"id" dc:"ResetPassword 接口需要的key"`
}

type ResetPasswordReq struct {
	g.Meta          `path:"/resetPassword" method:"post" summary:"重置密码" tags:"鉴权"`
	Password        string `json:"password" v:"required#请输入密码" dc:"登录密码"`
	ConfirmPassword string `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"确认密码"`
	IdKey           string `json:"idKey" v:"required#请输入KEY" dc:"KEY，通过ForgotPassword结构获取"`
}
