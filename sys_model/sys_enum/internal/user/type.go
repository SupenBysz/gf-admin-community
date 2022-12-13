package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type TypeEnum enum.Code

type userType struct {
	Anonymous   TypeEnum
	User        TypeEnum
	WeBusiness  TypeEnum
	Merchant    TypeEnum
	Advertiser  TypeEnum
	Facilitator TypeEnum
	Operator    TypeEnum
	SuperAdmin  TypeEnum
}

var Type = userType{
	Anonymous:   enum.New(0, "匿名"),
	User:        enum.New(1, "用户"),
	WeBusiness:  enum.New(2, "微商"),
	Merchant:    enum.New(4, "商户"),
	Advertiser:  enum.New(8, "广告主"),
	Facilitator: enum.New(16, "服务商"),
	Operator:    enum.New(32, "运营商"),
	SuperAdmin:  enum.New(-1, "超级管理员"),
}

func (e userType) New(code int, description string) TypeEnum {
	if (code&Type.Anonymous.Code()) == Type.Anonymous.Code() ||
		(code&Type.User.Code()) == Type.User.Code() ||
		(code&Type.WeBusiness.Code()) == Type.WeBusiness.Code() ||
		(code&Type.Merchant.Code()) == Type.Merchant.Code() ||
		(code&Type.Advertiser.Code()) == Type.Advertiser.Code() ||
		(code&Type.Facilitator.Code()) == Type.Facilitator.Code() ||
		(code&Type.Operator.Code()) == Type.Operator.Code() ||
		(code&Type.SuperAdmin.Code()) == Type.SuperAdmin.Code() {
		return enum.New(code, description)
	}
	panic("User.Type.New: error")
}
