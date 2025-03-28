package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetAnnouncementByIdReq struct {
	g.Meta `path:"/getAnnouncementById" method:"post" summary:"根据id查询公告｜信息" tags:"公告/管理"`
	Id     int64 `json:"id"  v:"required#公告Id不能为空"`
}

type CreateAnnouncementReq struct {
	g.Meta `path:"/createAnnouncement" method:"post" summary:"添加公告｜信息" tags:"公告/管理"`

	sys_model.SysAnnouncement
}

type UpdateAnnouncementReq struct {
	g.Meta `path:"/updateAnnouncement" method:"post" summary:"编辑公告｜信息" tags:"公告/管理"`
	Id     int64 `json:"id"  v:"required#公告Id不能为空"`

	sys_model.UpdateSysAnnouncement
}

type DeleteAnnouncementReq struct {
	g.Meta `path:"/deleteAnnouncement" method:"post" summary:"删除公告" tags:"公告/管理"`
	Id     int64 `json:"id" v:"required#公告Id不能为空"`
}

type QueryAnnouncementReq struct {
	g.Meta `path:"/queryAnnouncement" method:"post" summary:"查询公告｜列表" tags:"公告/管理"`
	base_model.SearchParams
	IsExport bool
}

// 新增API接口定义 - 公告分类相关

type GetAnnouncementCategoryByIdReq struct {
	g.Meta `path:"/getCategoryById" method:"post" summary:"根据id查询公告分类" tags:"公告/分类管理"`
	Id     int64 `json:"id"  v:"required#分类Id不能为空"`
}

type CreateAnnouncementCategoryReq struct {
	g.Meta      `path:"/createCategory" method:"post" summary:"创建公告分类" tags:"公告/分类管理"`
	Name        string `json:"name" v:"required#分类名称不能为空"`
	Code        string `json:"code" v:"required#分类编码不能为空"`
	Sort        int    `json:"sort" dc:"排序值"`
	Description string `json:"description" dc:"分类描述"`
}

type UpdateAnnouncementCategoryReq struct {
	g.Meta      `path:"/updateCategory" method:"post" summary:"更新公告分类" tags:"公告/分类管理"`
	Id          int64  `json:"id"  v:"required#分类Id不能为空"`
	Name        string `json:"name" dc:"分类名称"`
	Code        string `json:"code" dc:"分类编码"`
	Sort        int    `json:"sort" dc:"排序值"`
	Description string `json:"description" dc:"分类描述"`
}

type DeleteAnnouncementCategoryReq struct {
	g.Meta `path:"/deleteCategory" method:"post" summary:"删除公告分类" tags:"公告/分类管理"`
	Id     int64 `json:"id" v:"required#分类Id不能为空"`
}

type QueryAnnouncementCategoryReq struct {
	g.Meta `path:"/queryCategory" method:"post" summary:"查询公告分类列表" tags:"公告/分类管理"`
	base_model.SearchParams
	IsExport bool
}

// 新增API接口定义 - 公告确认相关

type ConfirmAnnouncementReq struct {
	g.Meta  `path:"/confirmAnnouncement" method:"post" summary:"确认公告" tags:"公告/用户公告"`
	Id      int64  `json:"id"  v:"required#公告Id不能为空"`
	Comment string `json:"comment" dc:"确认评论"`
}

type GetAnnouncementConfirmInfoReq struct {
	g.Meta `path:"/getAnnouncementConfirmInfo" method:"post" summary:"获取公告确认信息" tags:"公告/用户公告"`
	Id     int64 `json:"id" v:"required#公告Id不能为空"`
}

type QueryAnnouncementConfirmsReq struct {
	g.Meta `path:"/queryAnnouncementConfirms" method:"post" summary:"查询公告确认记录" tags:"公告/管理"`
	Id     int64 `json:"id" v:"required#公告Id不能为空"`
	base_model.SearchParams
	IsExport bool
}

// 新增API接口定义 - 公告统计相关

type GetAnnouncementStatisticsReq struct {
	g.Meta `path:"/getAnnouncementStatistics" method:"post" summary:"获取公告统计信息" tags:"公告/管理"`
	Id     int64 `json:"id" v:"required#公告Id不能为空"`
}

// 新增API接口定义 - 批量操作相关

type BatchMarkReadReq struct {
	g.Meta `path:"/batchMarkRead" method:"post" summary:"批量标记公告为已读" tags:"公告/用户公告"`
	Ids    []int64 `json:"ids" v:"required#公告Id列表不能为空"`
}

// 新增API接口定义 - 撤销公告

type RevokedAnnouncementReq struct {
	g.Meta `path:"/revokedAnnouncement" method:"post" summary:"撤销公告" tags:"公告/管理"`
	Id     int64 `json:"id" v:"required#公告Id不能为空"`
}
