// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	IAnnouncement interface {
		// GetAnnouncementById 根据id查询公告｜信息
		GetAnnouncementById(ctx context.Context, id int64, userId ...int64) (*sys_model.SysAnnouncementRes, error)
		// CreateAnnouncement 添加公告｜信息
		CreateAnnouncement(ctx context.Context, info *sys_model.SysAnnouncement, unionMainId int64, userId int64) (*sys_model.SysAnnouncementRes, error)
		// UpdateAnnouncement 编辑公告｜信息
		UpdateAnnouncement(ctx context.Context, info *sys_model.UpdateSysAnnouncement, unionMainId int64, userId int64) (*sys_model.SysAnnouncementRes, error)
		// DeleteAnnouncement 删除公告
		DeleteAnnouncement(ctx context.Context, id int64, unionMainId int64, userId int64) (bool, error)
		// QueryAnnouncement 查询公告｜列表
		QueryAnnouncement(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementListRes, error)
		// MarkRead 标记已读｜公告
		MarkRead(ctx context.Context, announcementId int64, userId int64) (bool, error)
		// MarkUnRead 标记未读｜公告
		MarkUnRead(ctx context.Context, announcementId int64, userId int64) (bool, error)
		// HasUnReadAnnouncement 获取未读公告数量
		HasUnReadAnnouncement(ctx context.Context, userId int64, unionMainId int64) (int, error)
		// QueryAnnouncementListByUser 查询用户的公告｜列表 （qType：0未读，1已读、2全部）
		QueryAnnouncementListByUser(ctx context.Context, userId int64, unionMainId int64, qType int, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementListRes, error)
	}
)

var (
	localAnnouncement IAnnouncement
)

func Announcement() IAnnouncement {
	if localAnnouncement == nil {
		panic("implement not found for interface IAnnouncement, forgot register?")
	}
	return localAnnouncement
}

func RegisterAnnouncement(i IAnnouncement) {
	localAnnouncement = i
}
