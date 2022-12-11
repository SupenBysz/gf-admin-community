package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysConfig 鉴权
var SysConfig = cSysConfig{}

type cSysConfig struct{}

// GetBaiduSdkConfList 获取百度SDK应用配置|列表
func (s *cSysConfig) GetBaiduSdkConfList(ctx context.Context, _ *sysapi.GetBaiduSdkConfListReq) (*sysapi.BaiduSdkConfListRes, error) {
	result := &sysapi.BaiduSdkConfListRes{
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := sys_service.SdkBaidu().GetBaiduSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}

	return result, nil
}

// GetBaiduSdkConf 查询百度SDK应用配置|信息
func (s *cSysConfig) GetBaiduSdkConf(ctx context.Context, req *sysapi.GetBaiduSdkConfReq) (*sysapi.BaiduSdkConfRes, error) {
	result, err := sys_service.SdkBaidu().GetBaiduSdkConf(ctx, req.Identifier)
	return (*sysapi.BaiduSdkConfRes)(result), err
}

// CreateBaiduSdkConf 创建百度SDK应用配置|信息
func (s *cSysConfig) CreateBaiduSdkConf(ctx context.Context, req *sysapi.CreateBaiduSdkConfReq) (*sysapi.BaiduSdkConfRes, error) {
	result, err := sys_service.SdkBaidu().SaveBaiduSdkConf(ctx, req.BaiduSdkConf, true)
	return (*sysapi.BaiduSdkConfRes)(result), err
}

// UpdateBaiduSdkConf 更新百度SDK应用配置|信息
func (s *cSysConfig) UpdateBaiduSdkConf(ctx context.Context, req *sysapi.UpdateBaiduSdkConfReq) (*sysapi.BaiduSdkConfRes, error) {
	result, err := sys_service.SdkBaidu().SaveBaiduSdkConf(ctx, req.BaiduSdkConf, false)
	return (*sysapi.BaiduSdkConfRes)(result), err
}

// DeleteBaiduSdkConf 删除百度SDK应用配置
func (s *cSysConfig) DeleteBaiduSdkConf(ctx context.Context, req *sysapi.DeleteBaiduSdkConfReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SdkBaidu().DeleteBaiduSdkConf(ctx, req.Identifier)
	return result == true, err
}

