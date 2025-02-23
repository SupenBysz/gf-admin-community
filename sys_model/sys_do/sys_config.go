// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta    `orm:"table:sys_config, do:true"`
	Name      interface{} // 配置名称
	Value     interface{} // 配置信息
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
