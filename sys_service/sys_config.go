// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysConfig interface {
		// QueryList 获取列表
		QueryList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysConfigListRes, error)
		// GetByName  根据Name获取应用配置
		GetByName(ctx context.Context, name string) (*sys_model.SysConfigRes, error)
		// SaveConfig 保存应用配信息
		SaveConfig(ctx context.Context, info *sys_model.SysConfig) (*sys_model.SysConfigRes, error)
		// DeleteConfig 删除应用配置信息
		DeleteConfig(ctx context.Context, name string) (bool, error)
	}
)

var (
	localSysConfig ISysConfig
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
