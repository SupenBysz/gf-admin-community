package sys_enum_file

type file struct {
	PermissionType permissionType
}

var File = file{
	PermissionType: PermissionType,
}
