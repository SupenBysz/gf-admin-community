package sys_enum_audit

type audit struct {
	Action         auditState
	Event          eventState
	Category       category
	PermissionType permissionType
	AuditState     auditState
}

var Audit = audit{
	Action:         AuditState,
	Event:          Event,
	Category:       Category,
	PermissionType: PermissionType,
	AuditState:     AuditState,
}
