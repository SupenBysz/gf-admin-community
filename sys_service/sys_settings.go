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
	ISysSettings interface {
		// QueryList 获取列表
		QueryList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysSettingListRes, error)
		// GetByName 根据 name 查询百度SDK应用配置信息
		GetByName(ctx context.Context, name string, info *base_model.SearchParams) (*sys_model.SysSettingsRes, error)
		// Save 保存系统配置信息
		Save(ctx context.Context, info *sys_model.SysSettings) (*sys_model.SysSettingsRes, error)
		// Create  创建系统配置信息
		Create(ctx context.Context, info *sys_model.SysSettings) (*sys_model.SysSettingsRes, error)
		// Update  修改系统配置信息
		Update(ctx context.Context, info *sys_model.SysSettings) (*sys_model.SysSettingsRes, error)
		// Delete 删除
		Delete(ctx context.Context, name string, unionMainId int64) (bool, error)
	}
)

var (
	localSysSettings ISysSettings
)

func SysSettings() ISysSettings {
	if localSysSettings == nil {
		panic("implement not found for interface ISysSettings, forgot register?")
	}
	return localSysSettings
}

func RegisterSysSettings(i ISysSettings) {
	localSysSettings = i
}
