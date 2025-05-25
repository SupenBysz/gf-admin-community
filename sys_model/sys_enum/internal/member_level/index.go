package sys_member_level

type memberLevel struct {
	Event eventState
}

var MemberLevel = memberLevel{
	Event: Event,
}
