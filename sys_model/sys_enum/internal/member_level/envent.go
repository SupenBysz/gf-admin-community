package sys_member_level

import "github.com/kysion/base-library/utility/enum"

type EventEnum enum.IEnumCode[int]

type eventState struct {
	Created EventEnum
	Updated EventEnum
	Deleted EventEnum
}

var Event = eventState{
	Created: enum.New[EventEnum](1, "创建"),
	Updated: enum.New[EventEnum](2, "更新"),
	Deleted: enum.New[EventEnum](4, "删除"),
}

func (e eventState) New(code int, description string) EventEnum {
	if (code&Event.Created.Code()) == Event.Created.Code() ||
		(code&Event.Updated.Code()) == Event.Updated.Code() ||
		(code&Event.Deleted.Code()) == Event.Deleted.Code() {
		return enum.New[EventEnum](code, description)
	}
	panic("kyAudit.Event.New: error")
}
