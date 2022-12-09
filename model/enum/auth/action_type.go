package kyAuth

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type ActionTypeEnum kyEnum.Code

type actionType struct {
	Login    ActionTypeEnum
	Logout   ActionTypeEnum
	Register ActionTypeEnum
}

var ActionType = actionType{
	Login:    kyEnum.New(1, "登录"),
	Logout:   kyEnum.New(2, "退出"),
	Register: kyEnum.New(4, "注册"),
}

func (e actionType) New(code int, description string) ActionTypeEnum {
	if (code&ActionType.Login.Code()) == ActionType.Login.Code() ||
		(code&ActionType.Logout.Code()) == ActionType.Logout.Code() ||
		(code&ActionType.Register.Code()) == ActionType.Register.Code() {
		return kyEnum.NewT[ActionTypeEnum](code, description)
	}
	panic("kyAuth.ActionType.New: error")
}
