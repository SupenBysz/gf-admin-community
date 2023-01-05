package sys_enum_casbin

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type EventEnum enum.IEnumCode[int]

type event struct {
	Check EventEnum
}

var Event = event{
	Check: enum.New[EventEnum](1, "检验"),
}

func (e event) New(code int, description string) EventEnum {
	if (code & Event.Check.Code()) == Event.Check.Code() {
		return enum.New[EventEnum](code, description)
	}
	panic("Casbin.Event.New: error")
}
