package kyAudit

import "github.com/SupenBysz/gf-admin-community/model/enum"

var (
	Created   = kyEnum.New(1, "有新的审核申请")
	ReSubmit  = kyEnum.New(2, "有再次提交的申请")
	ExecAudit = kyEnum.New(4, "处理审核")
)

type EventState kyEnum.Code
