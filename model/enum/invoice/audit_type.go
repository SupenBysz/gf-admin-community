package kyInvoice

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type AuditTypeEnum kyEnum.Code

type auditType struct {
	Reject     AuditTypeEnum
	WaitReview AuditTypeEnum
	Approve    AuditTypeEnum
}

var AuditType = auditType{
	WaitReview: kyEnum.New(0, "待审核"),
	Approve:    kyEnum.New(1, "通过"),
	Reject:     kyEnum.New(-1, "不通过"),
}

func (e auditType) New(code int, description string) AuditTypeEnum {
	if (code&AuditType.Reject.Code()) == AuditType.Reject.Code() ||
		(code&AuditType.WaitReview.Code()) == AuditType.WaitReview.Code() ||
		(code&AuditType.Approve.Code()) == AuditType.Approve.Code() {
		return kyEnum.New(code, description)
	} else {
		panic("kyAudit.Action.New: error")
	}
}
