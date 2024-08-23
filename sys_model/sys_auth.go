package sys_model

import (
	"time"
)

// 登陆支持方式： [1用户名 2手机号 4邮箱]  用户名+密码+图形验证码  手机号+密码或验证码  邮箱+密码或验证码 （OK）

// 注册支持格式： [1用户名 2手机号 4邮箱]  用户名+密码+图形验证码   用户名+手机号+验证码  用户名+邮箱+验证码

// 找回密码方式： [1手机号 2邮箱]  +用户名+验证码 （OK）

type LoginInfo struct {
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"登录密码" v:"min-length:6#密码最短为6位"`
	Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}

type LoginByMobileInfo struct {
	Username string `json:"username"  dc:"登录账号,会检验该手机号有几个账号，多个会返回userList，针对多账号请求需要携带userName"`
	Mobile   string `json:"mobile" v:"required|phone#手机号不能为空" dc:"手机号"`
	Captcha  string `json:"captcha" dc:"验证码，和密码二选一"`
	Password string `json:"password"  dc:"密码，和验证码二选一" v:"min-length:6#密码最短为6位"`
}

type LoginByMobileRes struct {
	UserList SysUserListRes `json:"userList" dc:"手机号关联的用户列表"`
	TokenInfo
	User *SysUser `json:"user"`
}

type LoginByMailInfo struct {
	Username string `json:"username"  dc:"登录账号,会检验该邮箱有几个账号，多个会返回userList，针对多账号请求需要携带userName"`
	Mail     string `json:"mail" v:"required|email#邮箱不能为空" dc:"邮箱"`
	Captcha  string `json:"captcha" dc:"验证码，和密码二选一"`
	Password string `json:"password" dc:"密码" v:"min-length:6#密码最短为6位"`
}

type LoginByMailRes struct {
	UserList SysUserListRes `json:"userList" dc:"邮箱关联的用户列表"`
	TokenInfo
	User *SysUser `json:"user"`
}

type LoginRes struct {
	TokenInfo
	User *SysUser `json:"user" dc:"用户信息"`
}

type TokenInfo struct {
	Token    string    `json:"token" dc:"Token"`
	ExpireAt time.Time `json:"expireAt" dc:"Expire"`
}

type ForgotPassword struct {
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Captcha  string `json:"captcha" v:"required#验证吗不能为空" dc:"验证码"`
	Mobile   string `json:"mobile" v:"required-with:phone|required-with:email#邮箱或手机号至少写一个" dc:"邮箱或手机号"`
}
