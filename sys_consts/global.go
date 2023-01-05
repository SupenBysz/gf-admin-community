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
		// 用户管理权限树
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
		// 组织架构权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948649344204869,
				Name:       "组织架构",
				Identifier: "Organization",
				Type:       1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看
				sys_enum.Organization.PermissionType.View,
				// 查看列表
				sys_enum.Organization.PermissionType.List,
				// 更新
				sys_enum.Organization.PermissionType.Update,
				// 删除
				sys_enum.Organization.PermissionType.Delete,
				// 创建
				sys_enum.Organization.PermissionType.Create,
			},
		},
		// 角色管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948684761759818,
				Name:       "角色管理",
				Identifier: "Role",
				Type:       1,
			},
			Children: []*permission.SysPermissionTree{
				sys_enum.Role.PermissionType.View,
				sys_enum.Role.PermissionType.List,
				sys_enum.Role.PermissionType.Update,
				sys_enum.Role.PermissionType.Delete,
				sys_enum.Role.PermissionType.Create,
			},
		},
		// 权限管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948684761759818,
				Name:       "权限管理",
				Identifier: "Permission",
				Type:       1,
			},
			Children: []*permission.SysPermissionTree{
				sys_enum.Permissions.PermissionType.View,
				sys_enum.Permissions.PermissionType.List,
				sys_enum.Permissions.PermissionType.Update,
				sys_enum.Permissions.PermissionType.Delete,
				sys_enum.Permissions.PermissionType.Create,
			},
		},
	}
)
