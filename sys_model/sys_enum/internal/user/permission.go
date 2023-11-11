package sys_enum_user

import (
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail     PermissionTypeEnum
	ViewMoreDetail PermissionTypeEnum
	List           PermissionTypeEnum
	SetState       PermissionTypeEnum
	ResetPassword  PermissionTypeEnum
	ChangePassword PermissionTypeEnum
	Create         PermissionTypeEnum
	SetUsername    PermissionTypeEnum
	SetUserRole    PermissionTypeEnum
	SetPermission  PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail:     base_permission.New(5947175853095365, "ViewDetail", "查看详情", "查看某个用户登录信息"),
	ViewMoreDetail: base_permission.New(5947175853095366, "ViewMoreDetail", "查看更多详情", "含完整手机号"),
	List:           base_permission.New(5947176286288325, "List", "用户列表", "查看所有用户"),
	SetState:       base_permission.New(5947176737372613, "SetState", "设置状态", "设置某个用户的状态"),
	ResetPassword:  base_permission.New(5947177123969477, "ResetPassword", "重置密码", "重置某个用户的登录密码"),
	ChangePassword: base_permission.New(5947177469213125, "ChangePassword", "修改密码", "修改自己的登录密码"),
	Create:         base_permission.New(5949854362632261, "Create", "创建用户", "创建一个新用户"),
	SetUsername:    base_permission.New(5949854362632262, "Update", "修改用户名称", "修改用户登录账户名称信息"),
	SetUserRole:    base_permission.New(5949854362632264, "SetUserRole", "设置用户角色", "设置某一个用户的角色"),
	SetPermission:  base_permission.New(5949854362632265, "SetPermission", "设置用户权限", "设置某一个用户的权限"),
}
