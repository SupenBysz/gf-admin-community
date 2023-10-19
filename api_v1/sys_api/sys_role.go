package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type QueryRoleListReq struct {
	g.Meta `path:"/queryRoleList" method:"post" summary:"获取所有角色|列表" tags:"角色"`
	base_model.SearchParams
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

type SetRoleMemberReq struct {
	g.Meta  `path:"/setRoleMember" method:"post" summary:"设置角色成员" tags:"角色"`
	RoleId  int64   `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
	UserIds []int64 `json:"userIds" v:"required#用户ID校验失败" dc:"用户IDS"`
}

type RemoveRoleMemberReq struct {
	g.Meta `path:"/removeRoleMember" method:"post" summary:"移除角色成员" tags:"角色"`
	RoleId int64 `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
	UserId int64 `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
}

type GetRoleMemberReq struct {
	g.Meta `path:"/getRoleMemberList" method:"post" summary:"获取角色成员|列表" tags:"角色"`
	RoleId int64 `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
}

type GetRoleMemberIdsReq struct {
	g.Meta `path:"/getRoleMemberIdsList" method:"post" summary:"获取角色成员Ids|列表" tags:"角色"`
	RoleId int64 `json:"roleId" v:"required#角色ID校验失败" dc:"角色ID"`
}

type GetRoleByUserIdListReq struct {
	g.Meta `path:"/getRoleByUserIdList" method:"post" summary:"根据用户ID获取所有关联角色|列表" tags:"角色"`
	UserId int64 `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
}

type SetRolePermissionsReq struct {
	g.Meta        `path:"/setRolePermissions" method:"post" summary:"设置角色权限" tags:"角色"`
	Id            int64   `json:"id" v:"required#角色ID校验失败" dc:"角色ID"`
	PermissionIds []int64 `json:"permissionIds" v:"array#权限ID集合参数无效" dc:"权限ID集合"`
}

type GetRolePermissionsIdsReq struct {
	g.Meta `path:"/getRolePermissionIds" method:"post" summary:"获取角色权限Ids" tags:"角色"`
	Id     string `json:"id" v:"required#角色ID校验失败" dc:"角色ID"`
}

type RoleInfoRes sys_model.RoleInfo
type UserListRes base_model.CollectRes[*sys_model.SysUser]
