// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/entity"
)

type (
	IArea interface {
		GetAreaListByParentId(ctx context.Context, parentId int64) (*model.AreaListRes, error)
		GetAreaById(ctx context.Context, id int64) *entity.SysArea
		GetAreaByCode(ctx context.Context, areaCode string) *entity.SysArea
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