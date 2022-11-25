package kyAuth

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

var (
	ActionLogin    = kyEnum.New(1, "登录")
	ActionLogout   = kyEnum.New(2, "退出")
	ActionRegister = kyEnum.New(4, "注册")
)

type ActionType kyEnum.Code
