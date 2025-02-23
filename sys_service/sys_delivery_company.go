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
	ISysDeliveryCompany interface {
		// GetDeliveryCompanyById 根据ID获取快递公司信息
		GetDeliveryCompanyById(ctx context.Context, id int64) (*sys_model.SysDeliveryCompanyRes, error)
		// DeleteDeliveryCompanyById 根据ID删除快递公司
		DeleteDeliveryCompanyById(ctx context.Context, id int64) (api_v1.BoolRes, error)
		// SaveDeliveryCompany 保存快递公司
		SaveDeliveryCompany(ctx context.Context, info *sys_model.SysDeliveryCompany) (*sys_model.SysDeliveryCompanyRes, error)
		// QueryDeliveryCompany 查询快递公司
		QueryDeliveryCompany(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysDeliveryCompanyListRes, error)
	}
)

var (
	localSysDeliveryCompany ISysDeliveryCompany
)

func SysDeliveryCompany() ISysDeliveryCompany {
	if localSysDeliveryCompany == nil {
		panic("implement not found for interface ISysDeliveryCompany, forgot register?")
	}
	return localSysDeliveryCompany
}

func RegisterSysDeliveryCompany(i ISysDeliveryCompany) {
	localSysDeliveryCompany = i
}
