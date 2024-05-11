package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysConfig struct {
	Name  string `json:"name"      description:"配置名称"`
	Value string `json:"value"     description:"配置信息"`
}

type SysConfigRes sys_entity.SysConfig

type SysConfigListRes base_model.CollectRes[sys_entity.SysConfig]
