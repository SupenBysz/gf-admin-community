package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type CreateAudit struct {
	Id             int64       `json:"id"               dc:""`
	State          int         `json:"state"            dc:"审核状态：-1不通过，0待审核，1通过" v:"required|in:-1,0,1#审核状态错误" default:"0"`
	UnionMainId    int64       `json:"unionMainId"      dc:"关联业务主体ID" v:"required#关联业务主体ID不能为空"`
	Category       int         `json:"category"         dc:"业务类别：1个人资质审核、2主体资质审核、4数据审核" v:"required|in:1,2,4#分类类型错误"`
	AuditData      string      `json:"auditData"        dc:"待审核的业务数据包" v:"required|json#验证信息必须为json格式字符串"`
	ExpireAt       *gtime.Time `json:"expireAt"         dc:"审核服务时限，超过该时间后没有审核通过的需要重新申请审核"`
	DataIdentifier string      `json:"dataIdentifier"   dc:"数据标识"`
	UserId         int64       `json:"userId"           dc:"关联业务用户ID" `
	Summary        string      `json:"summary"          dc:"概述"`
	AuditGroup     string      `json:"auditGroup"       dc:"审核分组"`
}

type SetAudit struct {
	Id          int64  `json:"id"            dc:"ID" v:"required#审核编号错误"`
	State       int    `json:"state"         dc:"审核状态：-1不通过，1通过" v:"required|in:-1,1#审核状态错误"`
	Reply       string `json:"reply"         dc:"不通过时回复的审核不通过原因"`
	UnionMainId int64  `json:"unionMainId"   dc:"关联业务ID" v:"required#关联业务ID参数粗我"`
	Category    int    `json:"category"      dc:"分类：1个人资质审核、2主体资质审核、4数据审核" v:"required|in:1,2,4#分类类型错误"`
}

type Audit sys_entity.SysAudit

type AuditRes sys_entity.SysAudit
type AuditListRes base_model.CollectRes[sys_entity.SysAudit]
