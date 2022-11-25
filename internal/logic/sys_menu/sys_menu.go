package sys_menu

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"time"
)

type sSysMenu struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	service.RegisterSysMenu(New())
}

// New sSysMenu 菜单逻辑实现
func New() *sSysMenu {
	return &sSysMenu{
		CacheDuration: time.Hour,
		CachePrefix:   dao.SysMenu.Table() + "_",
	}
}

// GetMenuById 根据ID获取菜单信息
func (s *sSysMenu) GetMenuById(ctx context.Context, menuId int64) (*entity.SysMenu, error) {
	result := entity.SysMenu{}
	err := dao.SysMenu.Ctx(ctx).Scan(&result, do.SysMenu{Id: menuId})
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "菜单信息查询失败", dao.SysMenu.Table())
	}
	return &result, err
}

// SaveMenu 新增或保存菜单信息，并自动更新对应的权限信息
func (s *sSysMenu) SaveMenu(ctx context.Context, info model.SysMenu) (*entity.SysMenu, error) {
	data := entity.SysMenu{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级菜单信息是否存在
	if data.ParentId > 0 {
		permissionInfo, err := s.GetMenuById(ctx, data.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "父级菜单信息不存在", dao.SysMenu.Table())
		}
	}

	err := dao.SysMenu.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 保存菜单权限信息
		sysPermission, err := service.SysPermission().SavePermission(ctx, model.SysPermission{
			Id:          data.Id,
			ParentId:    data.ParentId,
			Name:        data.Title,
			Description: data.Description,
		})

		if err != nil {
			return service.SysLogs().ErrorSimple(ctx, err, "保存菜单信息失", dao.SysMenu.Table())
		}

		if data.Id <= 0 {
			data.Id = sysPermission.Id
			data.CreatedAt = gtime.Now()
			data.UpdatedAt = gtime.Now()

			_, err = tx.Model(dao.SysMenu).OmitEmptyWhere().Insert(data)

			if err != nil {
				return service.SysLogs().ErrorSimple(ctx, err, "新增菜单信息失", dao.SysMenu.Table())
			}

		} else {
			data.UpdatedAt = gtime.Now()
			_, err = dao.SysPermission.Ctx(ctx).
				OmitEmptyWhere().
				Save(data)

			if err != nil {
				return service.SysLogs().ErrorSimple(ctx, err, "菜单信息保存失败", dao.SysMenu.Table())
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &data, nil
}
