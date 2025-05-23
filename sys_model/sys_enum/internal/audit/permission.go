package sys_enum_audit

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
	List       PermissionTypeEnum
	Update     PermissionTypeEnum
	Cancel     PermissionTypeEnum
}

var (
	PermissionType = permissionType{
		ViewDetail: permission.New(5953151699124297, "ViewDetail", "查看审核信息", "查看某条审核信息"),
		List:       permission.New(5953151699124298, "List", "获取审核列表", "查看所有审核"),
		Update:     permission.New(5953151699124299, "Update", "更新审核信息", "更新某条审核信息"),
		Cancel:     permission.New(5953151699124300, "Cancel", "取消审核信息", "撤销某条审核信息"),
	}
	permissionTypeMap = gmap.NewStrAnyMapFrom(gconv.Map(PermissionType))
)
