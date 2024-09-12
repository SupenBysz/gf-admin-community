package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type SetUsernameByIdReq struct {
	g.Meta      `path:"/setUsername" method:"post" summary:"设置用户登陆名称" tags:"我的"`
	NewUsername string `json:"newUsername" v:"required#新用户名称" dc:"新的用户名称"`
}

type UpdateUserPasswordReq struct {
	g.Meta `path:"/updateUserPassword" method:"post" summary:"修改用户密码" tags:"我的"`
	sys_model.UpdateUserPassword
}

type SetUserMobileReq struct {
	g.Meta   `path:"/setUserMobile" method:"post" summary:"修改用户手机号" tags:"我的"`
	Mobile   string `json:"mobile" v:"required|phone#请输入手机号|手机号错误" dc:"手机号"`
	Captcha  string `json:"captcha" v:"required#请输入手机验证码"`
	Password string `json:"password" v:"required#请输入账号密码" dc:"登录密码" v:"min-length:6#密码最短为6位"`
}

type MyPermissionsReq struct {
	g.Meta `path:"/getPermissions" method:"post" summary:"我的权限|列表" tags:"我的"`
}

type MyMenusReq struct {
	g.Meta `path:"/getMenus" method:"post" summary:"我的菜单|树" tags:"我的"`
}

type MyPersonLicenseReq struct {
	g.Meta `path:"/myPersonLicense" method:"post" summary:"我的个人资质" tags:"我的" dc:"返回的是当前登陆用户，最新且正在生效的资质"`
}

type MyPersonLicenseAuditReq struct {
	g.Meta `path:"/myPersonLicenseAudit" method:"post" summary:"获取最后一次提交的我个人资质审核信息" tags:"我的" dc:"返回的是当前登陆用户，最后一次提交的我个人资质审核信息"`
}

type SetUserMailReq struct {
	g.Meta   `path:"/setUserMail" method:"post" summary:"修改用户邮箱" tags:"我的"`
	OldMail  string `json:"oldMail" v:"email#邮箱账号格式错误" dc:"原邮箱，首次设置原邮箱地址可为空"`
	NewMail  string `json:"newMail" v:"required|email#请输入新邮箱账号|邮箱账号格式错误" dc:"新邮箱"`
	Captcha  string `json:"captcha" v:"required#请输入邮箱验证码"`
	Password string `json:"password" v:"required#请输入账号密码" dc:"登录密码" v:"min-length:6#密码最短为6位"`
}
