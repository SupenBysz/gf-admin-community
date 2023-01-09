package sys_enum_user

type user struct {
	Event          event
	Type           userType
	State          state
	PermissionType permissionType
}

var User = user{
	Event:          Event,
	Type:           Type,
	State:          State,
	PermissionType: PermissionType,
}
