package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

// 身份证识别

// 身份实名认证。扫脸

// 手机号认证，校验手机号是否是本人的，

// 银行卡识别

// 银行卡认证，校验银行卡是否是本人的，

type GetLicenseByIdReq struct {
	g.Meta `path:"/getLicenseById" method:"post" summary:"根据ID获取个人资质|信息" tags:"个人资质"`
	Id     int64 `json:"id" v:"required#个人资质ID校验失败" dc:"ID"`
}

type QueryLicenseListReq struct {
	g.Meta `path:"/queryLicenseList" method:"post" summary:"查询个人资质认证|列表" tags:"个人资质"`
	base_model.SearchParams
}

// type CreateLicenseReq struct {
//	g.Meta     `path:"/createLicense" method:"post" summary:"新增个人资质认证｜信息" tags:"伙伴个人资质"`
//	OperatorId int64 `json:"operatorId" v:"required|in:-1,0,1#关联运营商ID校验失败|关联运营商ID参赛错误" dc:"运营商ID"`
//	model.License
// }

type UpdateLicenseReq struct {
	g.Meta `path:"/updateLicense" method:"post" summary:"更新个人资质认证｜信息" tags:"个人资质"`
	sys_model.PersonLicense
	Id int64 `json:"id" v:"required#个人资质ID校验失败" dc:"ID"`
}
type SetLicenseStateReq struct {
	g.Meta `path:"/setLicenseState" method:"post" summary:"设置个人资质信息状态" tags:"个人资质"`
	Id     int64 `json:"id" v:"required#个人资质ID校验失败" dc:"ID"`
	State  int   `json:"state" v:"required#伙伴个人资质状态校验失败" dc:"状态：-1冻结、0未认证、1正常"`
}

type DeleteLicenseReq struct {
	g.Meta `path:"/deleteLicense" method:"post" summary:"设置个人资质审核编号" tags:"个人资质"`
	Id     int64 `json:"id" v:"required#个人资质ID校验失败" dc:"ID"`
}
