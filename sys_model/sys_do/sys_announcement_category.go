// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncementCategory is the golang structure of table sys_announcement_category for DAO operations like Where/Data.
type SysAnnouncementCategory struct {
	g.Meta      `orm:"table:sys_announcement_category, do:true"`
	Id          interface{} //
	Name        interface{} // 分类名称
	Code        interface{} // 分类编码
	UnionMainId interface{} // 所属主体ID
	Sort        interface{} // 排序
	Description interface{} // 描述
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	DeletedAt   *gtime.Time //
	DeletedBy   interface{} //
}
