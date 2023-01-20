package sys_enum_user

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type TypeEnum enum.IEnumCode[int]

type userType struct {
	SuperAdmin TypeEnum
}

var Type = userType{
	SuperAdmin: enum.New[TypeEnum](-1, "超级管理员"),
}

func (e userType) New(code int, description string) TypeEnum {
	if (code & Type.SuperAdmin.Code()) == Type.SuperAdmin.Code() {
		return Type.SuperAdmin
	}
	return enum.New[TypeEnum](code, description)
}
