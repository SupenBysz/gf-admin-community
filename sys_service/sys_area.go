// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type (
	IArea interface {
		// GetAreaListByParentId 获取属于父级ID的地区列表
		GetAreaListByParentId(ctx context.Context, parentId int64) (*sys_model.AreaListRes, error)
		// GetAreaById 根据ID获取区域信息
		GetAreaById(ctx context.Context, id int64) *sys_entity.SysArea
		// GetAreaByCode 根据区域编号获取区域信息
		GetAreaByCode(ctx context.Context, areaCode string) *sys_entity.SysArea
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
