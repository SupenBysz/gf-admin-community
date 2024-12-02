package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

var SysDeliveryCompany = cSysDeliveryCompany{}

type cSysDeliveryCompany struct{}

// GetDeliveryCompanyById 根据id查询物流公司｜信息
func (c *cSysDeliveryCompany) GetDeliveryCompanyById(ctx context.Context, req *sys_api.GetDeliveryCompanyByIdReq) (*sys_model.SysDeliveryCompanyRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.DeliveryCompany.PermissionType.View); has != true {
		return nil, err
	}

	return sys_service.SysDeliveryCompany().GetDeliveryCompanyById(ctx, req.Id)
}

// DeleteDeliveryCompanyById 删除物流公司｜Boolean
func (c *cSysDeliveryCompany) DeleteDeliveryCompanyById(ctx context.Context, req *sys_api.DeleteDeliveryCompanyReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.DeliveryCompany.PermissionType.Delete); has != true {
		return false, err
	}

	return sys_service.SysDeliveryCompany().DeleteDeliveryCompanyById(ctx, req.Id)
}

// CreateDeliveryCompany 添加物流公司｜信息
func (c *cSysDeliveryCompany) CreateDeliveryCompany(ctx context.Context, req *sys_api.CreateDeliveryCompanyReq) (*sys_model.SysDeliveryCompanyRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.DeliveryCompany.PermissionType.Create); has != true {
		return nil, err
	}

	return sys_service.SysDeliveryCompany().SaveDeliveryCompany(ctx, &req.SysDeliveryCompany)
}

// UpdateDeliveryCompany 编辑物流公司｜信息
func (c *cSysDeliveryCompany) UpdateDeliveryCompany(ctx context.Context, req *sys_api.UpdateDeliveryCompanyReq) (*sys_model.SysDeliveryCompanyRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.DeliveryCompany.PermissionType.Update); has != true {
		return nil, err
	}

	return sys_service.SysDeliveryCompany().SaveDeliveryCompany(ctx, &req.SysDeliveryCompany)
}

// QueryDeliveryCompany 查询物流公司｜列表
func (c *cSysDeliveryCompany) QueryDeliveryCompany(ctx context.Context, req *sys_api.QueryDeliveryCompanyReq) (*sys_model.SysDeliveryCompanyListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.DeliveryCompany.PermissionType.List); has != true {
		return nil, err
	}

	return sys_service.SysDeliveryCompany().QueryDeliveryCompany(ctx, &req.SearchParams, req.IsExport)
}
