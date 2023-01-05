package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type permissionType struct {
	View           *permission.SysPermissionTree
	List           *permission.SysPermissionTree
	SetState       *permission.SysPermissionTree
	ResetPassword  *permission.SysPermissionTree
	ChangePassword *permission.SysPermissionTree
}

var PermissionType = permissionType{
	View:           permission.New(5947175853095365, "View", "查看用户", "查看某个用户登录账户"),
	List:           permission.New(5947176286288325, "List", "用户列表", "查看所有用户"),
	SetState:       permission.New(5947176737372613, "SetState", "设置状态", "设置某个用户的状态"),
	ResetPassword:  permission.New(5947177123969477, "ResetPassword", "重置密码", "重置某个用户的登录密码"),
	ChangePassword: permission.New(5947177469213125, "ChangePassword", "修改密码", "修改自己的登录密码"),
}
