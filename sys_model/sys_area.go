package sys_model

type Area struct {
	Id            int64  `json:"id"            description:"ID"`
	AreaCode      int    `json:"areaCode"      description:"地区编码"`
	AreaName      string `json:"areaName"      description:"地区名称"`
	Level         int    `json:"level"         description:"1:省份province,2:市city,3:区县district,4:街道street"`
	LatLongCenter string `json:"latLongCenter" description:"城市中心点（即经纬度）"`
	ParentId      int64  `json:"parentId"      description:"地区父节点"`
}

type AreaListRes CollectRes[Area]
