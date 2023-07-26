package organization

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
	ViewDetail: base_permission.New(5948649434447941, "ViewDetail", "查看组织架构", "查看某个组织架构"),
	List:       base_permission.New(5948649530392645, "List", "组织架构列表", "查看所有组织架构列表"),
	Update:     base_permission.New(5948649642721349, "Update", "更新组织架构", "更新某个组织架构"),
	Delete:     base_permission.New(5948649739583557, "Delete", "删除组织架构", "删除某个组织架构"),
	Create:     base_permission.New(5948649828712517, "Create", "创建组织架构", "创建组织架构"),
}
