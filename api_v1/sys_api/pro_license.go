package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetLicenseByIdReq struct {
	g.Meta `path:"/getLicenseById" method:"post" summary:"根据ID获取主体|信息" tags:"主体资质"`
	Id     int64 `json:"id" v:"required#主体ID校验失败" dc:"ID"`
}

type QueryLicenseListReq struct {
	g.Meta `path:"/queryLicenseList" method:"post" summary:"查询主体认证|列表" tags:"主体资质"`
	sys_model.SearchParams
}

// type CreateLicenseReq struct {
//	g.Meta     `path:"/createLicense" method:"post" summary:"新增主体认证｜信息" tags:"伙伴主体资质"`
//	OperatorId int64 `json:"operatorId" v:"required|in:-1,0,1#关联运营商ID校验失败|关联运营商ID参赛错误" dc:"运营商ID"`
//	model.License
// }

type UpdateLicenseReq struct {
	g.Meta `path:"/updateLicense" method:"post" summary:"更新主体认证｜信息" tags:"主体资质"`
	sys_model.License
	Id int64 `json:"id" v:"required#主体ID校验失败" dc:"ID"`
}
type SetLicenseStateReq struct {
	g.Meta `path:"/setLicenseState" method:"post" summary:"设置主体信息状态" tags:"主体资质"`
	Id     int64 `json:"id" v:"required#主体ID校验失败" dc:"ID"`
	State  int   `json:"state" v:"required#伙伴主体状态校验失败" dc:"状态：-1冻结、0未认证、1正常"`
}

type DeleteLicenseReq struct {
	g.Meta `path:"/deleteLicense" method:"post" summary:"设置主体神审核编号" tags:"主体资质"`
	Id     int64 `json:"id" v:"required#主体ID校验失败" dc:"ID"`
}

type LicenseRes sys_model.LicenseRes
type LicenseListRes sys_model.LicenseListRes
