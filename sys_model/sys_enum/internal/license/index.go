package sys_enum_license

type license struct {
	PermissionType permissionType
}

var License = license{
	PermissionType: PermissionType,
}
