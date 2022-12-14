// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysArea is the golang structure of table sys_area for DAO operations like Where/Data.
type SysArea struct {
	g.Meta        `orm:"table:sys_area, do:true"`
	Id            interface{} // ID
	AreaCode      interface{} // 地区编码
	AreaName      interface{} // 地区名称
	Level         interface{} // 1:省份province,2:市city,3:区县district,4:街道street
	CityCode      interface{} // 城市编码
	LatLongCenter interface{} // 城市中心点（即经纬度）
	ParentId      interface{} // 地区父节点
}
