// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

// SysArea is the golang structure for table sys_area.
type SysArea struct {
	Id            int64  `json:"id"            description:"ID"`
	AreaCode      int    `json:"areaCode"      description:"地区编码"`
	AreaName      string `json:"areaName"      description:"地区名称"`
	Level         int    `json:"level"         description:"1街道street、2区县district、4市city、8省份province、16大区region、32全国nation"`
	CityCode      string `json:"cityCode"      description:"城市编码"`
	LongLatCenter string `json:"longLatCenter" description:"城市中心点（即经纬度）"`
	ParentId      int64  `json:"parentId"      description:"地区父节点"`
	PinYin        string `json:"pinYin"        description:"地区拼音"`
}
