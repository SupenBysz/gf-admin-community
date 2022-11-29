package sys_permission

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sSysPermission struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	service.RegisterSysPermission(New())
}

// New sSysPermission 权限控制逻辑实现
func New() *sSysPermission {
	return &sSysPermission{
		CacheDuration: time.Hour,
		CachePrefix:   dao.SysPermission.Table() + "_",
	}
}

// GetPermissionById 根据权限ID获取权限信息
func (s *sSysPermission) GetPermissionById(ctx context.Context, permissionId int64) (*entity.SysPermission, error) {
	result := entity.SysPermission{}

	err := dao.SysPermission.Ctx(ctx).Where(do.SysPermission{Id: permissionId}).Scan(&result)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", dao.SysPermission.Table())
	}

	return &result, nil
}

// GetPermissionByName 根据权限Name获取权限信息
func (s *sSysPermission) GetPermissionByName(ctx context.Context, permissionName string) (*entity.SysPermission, error) {
	result := entity.SysPermission{}

	err := dao.SysPermission.Ctx(ctx).Where(do.SysPermission{Name: permissionName}).Scan(&result)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", dao.SysPermission.Table())
	}

	return &result, nil
}

// QueryPermissionList 查询权限
func (s *sSysPermission) QueryPermissionList(ctx context.Context, info model.SearchFilter) (*sysapi.SysPermissionInfoListRes, error) {
	result, err := daoctl.Query[entity.SysPermission](dao.SysPermission.Ctx(ctx), &info, &info.OrderBy, false)
	return (*sysapi.SysPermissionInfoListRes)(result), err
}

// GetPermissionList 根据ID获取下级权限信息，返回列表
func (s *sSysPermission) GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) (*[]entity.SysPermission, error) {
	result := make([]entity.SysPermission, 0)
	err := dao.SysPermission.Ctx(ctx).
		// 数据查询结果缓存起来
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     s.CachePrefix + gconv.String(parentId),
			Force:    true,
		}).
		Where(do.SysPermission{ParentId: parentId}).Scan(&result)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysPermission.Table())
	}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(result) > 0 {
		for _, sysPermissionItem := range result {
			var children *[]entity.SysPermission
			children, err = s.GetPermissionList(ctx, sysPermissionItem.Id, IsRecursive)

			if err != nil {
				return nil, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysPermission.Table())
			}

			if children == nil || len(*children) <= 0 {
				continue
			}

			for _, sysOrganization := range *children {
				result = append(result, sysOrganization)
			}
		}
	}

	return &result, nil
}

// GetPermissionTree 根据ID获取下级权限信息，返回列表树
func (s *sSysPermission) GetPermissionTree(ctx context.Context, parentId int64) (*[]model.SysPermissionTree, error) {
	result, err := s.GetPermissionList(ctx, parentId, false)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysPermission.Table())
	}

	response := make([]model.SysPermissionTree, 0)

	// 有数据，则递归加载
	if len(*result) > 0 {
		for _, sysPermissionItem := range *result {
			item := model.SysPermissionTree{}
			gconv.Struct(sysPermissionItem, &item)

			item.Children, err = s.GetPermissionTree(ctx, sysPermissionItem.Id)

			if err != nil {
				return nil, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysPermission.Table())
			}

			response = append(response, item)
		}
	}
	return &response, nil
}

func (s *sSysPermission) CreatePermission(ctx context.Context, info model.SysPermission) (*entity.SysPermission, error) {
	return s.SavePermission(ctx, info)
}

func (s *sSysPermission) UpdatePermission(ctx context.Context, info model.SysPermission) (*entity.SysPermission, error) {
	if info.Id <= 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", dao.SysPermission.Table())
	}
	return s.SavePermission(ctx, info)
}

// SavePermission 新增/保存权限信息
func (s *sSysPermission) SavePermission(ctx context.Context, info model.SysPermission) (*entity.SysPermission, error) {
	data := entity.SysPermission{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级权限信息是否存在
	if data.ParentId > 0 {
		permissionInfo, err := s.GetPermissionById(ctx, data.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级权限信息不存在"), "", dao.SysPermission.Table())
		}
	}

	if data.Id <= 0 {
		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := dao.SysPermission.Ctx(ctx).Insert(data)

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "新增权限信息失败", dao.SysPermission.Table())
		}
	} else {
		data.UpdatedAt = gtime.Now()
		_, err := dao.SysPermission.Ctx(ctx).Where(do.SysPermission{Id: data.Id}).Update(do.SysPermission{
			ParentId:    data.ParentId,
			Name:        data.Name,
			Description: data.Description,
			Identifier:  data.Identifier,
			Type:        data.Type,
		})

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "权限信息保存失败", dao.SysPermission.Table())
		}
	}

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(dao.SysPermission.DB(), dao.SysPermission.Table())
	return &data, nil
}

// DeletePermission 删除权限信息
func (s *sSysPermission) DeletePermission(ctx context.Context, permissionId int64) (bool, error) {
	_, err := s.GetPermissionById(ctx, permissionId)

	if err != nil {
		return false, err
	}

	_, err = dao.SysPermission.Ctx(ctx).Delete(do.SysPermission{Id: permissionId})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "删除权限信息失败", dao.SysPermission.Table())
	}

	// 删除权限定义
	dao.SysCasbin.Ctx(ctx).Delete(do.SysCasbin{Ptype: "p", V2: permissionId})

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(dao.SysPermission.DB(), dao.SysPermission.Table())
	return true, nil
}
