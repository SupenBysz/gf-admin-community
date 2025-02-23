package sys_enum_category

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
	Update     PermissionTypeEnum
	Delete     PermissionTypeEnum
	Create     PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: permission.New(614206186025041, "ViewDetail", "查看分类信息", "查看某个分类信息"),
	Update:     permission.New(614206186025042, "Update", "更新分类信息", "更新某个分类信息"),
	Delete:     permission.New(614206186025043, "Delete", "删除分类信息", "删除某个分类信息"),
	Create:     permission.New(614206186025044, "Create", "创建分类信息", "创建分类信息"),
}
