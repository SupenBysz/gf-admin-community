package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysAudit 审核记录
var SysAudit = cSysAudit{}

type cSysAudit struct{}

// GetAuditLogList 获取审核信息|列表
func (c *cSysAudit) GetAuditLogList(ctx context.Context, req *sys_api.GetAuditListReq) (*sys_api.AuditListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.List); has != true {
		return nil, err
	}

	result, err := sys_service.SysAudit().GetAuditList(ctx, req.Category, req.State, &req.Pagination)

	return (*sys_api.AuditListRes)(result), err
}

// SetAuditApprove 审批通过
func (c *cSysAudit) SetAuditApprove(ctx context.Context, req *sys_api.SetAuditApproveReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.Update); has != true {
		return false, err
	}

	result, err := sys_service.SysAudit().UpdateAudit(ctx, req.Id, sys_enum.Audit.Action.Approve.Code(), "")

	return result == true, err
}

// SetAuditReject 审批不通过
func (c *cSysAudit) SetAuditReject(ctx context.Context, req *sys_api.SetAuditRejectReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.Update); has != true {
		return false, err
	}

	result, err := sys_service.SysAudit().UpdateAudit(ctx, req.Id, sys_enum.Audit.Action.Reject.Code(), req.Reply)
	return result == true, err
}

// GetAuditById 根据ID获取资质审核信息
func (c *cSysAudit) GetAuditById(ctx context.Context, req *sys_api.GetAuditByIdReq) (*sys_api.AuditRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result := sys_service.SysAudit().GetAuditById(ctx, req.Id)
	return (*sys_api.AuditRes)(result), nil
}
