package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type QueryRoleListReq struct {
	g.Meta `path:"/queryRoleList" method:"post" summary:"获取所有角色|列表" tags:"角色"`
	sys_model.SearchParams
}

type CreateRoleInfoReq struct {
	g.Meta `path:"/createRole" method:"post" summary:"新增角色|信息" tags:"角色"`
	sys_model.SysRole
}

type UpdateRoleInfoReq struct {
	g.Meta `path:"/updateRole" method:"post" summary:"更新角色|信息" tags:"角色"`
	sys_model.SysRole
}

type DeleteRoleInfoReq struct {
	g.Meta `path:"/deleteRole" method:"post" summary:"删除角色" tags:"角色"`
	Id     int64 `json:"id" v:"required#角色ID校验失败" dc:"角色ID"`
}

type SetRoleForUserReq struct {
	g.Meta `path:"/setRoleForUser" method:"post" summary:"设置角色用户" tags:"角色"`
	RoleId int64 `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
	UserId int64 `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
}

type RemoveRoleForUserReq struct {
	g.Meta `path:"/removeRoleForUser" method:"post" summary:"移除用户所拥有的角色" tags:"角色"`
	RoleId int64 `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
	UserId int64 `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
}

type GetRoleUsersReq struct {
	g.Meta `path:"/getRoleUserList" method:"post" summary:"获取角色下的所有用户|列表" tags:"角色"`
	RoleId int64 `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
}

type GetUserRolesReq struct {
	g.Meta `path:"/getUserRoleList" method:"post" summary:"获取用户拥有的所有角色|列表" tags:"角色"`
	UserId int64 `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
}

type SetRolePermissionsReq struct {
	g.Meta        `path:"/setRolePermissions" method:"post" summary:"设置角色权限" tags:"角色"`
	Id            int64   `json:"id" v:"required#角色ID校验失败" dc:"角色ID"`
	PermissionIds []int64 `json:"permissionIds" v:"array#权限ID集合参数无效" dc:"权限ID集合"`
}

type GetRolePermissionsReq struct {
	g.Meta `path:"/getRolePermissionIds" method:"post" summary:"获取角色权限Ids" tags:"角色"`
	Id     int64 `json:"id" v:"required#角色ID校验失败" dc:"角色ID"`
}

type RoleInfoRes sys_entity.SysRole
type UserListRes sys_model.CollectRes[sys_model.SysUser]
