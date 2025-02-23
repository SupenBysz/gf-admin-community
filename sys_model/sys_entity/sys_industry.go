// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysIndustry is the golang structure for table sys_industry.
type SysIndustry struct {
	Id           int64       `json:"id"           orm:"id"            description:"ID"`
	CategoryId   int64       `json:"categoryId"   orm:"category_id"   description:"行业ID"`
	CategoryName string      `json:"categoryName" orm:"category_name" description:"行业名称"`
	CategoryDesc string      `json:"categoryDesc" orm:"category_desc" description:"行业描述"`
	Rate         int         `json:"rate"         orm:"rate"          description:"费率"`
	ParentId     int64       `json:"parentId"     orm:"parent_id"     description:"父级ID"`
	Sort         int         `json:"sort"         orm:"sort"          description:"排序"`
	State        int         `json:"state"        orm:"state"         description:"状态：0隐藏，1显示"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
