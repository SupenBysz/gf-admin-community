package delivery

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	View   PermissionTypeEnum
	List   PermissionTypeEnum
	Update PermissionTypeEnum
	Delete PermissionTypeEnum
	Create PermissionTypeEnum
}

var PermissionType = permissionType{
	View:   permission.New(618607339090001, "ViewDetail", "查看物流公司", "查看某个物流公司"),
	List:   permission.New(618607339090002, "List", "物流公司列表", "查看所有物流公司"),
	Update: permission.New(618607339090003, "Update", "更新物流公司信息", "更新某个物流公司信息"),
	Delete: permission.New(618607339090004, "Delete", "删除物流公司", "删除某个物流公司"),
	Create: permission.New(618607339090005, "Create", "创建物流公司", "创建一个新物流公司"),
}
