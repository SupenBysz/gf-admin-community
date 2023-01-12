package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type QueryOrganizationListReq struct {
	g.Meta `path:"/queryOrganizationList" method:"post" summary:"获取组织架构|列表" tags:"组织架构"`
	sys_model.SearchParams
}

type GetOrganizationListReq struct {
	g.Meta      `path:"/getOrganizationList" method:"post" summary:"根据ID获取下级组织架构|列表" tags:"组织架构"`
	Id          int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
	IsRecursive bool  `json:"isRecursive" dc:"是否递归"`
}

type GetOrganizationTreeReq struct {
	g.Meta `path:"/getOrganizationTree" method:"post" summary:"根据ID获取下级组织架构|树" tags:"组织架构"`
	Id     int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
}

type CreateOrganizationInfoReq struct {
	g.Meta `path:"/createOrganization" method:"post" summary:"创建组织架构|信息" tags:"组织架构"`
	sys_model.SysOrganizationInfo
}

type UpdateOrganizationInfoReq struct {
	g.Meta `path:"/updateOrganization" method:"post" summary:"更新组织架构|信息" tags:"组织架构"`
	sys_model.SysOrganizationInfo
}

type GetOrganizationInfoReq struct {
	g.Meta `path:"/getOrganizationInfo" method:"post" summary:"获取组织架构|信息" tags:"组织架构"`
	Id     int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
}

type DeleteOrganizationInfoReq struct {
	g.Meta `path:"/deleteOrganizationInfo" method:"post" summary:"根据ID删除组织架构|信息" tags:"组织架构"`
	Id     int64 `json:"id" v:"required#缺少ID参数" dc:"组织架构ID"`
}

type OrganizationInfoRes sys_model.OrganizationInfo

type OrganizationInfoTreeRes sys_model.SysOrganizationTree
type OrganizationInfoTreeListRes []*sys_model.SysOrganizationTree
