package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetDeliveryCompanyByIdReq struct {
	g.Meta `path:"/getDeliveryCompanyById" method:"post" summary:"根据id查询物流公司｜信息" tags:"配送设置/物流公司"`
	Id     int64 `json:"id"  v:"required#物流公司Id不能为空"`
}

type CreateDeliveryCompanyReq struct {
	g.Meta `path:"/createDeliveryCompany" method:"post" summary:"添加物流公司｜信息" tags:"配送设置/物流公司"`

	sys_model.SysDeliveryCompany
}

type UpdateDeliveryCompanyReq struct {
	g.Meta `path:"/updateDeliveryCompany" method:"post" summary:"编辑物流公司｜信息" tags:"配送设置/物流公司"`
	Id     int64 `json:"id"  v:"required#物流公司Id不能为空"`

	sys_model.SysDeliveryCompany
}

type DeleteDeliveryCompanyReq struct {
	g.Meta `path:"/deleteDeliveryCompany" method:"post" summary:"删除物流公司" tags:"配送设置/物流公司"`
	Id     int64 `json:"id" v:"required#物流公司Id不能为空"`
}

type QueryDeliveryCompanyReq struct {
	g.Meta `path:"/queryDeliveryCompany" method:"post" summary:"查询物流公司｜列表" tags:"配送设置/物流公司"`
	base_model.SearchParams
	IsExport bool
}
