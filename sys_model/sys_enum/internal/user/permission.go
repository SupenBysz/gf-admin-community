package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type PermissionTypeEnum = *permission.SysPermissionTree

type permissionType struct {
	ViewDetail     PermissionTypeEnum
	List           PermissionTypeEnum
	SetState       PermissionTypeEnum
	ResetPassword  PermissionTypeEnum
	ChangePassword PermissionTypeEnum
	// Create         PermissionTypeEnum
	SetUsername   PermissionTypeEnum
	SetUserRole   PermissionTypeEnum
	SetPermission PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail:     permission.New(5947175853095365, "ViewDetail", "查看用户", "查看某个用户登录账户"),
	List:           permission.New(5947176286288325, "List", "用户列表", "查看所有用户"),
	SetState:       permission.New(5947176737372613, "SetState", "设置状态", "设置某个用户的状态"),
	ResetPassword:  permission.New(5947177123969477, "ResetPassword", "重置密码", "重置某个用户的登录密码"),
	ChangePassword: permission.New(5947177469213125, "ChangePassword", "修改密码", "修改自己的登录密码"),
	// Create:         permission.New(5949854362632261, "Create", "创建用户", "创建一个新用户"),
	SetUsername:   permission.New(5949854362632262, "Update", "修改用户名称", "修改用户登录账户名称信息"),
	SetUserRole:   permission.New(5949854362632264, "SetUserRole", "设置用户角色", "设置某一个用户的角色"),
	SetPermission: permission.New(5949854362632265, "SetPermission", "设置用户权限", "设置某一个用户的权限"),
}
