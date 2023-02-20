package sys_organization

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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sSysOrganization struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysOrganization(New())
}

// New sSysOrganization 权限控制逻辑实现
func New() *sSysOrganization {
	return &sSysOrganization{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

// QueryOrganizationList 获取组织架构信息列表
func (s *sSysOrganization) QueryOrganizationList(ctx context.Context, info base_model.SearchParams) (*sys_model.OrganizationInfoListRes, error) {
	result, err := daoctl.Query[*sys_entity.SysOrganization](sys_dao.SysOrganization.Ctx(ctx), &info, false)

	return (*sys_model.OrganizationInfoListRes)(result), err
}

// GetOrganizationList 获取组织架构信息列表
func (s *sSysOrganization) GetOrganizationList(ctx context.Context, parentId int64, IsRecursive bool) ([]*sys_entity.SysOrganization, int, error) {
	result := make([]*sys_entity.SysOrganization, 0)
	err := sys_dao.SysOrganization.Ctx(ctx).Where(sys_do.SysOrganization{ParentId: parentId}).Scan(&result)

	if err != nil {
		return nil, -1, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysOrganization.Table())
	}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(result) > 0 {
		for _, organization := range result {
			var children []*sys_entity.SysOrganization
			children, count, err := s.GetOrganizationList(ctx, organization.Id, IsRecursive)

			if err != nil {
				return nil, count, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysOrganization.Table())
			}

			if children == nil || len(children) <= 0 {
				continue
			}

			for _, sSysOrganization := range children {
				result = append(result, sSysOrganization)
			}
		}
	}

	return result, len(result), nil
}

// GetOrganizationTree 获取组织架构信息树
func (s *sSysOrganization) GetOrganizationTree(ctx context.Context, parentId int64) ([]*sys_model.SysOrganizationTree, error) {
	result, _, err := s.GetOrganizationList(ctx, parentId, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysOrganization.Table())
	}

	response := make([]*sys_model.SysOrganizationTree, 0)

	// 有数据，则递归加载
	if len(result) > 0 {
		for _, organization := range result {
			item := &sys_model.SysOrganizationTree{}
			gconv.Struct(organization, &item)

			item.Children, err = s.GetOrganizationTree(ctx, organization.Id)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysOrganization.Table())
			}

			response = append(response, item)
		}
	}
	return response, nil
}

// CreateOrganizationInfo 创建组织架构信息
func (s *sSysOrganization) CreateOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error) {
	return s.SaveOrganizationInfo(ctx, info)
}

// UpdateOrganizationInfo 更新组织架构信息
func (s *sSysOrganization) UpdateOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", sys_dao.SysOrganization.Table())
	}
	return s.SaveOrganizationInfo(ctx, info)
}

