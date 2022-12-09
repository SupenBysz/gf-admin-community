package kyAudit

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type ActionEnum kyEnum.Code

type action struct {
	Reject     ActionEnum
	WaitReview ActionEnum
	Approve    ActionEnum
}

var Action = action{
	Reject:     kyEnum.NewT[ActionEnum](-1, "不通过"),
	WaitReview: kyEnum.NewT[ActionEnum](0, "待审核"),
	Approve:    kyEnum.NewT[ActionEnum](1, "通过"),
}

func (e action) New(code int, description string) ActionEnum {
	if (code&Action.Reject.Code()) == Action.Reject.Code() ||
		(code&Action.WaitReview.Code()) == Action.WaitReview.Code() ||
		(code&Action.Approve.Code()) == Action.Approve.Code() {
		return kyEnum.NewT[ActionEnum](code, description)
	} else {
		panic("kyAudit.Action.New: error")
	}
}
