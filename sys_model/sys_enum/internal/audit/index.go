package sys_enum_audit

type audit struct {
	Action         action
	Event          eventState
	Category       category
	PermissionType permissionType
	AuditState     auditState
}

var Audit = audit{
	Action:         Action,
	Event:          Event,
	Category:       Category,
	PermissionType: PermissionType,
	AuditState:     AuditState,
}
