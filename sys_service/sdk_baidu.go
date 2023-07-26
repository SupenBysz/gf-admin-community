// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type (
	ISdkBaidu interface {
		// GetBaiduSdkConfToken 根据 identifier 查询百度SDK应用配置和Token信息
		GetBaiduSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.BaiduSdkConfToken, err error)
		// GetBaiduSdkConfList 获取百度SDK应用配置列表
		GetBaiduSdkConfList(ctx context.Context) ([]*sys_model.BaiduSdkConf, error)
		// GetBaiduSdkConf 根据 identifier 查询百度SDK应用配置信息
		GetBaiduSdkConf(ctx context.Context, identifier string) (*sys_model.BaiduSdkConf, error)
		// SaveBaiduSdkConf 保存百度SDK应用配信息
		SaveBaiduSdkConf(ctx context.Context, info *sys_model.BaiduSdkConf, isCreate bool) (*sys_model.BaiduSdkConf, error)
		// DeleteBaiduSdkConf 删除百度SDK应用配置信息
		DeleteBaiduSdkConf(ctx context.Context, identifier string) (bool, error)
		// OCRBankCard OCR识别银行卡
		OCRBankCard(ctx context.Context, imageBase64 string) (*sys_model.OCRBankCard, error)
		// OCRIDCard OCR识别身份证
		OCRIDCard(ctx context.Context, imageBase64 string, detectRisk string, idCardSide string) (*sys_model.BaiduSdkOCRIDCard, error)
		// OCRBusinessLicense OCR识别营业执照
		OCRBusinessLicense(ctx context.Context, imageBase64 string) (*sys_model.BusinessLicenseOCR, error)
	}
)

var (
	localSdkBaidu ISdkBaidu
)

func SdkBaidu() ISdkBaidu {
	if localSdkBaidu == nil {
		panic("implement not found for interface ISdkBaidu, forgot register?")
	}
	return localSdkBaidu
}

func RegisterSdkBaidu(i ISdkBaidu) {
	localSdkBaidu = i
}
