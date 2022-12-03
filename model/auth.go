package model

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyAuth "github.com/SupenBysz/gf-admin-community/model/enum/auth"
	userType "github.com/SupenBysz/gf-admin-community/model/enum/user_type"
	"time"
)

type LoginInfo struct {
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"登录密码"`
	Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}
type LoginByMobileInfo struct {
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Mobile   string `json:"mobile" v:"phone|required-without:email#邮箱或手机号至少写一个" dc:"手机号"`
	Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}

type TokenInfo struct {
	Token    string    `json:"token" dc:"Token"`
	ExpireAt time.Time `json:"expireAt" dc:"Expire"`
}

type ForgotPassword struct {
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Captcha  string `json:"captcha" v:"required#验证吗不能为空" dc:"验证码"`
	Mobile   string `json:"mobile" v:"phone|required-without:email#邮箱或手机号至少写一个" dc:"手机号"`
	Email    string `json:"email" v:"email|required-without:mobile#邮箱或手机号至少写一个" dc:"邮箱"'`
}

type AuthHookFunc func(ctx context.Context, state kyAuth.ActionType, info entity.SysUser) error
type AuthHookInfo struct {
	Key      kyAuth.ActionType
	Value    AuthHookFunc
	UserType userType.Code `json:"userType" dc:"用户类型"`
}
