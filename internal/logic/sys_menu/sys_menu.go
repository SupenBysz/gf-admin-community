package sys_menu

import (
	"context"
	"sort"

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
)

// 设计修改菜单的操作只有超级管理员-1有权限，其他用户只能看看

type sSysMenu struct {
}

func init() {
	sys_service.RegisterSysMenu(New())
}

// New sSysMenu 菜单逻辑实现
func New() sys_service.ISysMenu {
	return &sSysMenu{}
}

// GetMenuById 根据ID获取菜单信息
func (s *sSysMenu) GetMenuById(ctx context.Context, menuId int64) (*sys_entity.SysMenu, error) {
	result := sys_entity.SysMenu{}
	err := sys_dao.SysMenu.Ctx(ctx).Where(sys_do.SysMenu{Id: menuId}).Scan(&result)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_menu_query_failed", sys_dao.SysMenu.Table())
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "error_menu_id_parameter_incorrect"), "", sys_dao.SysMenu.Table())
	}
	data := kconv.Struct(info, &sys_model.SysMenu{})

	return s.SaveMenu(ctx, data)
}

// SaveMenu 新增或保存菜单信息，并自动更新对应的权限信息
func (s *sSysMenu) SaveMenu(ctx context.Context, info *sys_model.SysMenu) (*sys_entity.SysMenu, error) {
	// 初始化数据对象
	data := sys_do.SysMenu{}

	// 验证 info 是否为空
	if info == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeInvalidParameter, "menu info cannot be nil"), "", sys_dao.SysMenu.Table())
	}

	// 转换结构体并捕获错误
	err := gconv.Struct(info, &data)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "failed to convert menu info to database model", sys_dao.SysMenu.Table())
	}

	// 如果父级ID大于0，则校验父级菜单信息是否存在
	if info.ParentId != nil && *info.ParentId > 0 {
		permissionInfo, err := s.GetMenuById(ctx, gconv.Int64(data.ParentId))
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_parent_menu_not_exists", sys_dao.SysMenu.Table())
		}
	}

	err = sys_dao.SysMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if info.Id <= 0 { // 插入
			if data.ParentId == nil {
				*info.ParentId = 0
				data.ParentId = 0
			}
			// 保存菜单权限信息，更新操作不重新插入权限
			data, err = s.createOrUpdatePermission(ctx, info, &data, true)
			if err != nil {
				return err
			}

			// 菜单id = 权限id
			data.CreatedAt = gtime.Now()

			_, err = sys_dao.SysMenu.Ctx(ctx).Insert(&data)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_menu_add_failed", sys_dao.SysMenu.Table())
			}

		} else { // 更新
			data.UpdatedAt = gtime.Now()
			data.Id = nil
			_, err := sys_dao.SysMenu.Ctx(ctx).
				OmitNilData().Where(sys_dao.SysMenu.Columns().Id, info.Id).Update(&data)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_menu_info_save_failed", sys_dao.SysMenu.Table())
			}

			data, err = s.createOrUpdatePermission(ctx, info, &data, false)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetMenuById(ctx, info.Id)
}

// createOrUpdatePermission 创建或更新权限信息
func (s *sSysMenu) createOrUpdatePermission(ctx context.Context, info *sys_model.SysMenu, data *sys_do.SysMenu, isCreate bool) (sys_do.SysMenu, error) {
	var err error
	if isCreate {
		sysPermission, err := sys_service.SysPermission().CreatePermission(ctx, sys_model.SysPermission{
			Id:          info.Id,
			ParentId:    *info.ParentId,
			Name:        *info.Title,
			Type:        2,
			Description: gconv.String(data.Description),
			Identifier:  "Menu::" + *info.Name,
		})
		if err != nil {
			return *data, sys_service.SysLogs().ErrorSimple(ctx, err, "error_menu_info_save_failed", sys_dao.SysMenu.Table())
		}
		data.Id = sysPermission.Id
		info.Id = sysPermission.Id
		data.LimitHiddenRoleIds = "[]"
		data.DepPermissionIds = "[]"
		data.LimitHiddenUserTypes = "[]"
		data.LimitHiddenUserIds = "[]"
	} else {
		permissionInfo := sys_model.UpdateSysPermission{
			Id:          info.Id,
			Name:        info.Title,
			Type:        nil,
			Description: info.Description,
			Identifier:  nil,
		}

		if data.Name != nil {
			identifier := "Menu::" + *info.Name
			permissionInfo.Identifier = &identifier
		}

		// 更新菜单权限信息
		_, err = sys_service.SysPermission().UpdatePermission(ctx, &permissionInfo)
		if err != nil {
			return *data, sys_service.SysLogs().ErrorSimple(ctx, err, "error_menu_permission_update_failed", sys_dao.SysMenu.Table())
		}
	}
	return *data, err
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
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_menu_has_children", sys_dao.SysMenu.Table())
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

	// 有更新的时候，把所有菜单树缓存清除，然后重构

	return true, nil
}

