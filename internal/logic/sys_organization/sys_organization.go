package sys_organization

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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sSysOrganization struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	service.RegisterSysOrganization(New())
}

// New sSysOrganization 权限控制逻辑实现
func New() *sSysOrganization {
	return &sSysOrganization{
		CacheDuration: time.Hour,
		CachePrefix:   dao.SysOrganization.Table() + "_",
	}
}

// QueryOrganizationList 获取组织架构信息列表
func (s *sSysOrganization) QueryOrganizationList(ctx context.Context, info model.SearchFilter) (*sysapi.OrganizationInfoListRes, error) {
	result, err := daoctl.Query[entity.SysOrganization](dao.SysOrganization.Ctx(ctx), &info, &info.OrderBy, false)

	return (*sysapi.OrganizationInfoListRes)(result), err
}

// GetOrganizationList 获取组织架构信息列表
func (s *sSysOrganization) GetOrganizationList(ctx context.Context, parentId int64, IsRecursive bool) (*[]entity.SysOrganization, int, error) {
	result := make([]entity.SysOrganization, 0)
	err := dao.SysOrganization.Ctx(ctx).
		// 数据查询结果缓存起来
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     s.CachePrefix + gconv.String(parentId),
			Force:    true,
		}).
		Where(do.SysOrganization{ParentId: parentId}).Scan(&result)

	if err != nil {
		return nil, -1, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysOrganization.Table())
	}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(result) > 0 {
		for _, organization := range result {
			var children *[]entity.SysOrganization
			children, count, err := s.GetOrganizationList(ctx, organization.Id, IsRecursive)

			if err != nil {
				return nil, count, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysOrganization.Table())
			}

			if children == nil || len(*children) <= 0 {
				continue
			}

			for _, sSysOrganization := range *children {
				result = append(result, sSysOrganization)
			}
		}
	}

	return &result, len(result), nil
}

// GetOrganizationTree 获取组织架构信息树
func (s *sSysOrganization) GetOrganizationTree(ctx context.Context, parentId int64) (*[]model.SysOrganizationTree, error) {
	result, _, err := s.GetOrganizationList(ctx, parentId, false)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysOrganization.Table())
	}

	response := make([]model.SysOrganizationTree, 0)

	// 有数据，则递归加载
	if len(*result) > 0 {
		for _, organization := range *result {
			item := model.SysOrganizationTree{}
			gconv.Struct(organization, &item)

			item.Children, err = s.GetOrganizationTree(ctx, organization.Id)

			if err != nil {
				return nil, service.SysLogs().ErrorSimple(ctx, err, "查询失败", dao.SysOrganization.Table())
			}

			response = append(response, item)
		}
	}
	return &response, nil
}

// CreateOrganizationInfo 创建组织架构信息
func (s *sSysOrganization) CreateOrganizationInfo(ctx context.Context, info model.SysOrganizationInfo) (*entity.SysOrganization, error) {
	return s.SaveOrganizationInfo(ctx, info)
}

// UpdateOrganizationInfo 更新组织架构信息
func (s *sSysOrganization) UpdateOrganizationInfo(ctx context.Context, info model.SysOrganizationInfo) (*entity.SysOrganization, error) {
	if info.Id <= 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", dao.SysOrganization.Table())
	}
	return s.SaveOrganizationInfo(ctx, info)
}

