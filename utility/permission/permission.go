package permission

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/kysion/base-library/utility/base_permission"
)

func initFactory() {
	base_permission.InitializePermissionFactory(func() base_permission.IPermission {
		return &sys_model.SysPermissionTree{
			SysPermission: &sys_entity.SysPermission{},
		}
	})
}

func New(id int64, identifier string, name string, description ...string) base_permission.IPermission {
	initFactory()
	return base_permission.New(id, identifier, name, description...)
}

// NewInIdentifier 构造权限信息
func NewInIdentifier(identifier string, name string, description ...string) base_permission.IPermission {
	initFactory()
	return base_permission.New(idgen.NextId(), identifier, name, description...)
}
