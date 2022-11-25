package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/service"
)

// SysConfig 鉴权
var SysConfig = cSysConfig{}

type cSysConfig struct{}

// GetBaiduSdkConfList 获取百度SDK应用配置|列表
func (s *cSysConfig) GetBaiduSdkConfList(ctx context.Context, _ *sysapi.GetBaiduSdkConfListReq) (*sysapi.BaiduSdkConfListRes, error) {
	result := &sysapi.BaiduSdkConfListRes{
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := service.SdkBaidu().GetBaiduSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}

	return result, nil
}

// GetBaiduSdkConf 查询百度SDK应用配置|信息
func (s *cSysConfig) GetBaiduSdkConf(ctx context.Context, req *sysapi.GetBaiduSdkConfReq) (*sysapi.BaiduSdkConfRes, error) {
	result, err := service.SdkBaidu().GetBaiduSdkConf(ctx, req.Identifier)
	return (*sysapi.BaiduSdkConfRes)(result), err
}

// CreateBaiduSdkConf 创建百度SDK应用配置|信息
func (s *cSysConfig) CreateBaiduSdkConf(ctx context.Context, req *sysapi.CreateBaiduSdkConfReq) (*sysapi.BaiduSdkConfRes, error) {
	result, err := service.SdkBaidu().SaveBaiduSdkConf(ctx, req.BaiduSdkConf, true)
	return (*sysapi.BaiduSdkConfRes)(result), err
}

// UpdateBaiduSdkConf 更新百度SDK应用配置|信息
func (s *cSysConfig) UpdateBaiduSdkConf(ctx context.Context, req *sysapi.UpdateBaiduSdkConfReq) (*sysapi.BaiduSdkConfRes, error) {
	result, err := service.SdkBaidu().SaveBaiduSdkConf(ctx, req.BaiduSdkConf, false)
	return (*sysapi.BaiduSdkConfRes)(result), err
}

// DeleteBaiduSdkConf 删除百度SDK应用配置
func (s *cSysConfig) DeleteBaiduSdkConf(ctx context.Context, req *sysapi.DeleteBaiduSdkConfReq) (api_v1.BoolRes, error) {
	result, err := service.SdkBaidu().DeleteBaiduSdkConf(ctx, req.Identifier)
	return result == true, err
}
