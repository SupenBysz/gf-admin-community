package sys_settings

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

type sSysSettings struct {
}

func init() {
	sys_service.RegisterSysSettings(New())
}

func New() *sSysSettings {
	return &sSysSettings{}
}

// QueryList 获取列表
func (s *sSysSettings) QueryList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysSettingListRes, error) {
	result, err := daoctl.Query[sys_entity.SysSettings](sys_dao.SysSettings.Ctx(ctx), params, isExport)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New(g.I18n().T(ctx, "error_settings_list_query_failed")), "", sys_dao.SysSettings.Table())
	}

	return (*sys_model.SysSettingListRes)(result), err
}

// GetByName 根据 name 查询配置信息
func (s *sSysSettings) GetByName(ctx context.Context, name string, info *base_model.SearchParams) (*sys_model.SysSettingsRes, error) {
	items, err := s.QueryList(ctx, info, true)
	if err != nil {
		return nil, err
	}

	for _, setting := range items.Records {
		if setting.Name == name {
			return (*sys_model.SysSettingsRes)(&setting), nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_settings_get_by_name_failed"), sys_dao.SysSettings.Table())
}

// save 保存系统配置信息
func (s *sSysSettings) save(ctx context.Context, info *sys_model.SysSettings) (*sys_model.SysSettingsRes, error) {
	data := kconv.Struct(info, &sys_do.SysSettings{})

	selectInfo, err := daoctl.ScanWithError[sys_entity.SysSettings](sys_dao.SysSettings.Ctx(ctx).Where(sys_do.SysSettings{Name: info.Name}))

	if selectInfo != nil {
		if selectInfo.UnionMainId != info.UnionMainId && info.UnionMainId > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New(g.I18n().T(ctx, "error_settings_subject_mismatch")), "", sys_dao.SysSettings.Table()+":"+info.Name)
		}

		_, err = sys_dao.SysSettings.Ctx(ctx).Where(sys_do.SysSettings{Name: info.Name, UnionMainId: info.UnionMainId}).OmitNilData().Update(sys_do.SysSettings{Values: data.Values, Desc: data.Desc})
	} else {
		_, err = sys_dao.SysSettings.Ctx(ctx).Insert(data)
	}

	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_settings_save_failed"), sys_dao.SysSettings.Table()+":"+info.Name)
	}

	return s.GetByName(ctx, info.Name, nil)
}

// Create  创建系统配置信息
func (s *sSysSettings) Create(ctx context.Context, info *sys_model.SysSettings) (*sys_model.SysSettingsRes, error) {
	return s.save(ctx, info)
}

// Update  修改系统配置信息
func (s *sSysSettings) Update(ctx context.Context, info *sys_model.SysSettings) (*sys_model.SysSettingsRes, error) {
	return s.save(ctx, info)
}

// Delete 删除
func (s *sSysSettings) Delete(ctx context.Context, name string, unionMainId int64) (bool, error) {

	affected, err := daoctl.DeleteWithError(sys_dao.SysSettings.Ctx(ctx).Where(sys_do.SysSettings{Name: name, UnionMainId: unionMainId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_settings_delete_failed"), sys_dao.SysSettings.Table())
	}

	return affected > 0, err
}
