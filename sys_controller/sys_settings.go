package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_model"
)

var SysSettings = cSysSettings{}

type cSysSettings struct{}

func (c *cSysSettings) QueryList(ctx context.Context, req *sys_api.QuerySettingListReq) (*sys_model.SysSettingListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	req.Filter = append(req.Filter, base_model.FilterInfo{
		Field: sys_dao.SysSettings.Columns().UnionMainId,
		Where: "=",
		Value: user.UnionMainId,
	})

	ret, err := sys_service.SysSettings().QueryList(ctx, &req.SearchParams, false)

	return ret, err
}

func (c *cSysSettings) GetByName(ctx context.Context, req *sys_api.GetSettingByNameReq) (*sys_model.SysSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	req.Filter = append(req.Filter, base_model.FilterInfo{
		Field: sys_dao.SysSettings.Columns().UnionMainId,
		Where: "=",
		Value: user.UnionMainId,
	})

	ret, err := sys_service.SysSettings().GetByName(ctx, req.Name, &req.SearchParams)

	return ret, err
}

func (c *cSysSettings) Update(ctx context.Context, req *sys_api.UpdateSettingReq) (*sys_model.SysSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.UnionMainId = user.UnionMainId

	ret, err := sys_service.SysSettings().Update(ctx, &req.SysSettings)

	return ret, err
}

func (c *cSysSettings) Delete(ctx context.Context, req *sys_api.DeleteSettingReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	ret, err := sys_service.SysSettings().Delete(ctx, req.Name, user.UnionMainId)

	return ret == true, err
}
