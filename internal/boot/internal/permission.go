package internal

import (
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

// InitPermission 初始化权限树结构
func InitPermission() []*sys_model.SysPermissionTree {
	sys_consts.Global.PermissionTree = []*sys_model.SysPermissionTree{
		// 用户管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947106208184773,
				Name:       "用户管理",
				Identifier: "User",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看用户，查看某个用户登录账户
				sys_enum.User.PermissionType.ViewDetail,
				// 查看更多详情，含完整手机号
				sys_enum.User.PermissionType.ViewMoreDetail,
				// 用户列表，查看所有用户
				sys_enum.User.PermissionType.List,
				// 重置密码，重置某个用户的登录密码
				sys_enum.User.PermissionType.ResetPassword,
				// 设置状态，设置某个用户的状态
				sys_enum.User.PermissionType.SetState,
				// 修改密码，修改自己的登录密码
				sys_enum.User.PermissionType.ChangePassword,
				// "创建用户，创建一个新用户"
				// sys_enum.User.PermissionType.Create,
				// 修改用户名称，修改用户登录账户名称信息
				sys_enum.User.PermissionType.SetUsername,
				// 设置用户角色，设置某一个用户的角色
				sys_enum.User.PermissionType.SetUserRole,
				// 设置用户权限，设置某一个用户的权限
				sys_enum.User.PermissionType.SetPermission,
			},
		},
		// 组织架构权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948649344204869,
				Name:       "组织架构",
				Identifier: "Organization",
				Type:       1,
				IsShow:     0,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看，查看某个组织架构
				sys_enum.Organization.PermissionType.ViewDetail,
				// 查看列表，查看所有组织架构列表
				sys_enum.Organization.PermissionType.List,
				// 更新，更新某个组织架构
				sys_enum.Organization.PermissionType.Update,
				// 删除，删除某个组织架构
				sys_enum.Organization.PermissionType.Delete,
				// 创建，创建组织架构
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
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看角色，查看某个角色
				sys_enum.Role.PermissionType.ViewDetail,
				// 角色列表，查看所有角色
				sys_enum.Role.PermissionType.List,
				// 更新角色信息，更新某个角色信息
				sys_enum.Role.PermissionType.Update,
				// 删除角色，删除某个角色
				sys_enum.Role.PermissionType.Delete,
				// 创建角色，创建一个新角色
				sys_enum.Role.PermissionType.Create,
				// 设置角色成员，增加或移除角色成员
				sys_enum.Role.PermissionType.SetMember,
				// 设置角色权限，设置某个角色的权限
				sys_enum.Role.PermissionType.SetPermission,
			},
		},
		// 权限管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5950408166668741,
				Name:       "权限管理",
				Identifier: "Permission",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看权限，查看某个权限
				sys_enum.Permissions.PermissionType.ViewDetail,
				// 权限列表，查看所有权限
				sys_enum.Permissions.PermissionType.List,
				// 更新权限，更新某个权限
				sys_enum.Permissions.PermissionType.Update,
				// 删除权限，删除某个权限
				sys_enum.Permissions.PermissionType.Delete,
				// 创建权限，创建权限
				sys_enum.Permissions.PermissionType.Create,
			},
		},
		// 菜单管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5950408166676321,
				Name:       "菜单管理",
				Identifier: "Menu",
				Type:       1,
				IsShow:     0, // 默认隐藏
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看菜单，查看某个菜单
				sys_enum.Menu.PermissionType.ViewDetail,
				// 菜单树，查看菜单树
				sys_enum.Menu.PermissionType.Tree,
				// 更新菜单，更新某个菜单
				sys_enum.Menu.PermissionType.Update,
				// 删除菜单，删除某个菜单
				sys_enum.Menu.PermissionType.Delete,
				// 创建菜单，创建菜单
				sys_enum.Menu.PermissionType.Create,
			},
		},
		// sms

		// oss
	}

	// 添加个人资质和审核权限树
	licensePermission := initAuditAndLicensePermission()
	sys_consts.Global.PermissionTree = append(sys_consts.Global.PermissionTree, licensePermission...)

	return sys_consts.Global.PermissionTree
}

func initAuditAndLicensePermission() []*sys_model.SysPermissionTree {
	result := []*sys_model.SysPermissionTree{

		// 资质权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5953153121845333,
				Name:       "个人资质",
				Identifier: "PersonLicense",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看资质信息，查看某条资质信息
				sys_enum.License.PermissionType.ViewDetail,
				// 资质列表，查看所有资质信息
				sys_enum.License.PermissionType.List,
				// 更新资质审核信息，更新某条资质审核信息
				sys_enum.License.PermissionType.Update,
				// 创建资质，创建资质信息
				sys_enum.License.PermissionType.Create,
				// 设置资质状态，设置某资质认证状态
				sys_enum.License.PermissionType.SetState,
			},
		},
		// 审核管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5953151699124300,
				Name:       "个人资质审核管理",
				Identifier: "PersonAudit",
				Type:       1,
				IsShow:     1,
			},
			Children: []*sys_model.SysPermissionTree{
				// 查看审核信息，查看某条资质审核信息
				sys_enum.Audit.PermissionType.ViewDetail,
				// 资质审核列表，查看所有资质审核
				sys_enum.Audit.PermissionType.List,
				// 更新资质审核信息，更新某条资质审核信息
				sys_enum.Audit.PermissionType.Update,
			},
		},
	}
	return result
}
