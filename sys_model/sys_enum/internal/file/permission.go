package sys_enum_file

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type PermissionTypeEnum = *sys_model.SysPermissionTree

type permissionType struct {
	ViewDetail PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: permission.New(5970726323419589, "ViewDetail", "查看资源文件", "查看所有资源文件"),
}
