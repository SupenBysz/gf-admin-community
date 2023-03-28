package sys_enum_user

import "github.com/kysion/base-library/utility/enum"

type TypeEnum enum.IEnumCode[int]

type userType struct {
    Admin TypeEnum

    SuperAdmin TypeEnum
}

var Type = userType{
    Admin:      enum.New[TypeEnum](64, "后台"),
    SuperAdmin: enum.New[TypeEnum](-1, "超级管理员"),
}

func (e userType) New(code int, description string) TypeEnum {
    if (code & Type.SuperAdmin.Code()) == Type.SuperAdmin.Code() {
        return Type.SuperAdmin
    }
    if (code & Type.Admin.Code()) == Type.Admin.Code() {
        return Type.Admin
    }
    return enum.New[TypeEnum](code, description)
}
