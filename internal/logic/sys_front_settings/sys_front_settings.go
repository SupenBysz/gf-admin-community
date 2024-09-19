package sys_front_settings

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gerror"
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("前端配置列表获取失败"), "", sys_dao.SysFrontSettings.Table())
	}

	return (*sys_model.SysFrontSettingsListRes)(result), err
}

// GetByName 根据 name 查询前端配置信息
func (s *sSysFrontSettings) GetByName(ctx context.Context, name string, info *base_model.SearchParams) (*sys_model.SysFrontSettingsRes, error) {
	items, err := s.QueryList(ctx, info, true)
	if err != nil {
		return nil, err
	}

	for _, setting := range items.Records {
		if setting.Name == name {
			return (*sys_model.SysFrontSettingsRes)(&setting), nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据 name 查询前端配置信息失败", sys_dao.SysFrontSettings.Table())
}

// save 保存系统前端配置信息
func (s *sSysFrontSettings) save(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error) {
	data := kconv.Struct(info, &sys_do.SysFrontSettings{})

	count, err := sys_dao.SysFrontSettings.Ctx(ctx).Count(sys_do.SysFrontSettings{Name: info.Name})

	if count > 0 {
		_, err = sys_dao.SysFrontSettings.Ctx(ctx).Where(sys_do.SysFrontSettings{Name: info.Name, UnionMainId: info.UnionMainId}).OmitNilData().Update(sys_do.SysFrontSettings{Values: data.Values, Desc: data.Desc})
	} else {
		_, err = sys_dao.SysFrontSettings.Ctx(ctx).Insert(data)
	}

	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "前端配置保存失败", sys_dao.SysFrontSettings.Table()+":"+info.Name)
	}

	return s.GetByName(ctx, info.Name, nil)
}

// Create  创建系统前端配置信息
func (s *sSysFrontSettings) Create(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error) {
	return s.save(ctx, info)
}

// Update  修改系统前端配置信息
func (s *sSysFrontSettings) Update(ctx context.Context, info *sys_model.SysFrontSettings) (*sys_model.SysFrontSettingsRes, error) {
	return s.save(ctx, info)
}

// Delete 删除
func (s *sSysFrontSettings) Delete(ctx context.Context, name string, unionMainId int64) (bool, error) {
	affected, err := daoctl.DeleteWithError(sys_dao.SysFrontSettings.Ctx(ctx).Where(sys_do.SysFrontSettings{Name: name, UnionMainId: unionMainId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除前端配置信息失败", sys_dao.SysFrontSettings.Table())
	}

	return affected > 0, err
}
