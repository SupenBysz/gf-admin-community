package sys_enum_license

type license struct {
	PermissionType permissionType
	State          state
}

var License = license{
	PermissionType: PermissionType,
	State:          State,
}
