package sys_enum_audit

import "github.com/kysion/base-library/utility/enum"

type AuditStateEnum enum.IEnumCode[int]

type auditState struct {
	Cancel                   AuditStateEnum
	WaitReview               AuditStateEnum
	Approved                 AuditStateEnum
	Rejected                 AuditStateEnum
	Reviewing                AuditStateEnum
	WaitingSupplementaryInfo AuditStateEnum
}

var AuditState = auditState{
	Cancel:                   enum.New[AuditStateEnum](-1, "已取消"),
	WaitReview:               enum.New[AuditStateEnum](0, "待审核"),
	Approved:                 enum.New[AuditStateEnum](1, "通过"),
	Rejected:                 enum.New[AuditStateEnum](2, "拒绝通过"),
	Reviewing:                enum.New[AuditStateEnum](3, "审核中（人工复审）"),
	WaitingSupplementaryInfo: enum.New[AuditStateEnum](4, "补充资料待审核"),
}

func (e auditState) New(code int, description string) AuditStateEnum {
	switch code {
	case e.Rejected.Code():
		return e.Rejected
	case e.WaitReview.Code():
		return e.WaitReview
	case e.Approved.Code():
		return e.Approved
	case e.Reviewing.Code():
		return e.Reviewing
	case e.WaitingSupplementaryInfo.Code():
		return e.WaitingSupplementaryInfo
	default:
		return enum.New[AuditStateEnum](code, description)
	}
}
