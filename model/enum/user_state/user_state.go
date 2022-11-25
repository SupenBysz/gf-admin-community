package userState

import "github.com/SupenBysz/gf-admin-community/model/enum"

var (
	Unactivated = kyEnum.New(0, "未激活")
	Normal      = kyEnum.New(1, "正常")
	Suspended   = kyEnum.New(-1, "封号")
	Abnormality = kyEnum.New(-2, "异常")
	Canceled    = kyEnum.New(-3, "已注销")
)
