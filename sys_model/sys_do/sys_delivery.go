// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDelivery is the golang structure of table sys_delivery for DAO operations like Where/Data.
type SysDelivery struct {
	g.Meta                   `orm:"table:sys_delivery, do:true"`
	Id                       interface{} //
	Name                     interface{} // 物流公司
	Logo                     interface{} // LOGO
	Site                     interface{} // 网址
	ExpressNo                interface{} // 物流跟踪编号
	ExpressNoElectronicSheet interface{} // 电子面单编号
	PrintStyleJson           interface{} // 打印模板样式
	ExpTypeJson              interface{} // 业务类型
	UpdatedAt                *gtime.Time //
	CreatedAt                *gtime.Time //
}
