package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	kyAudit "github.com/SupenBysz/gf-admin-community/model/enum/audit"
	"github.com/SupenBysz/gf-admin-community/service"
)

// SysAudit 审核记录
var SysAudit = cSysAudit{}

type cSysAudit struct{}

// GetAuditLogList 獲取审核信息|列表
func (c *cSysAudit) GetAuditLogList(ctx context.Context, req *sysapi.GetAuditListReq) (*sysapi.AuditListRes, error) {
	result, err := service.SysAudit().GetAuditList(ctx, req.Category, req.State, &req.Pagination)

	return (*sysapi.AuditListRes)(result), err
}

// SetAuditApprove 审批通过
func (c *cSysAudit) SetAuditApprove(ctx context.Context, req *sysapi.SetAuditApproveReq) (api_v1.BoolRes, error) {
	result, err := service.SysAudit().UpdateAudit(ctx, req.Id, kyAudit.Approve.Code(), "")
	return result == true, err
}

// SetAuditReject 审批不通过
func (c *cSysAudit) SetAuditReject(ctx context.Context, req *sysapi.SetAuditRejectReq) (api_v1.BoolRes, error) {
	result, err := service.SysAudit().UpdateAudit(ctx, req.Id, kyAudit.Reject.Code(), req.Replay)
	return result == true, err
}

// GetAuditById 根据ID获取资质审核信息
func (c *cSysAudit) GetAuditById(ctx context.Context, req *sysapi.GetAuditByIdReq) (*sysapi.AuditRes, error) {
	result := service.SysAudit().GetAuditById(ctx, req.Id)
	return (*sysapi.AuditRes)(result), nil
}
