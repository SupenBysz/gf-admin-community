package sys_menu

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"sort"
)

// 设计修改菜单的操作只有超级管理员-1有权限，其他用户只能看看

type sSysMenu struct {
}

func init() {
	sys_service.RegisterSysMenu(New())
}

// New sSysMenu 菜单逻辑实现
func New() *sSysMenu {
	return &sSysMenu{}
}

// GetMenuById 根据ID获取菜单信息
func (s *sSysMenu) GetMenuById(ctx context.Context, menuId int64) (*sys_entity.SysMenu, error) {
	result := sys_entity.SysMenu{}
	err := sys_dao.SysMenu.Ctx(ctx).Where(sys_do.SysMenu{Id: menuId}).Scan(&result)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "菜单信息查询失败", sys_dao.SysMenu.Table())
	}
	return &result, err
}

// CreateMenu 创建菜单
func (s *sSysMenu) CreateMenu(ctx context.Context, info *sys_model.SysMenu) (*sys_entity.SysMenu, error) {
	return s.SaveMenu(ctx, info)
}

// UpdateMenu 更新菜单
func (s *sSysMenu) UpdateMenu(ctx context.Context, info *sys_model.UpdateSysMenu) (*sys_entity.SysMenu, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", sys_dao.SysMenu.Table())
	}
	data := kconv.Struct(info, &sys_model.SysMenu{})

	return s.SaveMenu(ctx, data)
}

// SaveMenu 新增或保存菜单信息，并自动更新对应的权限信息
func (s *sSysMenu) SaveMenu(ctx context.Context, info *sys_model.SysMenu) (*sys_entity.SysMenu, error) {
	data := sys_do.SysMenu{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级菜单信息是否存在
	if info.ParentId != nil && *info.ParentId > 0 {
		permissionInfo, err := s.GetMenuById(ctx, gconv.Int64(data.ParentId))
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "父级菜单信息不存在", sys_dao.SysMenu.Table())
		}
	}

	err := sys_dao.SysMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if info.Id <= 0 { // 插入
			if data.ParentId == nil {
				*info.ParentId = 0
				data.ParentId = 0
			}
			// 保存菜单权限信息，更新操作不重新插入权限
			sysPermission, err := sys_service.SysPermission().CreatePermission(ctx, sys_model.SysPermission{
				Id:          info.Id,
				ParentId:    *info.ParentId,
				Name:        *info.Title,
				Type:        2,
				Description: gconv.String(data.Description),
				Identifier:  "Menu::" + *info.Name,
			})

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "保存菜单信息失败", sys_dao.SysMenu.Table())
			}

			// 菜单id = 权限id
			data.Id = sysPermission.Id
			info.Id = sysPermission.Id
			data.CreatedAt = gtime.Now()

			_, err = sys_dao.SysMenu.Ctx(ctx).Insert(&data)

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "新增菜单信息失败", sys_dao.SysMenu.Table())
			}

		} else { // 更新
			data.UpdatedAt = gtime.Now()
			data.Id = nil
			_, err := sys_dao.SysMenu.Ctx(ctx).
				OmitNilData().Where(sys_dao.SysMenu.Columns().Id, info.Id).Update(&data)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "菜单信息保存失败", sys_dao.SysMenu.Table())
			}

			permisionInfo := sys_model.UpdateSysPermission{
				Id:          info.Id,
				Name:        info.Title,
				Type:        nil,
				Description: info.Description,
				Identifier:  nil,
			}

			if data.Name != nil {
				identifier := "Menu::" + *info.Name
				permisionInfo.Identifier = &identifier
			}

			// 更新菜单权限信息
			_, err = sys_service.SysPermission().UpdatePermission(ctx, &permisionInfo)

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "菜单权限更新失败", sys_dao.SysMenu.Table())
			}

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	//go s.makeMenuTeee(ctx, 0)
	return s.GetMenuById(ctx, info.Id)
}

// DeleteMenu 删除菜单，删除的时候要关联删除sys_permission,有子菜单时禁止删除。
func (s *sSysMenu) DeleteMenu(ctx context.Context, id int64) (bool, error) {
	_, err := s.GetMenuById(ctx, id)

	if err != nil {
		return false, err
	}

	// 判断是否具备子菜单
	count, _ := sys_dao.SysMenu.Ctx(ctx).Where(sys_dao.SysMenu.Columns().ParentId, id).Count()
	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "该菜单具备子菜单，请先移除子菜单再进行操作", sys_dao.SysMenu.Table())
	}

	err = sys_dao.SysMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除关联sys_permission权限记录 		菜单id = 权限id
		ret, err := sys_service.SysPermission().DeletePermission(ctx, id)
		if err != nil || ret == false {
			return err
		}

		// 删除菜单记录
		_, err = daoctl.DeleteWithError(sys_dao.SysMenu.Ctx(ctx).Where(sys_dao.SysMenu.Columns().Id, id))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	//g.DB().GetCache().Remove(ctx, "getMenuTreeCacheById_"+gconv.String(id))

	//daoctl.RemoveQueryCache(sys_dao.SysMenu.DB(), sys_dao.SysMenu.Table())

	//go s.MakeMenuTree(ctx, 0)

	// 有更新的时候，把所有菜单树缓存清除，然后重构

	return true, nil
}

