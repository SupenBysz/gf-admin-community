package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

var SysMenu = cSysMenu{}

type cSysMenu struct{}

// GetMenuById 根据ID获取菜单信息
func (c *cSysMenu) GetMenuById(ctx context.Context, req *sys_api.GetMenuByIdReq) (*sys_model.SysMenuRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Menu.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	ret, err := sys_service.SysMenu().GetMenuById(ctx, req.Id)

	return (*sys_model.SysMenuRes)(ret), err
}

// CreateMenu 创建菜单
func (c *cSysMenu) CreateMenu(ctx context.Context, req *sys_api.CreateMenuReq) (*sys_model.SysMenuRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Menu.PermissionType.Create); has != true {
		return nil, err
	}

	ret, err := sys_service.SysMenu().CreateMenu(ctx, &req.SysMenu)

	return (*sys_model.SysMenuRes)(ret), err
}

// UpdateMenu 更新菜单
func (c *cSysMenu) UpdateMenu(ctx context.Context, req *sys_api.UpdateMenuReq) (*sys_model.SysMenuRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Menu.PermissionType.Update); has != true {
		return nil, err
	}

	//menu := sys_model.SysMenu{}
	//gconv.Struct(req.UpdateSysMenu, &menu)

	ret, err := sys_service.SysMenu().UpdateMenu(ctx, &req.UpdateSysMenu)

	return (*sys_model.SysMenuRes)(ret), err
}

// DeleteMenu 删除菜单，删除的时候要关联删除sys_permission,有子菜单时禁止删除。
func (c *cSysMenu) DeleteMenu(ctx context.Context, req *sys_api.DeleteMenuReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Menu.PermissionType.Delete); has != true {
		return false, err
	}

	ret, err := sys_service.SysMenu().DeleteMenu(ctx, req.Id)

	return ret == true, err
}

// GetMenuTree 根据ID获取下级菜单信息，返回菜单树
func (c *cSysMenu) GetMenuTree(ctx context.Context, req *sys_api.GetMenuTreeReq) (sys_model.SysMenuTreeListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Menu.PermissionType.Tree); has != true {
		return nil, err
	}

	ret, err := sys_service.SysMenu().GetMenuTree(ctx, req.Id)

	return ret, err
}

// GetMenuList 根据ID获取下级菜单列表，IsRecursive代表是否需要返回下级
//func (c *cSysMenu) GetMenuList(ctx context.Context, req *sys_api.GetMenuListReq) ([]*sys_entity.SysMenu, error) {
//	ret, err := sys_service.SysMenu().GetMenuList(ctx)Á
//	return ret, err
//}
