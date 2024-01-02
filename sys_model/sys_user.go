package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

// 注册支持格式： [1用户名 2手机号 4邮箱]  用户名+密码+图形验证码   用户名+手机号+验证码  用户名+邮箱+验证码

type SysUserRegister struct {
	Username        string `json:"username" v:"required|length:4,30#请输入用户名称|用户名称长度非法"  dc:"登陆账号"`
	Password        string `json:"password" v:"required|password#请输入密码|密码长度非法"  dc:"密码" v:"min-length:6#密码最短为6位"`
	ConfirmPassword string `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"密码" v:"min-length:6#密码最短为6位"`
	Captcha         string `json:"captcha" v:"required" dc:"验证码"`
	InviteCode      string `json:"inviteCode" dc:"邀约码"`
}

type UserInnerRegister struct {
	Username        string  `json:"username" v:"required|length:4,30#请输入用户名称|用户名称长度非法"  dc:"登陆账号"`
	Password        string  `json:"password" v:"required|password#请输入密码|密码长度非法"  dc:"密码" v:"min-length:6#密码最短为6位"`
	ConfirmPassword string  `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"密码" v:"min-length:6#密码最短为6位"`
	RoleIds         []int64 `json:"roleIds" dc:"所属角色，多个用逗号隔开"`
	Mobile          string  `json:"mobile"    dc:"手机号"`
	Email           string  `json:"email"     description:"邮箱"`
	InviteCode      string  `json:"inviteCode" dc:"邀约码"`
}

type SysUserRegisterByMobileOrMail struct {
	Username        string `json:"username" v:"required|length:4,30#请输入用户名称|用户名称长度非法"  dc:"登陆账号"`
	Password        string `json:"password" v:"required|password#请输入密码|密码长度非法"  dc:"密码" v:"min-length:6#密码最短为6位"`
	ConfirmPassword string `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"密码" v:"min-length:6#密码最短为6位"`

	MobileOrMail string `json:"mobileOrMail" v:"required-with:phone|required-with:email#邮箱或手机号至少写一个" dc:"邮箱或手机号"`
	Captcha      string `json:"captcha" v:"required" dc:"验证码"`
	InviteCode   string `json:"inviteCode" dc:"邀约码"`
}

type SysUserRegisterRes struct {
	UserInfo     SysUser               `json:"userInfo" dc:"用户信息"`
	RoleInfoList []*sys_entity.SysRole `json:"roleInfoList" dc:"角色信息列表"`
}

type SysUser struct {
	*sys_entity.SysUser
	Detail    *sys_entity.SysUserDetail `orm:"with:id" json:"detail"`
	RoleNames []string                  `json:"roleNames" dc:"所属角色"`
}

type UpdateUserPassword struct {
	OldPassword     string `json:"oldPassword" v:"required#请输入原始密码" dc:"旧密码" v:"min-length:6#密码最短为6位"`
	Password        string `json:"password" v:"required#请输入新密码" dc:"新密码" v:"min-length:6#密码最短为6位"`
	ConfirmPassword string `json:"confirmPassword" v:"required#请确认密码" dc:"确认密码" v:"min-length:6#密码最短为6位"`
}

type UserInfo SysUser
type UserInfoList base_model.CollectRes[*SysUser]
type SysUserListRes base_model.CollectRes[*SysUser]
