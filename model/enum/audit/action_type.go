package kyAudit

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

var (
	Reject     = kyEnum.New(-1, "不通过")
	WaitReview = kyEnum.New(0, "待审核")
	Approve    = kyEnum.New(1, "通过")
)

type ActionType kyEnum.Code
