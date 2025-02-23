package sys_enum_audit

import "github.com/kysion/base-library/utility/enum"

type ActionEnum enum.IEnumCode[int]

type action struct {
	Cancel     ActionEnum
	Reject     ActionEnum
	WaitReview ActionEnum
	Approve    ActionEnum
}

var Action = action{
	Cancel:     enum.New[ActionEnum](-2, "已取消"),
	Reject:     enum.New[ActionEnum](-1, "不通过"),
	WaitReview: enum.New[ActionEnum](0, "待审核"),
	Approve:    enum.New[ActionEnum](1, "通过"),
}

func (e action) New(code int, description string) ActionEnum {
	if code == Action.Reject.Code() ||
		code == Action.WaitReview.Code() ||
		code == Action.Approve.Code() {
		return enum.New[ActionEnum](code, description)
	} else {
		panic("kyAudit.Action.New: error")
	}
}
