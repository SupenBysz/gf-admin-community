package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
)

// SysOrganization 组织架
var SysOrganization = cSysOrganization{}

type cSysOrganization struct{}

// QueryOrganizationList 获取组织架构|列表
func (c *cSysOrganization) QueryOrganizationList(ctx context.Context, req *sys_api.QueryOrganizationListReq) (*sys_model.OrganizationInfoListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.List); has != true {
		return nil, err
	}

	return sys_service.SysOrganization().QueryOrganizationList(ctx, req.SearchParams)
}

// GetOrganizationList 根据ID获取下级组织架构|列表
func (c *cSysOrganization) GetOrganizationList(ctx context.Context, req *sys_api.GetOrganizationListReq) (*sys_model.OrganizationInfoListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.List); has != true {
		return nil, err
	}

	data, count, err := sys_service.SysOrganization().GetOrganizationList(ctx, req.Id, req.IsRecursive)
	if err != nil {
		return nil, err
	}

	return &sys_model.OrganizationInfoListRes{
		Records: data,
		PaginationRes: base_model.PaginationRes{
			Pagination: base_model.Pagination{
				PageNum:  1,
				PageSize: count,
			},
			PageTotal: 1,
			Total:     gconv.Int64(count),
		},
	}, err
}

// GetOrganizationTree 根据ID获取下级组织架构|树
func (c *cSysOrganization) GetOrganizationTree(ctx context.Context, req *sys_api.GetOrganizationTreeReq) (*sys_api.OrganizationInfoTreeListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.List); has != true {
		return nil, err
	}

	result, err := sys_service.SysOrganization().GetOrganizationTree(ctx, req.Id)

	return (*sys_api.OrganizationInfoTreeListRes)(&result), err
}

// CreateOrganizationInfo 创建或更新组织架构|信息
func (c *cSysOrganization) CreateOrganizationInfo(ctx context.Context, req *sys_api.CreateOrganizationInfoReq) (*sys_api.OrganizationInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.Create); has != true {
		return nil, err
	}

	result, err := sys_service.SysOrganization().CreateOrganizationInfo(ctx, req.SysOrganizationInfo)
	return (*sys_api.OrganizationInfoRes)(result), err
}

// UpdateOrganizationInfo 创建或更新组织架构|信息
func (c *cSysOrganization) UpdateOrganizationInfo(ctx context.Context, req *sys_api.UpdateOrganizationInfoReq) (*sys_api.OrganizationInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.Update); has != true {
		return nil, err
	}

	result, err := sys_service.SysOrganization().UpdateOrganizationInfo(ctx, req.SysOrganizationInfo)
	return (*sys_api.OrganizationInfoRes)(result), err
}

// GetOrganizationInfo 获取组织架构|信息
func (c *cSysOrganization) GetOrganizationInfo(ctx context.Context, req *sys_api.GetOrganizationInfoReq) (*sys_api.OrganizationInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result, err := sys_service.SysOrganization().GetOrganizationInfo(ctx, req.Id)

	return (*sys_api.OrganizationInfoRes)(result), err
}

// DeleteOrganizationInfo 根据ID删除组织架构
func (c *cSysOrganization) DeleteOrganizationInfo(ctx context.Context, req *sys_api.DeleteOrganizationInfoReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Organization.PermissionType.Delete); has != true {
		return false, err
	}

	result, err := sys_service.SysOrganization().DeleteOrganizationInfo(ctx, req.Id)

	return result == true, err
}
