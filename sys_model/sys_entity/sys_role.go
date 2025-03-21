// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id          int64       `json:"id"          orm:"id"            description:""`
	Name        string      `json:"name"        orm:"name"          description:"名称"`
	Description string      `json:"description" orm:"description"   description:"描述"`
	IsSystem    bool        `json:"isSystem"    orm:"is_system"     description:"是否默认角色，true仅能修改名称，不允许删除和修改"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`
	UnionMainId int64       `json:"unionMainId" orm:"union_main_id" description:"主体id"`
}
