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
	ISysIndustry interface {
		// GetIndustryById 根据ID获取行业类别信息
		GetIndustryById(ctx context.Context, menuId int64) (*sys_entity.SysIndustry, error)
		// CreateIndustry 创建行业类别
		CreateIndustry(ctx context.Context, info *sys_model.SysIndustry) (*sys_entity.SysIndustry, error)
		// UpdateIndustry 更新行业类别
		UpdateIndustry(ctx context.Context, info *sys_model.UpdateSysIndustry) (*sys_entity.SysIndustry, error)
		// SaveIndustry 新增或保存行业类别信息，并自动更新对应的权限信息
		SaveIndustry(ctx context.Context, info *sys_model.SysIndustry) (*sys_entity.SysIndustry, error)
		// DeleteIndustry 删除行业类别，删除的时候要关联删除sys_permission,有子行业类别时禁止删除。
		DeleteIndustry(ctx context.Context, id int64) (bool, error)
		// MakeIndustryTree 构建行业类别树
		MakeIndustryTree(ctx context.Context, parentId int64, isMakeNodeFun func(ctx context.Context, cruuentIndustry *sys_entity.SysIndustry) bool) ([]*sys_model.SysIndustryTreeRes, error)
		// GetIndustryTree 根据ID获取下级行业类别信息，返回行业类别树，并缓存
		GetIndustryTree(ctx context.Context, parentId int64) (sys_model.SysIndustryTreeListRes, error)
		// GetIndustryList 根据ID获取下级行业类别列表，IsRecursive代表是否需要返回下级
		GetIndustryList(ctx context.Context, parentId int64, IsRecursive bool, limitChildrenIds ...int64) ([]*sys_entity.SysIndustry, error)
	}
)

var (
	localSysIndustry ISysIndustry
)

func SysIndustry() ISysIndustry {
	if localSysIndustry == nil {
		panic("implement not found for interface ISysIndustry, forgot register?")
	}
	return localSysIndustry
}

func RegisterSysIndustry(i ISysIndustry) {
	localSysIndustry = i
}
