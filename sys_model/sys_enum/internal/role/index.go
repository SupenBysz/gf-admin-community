package organization

type role struct {
	PermissionType permissionType
	Change         roleMemberChange
}

var Role = role{
	PermissionType: PermissionType,
	Change:         Change,
}
