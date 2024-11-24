// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrganization is the golang structure of table sys_organization for DAO operations like Where/Data.
type SysOrganization struct {
	g.Meta      `orm:"table:sys_organization, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	ParentId    interface{} // 父级ID
	CascadeDeep interface{} // 级联深度
	Description interface{} // 描述
}
