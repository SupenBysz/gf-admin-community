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
	Mobile   int64  `json:"mobile" v:"required|phone#请数据手机号|手机号错误" dc:"手机号"`
	Captcha  string `json:"captcha" v:"required#请输入手机验证码"`
	Password string `json:"password" v:"required#请输入账号密码" dc:"登录密码"`
}
