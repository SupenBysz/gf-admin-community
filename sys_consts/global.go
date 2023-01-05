package sys_consts

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/garray"
)

type global struct {
	DefaultRegisterType      int
	NotAllowLoginUserTypeArr *garray.SortedIntArray
	LogLevelToDatabaseArr    *garray.SortedIntArray
	ApiPreFix                string
}

var (
	Global = global{
		DefaultRegisterType:      0,
		NotAllowLoginUserTypeArr: garray.NewSortedIntArray(),
		LogLevelToDatabaseArr:    garray.NewSortedIntArray(),
		ApiPreFix:                "",
	}
	// PermissionTree 权限信息定义
	PermissionTree = []*permission.SysPermissionTree{
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947106208184773,
				Name:       "用户管理",
				Identifier: "User",
				Type:       1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看用户，查看某个用户登录账户
				sys_enum.User.PermissionType.View,
				// 用户列表，查看所有用户
				sys_enum.User.PermissionType.List,
				// 重置密码，重置某个用户的登录密码
				sys_enum.User.PermissionType.ResetPassword,
				// 设置状态，设置某个用户的状态
				sys_enum.User.PermissionType.SetState,
				// 修改密码，修改自己的登录密码
				sys_enum.User.PermissionType.ChangePassword,
			},
		},
	}
)
