package sys_enum_user

import "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"

type StateEnum sys_enum.Code

type state struct {
	Unactivated StateEnum
	Normal      StateEnum
	Suspended   StateEnum
	Abnormality StateEnum
	Canceled    StateEnum
}

var State = state{
	Unactivated: sys_enum.New(0, "未激活"),
	Normal:      sys_enum.New(1, "正常"),
	Suspended:   sys_enum.New(-1, "封号"),
	Abnormality: sys_enum.New(-2, "异常"),
	Canceled:    sys_enum.New(-3, "已注销"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Unactivated.Code()) == State.Unactivated.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() ||
		(code&State.Suspended.Code()) == State.Suspended.Code() ||
		(code&State.Abnormality.Code()) == State.Abnormality.Code() ||
		(code&State.Canceled.Code()) == State.Canceled.Code() {
		return sys_enum.New(code, description)
	}
	panic("kyUser.UserState.New: error")
}
