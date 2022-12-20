package sys_area

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sArea struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterArea(New())
}

func New() *sArea {
	return &sArea{
		cachePrefix: "area_",
	}
}

// GetAreaListByParentId 获取属于父级ID的地区列表
func (s *sArea) GetAreaListByParentId(ctx context.Context, parentId int64) (*sys_model.AreaListRes, error) {
	if parentId <= 0 || parentId == 10000 {
		parentId = 0
	}

	{
		// 如果缓存有数据则直接从缓存加载
		ret, _ := gcache.Get(ctx, s.cachePrefix+gconv.String(parentId))
		if !ret.IsEmpty() {
			response := sys_model.AreaListRes{}
			if nil != ret.Struct(&response) {
				return &response, nil
			}
		}
	}

	result, _ := daoctl.Query[sys_entity.SysArea](sys_dao.SysArea.Ctx(ctx), &sys_model.SearchParams{
		Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
			Field:       sys_dao.SysArea.Columns().ParentId,
			Where:       "=",
			IsOrWhere:   false,
			Value:       parentId,
			IsNullValue: false,
		}, sys_model.FilterInfo{
			Field:       sys_dao.SysArea.Columns().Id,
			Where:       ">",
			IsOrWhere:   false,
			Value:       0,
			IsNullValue: false,
		}),
		Pagination: sys_model.Pagination{
			Page:     1,
			PageSize: 100,
		},
		OrderBy: append(make([]sys_model.OrderBy, 0), sys_model.OrderBy{
			Field: sys_dao.SysArea.Columns().Id,
		}),
	}, false)

	items := make([]sys_model.Area, 0)
	ret := &sys_model.AreaListRes{
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: 0,
			},
			PageTotal: 0,
		},
		List: &items,
	}

	if len(*result.List) == 0 {
		return ret, nil
	}

	for _, area := range *result.List {
		info := sys_model.Area{}
		if nil == gconv.Struct(area, &info) {
			items = append(items, info)
		}
	}

	ret.List = &items
	ret.Total = len(items)

	// 写入缓存
	gcache.Set(ctx, s.cachePrefix+gconv.String(parentId), ret, time.Hour*24*7)

	return ret, nil
}

// GetAreaById 根据ID获取区域信息
func (s *sArea) GetAreaById(ctx context.Context, id int64) *sys_entity.SysArea {
	return daoctl.GetById[sys_entity.SysArea](sys_dao.SysArea.Ctx(ctx), id)
}

// GetAreaByCode 根据区域编号获取区域信息
func (s *sArea) GetAreaByCode(ctx context.Context, areaCode string) *sys_entity.SysArea {
	result := sys_entity.SysArea{}
	if sys_dao.SysArea.Ctx(ctx).Scan(&result, sys_do.SysArea{AreaCode: areaCode}) != nil {
		return nil
	}
	return &result
}
