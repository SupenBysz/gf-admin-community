// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id          int64       `json:"id"          description:""`
	Name        string      `json:"name"        description:"名称"`
	Description string      `json:"description" description:"描述"`
	IsSystem    bool        `json:"isSystem"    description:"是否默认角色，true仅能修改名称"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
}