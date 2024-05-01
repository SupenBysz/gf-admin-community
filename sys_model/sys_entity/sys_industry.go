// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysIndustry is the golang structure for table sys_industry.
type SysIndustry struct {
	Id           int64       `json:"id"           description:"ID"`
	CategoryId   int64       `json:"categoryId"   description:"行业ID"`
	CategoryName string      `json:"categoryName" description:"行业名称"`
	CategoryDesc string      `json:"categoryDesc" description:"行业描述"`
	Rate         int         `json:"rate"         description:"费率"`
	ParentId     int64       `json:"parentId"     description:"父级ID"`
	Sort         int         `json:"sort"         description:"排序"`
	State        int         `json:"state"        description:"状态：0隐藏，1显示"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    description:""`
}