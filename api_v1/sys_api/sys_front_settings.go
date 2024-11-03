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

type GetFrontSettingReq struct {
	g.Meta `path:"/getFrontSetting" method:"POST" tags:"系统前端配置" summary:"获取配置|信息" dc:"系统配置：会限制主体获取"`

	Name        string `json:"name" dc:"系统配置的名称"`
	UnionMainId int64  `json:"unionMainId" dc:"关联的主体id，为0代表是平台配置" `
	UserId      int64  `json:"userId" dc:"关联的用户id，为0代表平台配置" `
	Sys         *int   `json:"sys" dc:"1除主体管理员外，主体下的其他用户仅有只读权限"`
}

type SaveFrontSettingReq struct {
	g.Meta `path:"/saveFrontSetting" method:"POST" tags:"系统前端配置" summary:"保存配置|信息" dc:"保存系统配置：强制当前登陆用户作为配置的拥有者"`

	sys_model.SysFrontSettings
}

type DeleteFrontSettingReq struct {
	g.Meta `path:"/deleteFrontSetting" method:"POST" tags:"系统前端配置" summary:"删除配置|信息" dc:"删除系统配置，会根据unionMainId进行判断"`

	UserId      int64  `json:"user_id" dc:"用户ID"`
	UnionMainId int64  `json:"union_main_id" dc:"主体ID"`
	Name        string `json:"name" dc:"系统配置的名称"`
}
