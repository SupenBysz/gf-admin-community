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
	ISdkAliyun interface {
		GetAliyunSdkToken(ctx context.Context, tokenInfo sys_model.AliyunSdkConfToken, err error)
		GetAliyunSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.AliyunSdkConfToken, err error)
		GetAliyunSdkConfList(ctx context.Context) (*[]sys_model.AliyunSdkConf, error)
		GetAliyunSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.AliyunSdkConf, err error)
		SaveAliyunSdkConf(ctx context.Context, info sys_model.AliyunSdkConf, isCreate bool) (*sys_model.AliyunSdkConf, error)
		DeleteAliyunSdkConf(ctx context.Context, identifier string) (bool, error)
	}
)

var (
	localSdkAliyun ISdkAliyun
)

func SdkAliyun() ISdkAliyun {
	if localSdkAliyun == nil {
		panic("implement not found for interface ISdkAliyun, forgot register?")
	}
	return localSdkAliyun
}

func RegisterSdkAliyun(i ISdkAliyun) {
	localSdkAliyun = i
}