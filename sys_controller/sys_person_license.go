package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysLicense 合作伙伴资质
var SysLicense = cSysLicense{}

type cSysLicense struct{}

// GetLicenseById 根据ID获取个人资质|信息
func (c *cSysLicense) GetLicenseById(ctx context.Context, req *sys_api.GetLicenseByIdReq) (*sys_api.LicenseRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.License.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result, err := sys_service.SysPersonLicense().GetLicenseById(ctx, req.Id)
	return (*sys_api.LicenseRes)(result), err
}

// QueryLicenseList 查询个人认证|列表
func (c *cSysLicense) QueryLicenseList(ctx context.Context, req *sys_api.QueryLicenseListReq) (*sys_api.LicenseListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.License.PermissionType.List); has != true {
		return nil, err
	}

	result, err := sys_service.SysPersonLicense().QueryLicenseList(ctx, req.SearchParams)

	if err != nil {
		return nil, err
	}

	return (*sys_api.LicenseListRes)(result), err
}

// CreateLicense 新增个人认证|信息
// func (c *cSysLicense) CreateLicense(ctx context.Context, req *sys_api.CreateLicenseReq) (*sys_api.LicenseRes, error) {
//	result, err := sys_service.SysPersonLicense().CreateLicense(ctx, req.License, req.OperatorId)
//	return (*sys_api.LicenseRes)(result), err
// }

// // UpdateLicense 更新个人资质|信息
// func (c *cSysLicense) UpdateLicense(ctx context.Context, req *sys_api.UpdateLicenseReq) (*sys_api.LicenseRes, error) {
// 	result, err := sys_service.SysPersonLicense().UpdateLicense(ctx, req.License, req.Id)
// 	return (*sys_api.LicenseRes)(result), err
// }

// // SetLicenseState 设置个人信息状态
// func (c *cSysLicense) SetLicenseState(ctx context.Context, req *sys_api.SetLicenseStateReq) (api_sys_api.BoolRes, error) {
//	result, err := sys_service.SysPersonLicense().SetLicenseState(ctx, req.Id, req.State)
//	return result == true, err
// }

// DeleteLicense 软删除个人资质
// func (c *cSysLicense) DeleteLicense(ctx context.Context, req *sys_api.DeleteLicenseReq) (api_sys_api.BoolRes, error) {
// 	result, err := sys_service.SysPersonLicense().DeleteLicense(ctx, req.Id, true)
// 	return result == true, err
// }
