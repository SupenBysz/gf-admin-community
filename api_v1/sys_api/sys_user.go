package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateUserReq struct {
	g.Meta `path:"/createUser" method:"post" summary:"新增用户|信息" tags:"用户"`
	sys_model.UserInnerRegister
}

type QueryUserListReq struct {
	g.Meta `path:"/queryUserList" method:"post" summary:"获取用户|列表" tags:"用户"`
	sys_model.SearchParams
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

type UserInfoRes sys_entity.SysUser
type UserInfoListRes sys_model.CollectRes[sys_entity.SysUser]

type ResetUserPasswordReq struct {
	g.Meta          `path:"/resetUserPasswordReq" method:"post" summary:"重置用户密码" tags:"用户"`
	Password        string `json:"password" v:"required#请输入密码" dc:"登录密码"`
	ConfirmPassword string `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"确认密码"`
	UserId          int64  `json:"userId" v:"required#请输入用户id" dc:"用户ID"`
}
