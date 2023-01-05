package organization

type role struct {
	PermissionType permissionType
}

var Role = role{
	PermissionType: PermissionType,
}
