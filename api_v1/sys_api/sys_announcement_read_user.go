package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

// 记录已读 - 在获取公告详情的时候，就增加一条公告的已读用户记录
//type MarkReadReq struct {
//	g.Meta `path:"/getMessageById" method:"post" summary:"根据id查询公告｜信息" tags:"公告/用户公告"`
//	Id     int64 `json:"id"  v:"required#消息Id不能为空"`
//}

type MarkUnReadReq struct {
	g.Meta `path:"/markUnRead" method:"post" summary:"标记未读｜公告" tags:"公告/用户公告"`
	Id     int64 `json:"id"  v:"required#公告Id不能为空"`
}

type HasUnReadAnnouncementReq struct {
	g.Meta `path:"/hasUnReadAnnouncement" method:"post" summary:"获取未读公告数量" tags:"公告/用户公告"`
	//Type   int `json:"type" dc:"公告类型"`
}

type QueryAnnouncementListByUserReq struct {
	g.Meta `path:"/queryAnnouncementListByUser" method:"post" summary:"查询用户的公告｜列表" tags:"公告/用户公告" dc:"查询标识type：0未读，1已读、2全部" default:"2"`
	base_model.SearchParams
	Type     int `json:"type" dc:"查询标识type：0未读，1已读、2全部"`
	IsExport bool
}

//type QueryAnnouncementReadUserListReq struct {
//	g.Meta `path:"/queryAnnouncementReadUserList" method:"post" summary:"查询公告的已读用户｜列表" tags:"公告/用户公告"`
//	base_model.SearchParams
//	IsExport bool
//}
