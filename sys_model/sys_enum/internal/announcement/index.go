package announcement

type announcement struct {
	State    state
	FlagRead flagRead
}

var Announcement = announcement{
	State:    State,
	FlagRead: FlagRead,
}
