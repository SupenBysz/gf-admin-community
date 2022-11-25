package model

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyAudit "github.com/SupenBysz/gf-admin-community/model/enum/audit"
)

type CreateSysAudit struct {
	Id        int64       `json:"id"            description:""`
	State     int         `json:"state"         description:"审核状态：-1不通过，0待审核，1通过" v:"required|in:-1,0,1#审核状态错误" default:"0"`
	UnionId   int64       `json:"unionId"       description:"关联业务ID" v:"required#关联业务ID参数粗我"`
	Category  int         `json:"category"      description:"分类：1运营商主体资质审核，2服务商主体资质审核、4消费者实名审核" v:"required|in:1,2,4#分类类型错误"`
	AuditData string      `json:"auditData"     description:"待审核的业务数据包" v:"required|json#验证信息必须为json格式字符串"`
	ExpireAt  *gtime.Time `json:"expireAt"      description:"审核服务时限，超过该时间后没有审核通过的需要重新申请审核"`
}

type SetAudit struct {
	Id       int64  `json:"id"            description:"ID" v:"required#审核编号错误"`
	State    int    `json:"state"         description:"审核状态：-1不通过，1通过" v:"required|in:-1,1#审核状态错误"`
	Replay   string `json:"replay"        description:"不通过时回复的审核不通过原因"`
	UnionId  int64  `json:"unionId"       description:"关联业务ID" v:"required#关联业务ID参数粗我"`
	Category int    `json:"category"      description:"分类：1运营商主体资质审核，2服务商主体资质审核、4消费者实名审核" v:"required|in:1,2,4#分类类型错误"`
}

type AuditRes struct {
	Id            int64       `json:"auditId"       description:""`
	State         int         `json:"state"         description:"审核状态：-1不通过，0待审核，1通过"`
	Replay        string      `json:"replay"        description:"不通过时回复的审核不通过原因"`
	ExpireAt      *gtime.Time `json:"expireAt"      description:"服务时限"`
	AuditReplayAt *gtime.Time `json:"auditReplayAt" description:"审核回复时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
}

type SysAuditListRes CollectRes[entity.SysAudit]

type AuditHookFunc func(ctx context.Context, state kyAudit.EventState, info entity.SysAudit) error
type AuditHookInfo struct {
	Key      kyAudit.EventState
	Value    AuditHookFunc
	Category int `json:"category" dc:"业务类型"`
}
