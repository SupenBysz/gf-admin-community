package model

import (
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

type FdInvoiceRegister struct {
	Name           string      `json:"name"           v:"required#请输入发票抬头名称" dc:"发票抬头名称"`
	TaxId          string      `json:"taxId"          v:"required#请输入纳税识别号" dc:"纳税识别号"`
	Addr           string      `json:"addr"           dc:"发票收件地址，限纸质"`
	Email          string      `json:"email"          dc:"发票收件邮箱，限电子发票"`
	UserId         int64       `json:"userId"         v:"required#请输入申请人UserID" dc:"申请人UserID"`
	AuditUserId    int64       `json:"auditUserId"    v:"required#请输入审核人UserID" dc:"审核人UserID"`
	AuditReplayMsg string      `json:"auditReplayMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt        *gtime.Time `json:"auditAt"        v:"required#审核时间不能为空" dc:"审核时间"`
	State          int         `json:"state"          v:"required|in:1,2,3#请输入发票审核状态" dc:"状态：1待审核、2已通过、3不通过"`
}

type FdInvoiceInfo entity.FdInvoice

type FdInvoiceListRes CollectRes[entity.FdInvoice]
