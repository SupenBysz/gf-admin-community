package sys_announcement

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
)

// BatchMarkRead 批量标记公告为已读
func (s *sAnnouncement) BatchMarkRead(ctx context.Context, announcementIds []int64, userId int64) (bool, error) {
	if len(announcementIds) == 0 {
		return true, nil
	}

	// 检查用户是否存在
	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, userId)
	if err != nil || sysUser == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_user_not_exists}", sys_dao.SysUser.Table())
	}

	// 开始事务处理
	err = sys_dao.SysAnnouncementReadUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, announcementId := range announcementIds {
			// 检查公告是否存在
			announcement, err := s.GetAnnouncementById(ctx, announcementId)
			if err != nil || announcement == nil {
				// 跳过不存在的公告
				continue
			}

			// 检查是否已读
			readRecord, _ := daoctl.ScanWithError[sys_do.SysAnnouncementReadUser](
				sys_dao.SysAnnouncementReadUser.Ctx(ctx).
					Where(sys_do.SysAnnouncementReadUser{
						UserId:             userId,
						ReadAnnouncementId: gconv.String(announcementId),
					}),
			)

			if readRecord == nil {
				// 创建已读记录
				data := &sys_do.SysAnnouncementReadUser{
					Id:                 gconv.Int64(idgen.NextId()),
					UserId:             userId,
					ReadAnnouncementId: gconv.String(announcementId),
					ReadAt:             gtime.Now(),
					FlagRead:           sys_enum.Announcement.FlagRead.Readed.Code(),
				}

				affected, err := daoctl.InsertWithError(
					sys_dao.SysAnnouncementReadUser.Ctx(ctx).OmitNilData().Data(data),
				)

				if affected == 0 || err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_batch_mark_read_failed}", sys_dao.SysAnnouncementReadUser.Table())
				}

				// 更新公告阅读次数
				_, err = sys_dao.SysAnnouncement.Ctx(ctx).
					Where(sys_do.SysAnnouncement{Id: announcementId}).
					Increment(sys_dao.SysAnnouncement.Columns().ReadCount, 1)

				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_update_read_count_failed}", sys_dao.SysAnnouncement.Table())
				}
			} else if readRecord.FlagRead != sys_enum.Announcement.FlagRead.Readed.Code() {
				// 更新为已读状态
				affected, err := daoctl.UpdateWithError(
					sys_dao.SysAnnouncementReadUser.Ctx(ctx).
						Where(sys_do.SysAnnouncementReadUser{Id: readRecord.Id}).
						OmitNilData().Data(&sys_do.SysAnnouncementReadUser{
						ReadAt:   gtime.Now(),
						FlagRead: sys_enum.Announcement.FlagRead.Readed.Code(),
					}),
				)

				if affected == 0 || err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_batch_mark_read_failed}", sys_dao.SysAnnouncementReadUser.Table())
				}

				// 更新公告阅读次数
				_, err = sys_dao.SysAnnouncement.Ctx(ctx).
					Where(sys_do.SysAnnouncement{Id: announcementId}).
					Increment(sys_dao.SysAnnouncement.Columns().ReadCount, 1)

				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_update_read_count_failed}", sys_dao.SysAnnouncement.Table())
				}
			}
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetAnnouncementStatistics 获取公告统计信息
func (s *sAnnouncement) GetAnnouncementStatistics(ctx context.Context, announcementId int64) (*sys_model.SysAnnouncementStatistics, error) {
	// 检查公告是否存在
	announcement, err := s.GetAnnouncementById(ctx, announcementId)
	if err != nil || announcement == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_not_exists}", sys_dao.SysAnnouncement.Table())
	}

	// 获取有权限查看此公告的用户总数
	totalUserCount, _ := sys_dao.SysUser.Ctx(ctx).
		WhereOr(
			sys_dao.SysUser.Ctx(ctx).Builder().WhereNull(sys_dao.SysUser.Columns().Type),
			sys_dao.SysUser.Ctx(ctx).Builder().Where(sys_dao.SysUser.Columns().Type+" & ?", announcement.UserTypeScope),
		).Count()

	// 获取已读用户数量
	readCount, _ := sys_dao.SysAnnouncementReadUser.Ctx(ctx).
		Where(sys_do.SysAnnouncementReadUser{
			ReadAnnouncementId: gconv.String(announcementId),
			FlagRead:           sys_enum.Announcement.FlagRead.Readed.Code(),
		}).Count()

	// 获取确认用户数量
	confirmCount, _ := sys_dao.SysAnnouncementConfirm.Ctx(ctx).
		Where(sys_do.SysAnnouncementConfirm{
			AnnouncementId: announcementId,
		}).Count()

	// 构建统计信息
	statistics := &sys_model.SysAnnouncementStatistics{
		AnnouncementId: announcementId,
		ReadCount:      gconv.Int(readCount),
		UnreadCount:    gconv.Int(totalUserCount - readCount),
		ConfirmCount:   gconv.Int(confirmCount),
		TotalUserCount: gconv.Int(totalUserCount),
	}

	// 计算比率
	if totalUserCount > 0 {
		statistics.ReadRate = gconv.Float64(readCount) / gconv.Float64(totalUserCount) * 100
		statistics.ConfirmRate = gconv.Float64(confirmCount) / gconv.Float64(totalUserCount) * 100
	}

	return statistics, nil
}
