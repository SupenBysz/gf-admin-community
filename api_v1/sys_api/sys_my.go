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
