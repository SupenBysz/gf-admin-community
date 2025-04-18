package sys_enum_audit

import "github.com/kysion/base-library/utility/enum"

type AuditStateEnum enum.IEnumCode[int]

type auditState struct {
	Cancel        AuditStateEnum
	Reject        AuditStateEnum
	WaitReview    AuditStateEnum
	Approve       AuditStateEnum
	Reviewing     AuditStateEnum
	PendingReview AuditStateEnum
}

var AuditState = auditState{
	Cancel:        enum.New[AuditStateEnum](-2, "已取消"),
	Reject:        enum.New[AuditStateEnum](-1, "不通过"),
	WaitReview:    enum.New[AuditStateEnum](0, "待审核"),
	Approve:       enum.New[AuditStateEnum](1, "通过"),
	Reviewing:     enum.New[AuditStateEnum](2, "审核中"),
	PendingReview: enum.New[AuditStateEnum](3, "补充资料待审核"),
}

func (e auditState) New(code int, description string) AuditStateEnum {
	switch code {
	case e.Reject.Code():
		return e.Reject
	case e.WaitReview.Code():
		return e.WaitReview
	case e.Approve.Code():
		return e.Approve
	case e.Reviewing.Code():
		return e.Reviewing
	case e.PendingReview.Code():
		return e.PendingReview
	default:
		return enum.New[AuditStateEnum](code, description)
	}
}
