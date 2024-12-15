package sys_enum_comment

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
	ViewDetail: permission.New(622985990934602, "ViewDetail", "查看评论信息", "查看某个评论信息"),
	Update:     permission.New(622985990934603, "Update", "更新评论信息", "更新某个评论信息"),
	Delete:     permission.New(622985990934604, "Delete", "删除评论信息", "删除某个评论信息"),
	Create:     permission.New(622985990934605, "Create", "创建评论信息", "创建评论信息"),
}
