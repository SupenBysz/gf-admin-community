package message

import "github.com/kysion/base-library/utility/enum"

// TypeEnum 类型：1系统消息，
type TypeEnum enum.IEnumCode[int]

type messageType struct {
	System TypeEnum
}

var MessageType = messageType{
	System: enum.New[TypeEnum](1, "系统消息"),
}

func (e messageType) New(code int, description string) TypeEnum {
	if (code & MessageType.System.Code()) == MessageType.System.Code() {
		return MessageType.System
	}

	// 支持自定义拓展
	return enum.New[TypeEnum](code, description)
}
