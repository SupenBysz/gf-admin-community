// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
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
	Level         interface{} // 1街道street、2区县district、4市city、8省份province、16大区region、32全国nation
	CityCode      interface{} // 城市编码
	LongLatCenter interface{} // 城市中心点（即经纬度）
	ParentId      interface{} // 地区父节点
	PinYin        interface{} // 地区拼音
}
