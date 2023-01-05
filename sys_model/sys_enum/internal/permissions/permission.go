package permissions

import "github.com/SupenBysz/gf-admin-community/utility/permission"

type permissionType struct {
	View   *permission.SysPermissionTree
	List   *permission.SysPermissionTree
	Update *permission.SysPermissionTree
	Delete *permission.SysPermissionTree
	Create *permission.SysPermissionTree
}

var PermissionType = permissionType{
	View:   permission.New(5948682180886598, "View", "查看权限", "查看某个权限"),
	List:   permission.New(5948682180886599, "List", "权限列表", "查看所有权限"),
	Update: permission.New(5948682180886600, "Update", "更新权限", "更新某个权限"),
	Delete: permission.New(5948682180886601, "Delete", "删除权限", "删除某个权限"),
	Create: permission.New(5948682180886602, "Create", "创建权限", "创建权限"),
}
