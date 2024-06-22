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
	"github.com/kysion/base-library/utility/format_utils"
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

// GetAreaListByLevel 获取指定级级别地区列表
func (s *sArea) GetAreaListByLevel(ctx context.Context, level int) (*sys_model.AreaListRes, error) {
	result, err := s.QueryAreaList(ctx, &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0),
			base_model.FilterInfo{
				Field: sys_dao.SysArea.Columns().Level,
				Where: "=",
				Value: level,
			},
		),
	}, true)

	if err != nil {
		return &sys_model.AreaListRes{
			Records: make([]*sys_model.Area, 0),
		}, err
	}

	return result, err
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

	// 全部添加上拼音
	//if parentId == 0 {
	//	_, err := s.addPinyin(ctx, ret)
	//	fmt.Println(err)
	//}

	return ret, nil
}

// addPinyin 地区加上拼音
func (s *sArea) addPinyin(ctx context.Context, areaList *sys_model.AreaListRes) (bool, error) {

	//err := sys_dao.SysArea.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	// 省List
	for _, provinceItem := range areaList.Records {
		if provinceItem.PinYin == "" {
			//provinceItem.AreaName = provinceItem.AreaName[:len(provinceItem.AreaName)-1]
			provincePinyin := format_utils.ChineseToPinyin(provinceItem.AreaName)
			_, err := daoctl.UpdateWithError(sys_dao.SysArea.Ctx(ctx).
				Data(sys_do.SysArea{PinYin: provincePinyin}).
				Where(sys_do.SysArea{Id: provinceItem.Id, AreaCode: provinceItem.AreaCode}))
			if err != nil {
				continue
			}
		}
		// 市List
		cityList, _ := s.GetAreaListByParentId(ctx, provinceItem.Id)
		for _, cityItem := range cityList.Records {
			if cityItem.PinYin == "" {
				//cityItem.AreaName = cityItem.AreaName[:len(cityItem.AreaName)-1]
				cityPinyin := format_utils.ChineseToPinyin(cityItem.AreaName)
				_, err := daoctl.UpdateWithError(sys_dao.SysArea.Ctx(ctx).
					Data(sys_do.SysArea{PinYin: cityPinyin}).
					Where(sys_do.SysArea{Id: cityItem.Id, AreaCode: cityItem.AreaCode}))
				if err != nil {
					continue
				}
			}

			// 区县List
			regionList, _ := s.GetAreaListByParentId(ctx, cityItem.Id)
			for _, regionItem := range regionList.Records {
				if regionItem.PinYin == "" {
					//regionItem.AreaName = regionItem.AreaName[:len(regionItem.AreaName)-1]
					regionPinyin := format_utils.ChineseToPinyin(regionItem.AreaName)
					_, err := daoctl.UpdateWithError(sys_dao.SysArea.Ctx(ctx).
						Data(sys_do.SysArea{PinYin: regionPinyin}).
						Where(sys_do.SysArea{Id: regionItem.Id, AreaCode: regionItem.AreaCode}))
					if err != nil {
						continue
					}
				}
			}

		}

	}
	//	return nil
	//})

	//if err != nil {
	//	return false, nil
	//}

	return false, nil
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

// GetAreaByLongAndLong 根据区域经纬度获取区域信息
func (s *sArea) GetAreaByLongAndLong(ctx context.Context, long float64, lat float64) *sys_entity.SysArea {
	result := sys_entity.SysArea{}

	longLatStr := gconv.String(long) + "," + gconv.String(lat)

	if sys_dao.SysArea.Ctx(ctx).Scan(&result, sys_do.SysArea{LongLatCenter: longLatStr}) != nil {
		return nil
	}

	return &result
}

// GetAreaByName 根据区域名称获取区域信息
func (s *sArea) GetAreaByName(ctx context.Context, areaName string) *sys_entity.SysArea {
	result := sys_entity.SysArea{}
	//if sys_dao.SysArea.Ctx(ctx).Scan(&result, sys_do.SysArea{AreaName: areaName}) != nil {
	//	return nil
	//}

	//if sys_dao.SysArea.Ctx(ctx).WhereLike(sys_dao.SysArea.Columns().AreaName, "%"+areaName+"%").Scan(&result, sys_do.SysArea{AreaName: areaName}) != nil {
	if sys_dao.SysArea.Ctx(ctx).WhereLike(sys_dao.SysArea.Columns().AreaName, "%"+areaName+"%").Scan(&result) != nil {
		return nil
	}

	return &result
}

// QueryAreaList 查询区域列表
func (s *sArea) QueryAreaList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.AreaListRes, error) {
	result, err := daoctl.Query[*sys_model.Area](sys_dao.SysArea.Ctx(ctx), params, isExport)
	if err != nil {
		return &sys_model.AreaListRes{
			Records: make([]*sys_model.Area, 0),
		}, err
	}
	return (*sys_model.AreaListRes)(result), err
}
