package sys_enum_permission

import (
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
	List       PermissionTypeEnum
	Update     PermissionTypeEnum
	Delete     PermissionTypeEnum
	Create     PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: base_permission.New(5948682180886598, "ViewDetail", "查看权限", "查看某个权限"),
	List:       base_permission.New(5948682180886599, "List", "权限列表", "查看所有权限"),
	Update:     base_permission.New(5948682180886600, "Update", "更新权限", "更新某个权限"),
	Delete:     base_permission.New(5948682180886601, "Delete", "删除权限", "删除某个权限"),
	Create:     base_permission.New(5948682180886602, "Create", "创建权限", "创建权限"),
}
