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
	View:   permission.New(5948649434447941, "View", "查看组织架构", "查看某个组织架构"),
	List:   permission.New(5948649530392645, "List", "组织架构列表", "查看所有组织架构列表"),
	Update: permission.New(5948649642721349, "Update", "更新组织架构", "更新某个组织架构"),
	Delete: permission.New(5948649739583557, "Delete", "删除组织架构", "删除某个组织架构"),
	Create: permission.New(5948649828712517, "Create", "创建组织架构", "创建组织架构"),
}
