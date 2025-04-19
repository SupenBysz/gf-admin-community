package sys_enum_license

import "github.com/kysion/base-library/utility/enum"

// StateEnum  资质状态：0失效、1正常、-1未认证

type StateEnum enum.IEnumCode[int]

type state struct {
	Invalid    StateEnum
	Normal     StateEnum
	UnVerified StateEnum
}

var State = state{
	Invalid:    enum.New[StateEnum](-1, "失效"),
	Normal:     enum.New[StateEnum](1, "正常"),
	UnVerified: enum.New[StateEnum](0, "未认证"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Invalid.Code()) == State.Invalid.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() ||
		(code&State.UnVerified.Code()) == State.UnVerified.Code() {
		return enum.New[StateEnum](code, description)
	} else {
		panic("License.State.New: error")
	}
}
