package sys_enum_business

import (
	"github.com/kysion/base-library/utility/enum"
)

type TypeEnum enum.IEnumCode[int]

type businessType struct {
	UnKnow      TypeEnum
	Operator    TypeEnum
	Facilitator TypeEnum
	Member      TypeEnum
	WeBusiness  TypeEnum
	Merchant    TypeEnum
}

var Type = businessType{
	UnKnow:      enum.New[TypeEnum](0, "未知业务"),
	Operator:    enum.New[TypeEnum](1, "运营商"),
	Facilitator: enum.New[TypeEnum](2, "服务商"),
	Member:      enum.New[TypeEnum](4, "消费者"),
	WeBusiness:  enum.New[TypeEnum](8, "微商"),
	Merchant:    enum.New[TypeEnum](16, "商户"),
}

func (e businessType) New(code int, description string) TypeEnum {
	if (code&Type.UnKnow.Code()) == Type.UnKnow.Code() ||
		(code&Type.Operator.Code()) == Type.Operator.Code() ||
		(code&Type.Facilitator.Code()) == Type.Facilitator.Code() ||
		(code&Type.Member.Code()) == Type.Member.Code() ||
		(code&Type.WeBusiness.Code()) == Type.WeBusiness.Code() ||
		(code&Type.Merchant.Code()) == Type.Merchant.Code() {
		return enum.New[TypeEnum](code, description)
	}
	panic("pro_enum_business.Type.New: error")
}
