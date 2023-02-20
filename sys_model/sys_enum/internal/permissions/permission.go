package permissions

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type PermissionTypeEnum = *sys_model.SysPermissionTree

type permissionType struct {
	ViewDetail PermissionTypeEnum
	List       PermissionTypeEnum
	Update     PermissionTypeEnum
	Delete     PermissionTypeEnum
	Create     PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: permission.New(5948682180886598, "ViewDetail", "查看权限", "查看某个权限"),
	List:       permission.New(5948682180886599, "List", "权限列表", "查看所有权限"),
	Update:     permission.New(5948682180886600, "Update", "更新权限", "更新某个权限"),
	Delete:     permission.New(5948682180886601, "Delete", "删除权限", "删除某个权限"),
	Create:     permission.New(5948682180886602, "Create", "创建权限", "创建权限"),
}
