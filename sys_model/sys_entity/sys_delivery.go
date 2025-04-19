// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDelivery is the golang structure for table sys_delivery.
type SysDelivery struct {
	Id                       int64       `json:"id"                       orm:"id"                          description:""`
	Name                     string      `json:"name"                     orm:"name"                        description:"物流公司"`
	Logo                     string      `json:"logo"                     orm:"logo"                        description:"LOGO"`
	Site                     string      `json:"site"                     orm:"site"                        description:"网址"`
	ExpressNo                string      `json:"expressNo"                orm:"express_no"                  description:"物流跟踪编号"`
	ExpressNoElectronicSheet string      `json:"expressNoElectronicSheet" orm:"express_no_electronic_sheet" description:"电子面单编号"`
	PrintStyleJson           string      `json:"printStyleJson"           orm:"print_style_json"            description:"打印模板样式"`
	ExpTypeJson              string      `json:"expTypeJson"              orm:"exp_type_json"               description:"业务类型"`
	UpdatedAt                *gtime.Time `json:"updatedAt"                orm:"updated_at"                  description:""`
	CreatedAt                *gtime.Time `json:"createdAt"                orm:"created_at"                  description:""`
}
