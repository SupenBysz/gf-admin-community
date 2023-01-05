package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/frame/g"
)

type QueryOrganizationListReq struct {
	g.Meta `path:"/queryOrganizationList?cp=5948649530392645" method:"post" summary:"获取组织架构|列表" tags:"组织架构"`
	sys_model.SearchParams
}

type GetOrganizationListReq struct {
	g.Meta      `path:"/getOrganizationList?cp=5948649530392645" method:"post" summary:"根据ID获取下级组织架构|列表" tags:"组织架构"`
	Id          int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
	IsRecursive bool  `json:"isRecursive" dc:"是否递归"`
}

type GetOrganizationTreeReq struct {
	g.Meta `path:"/getOrganizationTree?cp=5948649530392645" method:"post" summary:"根据ID获取下级组织架构|树" tags:"组织架构"`
	Id     int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
}

type CreateOrganizationInfoReq struct {
	g.Meta `path:"/createOrganization?cp=5948649828712517" method:"post" summary:"创建组织架构|信息" tags:"组织架构"`
	sys_model.SysOrganizationInfo
}

type UpdateOrganizationInfoReq struct {
	g.Meta `path:"/updateOrganization?cp=5948649642721349" method:"post" summary:"更新组织架构|信息" tags:"组织架构"`
	sys_model.SysOrganizationInfo
}

type GetOrganizationInfoReq struct {
	g.Meta `path:"/getOrganizationInfo?cp=5948649434447941" method:"post" summary:"获取组织架构|信息" tags:"组织架构"`
	Id     int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
}

type DeleteOrganizationInfoReq struct {
	g.Meta `path:"/deleteOrganizationInfo?cp=5948649739583557" method:"post" summary:"根据ID删除组织架构|信息" tags:"组织架构"`
	Id     int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
}

type OrganizationInfoRes sys_entity.SysOrganization

type OrganizationInfoTreeRes sys_model.SysOrganizationTree
type OrganizationInfoTreeListRes []sys_model.SysOrganizationTree
