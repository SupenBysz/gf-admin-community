package userType

import "github.com/SupenBysz/gf-admin-community/model/enum"

var (
	Anonymous   = kyEnum.New(0, "匿名")
	User        = kyEnum.New(1, "用户")
	WeBusiness  = kyEnum.New(2, "微商")
	Merchant    = kyEnum.New(4, "商户")
	Advertiser  = kyEnum.New(8, "广告主")
	Facilitator = kyEnum.New(16, "服务商")
	Operator    = kyEnum.New(32, "运营商")
	SuperAdmin  = kyEnum.New(-1, "超级管理员")
)

type Code kyEnum.Code
