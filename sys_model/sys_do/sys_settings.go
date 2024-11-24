// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSettings is the golang structure of table sys_settings for DAO operations like Where/Data.
type SysSettings struct {
	g.Meta      `orm:"table:sys_settings, do:true"`
	Name        interface{} // 配置名称
	Values      interface{} // 配置信息JSON格式
	Desc        interface{} // 描述
	UnionMainId interface{} // 关联的主体id，为0代表是平台配置
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
