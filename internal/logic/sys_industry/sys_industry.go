package sys_menu

import (
	"context"
	"sort"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

// 设计修改行业类别的操作只有超级管理员-1有权限，其他用户只能看看

type sSysIndustry struct {
}

func init() {
	sys_service.RegisterSysIndustry(New())
}

// New sSysIndustry 行业类别逻辑实现
func New() *sSysIndustry {
	return &sSysIndustry{}
}

// GetIndustryById 根据ID获取行业类别信息
func (s *sSysIndustry) GetIndustryById(ctx context.Context, id int64) (*sys_entity.SysIndustry, error) {
	result := sys_entity.SysIndustry{}
	err := sys_dao.SysIndustry.Ctx(ctx).Where(sys_do.SysIndustry{Id: id}).Scan(&result)
	//res, err := daoctl.GetByIdWithError[sys_entity.SysIndustry](sys_dao.SysIndustry.Ctx(ctx), id)
	//res, err := daoctl.ScanWithError[sys_entity.SysIndustry](sys_dao.SysIndustry.Ctx(ctx).Where(sys_do.SysIndustry{Id: id}))

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_query_failed", sys_dao.SysIndustry.Table())
	}
	return &result, err
}

// CreateIndustry 创建行业类别
func (s *sSysIndustry) CreateIndustry(ctx context.Context, info *sys_model.SysIndustry) (*sys_entity.SysIndustry, error) {
	return s.SaveIndustry(ctx, info)
}

// UpdateIndustry 更新行业类别
func (s *sSysIndustry) UpdateIndustry(ctx context.Context, info *sys_model.UpdateSysIndustry) (*sys_entity.SysIndustry, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "error_industry_id_parameter_incorrect"), "", sys_dao.SysIndustry.Table())
	}
	data := kconv.Struct(info, &sys_model.SysIndustry{})

	return s.SaveIndustry(ctx, data)
}

// SaveIndustry 新增或保存行业类别信息，并自动更新对应的权限信息
func (s *sSysIndustry) SaveIndustry(ctx context.Context, info *sys_model.SysIndustry) (*sys_entity.SysIndustry, error) {
	data := sys_do.SysIndustry{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级行业类别信息是否存在
	if info.ParentId != nil && *info.ParentId > 0 {
		permissionInfo, err := s.GetIndustryById(ctx, gconv.Int64(data.ParentId))
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_parent_industry_not_exists", sys_dao.SysIndustry.Table())
		}
	}

	err := sys_dao.SysIndustry.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if info.Id <= 0 { // 插入
			if data.ParentId == nil {
				*info.ParentId = 0
				data.ParentId = 0
			}
			id := idgen.NextId()
			data.Id = id
			info.Id = id
			data.CreatedAt = gtime.Now()

			_, err := sys_dao.SysIndustry.Ctx(ctx).Insert(&data)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_add_failed", sys_dao.SysIndustry.Table())
			}

		} else { // 更新
			data.UpdatedAt = gtime.Now()
			data.Id = nil
			_, err := sys_dao.SysIndustry.Ctx(ctx).
				OmitNilData().Where(sys_dao.SysIndustry.Columns().Id, info.Id).Update(&data)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_save_failed", sys_dao.SysIndustry.Table())
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	//go s.makeIndustryTeee(ctx, 0)
	return s.GetIndustryById(ctx, info.Id)
}

// DeleteIndustry 删除行业类别，删除的时候要关联删除sys_permission,有子行业类别时禁止删除。
func (s *sSysIndustry) DeleteIndustry(ctx context.Context, id int64) (bool, error) {
	_, err := s.GetIndustryById(ctx, id)

	if err != nil {
		return false, err
	}

	// 判断是否具备子行业类别
	count, _ := sys_dao.SysIndustry.Ctx(ctx).Where(sys_dao.SysIndustry.Columns().ParentId, id).Count()
	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_industry_has_children", sys_dao.SysIndustry.Table())
	}

	err = sys_dao.SysIndustry.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除关联sys_permission权限记录 		行业类别id = 权限id
		ret, err := sys_service.SysPermission().DeletePermission(ctx, id)
		if err != nil || ret == false {
			return err
		}

		// 删除行业类别记录
		_, err = daoctl.DeleteWithError(sys_dao.SysIndustry.Ctx(ctx).Where(sys_dao.SysIndustry.Columns().Id, id))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	//g.DB().GetCache().Remove(ctx, "getIndustryTreeCacheById_"+gconv.String(id))

	//daoctl.RemoveQueryCache(sys_dao.SysIndustry.DB(), sys_dao.SysIndustry.Table())

	//go s.MakeIndustryTree(ctx, 0)

	// 有更新的时候，把所有行业类别树缓存清除，然后重构

	return true, nil
}

