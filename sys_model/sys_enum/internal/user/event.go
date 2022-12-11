package sys_enum_user

import "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"

type EventEnum sys_enum.Code

type event struct {
	BeforeCreate EventEnum
	AfterCreate  EventEnum
}

var Event = event{
	BeforeCreate: sys_enum.New(2, "创建前"),
	AfterCreate:  sys_enum.New(4, "创建后"),
}

func (e event) New(code int, description string) EventEnum {
	if (code&Event.BeforeCreate.Code()) == Event.BeforeCreate.Code() ||
		(code&Event.AfterCreate.Code()) == Event.AfterCreate.Code() {
		return sys_enum.New(code, description)
	}
	panic("kyUser.Event.New: error")
}
