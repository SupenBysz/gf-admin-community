// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMemberLevel is the golang structure of table sys_member_level for DAO operations like Where/Data.
type SysMemberLevel struct {
	g.Meta      `orm:"table:sys_member_level, do:true"`
	Id          interface{} // ID
	Name        interface{} // 名称
	Desc        interface{} // 描述
	Identifier  interface{} // 级别标识符
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	CreatedBy   interface{} //
	UnionMainId interface{} //
}
