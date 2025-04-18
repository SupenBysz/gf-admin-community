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
	switch code {
	case e.Reject.Code():
		return e.Reject
	case e.WaitReview.Code():
		return e.WaitReview
	case e.Approve.Code():
		return e.Approve
	default:
		return enum.New[ActionEnum](code, description)
	}
}
