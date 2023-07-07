package sys_model

import (
	"time"
)

type LoginInfo struct {
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"登录密码"`
	Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}

type LoginByMobileInfo struct {
	Username string `json:"username"  dc:"登录账号,会检验该手机号有几个账号，多个会返回userList，针对多账号请求需要携带userName"`
	Mobile   string `json:"mobile" v:"phone|required-without:email#邮箱或手机号至少写一个" dc:"手机号"`
	Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}

type LoginByMobileRes struct {
	SysUserListRes
	TokenInfo
}

type LoginByMailRes struct {
	SysUserListRes
	TokenInfo
}

type LoginByMailInfo struct {
	Username string `json:"username"  dc:"登录账号,会检验该邮箱有几个账号，多个会返回userList，针对多账号请求需要携带userName"`
	Mail     string `json:"mail" v:"required|email#邮箱不能为空" dc:"邮箱"`
	PassWord string `json:"passWord" v:"required#请输入密码" dc:"密码"`
}

type TokenInfo struct {
	Token    string    `json:"token" dc:"Token"`
	ExpireAt time.Time `json:"expireAt" dc:"Expire"`
}

type ForgotPassword struct {
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Captcha  string `json:"captcha" v:"required#验证吗不能为空" dc:"验证码"`
	Mobile   string `json:"mobile" v:"phone|required-without:email#邮箱或手机号至少写一个" dc:"手机号"`
}
