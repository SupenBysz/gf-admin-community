package invite

type invite struct {
	Type  inviteType
	State inviteState
}

var Invite = invite{
	Type:  Type,
	State: State,
}
