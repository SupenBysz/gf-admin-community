package sys_enum_audit

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type EventEnum enum.IEnumCode[int]

type eventState struct {
	Created   EventEnum
	ReSubmit  EventEnum
	ExecAudit EventEnum
}

var Event = eventState{
	Created:   enum.New[EventEnum](1, "有新的审核申请"),
	ReSubmit:  enum.New[EventEnum](2, "有再次提交的申请"),
	ExecAudit: enum.New[EventEnum](4, "处理审核"),
}

func (e eventState) New(code int, description string) EventEnum {
	if (code&Event.Created.Code()) == Event.Created.Code() ||
		(code&Event.ReSubmit.Code()) == Event.ReSubmit.Code() ||
		(code&Event.ExecAudit.Code()) == Event.ExecAudit.Code() {
		return enum.New[EventEnum](code, description)
	}
	panic("kyAudit.Event.New: error")
}
