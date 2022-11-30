package sysapi

import (
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateUserReq struct {
	g.Meta `path:"/createUser" method:"post" summary:"新增用户|信息" tags:"用户"`
	model.UserInnerRegister
}

type QueryUserListReq struct {
	g.Meta `path:"/queryUserList" method:"post" summary:"获取用户|列表" tags:"用户"`
	model.SearchParams
}

type SetUserRoleIdsReq struct {
	g.Meta  `path:"/setUserRoleIds" method:"post" summary:"设置用户角色" tags:"用户"`
	RoleIds []int64 `json:"roleIds" v:"required#角色ID校验失败" dc:"角色ID数组"`
	UserId  int64   `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
}

type SetUserPermissionIdsReq struct {
	g.Meta        `path:"/setUserPermissionIds" method:"post" summary:"设置用户权限" tags:"用户"`
	Id            int64   `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
	PermissionIds []int64 `json:"permissionIds" v:"array#权限ID集合参数无效" dc:"权限ID集合"`
}

type GetUserPermissionIdsReq struct {
	g.Meta `path:"/getUserPermissionIds" method:"post" summary:"获取用户权限|ID数组" tags:"用户"`
	Id     int64 `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
}

type SetUsernameByIdReq struct {
	g.Meta      `path:"/setUsername" method:"post" summary:"设置用户登陆名称" tags:"用户"`
	NewUsername string `json:"newUsername" v:"required#新用户名称" dc:"新的用户名称"`
}

type UserInfoRes entity.SysUser
type UserInfoListRes model.CollectRes[entity.SysUser]
