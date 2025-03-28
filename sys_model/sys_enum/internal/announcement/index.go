package announcement

type announcement struct {
	State          state
	FlagRead       flagRead
	PermissionType permissionType
}

var Announcement = announcement{
	State:          State,
	FlagRead:       FlagRead,
	PermissionType: PermissionType,
}
