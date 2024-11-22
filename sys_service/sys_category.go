// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysCategory interface {
		// GetCategoryById 根据ID查下分类
		GetCategoryById(ctx context.Context, id int64) (*sys_model.SysCategoryRes, error)
		// SaveCategory 保存分类
		SaveCategory(ctx context.Context, info *sys_model.SysCategory) (*sys_model.SysCategoryRes, error)
		// DeleteCategory 删除分类
		DeleteCategory(ctx context.Context, id int64) (api_v1.BoolRes, error)
		// QueryCategory 查询分类
		QueryCategory(ctx context.Context, search *base_model.SearchParams) (*sys_model.SysCategoryListRes, error)
	}
)

var (
	localSysCategory ISysCategory
)

func SysCategory() ISysCategory {
	if localSysCategory == nil {
		panic("implement not found for interface ISysCategory, forgot register?")
	}
	return localSysCategory
}

func RegisterSysCategory(i ISysCategory) {
	localSysCategory = i
}
