package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysMemberLevel struct {
	Name       string `json:"name"        orm:"name"          dc:"名称" v:"required#会员等级名称不能为空"`
	Desc       string `json:"desc"        orm:"desc"          dc:"描述"`
	Identifier string `json:"identifier"  orm:"identifier"    dc:"级别标识符" v:"required#会员等级标识符不能为空"`
}

type UpdateSysMemberLevel struct {
	Id         int64   `json:"id"               orm:"id"                  dc:"ID"  v:"required#会员等级名称不能为空"`
	Name       *string `json:"name"        orm:"name"          dc:"名称" `
	Desc       *string `json:"desc"        orm:"desc"          dc:"描述"`
	Identifier *string `json:"identifier"  orm:"identifier"    dc:"级别标识符"`
}

type SysMemberLevelRes sys_entity.SysMemberLevel
type SysMemberLevelListRes base_model.CollectRes[SysMemberLevelRes]

type SysMemberLevelUserRes sys_entity.SysMemberLevelUser
type SysMemberLevelUserListRes []SysMemberLevelUserRes
