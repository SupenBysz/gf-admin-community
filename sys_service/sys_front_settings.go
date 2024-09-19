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
	ISysFrontSettings interface {
		// QueryList 获取列表
		QueryList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysFrontSettingsListRes, error)
		// GetByName 根据 name 查询配置信息
		GetByName(ctx context.Context, name string, info *base_model.SearchParams) (*sys_model.SysFrontSettingsRes, error)
		// Create  创建系统配置信息
		Create(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error)
		// Update  修改系统配置信息
		Update(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error)
		// Delete 删除
		Delete(ctx context.Context, name string, unionMainId int64) (bool, error)
	}
)

var (
	localSysFrontSettings ISysFrontSettings
)

func SysFrontSettings() ISysFrontSettings {
	if localSysFrontSettings == nil {
		panic("implement not found for interface ISysFrontSettings, forgot register?")
	}
	return localSysFrontSettings
}

func RegisterSysFrontSettings(i ISysFrontSettings) {
	localSysFrontSettings = i
}
