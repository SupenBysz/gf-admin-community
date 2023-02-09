package sys_menu

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sSysMenu struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysMenu(New())
}

// New sSysMenu 菜单逻辑实现
func New() *sSysMenu {
	return &sSysMenu{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

// GetMenuById 根据ID获取菜单信息
func (s *sSysMenu) GetMenuById(ctx context.Context, menuId int64) (*sys_entity.SysMenu, error) {
	result := sys_entity.SysMenu{}
	err := sys_dao.SysMenu.Ctx(ctx).Hook(daoctl.HookHandler).Scan(&result, sys_do.SysMenu{Id: menuId})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "菜单信息查询失败", sys_dao.SysMenu.Table())
	}
	return &result, err
}

// SaveMenu 新增或保存菜单信息，并自动更新对应的权限信息
func (s *sSysMenu) SaveMenu(ctx context.Context, info sys_model.SysMenu) (*sys_entity.SysMenu, error) {
	data := sys_entity.SysMenu{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级菜单信息是否存在
	if data.ParentId > 0 {
		permissionInfo, err := s.GetMenuById(ctx, data.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "父级菜单信息不存在", sys_dao.SysMenu.Table())
		}
	}

	err := sys_dao.SysMenu.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 保存菜单权限信息
		sysPermission, err := sys_service.SysPermission().SavePermission(ctx, sys_model.SysPermission{
			Id:          data.Id,
			ParentId:    data.ParentId,
			Name:        data.Title,
			Description: data.Description,
		})

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "保存菜单信息失", sys_dao.SysMenu.Table())
		}

		if data.Id <= 0 {
			data.Id = sysPermission.Id
			data.CreatedAt = gtime.Now()
			data.UpdatedAt = gtime.Now()

			_, err = tx.Model(sys_dao.SysMenu).Hook(daoctl.HookHandler).OmitEmptyWhere().Insert(data)

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "新增菜单信息失", sys_dao.SysMenu.Table())
			}

		} else {
			data.UpdatedAt = gtime.Now()
			_, err = sys_dao.SysPermission.Ctx(ctx).
				OmitEmptyWhere().
				Save(data)

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "菜单信息保存失败", sys_dao.SysMenu.Table())
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &data, nil
}
