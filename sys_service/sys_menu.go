// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type (
	ISysMenu interface {
		// GetMenuById 根据ID获取菜单信息
		GetMenuById(ctx context.Context, menuId int64) (*sys_entity.SysMenu, error)
		// CreateMenu 创建菜单
		CreateMenu(ctx context.Context, info *sys_model.SysMenu) (*sys_entity.SysMenu, error)
		// UpdateMenu 更新菜单
		UpdateMenu(ctx context.Context, info *sys_model.SysMenu) (*sys_entity.SysMenu, error)
		// SaveMenu 新增或保存菜单信息，并自动更新对应的权限信息
		SaveMenu(ctx context.Context, info *sys_model.SysMenu) (*sys_entity.SysMenu, error)
		// DeleteMenu 删除菜单，删除的时候要关联删除sys_permission,有子菜单时禁止删除。
		DeleteMenu(ctx context.Context, id int64) (bool, error)
		// MakeMenuTree 构建菜单树
		MakeMenuTree(ctx context.Context, parentId int64, isMakeNodeFun func(ctx context.Context, cruuentMenu *sys_entity.SysMenu) bool) ([]*sys_model.SysMenuTreeRes, error)
		// GetMenuTree 根据ID获取下级菜单信息，返回菜单树，并缓存
		GetMenuTree(ctx context.Context, parentId int64) ([]*sys_model.SysMenuTreeRes, error)
		// GetMenuList 根据ID获取下级菜单列表，IsRecursive代表是否需要返回下级
		GetMenuList(ctx context.Context, parentId int64, IsRecursive bool, limitChildrenIds ...int64) ([]*sys_entity.SysMenu, error)
	}
)

var (
	localSysMenu ISysMenu
)

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}
