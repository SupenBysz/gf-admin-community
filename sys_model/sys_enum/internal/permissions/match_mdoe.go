package sys_enum_permission

import "github.com/kysion/base-library/utility/enum"

type MatchModeEnum enum.IEnumCode[int]

type matchMode struct {
	Id         MatchModeEnum
	Identifier MatchModeEnum
}

var MatchMode = matchMode{
	Id:         enum.New[MatchModeEnum](0, "通过ID匹配校验权限"),
	Identifier: enum.New[MatchModeEnum](1, "通过标识符匹配校验权限"),
}
