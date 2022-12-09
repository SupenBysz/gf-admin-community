package kyUser

import "github.com/SupenBysz/gf-admin-community/model/enum"

type EventEnum kyEnum.Code

type event struct {
	BeforeCreate EventEnum
	AfterCreate  EventEnum
}

var Event = event{
	BeforeCreate: kyEnum.NewT[EventEnum](2, "创建前"),
	AfterCreate:  kyEnum.NewT[EventEnum](4, "创建后"),
}

func (e event) New(code int, description string) EventEnum {
	if (code&Event.BeforeCreate.Code()) == Event.BeforeCreate.Code() ||
		(code&Event.AfterCreate.Code()) == Event.AfterCreate.Code() {
		return kyEnum.NewT[EventEnum](code, description)
	}
	panic("kyUser.Event.New: error")
}
