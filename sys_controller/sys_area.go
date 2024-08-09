package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysArea 地区
var SysArea = cSysArea{}

type cSysArea struct{}

// GetAreaListByParentId 获取属于父级ID的地区列表
func (c *cSysArea) GetAreaListByParentId(ctx context.Context, req *sys_api.GetAreaListByParentIdReq) (*sys_model.AreaListRes, error) {
	return sys_service.Area().GetAreaListByParentId(ctx, req.ParentId)
}

// GetAreaListByLevel 获取指定级别的地区列表
func (c *cSysArea) GetAreaListByLevel(ctx context.Context, req *sys_api.GetAreaListByLevelReq) (*sys_model.AreaListRes, error) {
	return sys_service.Area().GetAreaListByLevel(ctx, req.Level)
}
