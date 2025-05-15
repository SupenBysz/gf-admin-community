package sys_front_settings

import (
	"context"
	"database/sql"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

/*
	前端配置信息
*/

type sSysFrontSettings struct {
}

func init() {
	sys_service.RegisterSysFrontSettings(New())
}

func New() sys_service.ISysFrontSettings {
	return &sSysFrontSettings{}
}

// QueryList 获取列表
func (s *sSysFrontSettings) QueryList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysFrontSettingsListRes, error) {

	result, err := daoctl.Query[sys_model.SysFrontSettingsRes](sys_dao.SysFrontSettings.Ctx(ctx), params, isExport)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_front_settings_list_get_failed", sys_dao.SysFrontSettings.Table())
	}

	return (*sys_model.SysFrontSettingsListRes)(result), err
}

// GetFrontSetting 获取前端配置信息
func (s *sSysFrontSettings) GetFrontSetting(ctx context.Context, name string, unionMainId int64, userId int64) (*sys_model.SysFrontSettingsRes, error) {
	data, err := sys_dao.SysFrontSettings.Ctx(ctx).Where(sys_do.SysFrontSettings{
		Name:        name,
		UnionMainId: unionMainId,
		UserId:      userId,
	}).One()

	if data == nil && err != sql.ErrNoRows {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_front_settings_get_failed", sys_dao.SysFrontSettings.Table())
	}

	frontSetting := sys_model.SysFrontSettingsRes{}
	err = data.Struct(&frontSetting)

	return &frontSetting, err
}

// Create  创建系统前端配置信息
func (s *sSysFrontSettings) Create(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error) {
	data := kconv.Struct(info, &sys_do.SysFrontSettings{})

	_, err := sys_dao.SysFrontSettings.Ctx(ctx).Insert(data)

	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_front_settings_save_failed", sys_dao.SysFrontSettings.Table()+":"+info.Name)
	}

	return s.GetFrontSetting(ctx, info.Name, info.UnionMainId, info.UserId)
}

// Update  修改系统前端配置信息
func (s *sSysFrontSettings) Update(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error) {
	data := kconv.Struct(info, &sys_do.SysFrontSettings{})
	data.UnionMainId = nil
	data.UserId = nil
	data.Name = nil

	selectInfo, err := s.GetFrontSetting(ctx, info.Name, info.UnionMainId, info.UserId)

	if selectInfo == nil {
		return nil, err
	}

	if _, err := sys_dao.SysFrontSettings.Ctx(ctx).Where(sys_do.SysFrontSettings{Name: info.Name, UnionMainId: info.UnionMainId}).OmitNilData().Update(&data); err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_front_settings_save_failed", sys_dao.SysFrontSettings.Table()+":"+info.Name)
	}

	return s.GetFrontSetting(ctx, info.Name, info.UnionMainId, info.UserId)
}

// Delete 删除
func (s *sSysFrontSettings) Delete(ctx context.Context, name string, unionMainId int64, userId int64) (bool, error) {

	affected, err := daoctl.DeleteWithError(sys_dao.SysFrontSettings.Ctx(ctx).Where(sys_do.SysFrontSettings{
		Name:        name,
		UnionMainId: unionMainId,
		UserId:      userId,
	}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delete_front_settings_failed", sys_dao.SysFrontSettings.Table())
	}

	return affected > 0, err
}
