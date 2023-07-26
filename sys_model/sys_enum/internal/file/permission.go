package sys_enum_file

import (
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: base_permission.New(5970726323419589, "ViewDetail", "查看资源文件", "查看所有资源文件"),
}
