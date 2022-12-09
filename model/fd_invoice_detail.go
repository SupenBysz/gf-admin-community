package model

import (
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

type FdInvoiceDetailRegister struct {
	TaxNumber     string      `json:"taxNumber"     v:"required#请输入纳税识别号" dc:"纳税识别号"`
	TaxName       string      `json:"taxName"       v:"required#请输入纳税人名称"  dc:"纳税人名称"`
	BillIds       string      `json:"billIds"       v:"required#请输入开支账单"  dc:"账单ID组"`
	Amount        int64       `json:"amount"        v:"required#请输入开票金额"  dc:"开票金额，单位精度：分"`
	Rate          int         `json:"rate"          v:"required#请输入税率"  dc:"税率，如3% 则填入3"`
	RateMount     int64       `json:"rateMount"     v:"required#请输入税额"  dc:"税额，单位精度：分"`
	Remark        string      `json:"remark"        v:"required#请输入发票内容描述"  dc:"发票内容描述"`
	Type          int         `json:"type"          v:"required|in:1,2#请输入发票类型"  dc:"发票类型：1电子发票，2纸质发票"`
	State         int         `json:"state"         v:"required|in:1,2,4,8,16#请输入发票状态"  dc:"状态：1待审核、2待开票、4开票失败、8已开票、16已撤销"`
	AuditUserIds  int64       `json:"auditUserIds"  dc:"审核者UserID，多个用逗号隔开"`
	MakeType      int         `json:"makeType"      dc:"出票类型：1普通发票、2增值税专用发票、3专业发票"`
	MakeUserId    int64       `json:"makeUserId"    dc:"出票人UserID，如果是系统出票则默认-1"`
	MakeAt        *gtime.Time `json:"makeAt"        dc:"出票时间"`
	CourierName   string      `json:"courierName"   dc:"快递名称"`
	CourierNumber string      `json:"courierNumber" dc:"快递编号"`
	FdInvoiceId   int64       `json:"fdInvoiceId"   dc:"发票抬头ID"`
	AuditUserId   int64       `json:"auditUserId"   dc:"审核者UserID"`
	AuditReplyMsg string      `json:"auditReplyMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt       *gtime.Time `json:"auditAt"       dc:"审核时间"`
}

type FdInvoiceDetailInfo entity.FdInvoiceDetail
type FdInvoiceDetailListRes CollectRes[entity.FdInvoiceDetail]