// SaveOrganizationInfo 创建或更新组织架构信息
func (s *sSysOrganization) SaveOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error) {
	parentInfo := sys_entity.SysOrganization{}
	cascadeDeep := 1
	if info.ParentId > 0 {
		result, err := sys_dao.SysOrganization.Ctx(ctx).
			One(sys_do.SysOrganization{Id: info.ParentId})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "父级组织机构信息查询失败，请稍后再试", sys_dao.SysOrganization.Table())
		}

		if result.IsEmpty() {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级组织机构信息查询失败，请稍后再试"), "", sys_dao.SysOrganization.Table())
		}

		result.Struct(&parentInfo)
		cascadeDeep = parentInfo.CascadeDeep * 2
	} else {
		info.ParentId = 0
		cascadeDeep = 1
	}

	if info.Id <= 0 {
		result, err := sys_dao.SysOrganization.Ctx(ctx).
			One(sys_do.SysOrganization{ParentId: info.ParentId, Name: info.Name})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "添加组织机构信息失败", sys_dao.SysOrganization.Table())
		}

		if !result.IsEmpty() {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "添加组织机构信息失败，请更换其它名称"), "", sys_dao.SysOrganization.Table())
		}

		orgInfo := sys_entity.SysOrganization{}
		result.Struct(&orgInfo)

		info.Id = idgen.NextId()
		_, err = sys_dao.SysOrganization.Ctx(ctx).Insert(sys_do.SysOrganization{
			Id:          info.Id,
			Name:        info.Name,
			ParentId:    info.ParentId,
			CascadeDeep: cascadeDeep,
			Description: info.Description,
		})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "添加组织机构信息失败。"), "", sys_dao.SysOrganization.Table())
		}

		// 移除已缓存的数据
		daoctl.RemoveQueryCache(sys_dao.SysOrganization.DB(), sys_dao.SysOrganization.Table())
		return &sys_entity.SysOrganization{
			Id:          info.Id,
			Name:        info.Name,
			ParentId:    info.ParentId,
			CascadeDeep: cascadeDeep,
			Description: info.Description,
		}, nil
	} else {
		result, err := sys_dao.SysOrganization.Ctx(ctx).
			Where(sys_do.SysOrganization{ParentId: info.Id, Name: info.Name}).
			WhereNot(sys_dao.SysOrganization.Columns().Id, info.Id).
			One()

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "更新组织机构信息失败", sys_dao.SysOrganization.Table())
		}

		if !result.IsEmpty() {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "更新组织机构信息失败，请更换其它名称"), "", sys_dao.SysOrganization.Table())
		}

		result, err = sys_dao.SysOrganization.Ctx(ctx).Where(sys_do.SysOrganization{Id: info.Id}).One()

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "更新组织机构信息失败", sys_dao.SysOrganization.Table())
		}

		if result.IsEmpty() {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "更新组织机构信息失败，组织机构信息不存在"), "", sys_dao.SysOrganization.Table())
		}

		oldInfo := sys_entity.SysOrganization{}
		result.Struct(&oldInfo)

		info.ParentId = oldInfo.ParentId

		_, err = sys_dao.SysOrganization.Ctx(ctx).Where(sys_do.SysOrganization{Id: info.Id}).
			Update(sys_do.SysOrganization{Name: info.Name, CascadeDeep: oldInfo.CascadeDeep, Description: info.Description})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "更新组织机构信息失败", sys_dao.SysOrganization.Table())
		}

		// 移除已缓存的数据
		daoctl.RemoveQueryCache(sys_dao.SysOrganization.DB(), sys_dao.SysOrganization.Table())
		return &sys_entity.SysOrganization{
			Id:          info.Id,
			Name:        info.Name,
			ParentId:    info.ParentId,
			CascadeDeep: oldInfo.CascadeDeep,
			Description: info.Description,
		}, nil
	}
}

// GetOrganizationInfo 获取组织架构信息
func (s *sSysOrganization) GetOrganizationInfo(ctx context.Context, id int64) (*sys_entity.SysOrganization, error) {
	result, err := sys_dao.SysOrganization.Ctx(ctx).Where(sys_do.SysOrganization{Id: id}).One()

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询组织机构信息失败", sys_dao.SysOrganization.Table())
	}

	if result.IsEmpty() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "组织机构信息不存在"), "", sys_dao.SysOrganization.Table())
	}

	info := sys_entity.SysOrganization{}
	result.Struct(&info)

	return &info, nil
}

// DeleteOrganizationInfo 删除组织架构信息
func (s *sSysOrganization) DeleteOrganizationInfo(ctx context.Context, id int64) (bool, error) {

	count, err := sys_dao.SysOrganization.Ctx(ctx).Where(sys_do.SysOrganization{ParentId: id}).Count()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除组织机构信息失败", sys_dao.SysOrganization.Table())
	}

	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "存下级组织架构时禁止删除！"), "", sys_dao.SysOrganization.Table())
	}

	info := sys_entity.SysOrganization{}
	err = sys_dao.SysOrganization.Ctx(ctx).
		Where(sys_do.SysOrganization{Id: id}).Scan(&info)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "查询组织机构信息失败，请稍后再试", sys_dao.SysOrganization.Table())
	}

	if info.Id <= 0 {
		return true, nil
	}

	_, err = sys_dao.SysOrganization.Ctx(ctx).
		Cache(s.conf).
		Delete(sys_do.SysOrganization{Id: id})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除组织机构信息失败", sys_dao.SysOrganization.Table())
	}

	// 移除已缓存的数据
	daoctl.RemoveQueryCache(sys_dao.SysOrganization.DB(), sys_dao.SysOrganization.Table())
	return true, nil
}