// GetAllMenus 批量获取菜单列表
func (s *sSysMenu) GetAllMenus(ctx context.Context) ([]sys_entity.SysMenu, error) {
	var menus []sys_entity.SysMenu

	err := sys_dao.SysMenu.Ctx(ctx).Order(sys_dao.SysMenu.Columns().Sort).Scan(&menus)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysMenu.Table())
	}
	return menus, nil
}

// GetMenuTree 根据ID获取下级菜单信息，返回菜单树，并缓存
func (s *sSysMenu) GetMenuTree(ctx context.Context, parentId int64, filterIds ...int64) (sys_model.SysMenuTreeListRes, error) {
	menus, err := s.GetAllMenus(ctx)
	if err != nil {
		return nil, err
	}

	// 构建菜单映射
	menuMap := make(map[int64]*sys_model.SysMenuTreeRes)
	for _, menu := range menus {
		menuTree := &sys_model.SysMenuTreeRes{
			SysMenu:  &menu,
			Children: []*sys_model.SysMenuTreeRes{},
		}
		menuMap[menu.Id] = menuTree
	}

	// 过滤菜单
	var filteredMenus []sys_entity.SysMenu
	if len(filterIds) > 0 {
		filterSet := make(map[int64]struct{})
		for _, id := range filterIds {
			filterSet[id] = struct{}{}
		}
		for _, menu := range menus {
			if _, exists := filterSet[menu.Id]; exists {
				filteredMenus = append(filteredMenus, menu)
			}
		}
	} else {
		filteredMenus = menus
	}

	// 构建树结构
	var rootMenus []*sys_model.SysMenuTreeRes
	for _, menu := range filteredMenus {
		if menu.ParentId == parentId {
			rootMenus = append(rootMenus, menuMap[menu.Id])
		} else if parentMenu, exists := menuMap[menu.ParentId]; exists {
			parentMenu.Children = append(parentMenu.Children, menuMap[menu.Id])
		}
	}

	// 过滤掉不在 parentId 下的节点
	var validRootMenus []*sys_model.SysMenuTreeRes
	for _, rootMenu := range rootMenus {
		if rootMenu.SysMenu.ParentId == parentId {
			validRootMenus = append(validRootMenus, rootMenu)
		}
	}

	return validRootMenus, nil
}

// GetMenuList 根据ID获取下级菜单列表，IsRecursive代表是否需要返回下级
func (s *sSysMenu) GetMenuList(ctx context.Context, parentId int64, IsRecursive bool, limitChildrenIds ...int64) ([]*sys_entity.SysMenu, error) {
	dataArr := make([]*sys_entity.SysMenu, 0)

	// 查询数据库并填充数据
	model := sys_dao.SysMenu.Ctx(ctx).Where(sys_dao.SysMenu.Columns().ParentId, parentId)
	if len(limitChildrenIds) > 0 {
		model = model.WhereIn(sys_dao.SysMenu.Columns().Id, limitChildrenIds)
	}

	err := model.Scan(&dataArr)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysMenu.Table())
	}

	// 转换数据结构
	response := make([]*sys_entity.SysMenu, 0)
	queue := make([]*sys_entity.SysMenu, 0)
	queue = append(queue, dataArr...)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		response = append(response, current)

		if IsRecursive {
			children, err := s.GetMenuList(ctx, current.Id, false, limitChildrenIds...)
			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysMenu.Table())
			}
			queue = append(queue, children...)
		}
	}

	// 按排序字段排序
	sort.Slice(response, func(i, j int) bool {
		return response[i].Sort < response[j].Sort
	})

	return response, nil
}

// QueryMenuList 获取菜单列表
