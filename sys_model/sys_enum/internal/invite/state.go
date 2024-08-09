package invite

import "github.com/kysion/base-library/utility/enum"

// StateEnum 邀约状态： 0失效、1正常
type StateEnum enum.IEnumCode[int]

type inviteState struct {
	Invalid StateEnum
	Normal  StateEnum
}

var State = inviteState{
	Invalid: enum.New[StateEnum](0, "失效"),
	Normal:  enum.New[StateEnum](1, "正常"),
}

func (e inviteState) New(code int, description string) StateEnum {

	if code == State.Normal.Code() {
		return e.Normal
	}

	if (code & State.Invalid.Code()) == State.Invalid.Code() {
		return e.Invalid
	}

	panic("Invite.State.New: error")
}
