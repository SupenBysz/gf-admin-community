package announcement

import "github.com/kysion/base-library/utility/enum"

// FlagReadEnum 状态：0未读，1已读
type FlagReadEnum enum.IEnumCode[int]

type flagRead struct {
	UnRead FlagReadEnum
	Readed FlagReadEnum
}

var FlagRead = flagRead{
	UnRead: enum.New[FlagReadEnum](0, "未读"),
	Readed: enum.New[FlagReadEnum](1, "已读"),
}

func (e flagRead) New(code int) FlagReadEnum {
	if code == FlagRead.UnRead.Code() {
		return FlagRead.UnRead
	}
	if code == FlagRead.Readed.Code() {
		return FlagRead.Readed
	}

	panic("Announcement.Type.New: error")
}
