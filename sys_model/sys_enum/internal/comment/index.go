package sys_enum_comment

type comment struct {
	PermissionType permissionType
}

var Comment = comment{
	PermissionType: PermissionType,
}
