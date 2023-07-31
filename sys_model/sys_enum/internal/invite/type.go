package invite

import "github.com/kysion/base-library/utility/enum"

// TypeEnum 邀约类型：1注册、2加入团队、4加入角色， （复合类型）
type TypeEnum enum.IEnumCode[int]

type inviteType struct {
	Register TypeEnum
	JoinTeam TypeEnum
	JoinRole TypeEnum
}

var Type = inviteType{
	Register: enum.New[TypeEnum](1, "注册"),
	JoinTeam: enum.New[TypeEnum](2, "加入团队"),
	JoinRole: enum.New[TypeEnum](4, "加入角色"),
}

func (e inviteType) New(code int, description string) TypeEnum {
	if (code & Type.Register.Code()) == Type.Register.Code() {
		return e.Register
	}
	if (code & Type.JoinTeam.Code()) == Type.JoinTeam.Code() {
		return e.JoinTeam
	}
	if (code & Type.JoinRole.Code()) == Type.JoinRole.Code() {
		return e.JoinRole
	}

	return enum.New[TypeEnum](code, description)
}
