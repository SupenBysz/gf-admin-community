package sys_enum_audit

type audit struct {
	Action         action
	Event          eventState
	PermissionType permissionType
}

var Audit = audit{
	Action:         Action,
	Event:          Event,
	PermissionType: PermissionType,
}
