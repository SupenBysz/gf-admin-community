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
