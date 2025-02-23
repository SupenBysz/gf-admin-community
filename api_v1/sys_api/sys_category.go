package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetCategoryByIdReq struct {
	g.Meta `path:"/getCategoryById" method:"post" summary:"获取分类｜信息" tags:"分类管理"`
	Id     int64 `json:"id" dc:"分类ID" v:"required#请输入分类信息id"`
}

type CreateCategoryReq struct {
	g.Meta `path:"/createCategory" method:"post" summary:"创建获取分类｜信息" tags:"分类管理"`
	Name   string `json:"name"        v:"required#请输入分类名称"          description:"名称"`
	sys_model.SysCategory
}

type UpdateCategoryReq struct {
	g.Meta `path:"/updateCategory" method:"post" summary:"更新获取分类｜信息" tags:"分类管理"`
	Id     int64  `json:"id" v:"required#请输入分类信息id" dc:"分类ID"`
	Name   string `json:"name"        v:"required#请输入分类名称"          description:"名称"`
	sys_model.SysCategory
}

type DeleteCategoryReq struct {
	g.Meta `path:"/deleteCategory" method:"post" summary:"删除获取分类｜信息" tags:"分类管理"`
	Id     int64 `json:"id" dc:"分类ID" v:"required#请输入分类信息id"`
}

type QueryCategoryReq struct {
	g.Meta `path:"/queryCategory" method:"post" summary:"查询获取分类｜列表" tags:"分类管理"`
	base_model.SearchParams
}
