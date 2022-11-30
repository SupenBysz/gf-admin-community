package area

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sArea struct {
	cachePrefix string
}

func init() {
	service.RegisterArea(New())
}

func New() *sArea {
	return &sArea{
		cachePrefix: "area_",
	}
}

// GetAreaListByParentId 获取属于父级ID的地区列表
func (s *sArea) GetAreaListByParentId(ctx context.Context, parentId int64) (*model.AreaListRes, error) {
	if parentId <= 0 || parentId == 10000 {
		parentId = 0
	}

	{
		// 如果缓存有数据则直接从缓存加载
		ret, _ := gcache.Get(ctx, s.cachePrefix+gconv.String(parentId))
		if !ret.IsEmpty() {
			response := model.AreaListRes{}
			if nil != ret.Struct(&response) {
				return &response, nil
			}
		}
	}

	result, _ := daoctl.Query[entity.SysArea](dao.SysArea.Ctx(ctx), &model.SearchFilter{
		Fields: append(make([]model.SearchField, 0), model.SearchField{
			Field:       dao.SysArea.Columns().ParentId,
			Where:       "=",
			IsOrWhere:   false,
			Value:       parentId,
			IsNullValue: false,
		}, model.SearchField{
			Field:       dao.SysArea.Columns().Id,
			Where:       ">",
			IsOrWhere:   false,
			Value:       0,
			IsNullValue: false,
		}),
		Pagination: model.Pagination{
			Page:     1,
			PageSize: 100,
		},
		OrderBy: model.OrderBy{
			Columns: dao.SysArea.Columns().Id,
		},
	}, false)

	items := make([]model.Area, 0)
	ret := &model.AreaListRes{
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
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
		info := model.Area{}
		if nil == gconv.Struct(area, &info) {
			items = append(items, info)
		}
	}

	ret.List = &items

	// 写入缓存
	gcache.Set(ctx, s.cachePrefix+gconv.String(parentId), ret, time.Hour*24*7)

	return ret, nil
}

// GetAreaById 根据ID获取区域信息
func (s *sArea) GetAreaById(ctx context.Context, id int64) *entity.SysArea {
	return daoctl.GetById[entity.SysArea](dao.SysArea.Ctx(ctx), id)
}

// GetAreaByCode 根据区域编号获取区域信息
func (s *sArea) GetAreaByCode(ctx context.Context, areaCode string) *entity.SysArea {
	result := entity.SysArea{}
	if dao.SysArea.Ctx(ctx).Scan(&result, do.SysArea{AreaCode: areaCode}) != nil {
		return nil
	}
	return &result
}
