package sys_config

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

/*
	应用配置
*/

type sSysConfig struct {
	ConfigList []*sys_model.SysConfigRes
}

func init() {
	sys_service.RegisterSysConfig(New())
}

func New() sys_service.ISysConfig {
	return &sSysConfig{
		// 内存缓存 --- 暂时没有用上
		ConfigList: make([]*sys_model.SysConfigRes, 0),
	}
}

// QueryList 获取列表
func (s *sSysConfig) QueryList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysConfigListRes, error) {
	result, err := daoctl.Query[sys_entity.SysConfig](sys_dao.SysConfig.Ctx(ctx), params, isExport)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_sys_config_list_query_failed"), sys_dao.SysConfig.Table())
	}

	return (*sys_model.SysConfigListRes)(result), err
}

// GetByName  根据Name获取应用配置
func (s *sSysConfig) GetByName(ctx context.Context, name string) (*sys_model.SysConfigRes, error) {
	items, err := s.QueryList(ctx, nil, true)
	if err != nil {
		return nil, err
	}

	for _, config := range items.Records {
		if config.Name == name {
			return (*sys_model.SysConfigRes)(&config), nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_sys_config_get_by_name_failed"), sys_dao.SysConfig.Table()+":"+name)

}

// SaveConfig 保存应用配信息
func (s *sSysConfig) SaveConfig(ctx context.Context, info *sys_model.SysConfig) (*sys_model.SysConfigRes, error) {
	data := kconv.Struct(info, &sys_do.SysConfig{})

	count, err := sys_dao.SysConfig.Ctx(ctx).Count(sys_do.SysConfig{Name: info.Name})
	if count > 0 {
		_, err = daoctl.UpdateWithError(sys_dao.SysConfig.Ctx(ctx).Where(sys_do.SysConfig{Name: info.Name}).OmitNilData().Data(sys_do.SysConfig{Value: data.Value}))
	} else {
		_, err = daoctl.InsertWithError(sys_dao.SysConfig.Ctx(ctx).Data(data))
	}

	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_sys_config_save_failed"), sys_dao.SysConfig.Table()+":"+info.Name)
	}

	return s.GetByName(ctx, info.Name)
}

// DeleteConfig 删除应用配置信息
func (s *sSysConfig) DeleteConfig(ctx context.Context, name string) (bool, error) {
	affected, err := daoctl.DeleteWithError(sys_dao.SysConfig.Ctx(ctx).Where(sys_do.SysConfig{Name: name}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, g.I18n().T(ctx, "error_sys_config_delete_failed"), sys_dao.SysConfig.Table()+":"+name)
	}

	return affected > 0, err

}
