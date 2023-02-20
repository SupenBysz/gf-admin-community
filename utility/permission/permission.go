package permission

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// New 构造权限信息
func New(id int64, identifier string, name string, description ...string) *sys_model.SysPermissionTree {
	var desc string

	if len(description) > 0 {
		desc = description[0]
	}

	return &sys_model.SysPermissionTree{
		SysPermission: &sys_entity.SysPermission{
			Id:          id,
			Name:        name,
			Description: desc,
			Identifier:  identifier,
			Type:        1,
		},
	}
}

// NewInIdentifier 构造权限信息
func NewInIdentifier(identifier string, name string, description ...string) *sys_model.SysPermissionTree {
	var desc string

	if len(description) > 0 {
		desc = description[0]
	}

	return &sys_model.SysPermissionTree{
		SysPermission: &sys_entity.SysPermission{
			Id:          idgen.NextId(),
			Name:        name,
			Description: desc,
			Identifier:  identifier,
			Type:        1,
			MatchMode:   1,
		},
	}
}
