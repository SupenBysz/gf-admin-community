package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

/*
	前端配置信息
*/

type QueryFrontSettingListReq struct {
	g.Meta `path:"/queryFrontSettingList" method:"POST" tags:"系统前端配置" summary:"获取配置信息|列表" dc:"系统配置：会限制主体获取"`
	base_model.SearchParams
}

type GetFrontSettingByNameReq struct {
	g.Meta `path:"/getFrontSettingByName" method:"POST" tags:"系统前端配置" summary:"获取配置|信息" dc:"系统配置：会限制主体获取"`

	Name string `json:"name" dc:"系统配置的名称"`

	base_model.SearchParams
}

type SaveFrontSettingReq struct {
	g.Meta `path:"/saveFrontSetting" method:"POST" tags:"系统前端配置" summary:"保存配置|信息" dc:"保存系统配置：强制当前登陆用户作为配置的拥有者"`

	sys_model.SysFrontSettings
}

type DeleteFrontSettingReq struct {
	g.Meta `path:"/deleteFrontSetting" method:"POST" tags:"系统前端配置" summary:"删除配置|信息" dc:"删除系统配置，会根据unionMainId进行判断"`

	Name string `json:"name" dc:"系统配置的名称"`
}
