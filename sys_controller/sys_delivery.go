package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

var SysDelivery = cSysDelivery{}

type cSysDelivery struct{}

// GetDeliveryById 根据id查询物流公司｜信息
func (c *cSysDelivery) GetDeliveryById(ctx context.Context, req *sys_api.GetDeliveryByIdReq) (*sys_model.SysDeliveryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Delivery.PermissionType.View); has != true {
		return nil, err
	}

	return sys_service.SysDelivery().GetDeliveryById(ctx, req.Id)
}

// DeleteDeliveryById 删除物流公司｜Boolean
func (c *cSysDelivery) DeleteDeliveryById(ctx context.Context, req *sys_api.DeleteDeliveryReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Delivery.PermissionType.Delete); has != true {
		return false, err
	}

	return sys_service.SysDelivery().DeleteDeliveryById(ctx, req.Id)
}

// CreateDelivery 添加物流公司｜信息
func (c *cSysDelivery) CreateDelivery(ctx context.Context, req *sys_api.CreateDeliveryReq) (*sys_model.SysDeliveryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Delivery.PermissionType.Create); has != true {
		return nil, err
	}

	return sys_service.SysDelivery().SaveDelivery(ctx, &req.SysDelivery)
}

// UpdateDelivery 编辑物流公司｜信息
func (c *cSysDelivery) UpdateDelivery(ctx context.Context, req *sys_api.UpdateDeliveryReq) (*sys_model.SysDeliveryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Delivery.PermissionType.Update); has != true {
		return nil, err
	}

	return sys_service.SysDelivery().SaveDelivery(ctx, &req.SysDelivery)
}

// QueryDelivery 查询物流公司｜列表
func (c *cSysDelivery) QueryDelivery(ctx context.Context, req *sys_api.QueryDeliveryReq) (*sys_model.SysDeliveryListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Delivery.PermissionType.List); has != true {
		return nil, err
	}

	return sys_service.SysDelivery().QueryDelivery(ctx, &req.SearchParams, req.IsExport)
}