// SaveOrganizationInfo 创建或更新组织架构信息
func (s *sSysOrganization) SaveOrganizationInfo(ctx context.Context, info model.SysOrganizationInfo) (*entity.SysOrganization, error) {
	parentInfo := entity.SysOrganization{}
	cascadeDeep := 1
	if info.ParentId > 0 {
		result, err := dao.SysOrganization.Ctx(ctx).
			One(do.SysOrganization{Id: info.ParentId})

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "父级组织机构信息查询失败，请稍后再试", dao.SysOrganization.Table())
		}

		if result.IsEmpty() {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级组织机构信息查询失败，请稍后再试"), "", dao.SysOrganization.Table())
		}

		result.Struct(&parentInfo)
		cascadeDeep = parentInfo.CascadeDeep * 2
	} else {
		info.ParentId = 0
		cascadeDeep = 1
	}

	if info.Id <= 0 {
		result, err := dao.SysOrganization.Ctx(ctx).
			One(do.SysOrganization{ParentId: info.ParentId, Name: info.Name})

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "添加组织机构信息失败", dao.SysOrganization.Table())
		}

		if !result.IsEmpty() {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "添加组织机构信息失败，请更换其它名称"), "", dao.SysOrganization.Table())
		}

		orgInfo := entity.SysOrganization{}
		result.Struct(&orgInfo)

		info.Id = idgen.NextId()
		_, err = dao.SysOrganization.Ctx(ctx).Insert(do.SysOrganization{
			Id:          info.Id,
			Name:        info.Name,
			ParentId:    info.ParentId,
			CascadeDeep: cascadeDeep,
			Description: info.Description,
		})

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "添加组织机构信息失败。"), "", dao.SysOrganization.Table())
		}

		// 移除已缓存的数据
		daoctl.RemoveQueryCache(dao.SysOrganization.DB(), dao.SysOrganization.Table())
		return &entity.SysOrganization{
			Id:          info.Id,
			Name:        info.Name,
			ParentId:    info.ParentId,
			CascadeDeep: cascadeDeep,
			Description: info.Description,
		}, nil
	} else {
		result, err := dao.SysOrganization.Ctx(ctx).
			Where(do.SysOrganization{ParentId: info.Id, Name: info.Name}).
			WhereNot(dao.SysOrganization.Columns().Id, info.Id).
			One()

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "更新组织机构信息失败", dao.SysOrganization.Table())
		}

		if !result.IsEmpty() {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "更新组织机构信息失败，请更换其它名称"), "", dao.SysOrganization.Table())
		}

		result, err = dao.SysOrganization.Ctx(ctx).Where(do.SysOrganization{Id: info.Id}).One()

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "更新组织机构信息失败", dao.SysOrganization.Table())
		}

		if result.IsEmpty() {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "更新组织机构信息失败，组织机构信息不存在"), "", dao.SysOrganization.Table())
		}

		oldInfo := entity.SysOrganization{}
		result.Struct(&oldInfo)

		info.ParentId = oldInfo.ParentId

		_, err = dao.SysOrganization.Ctx(ctx).
			Where(do.SysOrganization{Id: info.Id}).
			Update(do.SysOrganization{Name: info.Name, CascadeDeep: oldInfo.CascadeDeep, Description: info.Description})

		if err != nil {
			return nil, service.SysLogs().ErrorSimple(ctx, err, "更新组织机构信息失败", dao.SysOrganization.Table())
		}

		// 移除已缓存的数据
		daoctl.RemoveQueryCache(dao.SysOrganization.DB(), dao.SysOrganization.Table())
		return &entity.SysOrganization{
			Id:          info.Id,
			Name:        info.Name,
			ParentId:    info.ParentId,
			CascadeDeep: oldInfo.CascadeDeep,
			Description: info.Description,
		}, nil
	}
}

// GetOrganizationInfo 获取组织架构信息
func (s *sSysOrganization) GetOrganizationInfo(ctx context.Context, id int64) (*entity.SysOrganization, error) {
	result, err := dao.SysOrganization.Ctx(ctx).
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     "Cache-Organization",
			Force:    false,
		}).
		Where(do.SysOrganization{Id: id}).One()

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "查询组织机构信息失败", dao.SysOrganization.Table())
	}

	if result.IsEmpty() {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "组织机构信息不存在"), "", dao.SysOrganization.Table())
	}

	info := entity.SysOrganization{}
	result.Struct(&info)

	return &info, nil
}

// DeleteOrganizationInfo 删除组织架构信息
func (s *sSysOrganization) DeleteOrganizationInfo(ctx context.Context, id int64) (bool, error) {

	count, err := dao.SysOrganization.Ctx(ctx).Where(do.SysOrganization{ParentId: id}).Count()

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "删除组织机构信息失败", dao.SysOrganization.Table())
	}

	if count > 0 {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "存下级组织架构时禁止删除！"), "", dao.SysOrganization.Table())
	}

	info := entity.SysOrganization{}
	err = dao.SysOrganization.Ctx(ctx).
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     "Cache-Organization",
			Force:    false,
		}).
		Where(do.SysOrganization{Id: id}).Scan(&info)

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "查询组织机构信息失败，请稍后再试", dao.SysOrganization.Table())
	}

	if info.Id <= 0 {
		return true, nil
	}

	_, err = dao.SysOrganization.Ctx(ctx).
		Cache(gdb.CacheOption{
			Duration: -1,
			Name:     "Cache-Organization",
			Force:    false,
		}).
		Delete(do.SysOrganization{Id: id})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "删除组织机构信息失败", dao.SysOrganization.Table())
	}

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(dao.SysOrganization.DB(), dao.SysOrganization.Table())
	return true, nil
}
