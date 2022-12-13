package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type EventEnum enum.Code

type event struct {
	BeforeCreate   EventEnum
	AfterCreate    EventEnum
	ChangePassword EventEnum
	ResetPassword  EventEnum
}

var Event = event{
	BeforeCreate:   enum.New(2, "创建前"),
	AfterCreate:    enum.New(4, "创建后"),
	ChangePassword: enum.New(8, "修改密码"),
	ResetPassword:  enum.New(16, "重置密码"),
}

func (e event) New(code int, description string) EventEnum {
	if (code&Event.BeforeCreate.Code()) == Event.BeforeCreate.Code() ||
		(code&Event.AfterCreate.Code()) == Event.AfterCreate.Code() ||
		(code&Event.ChangePassword.Code()) == Event.ChangePassword.Code() ||
		(code&Event.ResetPassword.Code()) == Event.ResetPassword.Code() {
		return enum.New(code, description)
	}
	panic("user.Event.New: error")
}
