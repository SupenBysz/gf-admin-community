package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type EventEnum enum.Code

type event struct {
	BeforeCreate EventEnum
	AfterCreate  EventEnum
}

var Event = event{
	BeforeCreate: enum.New(2, "创建前"),
	AfterCreate:  enum.New(4, "创建后"),
}

func (e event) New(code int, description string) EventEnum {
	if (code&Event.BeforeCreate.Code()) == Event.BeforeCreate.Code() ||
		(code&Event.AfterCreate.Code()) == Event.AfterCreate.Code() {
		return enum.New(code, description)
	}
	panic("kyUser.Event.New: error")
}
