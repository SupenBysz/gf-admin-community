// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type (
	ISdkTencent interface {
		GetTencentSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConfToken, err error)
		GetTencentSdkConfList(ctx context.Context) (*[]sys_model.TencentSdkConf, error)
		GetTencentSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConf, err error)
		SaveTencentSdkConf(ctx context.Context, info sys_model.TencentSdkConf, isCreate bool) (*sys_model.TencentSdkConf, error)
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