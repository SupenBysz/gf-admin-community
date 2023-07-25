// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/casbin/casbin/v2"
)

type (
	ICasbin interface {
		// InstallHook 安装Hook
		InstallHook(userType sys_enum.UserType, hookFunc sys_hook.CasbinHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除所有Hook
		CleanAllHook()
		Check() error
		Enforcer() *casbin.Enforcer
		// AddRoleForUserInDomain 添加用户角色关联关系
		AddRoleForUserInDomain(userName string, roleName string, domain string) (bool, error)
		// DeleteRoleForUserInDomain 删除用户角色关联关系
		DeleteRoleForUserInDomain(userName, roleName string, domain string) (bool, error)
		// DeleteRolesForUser 清空用户角色关联关系
		DeleteRolesForUser(userName string, domain string) (bool, error)
		// AddPermissionForUser 添加角色与资源关系
		AddPermissionForUser(roleName, path, method string) (bool, error)
		// AddPermissionsForUser 添加角色与资源关系
		AddPermissionsForUser(roleName string, path []string) (bool, error)
		// DeletePermissionForUser 删除角色与资源关系
		DeletePermissionForUser(roleName, path, method string) (bool, error)
		// DeletePermissionsForUser 清空角色与资源关系
		DeletePermissionsForUser(roleName string) (bool, error)
		// EnforceCheck 校验  确认访问权限
		EnforceCheck(userName, path, role, method interface{}) (bool, error)
	}
)

var (
	localCasbin ICasbin
)

func Casbin() ICasbin {
	if localCasbin == nil {
		panic("implement not found for interface ICasbin, forgot register?")
	}
	return localCasbin
}

func RegisterCasbin(i ICasbin) {
	localCasbin = i
}