// MakeIndustryTree 构建行业类别树
func (s *sSysIndustry) MakeIndustryTree(ctx context.Context, parentId int64, isMakeNodeFun func(ctx context.Context, cruuentIndustry *sys_entity.SysIndustry) bool) ([]*sys_model.SysIndustryTreeRes, error) {
	// 当前级行业类别树没加入

	// 获取下级行业类别列表
	result, err := s.GetIndustryList(ctx, parentId, false)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_query_failed", sys_dao.SysPermission.Table())
	}

	response := make([]*sys_model.SysIndustryTreeRes, 0)

	// 有数据，则递归加载
	if len(result) > 0 {
		// 我第一次进来的样子

		// 我进来是否具备上级
		for _, sysIndustryItem := range result {
			item := &sys_model.SysIndustryTreeRes{}
			gconv.Struct(sysIndustryItem, &item)

			tree, err := s.MakeIndustryTree(ctx, sysIndustryItem.Id, isMakeNodeFun)

			for _, childItem := range tree {
				if isMakeNodeFun(ctx, childItem.SysIndustry) || len(childItem.Children) > 0 {
					item.Children = append(item.Children, childItem)
				}
			}

			if len(tree) == 0 && !isMakeNodeFun(ctx, sysIndustryItem) && len(item.Children) == 0 {
				continue
			}

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_query_failed", sys_dao.SysIndustry.Table())
			}

			response = append(response, item)

		}
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_query_failed", sys_dao.SysIndustry.Table())
	}

	return response, nil
}

// GetIndustryTree 根据ID获取下级行业类别信息，返回行业类别树，并缓存
func (s *sSysIndustry) GetIndustryTree(ctx context.Context, parentId int64) (sys_model.SysIndustryTreeListRes, error) {
	// 先判断缓存中是否存在行业类别树，存在直接返回
	//res, _ := g.DB().GetCache().Get(ctx, "IndustryTreeCache_"+gconv.String(parentId))

	//if res.Val() != nil {
	//	data := make([]*sys_model.SysIndustryTreeRes, 0)
	//	gconv.Struct(res, &data)
	//	return data, nil
	//}

	// 将权限树缓存起来,
	// 注意：这种自定义缓存方式，Ctx内部的缓存机制不会自动帮我们删除该缓存，所以，需要每次有更新操作，手动删除缓存getIndustryTreeCacheById_的缓存
	tree, err := s.MakeIndustryTree(ctx, parentId, func(ctx context.Context, cruuentIndustry *sys_entity.SysIndustry) bool {
		return true
	})
	if err != nil {
		return nil, err
	}

	//g.DB().GetCache().Set(ctx, "IndustryTreeCache_"+gconv.String(parentId), tree, 0)

	return tree, nil
}

// GetIndustryList 根据ID获取下级行业类别列表，IsRecursive代表是否需要返回下级
func (s *sSysIndustry) GetIndustryList(ctx context.Context, parentId int64, IsRecursive bool, limitChildrenIds ...int64) ([]*sys_entity.SysIndustry, error) {
	dataArr := make([]*sys_entity.SysIndustry, 0)

	response := make([]*sys_entity.SysIndustry, 0)
	model := sys_dao.SysIndustry.Ctx(ctx).Where(sys_dao.SysIndustry.Columns().ParentId, parentId)

	if len(limitChildrenIds) > 0 {
		model = model.WhereIn(sys_dao.SysIndustry.Columns().Id, limitChildrenIds)
	}

	model.Scan(&dataArr)
	gconv.Struct(dataArr, &response)

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(dataArr) > 0 {
		for _, sysIndustry := range dataArr {
			var children []*sys_entity.SysIndustry
			children, err := s.GetIndustryList(ctx, sysIndustry.Id, IsRecursive, limitChildrenIds...)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_industry_query_failed", sys_dao.SysIndustry.Table())
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

// QueryIndustryList 获取行业类别列表
