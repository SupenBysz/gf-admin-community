// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type (
	IArea interface {
		// GetAreaListByLevel 获取指定级级别地区列表
		GetAreaListByLevel(ctx context.Context, level int) (*sys_model.AreaListRes, error)
		// GetAreaListByParentId 获取属于父级ID的地区列表
		GetAreaListByParentId(ctx context.Context, parentId int64) (*sys_model.AreaListRes, error)
		// GetAreaById 根据ID获取区域信息
		GetAreaById(ctx context.Context, id int64) *sys_entity.SysArea
		// GetAreaByCode 根据区域编号获取区域信息
		GetAreaByCode(ctx context.Context, areaCode string) *sys_entity.SysArea
		// GetAreaByLongAndLong 根据区域经纬度获取区域信息
		GetAreaByLongAndLong(ctx context.Context, long float64, lat float64) *sys_entity.SysArea
		// GetAreaByName 根据区域名称获取区域信息
		GetAreaByName(ctx context.Context, areaName string) *sys_entity.SysArea
		// QueryAreaList 查询区域列表
		QueryAreaList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.AreaListRes, error)
	}
)

var (
	localArea IArea
)

func Area() IArea {
	if localArea == nil {
		panic("implement not found for interface IArea, forgot register?")
	}
	return localArea
}

func RegisterArea(i IArea) {
	localArea = i
}