// GetAliyunSdkConfList 获取阿里云SDK应用配置|列表
func (s *cSysConfig) GetAliyunSdkConfList(ctx context.Context, _ *sysapi.GetAliyunSdkConfListReq) (*sysapi.AliyunSdkConfListRes, error) {
	result := &sysapi.AliyunSdkConfListRes{
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := sys_service.SdkAliyun().GetAliyunSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetAliyunSdkConf 查询阿里云SDK应用配置|信息
func (s *cSysConfig) GetAliyunSdkConf(ctx context.Context, req *sysapi.GetAliyunSdkConfReq) (*sysapi.AliyunSdkConfRes, error) {
	result, err := sys_service.SdkAliyun().GetAliyunSdkConf(ctx, req.Identifier)
	return (*sysapi.AliyunSdkConfRes)(result), err
}

// CreateAliyunSdkConf 创建阿里云SDK应用配置|信息
func (s *cSysConfig) CreateAliyunSdkConf(ctx context.Context, req *sysapi.CreateAliyunSdkConfReq) (*sysapi.AliyunSdkConfRes, error) {
	result, err := sys_service.SdkAliyun().SaveAliyunSdkConf(ctx, req.AliyunSdkConf, true)
	return (*sysapi.AliyunSdkConfRes)(result), err
}

// UpdateAliyunSdkConf 更新阿里云SDK应用配置|信息
func (s *cSysConfig) UpdateAliyunSdkConf(ctx context.Context, req *sysapi.UpdateAliyunSdkConfReq) (*sysapi.AliyunSdkConfRes, error) {
	result, err := sys_service.SdkAliyun().SaveAliyunSdkConf(ctx, req.AliyunSdkConf, false)
	return (*sysapi.AliyunSdkConfRes)(result), err
}

// DeleteAliyunSdkConf 删除阿里云SDK应用配置|信息
func (s *cSysConfig) DeleteAliyunSdkConf(ctx context.Context, req *sysapi.DeleteAliyunSdkConfReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SdkAliyun().DeleteAliyunSdkConf(ctx, req.Identifier)
	return result == true, err
}

// 华为云

// GetHuaweiSdkConfList 获取华为云SDK应用配置|列表
func (s *cSysConfig) GetHuaweiSdkConfList(ctx context.Context, _ *sysapi.GetHuaweiSdkConfListReq) (*sysapi.HuaweiSdkConfListRes, error) {
	result := &sysapi.HuaweiSdkConfListRes{
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := sys_service.SdkHuawei().GetHuaweiSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetHuaweiSdkConf 查询华为云SDK应用配置|信息
func (s *cSysConfig) GetHuaweiSdkConf(ctx context.Context, req *sysapi.GetHuaweiSdkConfReq) (*sysapi.HuaweiSdkConfRes, error) {
	result, err := sys_service.SdkHuawei().GetHuaweiSdkConf(ctx, req.Identifier)
	return (*sysapi.HuaweiSdkConfRes)(result), err
}

// CreateHuaweiSdkConf 创建华为云SDK应用配置|信息
func (s *cSysConfig) CreateHuaweiSdkConf(ctx context.Context, req *sysapi.CreateHuaweiSdkConfReq) (*sysapi.HuaweiSdkConfRes, error) {
	result, err := sys_service.SdkHuawei().SaveHuaweiSdkConf(ctx, req.HuaweiSdkConf, true)
	return (*sysapi.HuaweiSdkConfRes)(result), err
}

// UpdateHuaweiSdkConf 更新华为云SDK应用配置|信息
func (s *cSysConfig) UpdateHuaweiSdkConf(ctx context.Context, req *sysapi.UpdateHuaweiSdkConfReq) (*sysapi.HuaweiSdkConfRes, error) {
	result, err := sys_service.SdkHuawei().SaveHuaweiSdkConf(ctx, req.HuaweiSdkConf, false)
	return (*sysapi.HuaweiSdkConfRes)(result), err
}

// DeleteHuaweiSdkConf 删除华为云SDK应用配置|信息
func (s *cSysConfig) DeleteHuaweiSdkConf(ctx context.Context, req *sysapi.DeleteHuaweiSdkConfReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SdkHuawei().DeleteHuaweiSdkConf(ctx, req.Identifier)
	return result == true, err
}

// 腾讯云

// GetTencentSdkConfList 获取腾讯云SDK应用配置|列表
func (s *cSysConfig) GetTencentSdkConfList(ctx context.Context, _ *sysapi.GetTencentSdkConfListReq) (*sysapi.TencentSdkConfListRes, error) {
	result := &sysapi.TencentSdkConfListRes{
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := sys_service.SdkTencent().GetTencentSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetTencentSdkConf 查询腾讯云SDK应用配置|信息
func (s *cSysConfig) GetTencentSdkConf(ctx context.Context, req *sysapi.GetTencentSdkConfReq) (*sysapi.TencentSdkConfRes, error) {
	result, err := sys_service.SdkTencent().GetTencentSdkConf(ctx, req.Identifier)
	return (*sysapi.TencentSdkConfRes)(result), err
}

// CreateTencentSdkConf 创建腾讯云SDK应用配置|信息
func (s *cSysConfig) CreateTencentSdkConf(ctx context.Context, req *sysapi.CreateTencentSdkConfReq) (*sysapi.TencentSdkConfRes, error) {
	result, err := sys_service.SdkTencent().SaveTencentSdkConf(ctx, req.TencentSdkConf, true)
	return (*sysapi.TencentSdkConfRes)(result), err
}

// UpdateTencentSdkConf 更新腾讯云SDK应用配置|信息
func (s *cSysConfig) UpdateTencentSdkConf(ctx context.Context, req *sysapi.UpdateTencentSdkConfReq) (*sysapi.TencentSdkConfRes, error) {
	result, err := sys_service.SdkTencent().SaveTencentSdkConf(ctx, req.TencentSdkConf, false)
	return (*sysapi.TencentSdkConfRes)(result), err
}

// DeleteTengxunSdkConf 删除腾讯云SDK应用配置|信息
func (s *cSysConfig) DeleteTengxunSdkConf(ctx context.Context, req *sysapi.DeleteTencentSdkConfReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SdkTencent().DeleteTencentSdkConf(ctx, req.Identifier)
	return result == true, err
}

// 天翼云

// GetCtyunSdkConfList 获取天翼云SDK应用配置|列表
func (s *cSysConfig) GetCtyunSdkConfList(ctx context.Context, _ *sysapi.GetCtyunSdkConfListReq) (*sysapi.CtyunSdkConfListRes, error) {
	result := &sysapi.CtyunSdkConfListRes{
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: 20,
			},
			PageTotal: 1,
		},
	}

	if items, err := sys_service.SdkCtyun().GetCtyunSdkConfList(ctx); err != nil {
		return nil, err
	} else {
		result.List = items
		result.PageSize = len(*items)
	}
	return result, nil
}

// GetCtyunSdkConf 查询天翼云SDK应用配置|信息
func (s *cSysConfig) GetCtyunSdkConf(ctx context.Context, req *sysapi.GetCtyunSdkConfReq) (*sysapi.CtyunSdkConfRes, error) {
	result, err := sys_service.SdkCtyun().GetCtyunSdkConf(ctx, req.Identifier)
	return (*sysapi.CtyunSdkConfRes)(result), err
}

// CreateCtyunSdkConf 创建天翼云SDK应用配置|信息
func (s *cSysConfig) CreateCtyunSdkConf(ctx context.Context, req *sysapi.CreateCtyunSdkConfReq) (*sysapi.CtyunSdkConfRes, error) {
	result, err := sys_service.SdkCtyun().SaveCtyunSdkConf(ctx, req.CtyunSdkConf, true)
	return (*sysapi.CtyunSdkConfRes)(result), err
}

// UpdateCtyunSdkConf 更新天翼云SDK应用配置|信息
func (s *cSysConfig) UpdateCtyunSdkConf(ctx context.Context, req *sysapi.UpdateCtyunSdkConfReq) (*sysapi.CtyunSdkConfRes, error) {
	result, err := sys_service.SdkCtyun().SaveCtyunSdkConf(ctx, req.CtyunSdkConf, false)
	return (*sysapi.CtyunSdkConfRes)(result), err
}

// DeleteTengxunSdkConf 删除天翼云SDK应用配置|信息
func (s *cSysConfig) DeleteCtyunSdkConf(ctx context.Context, req *sysapi.DeleteCtyunSdkConfReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SdkCtyun().DeleteCtyunSdkConf(ctx, req.Identifier)
	return result == true, err
}
