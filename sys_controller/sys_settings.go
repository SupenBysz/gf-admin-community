package sys_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

var SysSettings = cSysSettings{}

type cSysSettings struct{}

func (c *cSysSettings) QueryList(ctx context.Context, req *sys_api.QuerySettingListReq) (*sys_model.SysSettingListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	queryUnionMainId := user.UnionMainId

	filter := make([]base_model.FilterInfo, 0)

	for i, info := range req.Filter {
		if info.Field == sys_dao.SysSettings.Columns().UnionMainId {
			queryUnionMainId = info.Value.(int64)
		} else {
			filter = append(filter, req.Filter[i])
		}
	}

	cfgValue, err := g.Cfg().Get(ctx, "service.superAdminMainId", 0)

	if cfgValue != nil && err == nil {
		superAdminMainId := cfgValue.Int64()

		if user.UnionMainId != superAdminMainId {
			queryUnionMainId = user.UnionMainId
		} else {
			if queryUnionMainId == -1 {
				queryUnionMainId = 0
			}
		}
	}

	if queryUnionMainId > 0 {
		filter = append(filter, base_model.FilterInfo{
			Field: sys_dao.SysSettings.Columns().UnionMainId,
			Where: "=",
			Value: queryUnionMainId,
		})
	}

	req.Filter = filter

	ret, err := sys_service.SysSettings().QueryList(ctx, &req.SearchParams, false)

	return ret, err
}

func (c *cSysSettings) GetByName(ctx context.Context, req *sys_api.GetSettingByNameReq) (*sys_model.SysSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	queryUnionMainId := user.UnionMainId
	filter := make([]base_model.FilterInfo, 0)
	for i, info := range req.Filter {
		if info.Field == sys_dao.SysSettings.Columns().UnionMainId {
			queryUnionMainId = gconv.Int64(info.Value)
		} else {
			filter = append(filter, req.Filter[i])
		}
	}

	cfgValue, err := g.Cfg().Get(ctx, "service.superAdminMainId", 0)
	if cfgValue != nil && err == nil {
		superAdminMainId := cfgValue.Int64()

		if user.UnionMainId != superAdminMainId && queryUnionMainId != -1 {
			queryUnionMainId = user.UnionMainId
		} else {
			if queryUnionMainId == -1 {
				queryUnionMainId = 0
			}
		}
	}

	if queryUnionMainId > 0 {
		filter = append(filter, base_model.FilterInfo{
			Field: sys_dao.SysSettings.Columns().UnionMainId,
			Where: "=",
			Value: queryUnionMainId,
		})
	}

	req.Filter = filter

	ret, err := sys_service.SysSettings().GetByName(ctx, req.Name, &req.SearchParams)

	return ret, err
}

func (c *cSysSettings) Save(ctx context.Context, req *sys_api.SaveSettingReq) (*sys_model.SysSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	cfgValue, err := g.Cfg().Get(ctx, "service.superAdminMainId", 0)

	if cfgValue != nil && err == nil {
		superAdminMainId := cfgValue.Int64()

		if user.UnionMainId != superAdminMainId || req.UnionMainId == 0 {
			req.UnionMainId = user.UnionMainId
		} else {
			if req.UnionMainId == -1 {
				req.UnionMainId = 0
			}
		}
	}

	ret, err := sys_service.SysSettings().Create(ctx, &req.SysSettings)

	return ret, err
}

func (c *cSysSettings) Delete(ctx context.Context, req *sys_api.DeleteSettingReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	selectInfo, err := daoctl.ScanWithError[sys_entity.SysFrontSettings](sys_dao.SysFrontSettings.Ctx(ctx).Where(sys_do.SysFrontSettings{Name: req.Name, UnionMainId: user.UnionMainId}))
	if selectInfo != nil && selectInfo.UnionMainId <= 0 && user.IsSuperAdmin != true {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_settings_platform_config_delete_forbidden", sys_dao.SysFrontSettings.Table())
	}

	ret, err := sys_service.SysSettings().Delete(ctx, req.Name, user.UnionMainId)

	return ret == true, err
}
