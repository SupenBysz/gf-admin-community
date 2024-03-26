package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetPermissionByIdReq struct {
	g.Meta `path:"/getPermissionById" method:"post" summary:"根据权限ID获取权限|信息" tags:"权限"`
	Id     int64 `json:"id" v:"required#权限ID校验失败" dc:"权限ID"`
}

type GetPermissionByIdentifierReq struct {
	g.Meta     `path:"/getPermissionByIdentifier" method:"post" summary:"根据权限标识符获取权限|信息" tags:"权限"`
	Identifier string `json:"identifier" v:"required|max-length:64#权限Name校验失败|仅支持最大字符长度64" dc:"权限标识符"`
}

type QueryPermissionListReq struct {
	g.Meta `path:"/queryPermissionList" method:"post" summary:"根据ID获取下级权限|列表，返回列表" tags:"权限"`
	base_model.SearchParams
}

type GetPermissionListReq struct {
	g.Meta      `path:"/getPermissionList" method:"post" summary:"根据ID获取下级权限|列表，返回列表" tags:"权限"`
	Id          int64 `json:"id" v:"min:0#权限ID校验失败" dc:"权限ID" default:"0"`
	IsRecursive bool  `json:"isRecursive" dc:"是否递归，true则显示所有子级"`
}

type GetPermissionTreeReq struct {
	g.Meta `path:"/getPermissionTree" method:"post" summary:"根据ID获取下级权限|树" tags:"权限"`
	Id     int64 `json:"id" v:"required#权限ID校验失败" dc:"权限ID"`
	Type   int   `json:"type" dc:"权限类型：1Api 2Menu"`
}

type CreatePermissionReq struct {
	g.Meta `path:"/createPermission" method:"post" summary:"新增权限|信息" tags:"权限"`
	sys_model.SysPermission
}

type UpdatePermissionReq struct {
	g.Meta `path:"/updatePermission" method:"post" summary:"保存权限|信息" tags:"权限"`
	sys_model.UpdateSysPermission
}

type DeletePermissionReq struct {
	g.Meta `path:"/deletePermission" method:"post" summary:"删除权限" tags:"权限"`
	Id     int64 `json:"id" v:"required#权限ID校验失败" dc:"权限ID"`
}
