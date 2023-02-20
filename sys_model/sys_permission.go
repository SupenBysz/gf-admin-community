package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysPermission struct {
	Id          int64  `json:"id"             dc:"ID" v:"integer"`
	ParentId    int64  `json:"parentId"       dc:"父级ID" v:"min:0#必须是正整数，该属性创建后不支持修改"`
	Name        string `json:"name"           dc:"名称" v:"max-length:64#仅支持最大字符长度64"`
	Description string `json:"description"    dc:"描述" v:"max-length:128#仅支持最大字符长度128"`
	Identifier  string `json:"identifier"     dc:"标识符"`
	Type        int    `json:"type"           dc:"类型：1api、2menu"`
	MatchMode   int    `json:"matchMode"      dc:"匹配模式：ID：0，标识符：1"`
	IsShow      int    `json:"isShow"         dc:"是否显示：0不显示 1显示"`
	Sort        int    `json:"sort"           dc:"排序"`
}

type SysPermissionTree struct {
	*sys_entity.SysPermission
	Children []*SysPermissionTree `json:"children"       dc:"下级权限"`
}

type SysPermissionInfoRes sys_entity.SysPermission
type SysPermissionInfoListRes base_model.CollectRes[*sys_entity.SysPermission]
type SysPermissionInfoTreeRes []*SysPermissionTree
