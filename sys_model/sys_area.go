package sys_model

import "github.com/kysion/base-library/base_model"

type Area struct {
	Id            int64  `json:"id"            description:"ID"`
	AreaCode      int    `json:"areaCode"      description:"地区编码"`
	AreaName      string `json:"areaName"      description:"地区名称"`
	Level         int    `json:"level"         description:"1区县district、2市city、4省份province、8大区region、16全国nation"`
	LatLongCenter string `json:"latLongCenter" description:"城市中心点（即经纬度）"`
	ParentId      int64  `json:"parentId"      description:"地区父节点"`
}

type AreaListRes base_model.CollectRes[*Area]
