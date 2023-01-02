package sys_permission

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
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
	sys_service.RegisterSysPermission(New())
}

// New sSysPermission 权限控制逻辑实现
func New() *sSysPermission {
	return &sSysPermission{
		CacheDuration: time.Hour,
		CachePrefix:   sys_dao.SysPermission.Table() + "_",
	}
}

// GetPermissionById 根据权限ID获取权限信息
func (s *sSysPermission) GetPermissionById(ctx context.Context, permissionId int64) (*sys_entity.SysPermission, error) {
	result := sys_entity.SysPermission{}

	err := sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Id: permissionId}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return &result, nil
}

// GetPermissionByName 根据权限Name获取权限信息
func (s *sSysPermission) GetPermissionByName(ctx context.Context, permissionName string) (*sys_entity.SysPermission, error) {
	result := sys_entity.SysPermission{}

	err := sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Name: permissionName}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return &result, nil
}

// QueryPermissionList 查询权限
func (s *sSysPermission) QueryPermissionList(ctx context.Context, info sys_model.SearchParams) (*sys_model.SysPermissionInfoListRes, error) {
	result, err := daoctl.Query[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx), &info, false)
	return (*sys_model.SysPermissionInfoListRes)(result), err
}

// GetPermissionList 根据ID获取下级权限信息，返回列表
func (s *sSysPermission) GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) (*[]sys_entity.SysPermission, error) {
	result := make([]sys_entity.SysPermission, 0)
	err := sys_dao.SysPermission.Ctx(ctx).
		// 数据查询结果缓存起来
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     s.CachePrefix + gconv.String(parentId),
			Force:    true,
		}).
		Where(sys_do.SysPermission{
			ParentId: parentId,
			IsShow:   1,
		}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(result) > 0 {
		for _, sysPermissionItem := range result {
			var children *[]sys_entity.SysPermission
			children, err = s.GetPermissionList(ctx, sysPermissionItem.Id, IsRecursive)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
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
func (s *sSysPermission) GetPermissionTree(ctx context.Context, parentId int64) (*[]sys_model.SysPermissionTree, error) {
	result, err := s.GetPermissionList(ctx, parentId, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}

	response := make([]sys_model.SysPermissionTree, 0)

	// 有数据，则递归加载
	if len(*result) > 0 {
		for _, sysPermissionItem := range *result {
			item := sys_model.SysPermissionTree{}
			gconv.Struct(sysPermissionItem, &item)

			item.Children, err = s.GetPermissionTree(ctx, sysPermissionItem.Id)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
			}

			response = append(response, item)
		}
	}
	return &response, nil
}

func (s *sSysPermission) CreatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	return s.SavePermission(ctx, info)
}

func (s *sSysPermission) UpdatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", sys_dao.SysPermission.Table())
	}
	return s.SavePermission(ctx, info)
}

// SavePermission 新增/保存权限信息
func (s *sSysPermission) SavePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	data := sys_entity.SysPermission{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级权限信息是否存在
	if data.ParentId > 0 {
		permissionInfo, err := s.GetPermissionById(ctx, data.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级权限信息不存在"), "", sys_dao.SysPermission.Table())
		}
	}

	if data.Id <= 0 {
		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := sys_dao.SysPermission.Ctx(ctx).Insert(data)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增权限信息失败", sys_dao.SysPermission.Table())
		}
	} else {
		data.UpdatedAt = gtime.Now()
		_, err := sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Id: data.Id}).Update(sys_do.SysPermission{
			ParentId:    data.ParentId,
			Name:        data.Name,
			Description: data.Description,
			Identifier:  data.Identifier,
			IsShow:      1,
			Type:        data.Type,
		})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息保存失败", sys_dao.SysPermission.Table())
		}
	}

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(sys_dao.SysPermission.DB(), s.CachePrefix)
	return &data, nil
}

// DeletePermission 删除权限信息
func (s *sSysPermission) DeletePermission(ctx context.Context, permissionId int64) (bool, error) {
	_, err := s.GetPermissionById(ctx, permissionId)

	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysPermission.Ctx(ctx).Delete(sys_do.SysPermission{Id: permissionId})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除权限信息失败", sys_dao.SysPermission.Table())
	}

	// 删除权限定义
	sys_dao.SysCasbin.Ctx(ctx).Delete(sys_do.SysCasbin{Ptype: "p", V2: permissionId})

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(sys_dao.SysPermission.DB(), s.CachePrefix)
	return true, nil
}

// GetPermissionTreeIdByUrl 根据请求URL去匹配权限树，返回权限
func (s *sSysPermission) GetPermissionTreeIdByUrl(ctx context.Context, path string) (*sys_entity.SysPermission, error) {
	if path == "" {
		return nil, gerror.New("传入的请求url为空")
	}

	result := sys_entity.SysPermission{}

	// 在权限树标识中匹标识后缀，|为标识符的分隔符
	path = "%|" + path

	err := sys_dao.SysPermission.Ctx(ctx).WhereLike(sys_dao.SysPermission.Columns().Identifier, path).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
