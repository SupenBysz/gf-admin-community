package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type SysMemberLevel struct {
	sys_entity.SysMemberLevel
	Id          int64       `json:"-"   orm:"-"`
	CreatedAt   *gtime.Time `json:"-"   orm:"-"`
	UpdatedAt   *gtime.Time `json:"-"   orm:"-"`
	CreatedBy   int64       `json:"-"   orm:"-"`
	UnionMainId int64       `json:"-"   orm:"-"`
}

type UpdateSysMemberLevel struct {
	sys_entity.SysMemberLevel
	Id          int64       `json:"id"  orm:"id" dc:"ID"  v:"required#会员等级名称不能为空"`
	CreatedAt   *gtime.Time `json:"-"   orm:"-"`
	UpdatedAt   *gtime.Time `json:"-"   orm:"-"`
	CreatedBy   int64       `json:"-"   orm:"-"`
	UnionMainId int64       `json:"-"   orm:"-"`
}

type SysMemberLevelRes sys_entity.SysMemberLevel
type SysMemberLevelListRes base_model.CollectRes[SysMemberLevelRes]

type SysMemberLevelUserRes sys_entity.SysMemberLevelUser
type SysMemberLevelUserListRes []SysMemberLevelUserRes
