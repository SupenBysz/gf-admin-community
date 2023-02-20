package sys_enum_user

import "github.com/kysion/base-library/utility/enum"

type EventEnum enum.IEnumCode[int]

type event struct {
	BeforeCreate   EventEnum
	AfterCreate    EventEnum
	ChangePassword EventEnum
	ResetPassword  EventEnum
}

var Event = event{
	BeforeCreate:   enum.New[EventEnum](2, "创建前"),
	AfterCreate:    enum.New[EventEnum](4, "创建后"),
	ChangePassword: enum.New[EventEnum](8, "修改密码"),
	ResetPassword:  enum.New[EventEnum](16, "重置密码"),
}

func (e event) New(code int, description string) EventEnum {
	if (code&Event.BeforeCreate.Code()) == Event.BeforeCreate.Code() ||
		(code&Event.AfterCreate.Code()) == Event.AfterCreate.Code() ||
		(code&Event.ChangePassword.Code()) == Event.ChangePassword.Code() ||
		(code&Event.ResetPassword.Code()) == Event.ResetPassword.Code() {
		return enum.New[EventEnum](code, description)
	}
	panic("user.Event.New: error")
}
