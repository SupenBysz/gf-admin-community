// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMemberLevel is the golang structure for table sys_member_level.
type SysMemberLevel struct {
	Id          int64       `json:"id"          orm:"id"            description:"ID"`
	Name        string      `json:"name"        orm:"name"          description:"名称"`
	Desc        string      `json:"desc"        orm:"desc"          description:"描述"`
	Identifier  string      `json:"identifier"  orm:"identifier"    description:"级别标识符"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:""`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"    description:""`
	UnionMainId int64       `json:"unionMainId" orm:"union_main_id" description:""`
	Level       int         `json:"level"       orm:"level"         description:"权益级别，要遵循数值越大，权益越多"`
}
