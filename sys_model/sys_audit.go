package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type CreateAudit struct {
	Id             int64       `json:"id"            description:""`
	State          int         `json:"state"         description:"审核状态：-1不通过，0待审核，1通过" v:"required|in:-1,0,1#审核状态错误" default:"0"`
	UnionMainId    int64       `json:"unionMainId"   description:"关联业务主体ID" v:"required#关联业务主体ID不能为空"`
	Category       int         `json:"category"      description:"业务类别：1个人资质审核、2主体资质审核、4数据审核" v:"required|in:1,2,4#分类类型错误"`
	AuditData      string      `json:"auditData"     description:"待审核的业务数据包" v:"required|json#验证信息必须为json格式字符串"`
	ExpireAt       *gtime.Time `json:"expireAt"      description:"审核服务时限，超过该时间后没有审核通过的需要重新申请审核"`
	DataIdentifier string      `json:"dataIdentifier" description:"数据标识"`
}

type SetAudit struct {
	Id          int64  `json:"id"            description:"ID" v:"required#审核编号错误"`
	State       int    `json:"state"         description:"审核状态：-1不通过，1通过" v:"required|in:-1,1#审核状态错误"`
	Reply       string `json:"reply"        description:"不通过时回复的审核不通过原因"`
	UnionMainId int64  `json:"unionMainId"       description:"关联业务ID" v:"required#关联业务ID参数粗我"`
	Category    int    `json:"category"      description:"分类：1运营商主体资质审核，2服务商主体资质审核、4消费者实名审核" v:"required|in:1,2,4#分类类型错误"`
}

type AuditRes struct {
	Id           int64       `json:"auditId"       description:""`
	State        int         `json:"state"         description:"审核状态：-1不通过，0待审核，1通过"`
	Reply        string      `json:"reply"        description:"不通过时回复的审核不通过原因"`
	ExpireAt     *gtime.Time `json:"expireAt"      description:"服务时限"`
	AuditReplyAt *gtime.Time `json:"auditReplyAt" description:"审核回复时间"`
	CreatedAt    *gtime.Time `json:"createdAt"     description:""`
}

type Audit sys_entity.SysAudit
type AuditListRes base_model.CollectRes[sys_entity.SysAudit]
