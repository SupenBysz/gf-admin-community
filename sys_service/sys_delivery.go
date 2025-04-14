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
	ISysDelivery interface {
		// GetDeliveryById 根据ID获取快递公司信息
		GetDeliveryById(ctx context.Context, id int64) (*sys_model.SysDeliveryRes, error)
		// DeleteDeliveryById 根据ID删除快递公司
		DeleteDeliveryById(ctx context.Context, id int64) (api_v1.BoolRes, error)
		// SaveDelivery 保存快递公司
		SaveDelivery(ctx context.Context, info *sys_model.SysDelivery) (*sys_model.SysDeliveryRes, error)
		// QueryDelivery 查询快递公司
		QueryDelivery(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysDeliveryListRes, error)
	}
)

var (
	localSysDelivery ISysDelivery
)

func SysDelivery() ISysDelivery {
	if localSysDelivery == nil {
		panic("implement not found for interface ISysDelivery, forgot register?")
	}
	return localSysDelivery
}

func RegisterSysDelivery(i ISysDelivery) {
	localSysDelivery = i
}
