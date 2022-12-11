package sys_enum_upload

type upload struct {
	EventState eventState
}

var Upload = upload{
	EventState: EventState,
}
