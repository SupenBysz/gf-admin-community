package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysCategory 类目管理
var SysCategory = cSysCategory{}

type cSysCategory struct{}

// GetCategoryById 根据ID查下分类
func (c *cSysCategory) GetCategoryById(ctx context.Context, req *sys_api.GetCategoryByIdReq) (*sys_model.SysCategoryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Category.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	return sys_service.SysCategory().GetCategoryById(ctx, req.Id)
}

// CreateCategory 新建分类
func (c *cSysCategory) CreateCategory(ctx context.Context, req *sys_api.CreateCategoryReq) (*sys_model.SysCategoryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Category.PermissionType.Create); has != true {
		return nil, err
	}

	return sys_service.SysCategory().SaveCategory(ctx, &req.SysCategory)
}

// UpdateCategory 保存分类
func (c *cSysCategory) UpdateCategory(ctx context.Context, req *sys_api.UpdateCategoryReq) (*sys_model.SysCategoryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Category.PermissionType.Update); has != true {
		return nil, err
	}

	return sys_service.SysCategory().SaveCategory(ctx, &req.SysCategory)
}

// DeleteCategory 删除分类
func (c *cSysCategory) DeleteCategory(ctx context.Context, req *sys_api.DeleteCategoryReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Category.PermissionType.Delete); has != true {
		return false, err
	}

	return sys_service.SysCategory().DeleteCategory(ctx, req.Id)
}

// QueryCategory 查询分类
func (c *cSysCategory) QueryCategory(ctx context.Context, req *sys_api.QueryCategoryReq) (*sys_model.SysCategoryListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Category.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	return sys_service.SysCategory().QueryCategory(ctx, &req.SearchParams)
}
