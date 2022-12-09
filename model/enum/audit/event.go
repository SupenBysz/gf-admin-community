package kyAudit

import "github.com/SupenBysz/gf-admin-community/model/enum"

type EventEnum kyEnum.Code

type eventState struct {
	Created   EventEnum
	ReSubmit  EventEnum
	ExecAudit EventEnum
}

var Event = eventState{
	Created:   kyEnum.NewT[EventEnum](1, "有新的审核申请"),
	ReSubmit:  kyEnum.NewT[EventEnum](2, "有再次提交的申请"),
	ExecAudit: kyEnum.NewT[EventEnum](4, "处理审核"),
}

func (e eventState) New(code int, description string) EventEnum {
	if (code&Event.Created.Code()) == Event.Created.Code() ||
		(code&Event.ReSubmit.Code()) == Event.ReSubmit.Code() ||
		(code&Event.ExecAudit.Code()) == Event.ExecAudit.Code() {
		return kyEnum.NewT[EventEnum](code, description)
	}
	panic("kyAudit.Event.New: error")
}
