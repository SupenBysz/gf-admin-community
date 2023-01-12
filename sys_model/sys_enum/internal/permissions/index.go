package permissions

type permissions struct {
	PermissionType permissionType
	MatchMode      matchMode
}

var Permissions = permissions{
	PermissionType: PermissionType,
	MatchMode:      MatchMode,
}
