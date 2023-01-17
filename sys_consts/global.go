package sys_consts

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/garray"
)

type global struct {
	DefaultRegisterType      int
	NotAllowLoginUserTypeArr *garray.SortedIntArray
	LogLevelToDatabaseArr    *garray.SortedIntArray
	ApiPreFix                string
	OrmCacheConf             []*sys_model.TableCacheConf
	PermissionTree           []*permission.SysPermissionTree // PermissionTree 权限信息定义
}

var (
	Global = global{
		DefaultRegisterType:      0,
		NotAllowLoginUserTypeArr: garray.NewSortedIntArray(),
		LogLevelToDatabaseArr:    garray.NewSortedIntArray(),
		ApiPreFix:                "",
		OrmCacheConf:             []*sys_model.TableCacheConf{},
		PermissionTree:           []*permission.SysPermissionTree{},
	}
)
