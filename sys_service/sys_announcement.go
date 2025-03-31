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
		// GetCategoryById 根据ID获取公告分类
		GetCategoryById(ctx context.Context, id int64) (*sys_model.SysAnnouncementCategoryRes, error)
		// CreateCategory 创建公告分类
		CreateCategory(ctx context.Context, info *sys_model.SysAnnouncementCategory, unionMainId int64, userId int64) (*sys_model.SysAnnouncementCategoryRes, error)
		// UpdateCategory 更新公告分类
		UpdateCategory(ctx context.Context, info *sys_model.SysAnnouncementCategory, unionMainId int64, userId int64) (*sys_model.SysAnnouncementCategoryRes, error)
		// DeleteCategory 删除公告分类
		DeleteCategory(ctx context.Context, id int64, unionMainId int64) (bool, error)
		// QueryCategory 查询公告分类列表
		QueryCategory(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementCategoryListRes, error)
		// CheckCategoryUsage 检查分类是否被使用
		CheckCategoryUsage(ctx context.Context, categoryId int64) (bool, int, error)
		// ConfirmAnnouncement 确认公告
		ConfirmAnnouncement(ctx context.Context, announcementId int64, userId int64, comment string) (bool, error)
		// IsAnnouncementConfirmed 检查公告是否已确认
		IsAnnouncementConfirmed(ctx context.Context, announcementId int64, userId int64) (bool, error)
		// GetAnnouncementConfirmInfo 获取公告确认信息
		GetAnnouncementConfirmInfo(ctx context.Context, announcementId int64, userId int64) (*sys_model.SysAnnouncementConfirmRes, error)
		// QueryAnnouncementConfirms 查询公告确认记录列表
		QueryAnnouncementConfirms(ctx context.Context, announcementId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementConfirmListRes, error)
		// MarkRead 标记已读｜公告
		MarkRead(ctx context.Context, announcementId int64, userId int64) (bool, error)
		// MarkUnRead 标记未读｜公告
		MarkUnRead(ctx context.Context, announcementId int64, userId int64) (bool, error)
		// HasUnReadAnnouncement 获取未读公告数量
		HasUnReadAnnouncement(ctx context.Context, userId int64, unionMainId int64) (int, error)
		// QueryAnnouncementListByUser 查询用户的公告｜列表 （qType：0未读，1已读、2全部）
		QueryAnnouncementListByUser(ctx context.Context, userId int64, unionMainId int64, qType int, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementListRes, error)
		// RevokedAnnouncement 撤销公告
		RevokedAnnouncement(ctx context.Context, announcementId int64) (bool, error)
		// SetAnnouncementState 设置公告状态
		SetAnnouncementState(ctx context.Context, announcementId int64, state int) (bool, error)
		// BatchMarkRead 批量标记公告为已读
		BatchMarkRead(ctx context.Context, announcementIds []int64, userId int64) (bool, error)
		// GetAnnouncementStatistics 获取公告统计信息
		GetAnnouncementStatistics(ctx context.Context, announcementId int64) (*sys_model.SysAnnouncementStatistics, error)
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
