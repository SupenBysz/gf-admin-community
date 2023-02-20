package sys_area

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"time"
)

type sArea struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterArea(New())
}

func New() *sArea {
	return &sArea{
		conf: gdb.CacheOption{
			Duration: time.Hour * 24 * 7,
			Force:    false,
		},
	}
}

// GetAreaListByParentId 获取属于父级ID的地区列表
func (s *sArea) GetAreaListByParentId(ctx context.Context, parentId int64) (*sys_model.AreaListRes, error) {
	if parentId <= 0 || parentId == 10000 {
		parentId = 0
	}

	{
		// 如果缓存有数据则直接从缓存加载
		// ret, _ := gcache.Get(ctx, s.cachePrefix+gconv.String(parentId))
		//    cache := sys_dao.SysArea.DB().GetCache()

		// if !ret.IsEmpty() {
		//	response := sys_model.AreaListRes{}
		//	if nil != ret.Struct(&response) {
		//		return &response, nil
		//	}
		// }
	}

	result, _ := daoctl.Query[*sys_entity.SysArea](sys_dao.SysArea.Ctx(ctx), &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field:       sys_dao.SysArea.Columns().ParentId,
			Where:       "=",
			IsOrWhere:   false,
			Value:       parentId,
			IsNullValue: false,
		}, base_model.FilterInfo{
			Field:       sys_dao.SysArea.Columns().Id,
			Where:       ">",
			IsOrWhere:   false,
			Value:       0,
			IsNullValue: false,
		}),
		Pagination: base_model.Pagination{
			PageNum:  1,
			PageSize: 100,
		},
		OrderBy: append(make([]base_model.OrderBy, 0), base_model.OrderBy{
			Field: sys_dao.SysArea.Columns().Id,
		}),
	}, false)

	items := make([]*sys_model.Area, 0)
	ret := &sys_model.AreaListRes{
		PaginationRes: base_model.PaginationRes{
			Pagination: base_model.Pagination{
				PageNum:  1,
				PageSize: 0,
			},
			PageTotal: 0,
		},
		Records: items,
	}

	if len(result.Records) == 0 {
		return ret, nil
	}

	for _, area := range result.Records {
		info := &sys_model.Area{}
		if nil == gconv.Struct(area, info) {
			items = append(items, info)
		}
	}

	ret.Records = items
	ret.Total = gconv.Int64(len(items))

	// 写入缓存
	// gcache.Set(ctx, s.cachePrefix+gconv.String(parentId), ret, time.Hour*24*7)

	return ret, nil
}

// GetAreaById 根据ID获取区域信息
func (s *sArea) GetAreaById(ctx context.Context, id int64) *sys_entity.SysArea {
	result, err := daoctl.GetByIdWithError[sys_entity.SysArea](sys_dao.SysArea.Ctx(ctx), id)

	if err != nil {
		return nil
	}

	return result
}

// GetAreaByCode 根据区域编号获取区域信息
func (s *sArea) GetAreaByCode(ctx context.Context, areaCode string) *sys_entity.SysArea {
	result := sys_entity.SysArea{}
	if sys_dao.SysArea.Ctx(ctx).Scan(&result, sys_do.SysArea{AreaCode: areaCode}) != nil {
		return nil
	}
	return &result
}
