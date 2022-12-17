package casbin

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type EventEnum enum.Code

type event struct {
	Check EventEnum
}

var Event = event{
	Check: enum.New(1, "检验"),
}

func (e event) New(code int, description string) EventEnum {
	if (code & Event.Check.Code()) == Event.Check.Code() {
		return enum.New(code, description)
	}
	panic("Casbin.Event.New: error")
}
