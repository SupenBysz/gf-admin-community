package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

// SysFrontSettings 前端配置信息
type SysFrontSettings struct {
	Name        string `json:"name"                  dc:"配置名称" v:"required#配置名称不能为空"`
	Values      string `json:"values"              dc:"配置信息JSON格式" `
	Desc        string `json:"desc"                  dc:"描述" `
	UnionMainId int64  `json:"unionMainId"  dc:"关联的主体id，为0代表是平台配置" `
	UserId      int64  `json:"userId"             dc:"关联的用户id，为0代表平台配置" `
}

type SysFrontSettingsRes sys_entity.SysFrontSettings

type SysFrontSettingsListRes base_model.CollectRes[SysFrontSettingsRes]
