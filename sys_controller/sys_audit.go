package sys_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysAudit 审核记录
var SysAudit = cSysAudit{}

type cSysAudit struct{}

// GetAuditLogList 获取审核信息|列表
func (c *cSysAudit) GetAuditLogList(ctx context.Context, req *sys_api.QueryAuditListReq) (*sys_model.AuditListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.List); has != true {
		return nil, err
	}

	result, err := sys_service.SysAudit().QueryAuditList(ctx, &req.SearchParams)

	return (*sys_model.AuditListRes)(result), err
}

// SetAuditApprove 审批通过
func (c *cSysAudit) SetAuditApprove(ctx context.Context, req *sys_api.SetAuditApproveReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.Update); has != true {
		return false, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result, err := sys_service.SysAudit().UpdateAudit(ctx, req.Id, sys_enum.Audit.Action.Approve.Code(), "", user.Id)

	return result == true, err
}

// SetAuditReject 审批不通过
func (c *cSysAudit) SetAuditReject(ctx context.Context, req *sys_api.SetAuditRejectReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.Update); has != true {
		return false, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result, err := sys_service.SysAudit().UpdateAudit(ctx, req.Id, sys_enum.Audit.Action.Reject.Code(), req.Reply, user.Id)
	return result == true, err
}

// GetAuditById 根据ID获取资质审核信息
func (c *cSysAudit) GetAuditById(ctx context.Context, req *sys_api.GetAuditByIdReq) (*sys_model.AuditRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result := sys_service.SysAudit().GetAuditById(ctx, req.Id)
	return (*sys_model.AuditRes)(result), nil
}

// GetAuditByDataIdentifier 根据数据标识符获取审核信息
func (c *cSysAudit) GetAuditByDataIdentifier(ctx context.Context, req *sys_api.GetAuditByDataIdentifierReq) (*sys_model.AuditRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result, err := sys_service.SysAudit().GetAuditByDataIdentifier(ctx, req.DataIdentifier, req.UserId, req.UnionMainId)
	return (*sys_model.AuditRes)(result), err
}

// CancelAudit 取消审核申请
func (s *cSysAudit) CancelAudit(ctx context.Context, req *sys_api.CancelAuditReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Audit.PermissionType.Cancel); has != true {
		return false, err
	}

	return sys_service.SysAudit().CancelAudit(ctx, req.Id)
}