// MakeMenuTree 构建菜单树
func (s *sSysMenu) MakeMenuTree(ctx context.Context, parentId int64, isMakeNodeFun func(ctx context.Context, cruuentMenu *sys_entity.SysMenu) bool) ([]*sys_model.SysMenuTreeRes, error) {
	// 当前级菜单树没加入

	// 获取下级菜单列表
	result, err := s.GetMenuList(ctx, parentId, false)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询菜单失败", sys_dao.SysPermission.Table())
	}

	response := make([]*sys_model.SysMenuTreeRes, 0)

	// 有数据，则递归加载
	if len(result) > 0 {
		// 我第一次进来的样子

		// 我进来是否具备上级
		for _, sysMenuItem := range result {
			item := &sys_model.SysMenuTreeRes{}
			gconv.Struct(sysMenuItem, &item)

			tree, err := s.MakeMenuTree(ctx, sysMenuItem.Id, isMakeNodeFun)

			for _, childItem := range tree {
				if isMakeNodeFun(ctx, childItem.SysMenu) || len(childItem.Children) > 0 {
					item.Children = append(item.Children, childItem)
				}
			}

			if len(tree) == 0 && !isMakeNodeFun(ctx, sysMenuItem) && len(item.Children) == 0 {
				continue
			}

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询菜单失败", sys_dao.SysMenu.Table())
			}

			response = append(response, item)

		}
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询菜单失败", sys_dao.SysMenu.Table())
	}

	return response, nil
}

// GetMenuTree 根据ID获取下级菜单信息，返回菜单树，并缓存
func (s *sSysMenu) GetMenuTree(ctx context.Context, parentId int64) (sys_model.SysMenuTreeListRes, error) {
	// 先判断缓存中是否存在菜单树，存在直接返回
	//res, _ := g.DB().GetCache().Get(ctx, "MenuTreeCache_"+gconv.String(parentId))

	//if res.Val() != nil {
	//	data := make([]*sys_model.SysMenuTreeRes, 0)
	//	gconv.Struct(res, &data)
	//	return data, nil
	//}

	// 将权限树缓存起来,
	// 注意：这种自定义缓存方式，Ctx内部的缓存机制不会自动帮我们删除该缓存，所以，需要每次有更新操作，手动删除缓存getMenuTreeCacheById_的缓存
	tree, err := s.MakeMenuTree(ctx, parentId, func(ctx context.Context, cruuentMenu *sys_entity.SysMenu) bool {
		return true
	})
	if err != nil {
		return nil, err
	}

	//g.DB().GetCache().Set(ctx, "MenuTreeCache_"+gconv.String(parentId), tree, 0)

	return tree, nil
}

// GetMenuList 根据ID获取下级菜单列表，IsRecursive代表是否需要返回下级
func (s *sSysMenu) GetMenuList(ctx context.Context, parentId int64, IsRecursive bool, limitChildrenIds ...int64) ([]*sys_entity.SysMenu, error) {
	dataArr := make([]*sys_entity.SysMenu, 0)

	response := make([]*sys_entity.SysMenu, 0)
	model := sys_dao.SysMenu.Ctx(ctx).Where(sys_dao.SysMenu.Columns().ParentId, parentId)

	if len(limitChildrenIds) > 0 {
		model = model.WhereIn(sys_dao.SysMenu.Columns().Id, limitChildrenIds)
	}

	model.Scan(&dataArr)
	gconv.Struct(dataArr, &response)

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(dataArr) > 0 {
		for _, sysMenu := range dataArr {
			var children []*sys_entity.SysMenu
			children, err := s.GetMenuList(ctx, sysMenu.Id, IsRecursive, limitChildrenIds...)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysMenu.Table())
			}

			if children == nil || len(children) <= 0 {
				continue
			}

			for _, menu := range children {
				response = append(response, menu)
			}
		}
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Sort < response[j].Sort
	})

	return response, nil
}

// QueryMenuList 获取菜单列表
