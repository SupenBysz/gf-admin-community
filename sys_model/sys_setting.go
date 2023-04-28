package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysSettings struct {
	Name        string `json:"name"        dc:"配置名称"`
	Values      string `json:"values"      dc:"配置信息JSON格式"`
	Desc        string `json:"desc"        dc:"描述"`
	UnionMainId int64  `json:"unionMainId" dc:"关联的主体id，为0代表是平台配置"`
}

type SysSettingsRes sys_entity.SysSettings

type SysSettingListRes base_model.CollectRes[sys_entity.SysSettings]
