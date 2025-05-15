package sys_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
)

/*
	前端配置信息
*/

var SysFrontSettings = cSysFrontSettings{}

type cSysFrontSettings struct{}

func (c *cSysFrontSettings) QueryList(ctx context.Context, req *sys_api.QueryFrontSettingListReq) (*sys_model.SysFrontSettingsListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	userId := int64(0)
	newFilter := make([]base_model.FilterInfo, 0)

	// 如果不是超级管理员则对查询条件做一定的限制
	if user.IsSuperAdmin == false {
		for _, info := range req.Filter {
			// 过滤掉 UnionMainId 和 UserId 的查询条件
			if info.Field != sys_dao.SysFrontSettings.Columns().UnionMainId && info.Field != sys_dao.SysFrontSettings.Columns().UserId {
				newFilter = append(newFilter, info)
			}
			// 提取 UserId 值
			if info.Field == sys_dao.SysFrontSettings.Columns().UserId {
				userId = gconv.Int64(info.Value)
			}
		}
		// 强制附加用户所属 UnionMainId 条件
		newFilter = append(newFilter, base_model.FilterInfo{
			Field: sys_dao.SysFrontSettings.Columns().UnionMainId,
			Where: "=",
			Value: user.UnionMainId,
		})
		// 如果不是主体管理员，则只循序查询自己的个人设置项
		if user.IsAdmin == false {
			// 强制附加 UserId 条件
			newFilter = append(newFilter, base_model.FilterInfo{
				Field: sys_dao.SysFrontSettings.Columns().UserId,
				Where: "=",
				Value: user.Id,
			})
		} else { // 否则是主体管理员，允许查询主体管理员下所有个人设置
			// 强制附加 UnionMainId 条件
			newFilter = append(newFilter, base_model.FilterInfo{
				Field: sys_dao.SysFrontSettings.Columns().UnionMainId,
				Where: "=",
				Value: userId,
			})
		}
	}

	ret, err := sys_service.SysFrontSettings().QueryList(ctx, &req.SearchParams, false)

	return ret, err
}

func (c *cSysFrontSettings) GetFrontSetting(ctx context.Context, req *sys_api.GetFrontSettingReq) (*sys_model.SysFrontSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if user.IsSuperAdmin == false && user.IsAdmin == false {
		req.UserId = user.Id
	}

	isSys := 0
	unionMainId := req.UnionMainId

	if req.Sys != nil {
		isSys = *req.Sys
	}

	if isSys == 1 {
		req.UserId = 0
	}

	ret, _ := sys_service.SysFrontSettings().GetFrontSetting(ctx, req.Name, unionMainId, req.UserId)

	return ret, nil
}

func (c *cSysFrontSettings) Save(ctx context.Context, req *sys_api.SaveFrontSettingReq) (*sys_model.SysFrontSettingsRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if user.IsSuperAdmin == false && user.IsAdmin == false && req.Sys != 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_permission_insufficient"), "", sys_dao.SysFrontSettings.Table())
	}

	req.UserId = user.Id
	req.UnionMainId = user.UnionMainId

	// 如果设置的是全局配置或主体配置项，则忽略用户条件
	if req.Sys == 1 {
		req.UserId = 0
	}

	// 如果不是管理员，仅限保持个人自己的配置设置
	if user.IsSuperAdmin == false && user.IsAdmin == false {
		req.UserId = user.Id
		req.Sys = 0
	}

	newFilter := make([]base_model.FilterInfo, 0)

	newFilter = append(newFilter, base_model.FilterInfo{
		Field: sys_dao.SysFrontSettings.Columns().UserId,
		Where: "=",
		Value: req.UserId,
	}, base_model.FilterInfo{
		Field: sys_dao.SysFrontSettings.Columns().UnionMainId,
		Where: "=",
		Value: req.UnionMainId,
	}, base_model.FilterInfo{
		Field: sys_dao.SysFrontSettings.Columns().Sys,
		Where: "=",
		Value: req.Sys,
	})

	info, _ := sys_service.SysFrontSettings().GetFrontSetting(ctx, req.Name, req.UnionMainId, req.UserId)

	if info != nil {
		return sys_service.SysFrontSettings().Update(ctx, &sys_model.SysFrontSettings{
			Name:        req.Name,
			UserId:      req.UserId,
			UnionMainId: req.UnionMainId,
			Values:      req.Values,
			Version:     req.Version,
			Sys:         req.Sys,
		})
	}

	return sys_service.SysFrontSettings().Create(ctx, &req.SysFrontSettings)
}

func (c *cSysFrontSettings) Delete(ctx context.Context, req *sys_api.DeleteFrontSettingReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	if user.IsSuperAdmin == false && user.IsAdmin == false {
		req.UserId = user.Id
	}

	ret, err := sys_service.SysFrontSettings().Delete(ctx, req.Name, user.UnionMainId, user.Id)

	return ret == true, err
}
