// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta      `orm:"table:sys_role, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Description interface{} // 描述
	IsSystem    interface{} // 是否默认角色，true仅能修改名称，不允许删除和修改
	UpdatedAt   *gtime.Time //
	CreatedAt   *gtime.Time //
	UnionMainId interface{} // 主体id
}
