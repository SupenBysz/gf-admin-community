package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type SysUserRegister struct {
	Username        string `json:"username" v:"required|length:4,30#请输入用户名称|用户名称长度非法"  dc:"登陆账号"`
	Password        string `json:"password" v:"required|password#请输入密码|密码长度非法"  dc:"密码"`
	ConfirmPassword string `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"密码"`
	Captcha         string `json:"captcha" v:"required" dc:"验证码"`
}

type UserInnerRegister struct {
	Username        string  `json:"username" v:"required|length:4,30#请输入用户名称|用户名称长度非法"  dc:"登陆账号"`
	Password        string  `json:"password" v:"required|password#请输入密码|密码长度非法"  dc:"密码"`
	ConfirmPassword string  `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"密码"`
	RoleIds         []int64 `json:"roleIds" dc:"所属角色，多个用逗号隔开"`
	Mobile          string  `json:"mobile"    dc:"手机号"`
}

type SysUserRegisterRes struct {
	UserInfo     SysUser               `json:"userInfo" dc:"用户信息"`
	RoleInfoList []*sys_entity.SysRole `json:"roleInfoList" dc:"角色信息列表"`
}

type SysUser struct {
	sys_entity.SysUser
	Detail    sys_entity.SysUserDetail `orm:"with:id" json:"detail"`
	RoleNames []string                 `json:"roleNames" dc:"所属角色"`
}

type UpdateUserPassword struct {
	OldPassword     string `json:"oldPassword" v:"required#请输入原始密码" dc:"旧密码"`
	Password        string `json:"password" v:"required#请输入新密码" dc:"新密码"`
	ConfirmPassword string `json:"confirmPassword" v:"required#请确认密码" dc:"确认密码"`
}

type UserInfo SysUser
type UserInfoList CollectRes[*SysUser]
type SysUserListRes CollectRes[*SysUser]

type UserHookFunc HookFunc[sys_enum.UserEvent, SysUser]
type UserHookInfo HookEventType[sys_enum.UserEvent, UserHookFunc]
