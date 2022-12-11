package sys_enum_auth

type auth struct {
	ActionType actionType
}

var Auth = auth{
	ActionType: ActionType,
}
