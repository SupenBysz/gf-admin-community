package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type StateEnum enum.Code

type state struct {
	Unactivated StateEnum
	Normal      StateEnum
	Suspended   StateEnum
	Abnormality StateEnum
	Canceled    StateEnum
}

var State = state{
	Unactivated: enum.New(0, "未激活"),
	Normal:      enum.New(1, "正常"),
	Suspended:   enum.New(-1, "封号"),
	Abnormality: enum.New(-2, "异常"),
	Canceled:    enum.New(-3, "已注销"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Unactivated.Code()) == State.Unactivated.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() ||
		(code&State.Suspended.Code()) == State.Suspended.Code() ||
		(code&State.Abnormality.Code()) == State.Abnormality.Code() ||
		(code&State.Canceled.Code()) == State.Canceled.Code() {
		return enum.New(code, description)
	}
	panic("User.State.New: error")
}
