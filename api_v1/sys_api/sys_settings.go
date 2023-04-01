package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type QuerySettingListReq struct {
	g.Meta `path:"/querySettingList" method:"POST" tags:"系统配置" summary:"获取配置信息|列表" dc:"系统配置：会限制主体获取"`
	base_model.SearchParams
}

type GetSettingByNameReq struct {
	g.Meta `path:"/getSettingByName" method:"POST" tags:"系统配置" summary:"获取配置|信息" dc:"系统配置：会限制主体获取"`

	Name string `json:"name" dc:"系统配置的名称"`

	base_model.SearchParams
}

type UpdateSettingReq struct {
	g.Meta `path:"/updateSetting" method:"POST" tags:"系统配置" summary:"修改配置|信息" dc:"修改系统配置"`

	sys_model.SysSettings
}

type DeleteSettingReq struct {
	g.Meta `path:"/deleteSetting" method:"POST" tags:"系统配置" summary:"删除配置|信息" dc:"删除系统配置，会根据unionMainId进行判断"`

	Name string `json:"name" dc:"系统配置的名称"`
}
