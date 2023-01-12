package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type SysPermission struct {
	Id          int64  `json:"id"             dc:"ID" v:"integer"`
	ParentId    int64  `json:"parentId"       dc:"父级ID" v:"min:0#必须是正整数，该属性创建后不支持修改"`
	Name        string `json:"name"           dc:"名称" v:"max-length:64#仅支持最大字符长度64"`
	Description string `json:"description"    dc:"描述" v:"max-length:128#仅支持最大字符长度128"`
	Identifier  string `json:"identifier"  description:"标识符"`
	Type        int    `json:"type"        description:"类型：1api、2menu"`
}

type SysPermissionInfoRes sys_entity.SysPermission
type SysPermissionInfoListRes CollectRes[*sys_entity.SysPermission]
type SysPermissionInfoTreeRes []*permission.SysPermissionTree
