package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysPermission 权限管理
var SysPermission = cSysPermission{}

type cSysPermission struct{}

// GetPermissionById 根据权限ID获取权限信|息
func (c *cSysPermission) GetPermissionById(ctx context.Context, req *sys_api.GetPermissionByIdReq) (*sys_model.SysPermissionInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result, err := sys_service.SysPermission().GetPermissionById(ctx, req.Id)
	return (*sys_model.SysPermissionInfoRes)(result), err
}

// GetPermissionByIdentifier 根据权限标识符获取权限|信息
func (c *cSysPermission) GetPermissionByIdentifier(ctx context.Context, req *sys_api.GetPermissionByIdentifierReq) (*sys_model.SysPermissionInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	result, err := sys_service.SysPermission().GetPermissionByIdentifier(ctx, req.Identifier)
	return (*sys_model.SysPermissionInfoRes)(result), err
}

// QueryPermissionListReq 查询权限|列表
func (c *cSysPermission) QueryPermissionListReq(ctx context.Context, req *sys_api.QueryPermissionListReq) (*sys_model.SysPermissionInfoListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.List); has != true {
		return nil, err
	}

	return sys_service.SysPermission().QueryPermissionList(ctx, req.SearchParams)
}

// GetPermissionList 根据ID获取下级权限|列表
// func (c *cSysPermission) GetPermissionList(ctx context.Context, req *sys_api.GetPermissionListReq) (*sys_api.SysPermissionInfoListRes, error) {
//	data, err := sys_service.SysPermission().GetPermissionList(ctx, req.Id, req.IsRecursive)
//	if err != nil {
//		return nil, err
//	}
//
//	count := len(*data)
//
//	return &sys_api.SysPermissionInfoListRes{
//		List: data,
//		PaginationRes: model.PaginationRes{
//			Pagination: model.Pagination{
//				Page:     1,
//				PageSize: count,
//			},
//			PageTotal: 1,
//		},
//	}, err
// }

// GetPermissionTree 根据ID获取下级权限|树
func (c *cSysPermission) GetPermissionTree(ctx context.Context, req *sys_api.GetPermissionTreeReq) (*sys_model.SysPermissionInfoTreeRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.List); has != true {
		return nil, err
	}

	result, err := sys_service.SysPermission().GetPermissionTree(ctx, req.Id, req.Type)

	return (*sys_model.SysPermissionInfoTreeRes)(&result), err

	//return kconv.Struct(result, &sys_model.SysPermissionInfoTreeRes{}), err
}

// CreatePermission 新增权限|信息
func (c *cSysPermission) CreatePermission(ctx context.Context, req *sys_api.CreatePermissionReq) (*sys_model.SysPermissionInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.Create); has != true {
		return nil, err
	}

	result, err := sys_service.SysPermission().CreatePermission(ctx, req.SysPermission)
	return (*sys_model.SysPermissionInfoRes)(result), err
}

// UpdatePermission 更新权限|信息
func (c *cSysPermission) UpdatePermission(ctx context.Context, req *sys_api.UpdatePermissionReq) (*sys_model.SysPermissionInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.Update); has != true {
		return nil, err
	}

	result, err := sys_service.SysPermission().UpdatePermission(ctx, &req.UpdateSysPermission)
	return (*sys_model.SysPermissionInfoRes)(result), err
}

// DeletePermission 删除权限
func (c *cSysPermission) DeletePermission(ctx context.Context, req *sys_api.DeletePermissionReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Permissions.PermissionType.Delete); has != true {
		return false, err
	}

	result, err := sys_service.SysPermission().DeletePermission(ctx, req.Id)
	return result == true, err
}
