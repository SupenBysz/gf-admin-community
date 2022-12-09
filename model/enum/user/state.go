package kyUser

import "github.com/SupenBysz/gf-admin-community/model/enum"

type StateEnum kyEnum.Code

type state struct {
	Unactivated StateEnum
	Normal      StateEnum
	Suspended   StateEnum
	Abnormality StateEnum
	Canceled    StateEnum
}

var State = state{
	Unactivated: kyEnum.New(0, "未激活"),
	Normal:      kyEnum.New(1, "正常"),
	Suspended:   kyEnum.New(-1, "封号"),
	Abnormality: kyEnum.New(-2, "异常"),
	Canceled:    kyEnum.New(-3, "已注销"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.Unactivated.Code()) == State.Unactivated.Code() ||
		(code&State.Normal.Code()) == State.Normal.Code() ||
		(code&State.Suspended.Code()) == State.Suspended.Code() ||
		(code&State.Abnormality.Code()) == State.Abnormality.Code() ||
		(code&State.Canceled.Code()) == State.Canceled.Code() {
		return kyEnum.NewT[StateEnum](code, description)
	}
	panic("kyUser.UserState.New: error")
}
