package organization

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type permissionType struct {
	View   *permission.SysPermissionTree
	List   *permission.SysPermissionTree
	Update *permission.SysPermissionTree
	Delete *permission.SysPermissionTree
	Create *permission.SysPermissionTree
}

var PermissionType = permissionType{
	View:   permission.New(5948684761759813, "View", "查看角色", "查看某个角色"),
	List:   permission.New(5948684761759814, "List", "角色列表", "查看所有角色"),
	Update: permission.New(5948684761759815, "Update", "更新角色", "更新某个角色"),
	Delete: permission.New(5948684761759816, "Delete", "删除角色", "删除某个角色"),
	Create: permission.New(5948684761759817, "Create", "创建角色", "创建角色"),
}
