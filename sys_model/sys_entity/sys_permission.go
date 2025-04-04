// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure for table sys_permission.
type SysPermission struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	ParentId    int64       `json:"parentId"    orm:"parent_id"   description:"父级ID"`
	Name        string      `json:"name"        orm:"name"        description:"名称"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Identifier  string      `json:"identifier"  orm:"identifier"  description:"标识符"`
	Type        int         `json:"type"        orm:"type"        description:"类型：1api，2menu"`
	MatchMode   int         `json:"matchMode"   orm:"match_mode"  description:"匹配模式：ID：0，标识符：1"`
	IsShow      int         `json:"isShow"      orm:"is_show"     description:"是否显示：0不显示 1显示"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}
