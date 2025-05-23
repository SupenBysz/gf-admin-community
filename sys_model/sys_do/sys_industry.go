// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysIndustry is the golang structure of table sys_industry for DAO operations like Where/Data.
type SysIndustry struct {
	g.Meta       `orm:"table:sys_industry, do:true"`
	Id           interface{} // ID
	CategoryId   interface{} // 行业ID
	CategoryName interface{} // 行业名称
	CategoryDesc interface{} // 行业描述
	Rate         interface{} // 费率
	ParentId     interface{} // 父级ID
	Sort         interface{} // 排序
	State        interface{} // 状态：0隐藏，1显示
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
