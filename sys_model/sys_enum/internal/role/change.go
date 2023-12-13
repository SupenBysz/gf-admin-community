package organization

import "github.com/kysion/base-library/utility/enum"

type ChangeEnum enum.IEnumCode[int]

type roleMemberChange struct {
	Add    ChangeEnum
	Remove ChangeEnum
}

var Change = roleMemberChange{
	Add:    enum.New[ChangeEnum](1, "增加成员"),
	Remove: enum.New[ChangeEnum](2, "移除成员"),
}

func (e roleMemberChange) New(code int, description string) ChangeEnum {
	if (code&Change.Add.Code()) == Change.Add.Code() ||
		(code&Change.Remove.Code()) == Change.Remove.Code() {
		return enum.New[ChangeEnum](code, description)
	}
	panic("user.Change.New: error")
}
