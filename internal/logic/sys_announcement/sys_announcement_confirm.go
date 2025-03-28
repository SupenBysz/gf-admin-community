package sys_announcement

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

// ConfirmAnnouncement 确认公告
func (s *sAnnouncement) ConfirmAnnouncement(ctx context.Context, announcementId int64, userId int64, comment string) (bool, error) {
	// 判断公告是否存在
	announcement, err := s.GetAnnouncementById(ctx, announcementId)
	if err != nil || announcement == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_not_exists}", sys_dao.SysAnnouncement.Table())
	}

	// 判断公告是否需要确认
	if announcement.ConfirmRequired != 1 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_confirm_not_required}", sys_dao.SysAnnouncement.Table())
	}

	// 判断是否已确认
	confirmed, err := s.IsAnnouncementConfirmed(ctx, announcementId, userId)
	if err != nil {
		return false, err
	}

	if confirmed {
		return true, nil // 已确认过，直接返回成功
	}

	// 开始事务处理
	err = sys_dao.SysAnnouncementConfirm.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 添加确认记录
		data := &sys_do.SysAnnouncementConfirm{
			Id:             idgen.NextId(),
			UserId:         userId,
			AnnouncementId: announcementId,
			ConfirmAt:      gtime.Now(),
			ConfirmComment: comment,
			CreatedBy:      userId,
			CreatedAt:      gtime.Now(),
		}

		affected, err := daoctl.InsertWithError(sys_dao.SysAnnouncementConfirm.Ctx(ctx).OmitNilData().Data(data))
		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_confirm_failed}", sys_dao.SysAnnouncementConfirm.Table())
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// IsAnnouncementConfirmed 检查公告是否已确认
func (s *sAnnouncement) IsAnnouncementConfirmed(ctx context.Context, announcementId int64, userId int64) (bool, error) {
	count, err := sys_dao.SysAnnouncementConfirm.Ctx(ctx).
		Where(sys_do.SysAnnouncementConfirm{
			AnnouncementId: announcementId,
			UserId:         userId,
		}).Count()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_confirm_check_failed}", sys_dao.SysAnnouncementConfirm.Table())
	}

	return count > 0, nil
}

// GetAnnouncementConfirmInfo 获取公告确认信息
func (s *sAnnouncement) GetAnnouncementConfirmInfo(ctx context.Context, announcementId int64, userId int64) (*sys_model.SysAnnouncementConfirmRes, error) {
	// 检查公告是否存在
	announcement, err := s.GetAnnouncementById(ctx, announcementId)
	if err != nil || announcement == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_not_exists}", sys_dao.SysAnnouncement.Table())
	}

	// 获取确认记录
	result, err := daoctl.ScanWithError[sys_model.SysAnnouncementConfirmRes](
		sys_dao.SysAnnouncementConfirm.Ctx(ctx).
			Where(sys_do.SysAnnouncementConfirm{
				AnnouncementId: announcementId,
				UserId:         userId,
			}),
	)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_confirm_info_query_failed}", sys_dao.SysAnnouncementConfirm.Table())
	}

	if result == nil {
		return nil, nil // 未确认
	}

	// 获取用户信息
	sysUser, _ := sys_service.SysUser().GetSysUserById(ctx, userId)
	if sysUser != nil {
		result.UserName = sysUser.Username
	}

	return result, nil
}

// QueryAnnouncementConfirms 查询公告确认记录列表
func (s *sAnnouncement) QueryAnnouncementConfirms(ctx context.Context, announcementId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementConfirmListRes, error) {
	// 检查公告是否存在
	announcement, err := s.GetAnnouncementById(ctx, announcementId)
	if err != nil || announcement == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_not_exists}", sys_dao.SysAnnouncement.Table())
	}

	// 添加公告ID过滤条件
	if params == nil {
		params = &base_model.SearchParams{}
	}

	params.Filter = append(params.Filter, base_model.FilterInfo{
		Field: sys_dao.SysAnnouncementConfirm.Columns().AnnouncementId,
		Where: "=",
		Value: announcementId,
	})

	result, err := daoctl.Query[sys_model.SysAnnouncementConfirmRes](
		sys_dao.SysAnnouncementConfirm.Ctx(ctx).OrderDesc(sys_dao.SysAnnouncementConfirm.Columns().ConfirmAt),
		params,
		isExport,
	)

	if err != nil {
		return &sys_model.SysAnnouncementConfirmListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_confirm_list_query_failed}", sys_dao.SysAnnouncementConfirm.Table())
	}

	// 填充用户信息
	for i, item := range result.Records {
		sysUser, _ := sys_service.SysUser().GetSysUserById(ctx, item.UserId)
		if sysUser != nil {
			result.Records[i].UserName = sysUser.Username
		}
	}

	return (*sys_model.SysAnnouncementConfirmListRes)(result), nil
}
