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

// GetAliyunSdkConfList 获取阿里云SDK应用配置|列表
func (s *cSysConfig) GetAliyunSdkConfList(ctx context.Context, _ *sysapi.GetAliyunSdkConfListReq) (*sysapi.AliyunSdkConfListRes, error) {
	result := &sysapi.AliyunSdkConfListRes{
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := service.SdkAliyun().GetAliyunSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetAliyunSdkConf 查询阿里云SDK应用配置|信息
func (s *cSysConfig) GetAliyunSdkConf(ctx context.Context, req *sysapi.GetAliyunSdkConfReq) (*sysapi.AliyunSdkConfRes, error) {
	result, err := service.SdkAliyun().GetAliyunSdkConf(ctx, req.Identifier)
	return (*sysapi.AliyunSdkConfRes)(result), err
}

// CreateAliyunSdkConf 创建阿里云SDK应用配置|信息
func (s *cSysConfig) CreateAliyunSdkConf(ctx context.Context, req *sysapi.CreateAliyunSdkConfReq) (*sysapi.AliyunSdkConfRes, error) {
	result, err := service.SdkAliyun().SaveAliyunSdkConf(ctx, req.AliyunSdkConf, true)
	return (*sysapi.AliyunSdkConfRes)(result), err
}

// UpdateAliyunSdkConf 更新阿里云SDK应用配置|信息
func (s *cSysConfig) UpdateAliyunSdkConf(ctx context.Context, req *sysapi.UpdateAliyunSdkConfReq) (*sysapi.AliyunSdkConfRes, error) {
	result, err := service.SdkAliyun().SaveAliyunSdkConf(ctx, req.AliyunSdkConf, false)
	return (*sysapi.AliyunSdkConfRes)(result), err
}

// DeleteAliyunSdkConf 删除阿里云SDK应用配置|信息
func (s *cSysConfig) DeleteAliyunSdkConf(ctx context.Context, req *sysapi.DeleteAliyunSdkConfReq) (api_v1.BoolRes, error) {
	result, err := service.SdkAliyun().DeleteAliyunSdkConf(ctx, req.Identifier)
	return result == true, err
}

// 华为云

// GetHuaWeiYunSdkConfList 获取华为云SDK应用配置|列表
func (s *cSysConfig) GetHuaWeiYunSdkConfList(ctx context.Context, _ *sysapi.GetHuaWeiYunSdkConfListReq) (*sysapi.HuaWeiYunSdkConfListRes, error) {
	result := &sysapi.HuaWeiYunSdkConfListRes{
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := service.SdkHuaWeiYun().GetHuaWeiYunSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetHuaWeiYunSdkConf 查询华为云SDK应用配置|信息
func (s *cSysConfig) GetHuaWeiYunSdkConf(ctx context.Context, req *sysapi.GetHuaWeiYunSdkConfReq) (*sysapi.HuaWeiYunSdkConfRes, error) {
	result, err := service.SdkHuaWeiYun().GetHuaWeiYunSdkConf(ctx, req.Identifier)
	return (*sysapi.HuaWeiYunSdkConfRes)(result), err
}

// CreateHuaWeiYunSdkConf 创建华为云SDK应用配置|信息
func (s *cSysConfig) CreateHuaWeiYunSdkConf(ctx context.Context, req *sysapi.CreateHuaWeiYunSdkConfReq) (*sysapi.HuaWeiYunSdkConfRes, error) {
	result, err := service.SdkHuaWeiYun().SaveHuaWeiYunSdkConf(ctx, req.HuaWeiYunSdkConf, true)
	return (*sysapi.HuaWeiYunSdkConfRes)(result), err
}

// UpdateHuaWeiYunSdkConf 更新华为云SDK应用配置|信息
func (s *cSysConfig) UpdateHuaWeiYunSdkConf(ctx context.Context, req *sysapi.UpdateHuaWeiYunSdkConfReq) (*sysapi.HuaWeiYunSdkConfRes, error) {
	result, err := service.SdkHuaWeiYun().SaveHuaWeiYunSdkConf(ctx, req.HuaWeiYunSdkConf, false)
	return (*sysapi.HuaWeiYunSdkConfRes)(result), err
}

// DeleteHuaWeiYunSdkConf 删除华为云SDK应用配置|信息
func (s *cSysConfig) DeleteHuaWeiYunSdkConf(ctx context.Context, req *sysapi.DeleteHuaWeiYunSdkConfReq) (api_v1.BoolRes, error) {
	result, err := service.SdkHuaWeiYun().DeleteHuaWeiYunSdkConf(ctx, req.Identifier)
	return result == true, err
}

// 腾讯云

// GetTengxunSdkConfList 获取腾讯云SDK应用配置|列表
func (s *cSysConfig) GetTengxunSdkConfList(ctx context.Context, _ *sysapi.GetTengxunSdkConfListReq) (*sysapi.TengxunSdkConfListRes, error) {
	result := &sysapi.TengxunSdkConfListRes{
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := service.SdkTengxun().GetTengxunSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetTengxunSdkConf 查询腾讯云SDK应用配置|信息
func (s *cSysConfig) GetTengxunSdkConf(ctx context.Context, req *sysapi.GetTengxunSdkConfReq) (*sysapi.TengxunSdkConfRes, error) {
	result, err := service.SdkTengxun().GetTengxunSdkConf(ctx, req.Identifier)
	return (*sysapi.TengxunSdkConfRes)(result), err
}

// CreateTengxunSdkConf 创建腾讯云SDK应用配置|信息
func (s *cSysConfig) CreateTengxunSdkConf(ctx context.Context, req *sysapi.CreateTengxunSdkConfReq) (*sysapi.TengxunSdkConfRes, error) {
	result, err := service.SdkTengxun().SaveTengxunSdkConf(ctx, req.TengxunSdkConf, true)
	return (*sysapi.TengxunSdkConfRes)(result), err
}

// UpdateTengxunSdkConf 更新腾讯云SDK应用配置|信息
func (s *cSysConfig) UpdateTengxunSdkConf(ctx context.Context, req *sysapi.UpdateTengxunSdkConfReq) (*sysapi.TengxunSdkConfRes, error) {
	result, err := service.SdkTengxun().SaveTengxunSdkConf(ctx, req.TengxunSdkConf, false)
	return (*sysapi.TengxunSdkConfRes)(result), err
}

// DeleteTengxunSdkConf 删除腾讯云SDK应用配置|信息
func (s *cSysConfig) DeleteTengxunSdkConf(ctx context.Context, req *sysapi.DeleteTengxunSdkConfReq) (api_v1.BoolRes, error) {
	result, err := service.SdkTengxun().DeleteTengxunSdkConf(ctx, req.Identifier)
	return result == true, err
}
