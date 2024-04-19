package industry

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
	Tree       PermissionTypeEnum
	Update     PermissionTypeEnum
	Delete     PermissionTypeEnum
	Create     PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: permission.New(5948649434487532, "ViewDetail", "查看行业类别", "查看某个行业类别"),
	Tree:       permission.New(5948649530309523, "Tree", "行业类别树", "查看行业类别树"),
	Update:     permission.New(5948649642787342, "Update", "更新行业类别", "更新某个行业类别"),
	Delete:     permission.New(5948649739509856, "Delete", "删除行业类别", "删除某个行业类别"),
	Create:     permission.New(5948649828767321, "Create", "创建行业类别", "创建行业类别"),
}

//
//func PermissionType() *permissionType {
//	return &permissionType{
//		ViewDetail: base_permission.New(5948649434487532, "ViewDetail", "查看行业类别", "查看某个行业类别"),
//		Tree:       base_permission.New(5948649530309523, "Tree", "行业类别树", "查看行业类别树"),
//		Update:     base_permission.New(5948649642787342, "Update", "更新行业类别", "更新某个行业类别"),
//		Delete:     base_permission.New(5948649739509856, "Delete", "删除行业类别", "删除某个行业类别"),
//		Create:     base_permission.New(5948649828767321, "Create", "创建行业类别", "创建行业类别"),
//	}
//}
