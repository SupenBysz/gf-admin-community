// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAudit is the golang structure for table sys_audit.
type SysAudit struct {
	Id             int64       `json:"id"             orm:"id"              description:""`
	State          int         `json:"state"          orm:"state"           description:"审核状态：0 - 待审核；1 - 审核通过；2 - 审核不通过；3 - 审核中（人工复审）；4 - 补充资料待审核"`
	Reply          string      `json:"reply"          orm:"reply"           description:"不通过时回复的审核不通过原因"`
	UnionMainId    int64       `json:"unionMainId"    orm:"union_main_id"   description:"关联主体ID"`
	Category       int         `json:"category"       orm:"category"        description:"业务类别：1个人资质审核、2主体资质审核、4数据审核"`
	AuditData      string      `json:"auditData"      orm:"audit_data"      description:"待审核的业务数据包"`
	ExpireAt       *gtime.Time `json:"expireAt"       orm:"expire_at"       description:"服务时限"`
	AuditReplyAt   *gtime.Time `json:"auditReplyAt"   orm:"audit_reply_at"  description:"审核回复时间"`
	HistoryItems   string      `json:"historyItems"   orm:"history_Items"   description:"历史申请记录"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	AuditUserId    int64       `json:"auditUserId"    orm:"audit_user_id"   description:"审核操作者id"`
	DataIdentifier string      `json:"dataIdentifier" orm:"data_identifier" description:"数据标识"`
	UserId         int64       `json:"userId"         orm:"user_id"         description:"关联用户ID"`
	Summary        string      `json:"summary"        orm:"summary"         description:"概述"`
	AuditGroup     string      `json:"auditGroup"     orm:"audit_group"     description:"审核分组"`
}
