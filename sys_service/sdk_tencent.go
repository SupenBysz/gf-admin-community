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
	ISdkTencent interface {
		// GetTencentSdkConfToken 根据 identifier 查询腾讯SDK应用配置和Token信息
		GetTencentSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConfToken, err error)
		// GetTencentSdkConfList 获取腾讯云SDK应用配置列表
		GetTencentSdkConfList(ctx context.Context) ([]*sys_model.TencentSdkConf, error)
		// GetTencentSdkConf 根据identifier标识获取SDK配置信息
		GetTencentSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConf, err error)
		// SaveTencentSdkConf 保存腾讯SDK应用配信息, isCreate判断是更新还是新建
		SaveTencentSdkConf(ctx context.Context, info *sys_model.TencentSdkConf, isCreate bool) (*sys_model.TencentSdkConf, error)
		// DeleteTencentSdkConf 删除腾讯SDK应用配置信息
		DeleteTencentSdkConf(ctx context.Context, identifier string) (bool, error)
	}
)

var (
	localSdkTencent ISdkTencent
)

func SdkTencent() ISdkTencent {
	if localSdkTencent == nil {
		panic("implement not found for interface ISdkTencent, forgot register?")
	}
	return localSdkTencent
}

func RegisterSdkTencent(i ISdkTencent) {
	localSdkTencent = i
}
