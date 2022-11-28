// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model"
)

type (
	ISdkHuawei interface {
		GetHuaweiSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.HuaweiSdkConfToken, err error)
		GetHuaweiSdkConfList(ctx context.Context) (*[]model.HuaweiSdkConf, error)
		GetHuaweiSdkConf(ctx context.Context, identifier string) (tokenInfo *model.HuaweiSdkConf, err error)
		SaveHuaweiSdkConf(ctx context.Context, info model.HuaweiSdkConf, isCreate bool) (*model.HuaweiSdkConf, error)
		DeleteHuaweiSdkConf(ctx context.Context, identifier string) (bool, error)
	}
)

var (
	localSdkHuawei ISdkHuawei
)

func SdkHuawei() ISdkHuawei {
	if localSdkHuawei == nil {
		panic("implement not found for interface ISdkHuawei, forgot register?")
	}
	return localSdkHuawei
}

func RegisterSdkHuawei(i ISdkHuawei) {
	localSdkHuawei = i
}