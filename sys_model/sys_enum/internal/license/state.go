package sys_enum_license

import "github.com/kysion/base-library/utility/enum"

// StateEnum  资质状态：0失效、1正常

type StateEnum enum.IEnumCode[int]

type state struct {
	Disabled StateEnum
	Normal   StateEnum
}

var State = state{
	Disabled: enum.New[StateEnum](0, "失效"),
	Normal:   enum.New[StateEnum](1, "正常"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Disabled.Code()) == State.Disabled.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() {
		return enum.New[StateEnum](code, description)
	} else {
		panic("License.State.New: error")
	}
}
