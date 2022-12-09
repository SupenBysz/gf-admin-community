package kyAudit

import "github.com/SupenBysz/gf-admin-community/model/enum"

type EventEnum kyEnum.Code

type eventState struct {
	Created   EventEnum
	ReSubmit  EventEnum
	ExecAudit EventEnum
}

var Event = eventState{
	Created:   kyEnum.New(1, "有新的审核申请"),
	ReSubmit:  kyEnum.New(2, "有再次提交的申请"),
	ExecAudit: kyEnum.New(4, "处理审核"),
}

func (e eventState) New(code int, description string) EventEnum {
	if (code&Event.Created.Code()) == Event.Created.Code() ||
		(code&Event.ReSubmit.Code()) == Event.ReSubmit.Code() ||
		(code&Event.ExecAudit.Code()) == Event.ExecAudit.Code() {
		return kyEnum.New(code, description)
	}
	panic("kyAudit.Event.New: error")
}
