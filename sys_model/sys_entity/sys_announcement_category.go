// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncementCategory is the golang structure for table sys_announcement_category.
type SysAnnouncementCategory struct {
	Id          int64       `json:"id"          orm:"id"            description:""`
	Name        string      `json:"name"        orm:"name"          description:"分类名称"`
	Code        string      `json:"code"        orm:"code"          description:"分类编码"`
	UnionMainId int64       `json:"unionMainId" orm:"union_main_id" description:"所属主体ID"`
	Sort        int         `json:"sort"        orm:"sort"          description:"排序"`
	Description string      `json:"description" orm:"description"   description:"描述"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:""`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"    description:""`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"    description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"    description:""`
	DeletedBy   int64       `json:"deletedBy"   orm:"deleted_by"    description:""`
}
