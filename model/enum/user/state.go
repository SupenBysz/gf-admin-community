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
	Unactivated: kyEnum.NewT[StateEnum](0, "未激活"),
	Normal:      kyEnum.NewT[StateEnum](1, "正常"),
	Suspended:   kyEnum.NewT[StateEnum](-1, "封号"),
	Abnormality: kyEnum.NewT[StateEnum](-2, "异常"),
	Canceled:    kyEnum.NewT[StateEnum](-3, "已注销"),
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
