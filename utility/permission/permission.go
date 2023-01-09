package permission

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type SysPermissionTree struct {
	*sys_entity.SysPermission
	Children []*SysPermissionTree `json:"children"       dc:"下级权限"`
}

// New 构造权限信息
func New(id int64, identifier string, name string, description ...string) *SysPermissionTree {
	var desc string

	if len(description) > 0 {
		desc = description[0]
	}

	return &SysPermissionTree{
		SysPermission: &sys_entity.SysPermission{
			Id:          id,
			Name:        name,
			Description: desc,
			Identifier:  identifier,
			Type:        1,
		},
	}
}
