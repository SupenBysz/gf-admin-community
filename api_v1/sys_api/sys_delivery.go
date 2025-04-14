package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetDeliveryByIdReq struct {
	g.Meta `path:"/getDeliveryById" method:"post" summary:"根据id查询物流公司｜信息" tags:"配送设置/物流公司"`
	Id     int64 `json:"id"  v:"required#物流公司Id不能为空"`
}

type CreateDeliveryReq struct {
	g.Meta `path:"/createDelivery" method:"post" summary:"添加物流公司｜信息" tags:"配送设置/物流公司"`

	sys_model.SysDelivery
}

type UpdateDeliveryReq struct {
	g.Meta `path:"/updateDelivery" method:"post" summary:"编辑物流公司｜信息" tags:"配送设置/物流公司"`
	Id     int64 `json:"id"  v:"required#物流公司Id不能为空"`

	sys_model.SysDelivery
}

type DeleteDeliveryReq struct {
	g.Meta `path:"/deleteDelivery" method:"post" summary:"删除物流公司" tags:"配送设置/物流公司"`
	Id     int64 `json:"id" v:"required#物流公司Id不能为空"`
}

type QueryDeliveryReq struct {
	g.Meta `path:"/queryDelivery" method:"post" summary:"查询物流公司｜列表" tags:"配送设置/物流公司"`
	base_model.SearchParams
	IsExport bool
}
