package message

import "github.com/kysion/base-library/utility/enum"

// StateEnum 状态：0未读，1已读
type StateEnum enum.IEnumCode[int]

type state struct {
	UnRead StateEnum
	Readed StateEnum
}

var State = state{
	UnRead: enum.New[StateEnum](0, "未读"),
	Readed: enum.New[StateEnum](1, "已读"),
}

func (e state) New(code int) StateEnum {
	if code == State.UnRead.Code() {
		return State.UnRead
	}
	if (code & State.Readed.Code()) == State.Readed.Code() {
		return State.Readed
	}

	panic("Message.Type.New: error")
}
