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

/*
	前端配置信息
*/

var SysFrontSettings = cSysFrontSettings{}

type cSysFrontSettings struct{}

func (c *cSysFrontSettings) QueryList(ctx context.Context, req *sys_api.QueryFrontSettingListReq) (*sys_model.SysFrontSettingsListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	req.Filter = append(req.Filter, base_model.FilterInfo{
		Field: sys_dao.SysFrontSettings.Columns().UnionMainId,
		Where: "=",
		Value: user.UnionMainId,
	})

	ret, err := sys_service.SysFrontSettings().QueryList(ctx, &req.SearchParams, false)

	return ret, err
}

func (c *cSysFrontSettings) GetByName(ctx context.Context, req *sys_api.GetFrontSettingByNameReq) (*sys_model.SysFrontSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	req.Filter = append(req.Filter, base_model.FilterInfo{
		Field: sys_dao.SysFrontSettings.Columns().UnionMainId,
		Where: "=",
		Value: user.UnionMainId,
	})

	ret, err := sys_service.SysFrontSettings().GetByName(ctx, req.Name, &req.SearchParams)

	return ret, err
}

func (c *cSysFrontSettings) Save(ctx context.Context, req *sys_api.SaveFrontSettingReq) (*sys_model.SysFrontSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.UnionMainId = user.UnionMainId
	req.UserId = user.Id

	ret, err := sys_service.SysFrontSettings().Create(ctx, &req.SysFrontSettings)

	return ret, err
}

func (c *cSysFrontSettings) Delete(ctx context.Context, req *sys_api.DeleteFrontSettingReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	ret, err := sys_service.SysFrontSettings().Delete(ctx, req.Name, user.UnionMainId)

	return ret == true, err
}
