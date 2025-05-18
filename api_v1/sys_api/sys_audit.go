package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type QueryAuditListReq struct {
	g.Meta `path:"/queryAuditList" method:"post" summary:"获取审核信息|列表" tags:"审核管理"` // 券 个人资质 主体等业务类型，根据下面过滤字段进行筛选
	base_model.SearchParams

	//base_model.Pagination
	//State int `json:"state" v:"in:-1,0,1#状态参数错误" dc:"过滤审核状态支持：-1不通过，0待审核，1通过" default:"0"`
	//Category int `json:"category" v:"in:0,1,2,4#业务类型参数错误" dc:"业务类型支持：1运营商主体资质审核，2服务商主体资质审核、4消费者实名审核，0所有类型" default:"0"`
}

type GetAuditByIdReq struct { //
	g.Meta `path:"/getAuditById" method:"post" summary:"根据审核ID获取审核|信息" tags:"审核管理"`
	Id     int64 `json:"id" dc:"审核ID"`
}

type GetAuditByDataIdentifierReq struct {
	g.Meta         `path:"/getAuditByDataIdentifier" method:"post" summary:"根据数据标识符获取审核|信息" tags:"审核管理"`
	DataIdentifier string `json:"dataIdentifier" v:"required#数据标识符不能为空" dc:"数据标识符"`
	UserId         int64  `json:"userId" dc:"用户ID" default:"0"`
	UnionMainId    int64  `json:"unionMainId" dc:"业务主体ID" default:"0"`
}

type CancelAuditReq struct { //
	g.Meta `path:"/cancelAudit" method:"post" summary:"撤销审核申请|Boolean" tags:"审核管理"`
	Id     int64 `json:"id" dc:"审核ID"`
}

type SetAuditApproveReq struct {
	g.Meta `path:"/setAuditApprove" method:"post" summary:"审批通过" tags:"审核管理"`
	Id     int64 `json:"id" v:"required|min:1#ID参数错误|ID必须大于0" dc:"审核ID"`
}

type SetAuditRejectReq struct {
	g.Meta `path:"/setAuditReject" method:"post" summary:"审批不通过" tags:"审核管理"`
	Id     int64  `json:"id" v:"required|min:1#ID参数错误|ID必须大于0" dc:"审核ID"`
	Reply  string `json:"reply" v:"required#请输入不通过原因" dc:"不通过原因"`
}
