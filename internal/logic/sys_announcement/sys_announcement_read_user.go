package sys_announcement

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/samber/lo"
)

// MarkRead 标记已读｜公告
func (s *sAnnouncement) MarkRead(ctx context.Context, announcementId, userId int64) (bool, error) {
	// 是否已存在已读记录
	info, _ := daoctl.ScanWithError[sys_entity.SysAnnouncementReadUser](sys_dao.SysAnnouncementReadUser.Ctx(ctx).
		Where(sys_do.SysAnnouncementReadUser{ReadAnnouncementId: announcementId, UserId: userId}))

	if info == nil {
		data := &sys_do.SysAnnouncementReadUser{
			Id:                 idgen.NextId(),
			UserId:             userId,
			ReadAnnouncementId: announcementId,
			ReadAt:             gtime.Now(),
			ExtDataJson:        nil,
			FlagRead:           sys_enum.Announcement.FlagRead.Readed.Code(),
		}

		affected, err := daoctl.InsertWithError(sys_dao.SysAnnouncementReadUser.Ctx(ctx).OmitNilData().Data(&data))
		if affected == 0 || err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "公告标记已读失败", sys_dao.SysAnnouncementReadUser.Table())
		}

	} else if info.FlagRead != sys_enum.Announcement.FlagRead.Readed.Code() && info.UserId != userId {
		data := &sys_do.SysAnnouncementReadUser{
			ReadAt:   gtime.Now(),
			FlagRead: sys_enum.Announcement.FlagRead.Readed.Code(),
		}

		affected, err := daoctl.UpdateWithError(sys_dao.SysAnnouncementReadUser.Ctx(ctx).Where(sys_do.SysAnnouncementReadUser{Id: info.Id}).OmitNilData().Data(&data))
		if affected == 0 || err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "公告标记已读失败", sys_dao.SysAnnouncementReadUser.Table())
		}
	}

	return true, nil
}

// MarkUnRead 标记未读｜公告
func (s *sAnnouncement) MarkUnRead(ctx context.Context, announcementId, userId int64) (bool, error) {
	// 是否存在公告
	announcement, _ := s.GetAnnouncementById(ctx, announcementId)
	if announcement == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "公告不存在", sys_dao.SysAnnouncement.Table())
	}

	// 是否已存在已读记录
	info, _ := daoctl.ScanWithError[sys_entity.SysAnnouncementReadUser](sys_dao.SysAnnouncementReadUser.Ctx(ctx).
		Where(sys_do.SysAnnouncementReadUser{ReadAnnouncementId: announcementId, UserId: userId}))

	if info == nil || info.FlagRead == sys_enum.Announcement.FlagRead.UnRead.Code() { // 未读
		return true, nil
	}

	if info.FlagRead == sys_enum.Announcement.FlagRead.Readed.Code() {
		data := &sys_do.SysAnnouncementReadUser{
			FlagRead: sys_enum.Announcement.FlagRead.UnRead.Code(),
		}

		affected, err := daoctl.UpdateWithError(sys_dao.SysAnnouncementReadUser.Ctx(ctx).Where(sys_do.SysAnnouncementReadUser{Id: info.Id}).OmitNilData().Data(&data))
		if affected == 0 || err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "公告标记未读失败", sys_dao.SysAnnouncementReadUser.Table())
		}
	}

	return true, nil

}

// 查询和我相关的公告列表
func (s *sAnnouncement) queryMyAnnouncementList(ctx context.Context, userId int64, unionMainId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementListRes, error) {
	res := sys_model.SysAnnouncementListRes{Records: make([]sys_model.SysAnnouncementRes, 0)}

	sysUserId, _ := sys_service.SysUser().GetSysUserById(ctx, userId)
	if sysUserId == nil {
		return &res, sys_service.SysLogs().ErrorSimple(ctx, nil, "用户不存在", sys_dao.SysUser.Table())
	}

	params.Filter = append(params.Filter,
		base_model.FilterInfo{ // 公告主体
			Field: sys_dao.SysAnnouncement.Columns().UnionMainId,
			Where: "in",
			Value: []int64{0, unionMainId}, // 不区分 本主体和平台的公告
		},
		base_model.FilterInfo{ // 公告用户范围
			Field: sys_dao.SysAnnouncement.Columns().UserTypeScope + " & " + gconv.String(sysUserId.Type),
			Where: "=",
			Value: sysUserId.Type,
		},
		base_model.FilterInfo{ // 已发布
			Field: sys_dao.SysAnnouncement.Columns().State,
			Where: "=",
			Value: sys_enum.Announcement.State.Published.Code(),
		},
		base_model.FilterInfo{ // 已达到发布时间
			Field: sys_dao.SysAnnouncement.Columns().PublicAt,
			Where: "<=",
			Value: gtime.Now(),
		},
		base_model.FilterInfo{ // 在过期时间之前
			Field: sys_dao.SysAnnouncement.Columns().ExpireAt,
			Where: ">=",
			Value: gtime.Now(),
		},
	)

	// 找到和用户相关的所有公告
	announcementList, err := s.QueryAnnouncement(ctx, params, isExport)
	if announcementList == nil || len(announcementList.Records) <= 0 || err != nil {
		return &res, nil
	}

	return announcementList, err
}

// HasUnReadAnnouncement 获取未读公告数量
func (s *sAnnouncement) HasUnReadAnnouncement(ctx context.Context, userId int64, unionMainId int64) (int, error) {

	// 统计未读数量
	ret := 0
	announcementIds := make([]int64, 0)
	//for _, item := range announcementList.Records {
	//	// 是否已存在已读记录
	//	result, _ := daoctl.ScanWithError[sys_entity.SysAnnouncementReadUser](sys_dao.SysAnnouncementReadUser.Ctx(ctx).
	//		Where(sys_do.SysAnnouncementReadUser{ReadAnnouncementId: item.Id, UserId: userId, FlagRead: sys_enum.Announcement.FlagRead.Readed.Code()}))
	//	if result == nil {
	//		ret++
	//	}
	//}

	announcementList, err := s.queryMyAnnouncementList(ctx, userId, unionMainId, &base_model.SearchParams{}, true)
	if announcementList == nil || len(announcementList.Records) <= 0 || err != nil {
		return 0, nil
	}

	for _, item := range announcementList.Records {
		announcementIds = append(announcementIds, item.Id)
	}

	// 已读数量
	readCount, _ := sys_dao.SysAnnouncementReadUser.Ctx(ctx).
		Where(sys_do.SysAnnouncementReadUser{UserId: userId, FlagRead: sys_enum.Announcement.FlagRead.Readed.Code()}).
		WhereIn(sys_dao.SysAnnouncementReadUser.Columns().ReadAnnouncementId, announcementIds).Count()

	// 未读数量 = 总量 - 已读数量
	ret = len(announcementList.Records) - readCount

	return ret, err
}

// QueryAnnouncementListByUser 查询用户的公告｜列表 （qType：0未读，1已读、2全部）
func (s *sAnnouncement) QueryAnnouncementListByUser(ctx context.Context, userId int64, unionMainId int64, qType int, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementListRes, error) {
	announcementList := &sys_model.SysAnnouncementListRes{}

	// 已读公告
	announcementIds, _ := sys_dao.SysAnnouncementReadUser.Ctx(ctx).
		Where(sys_do.SysAnnouncementReadUser{
			UserId:   userId,
			FlagRead: sys_enum.Announcement.FlagRead.Readed.Code(),
		}).
		Fields(sys_dao.SysAnnouncementReadUser.Columns().ReadAnnouncementId).All()

	// 已读的ids
	readIds := make([]int64, 0)
	for _, id := range announcementIds {
		readIds = append(readIds, gconv.Int64(id[sys_dao.SysAnnouncementReadUser.Columns().ReadAnnouncementId]))
	}

	if qType == 0 { // 未读
		// 所有公告
		allList, _ := s.queryMyAnnouncementList(ctx, userId, unionMainId, &base_model.SearchParams{}, true)

		// 未读的ids
		unreadIds := make([]int64, 0)
		for _, item := range allList.Records {
			if !lo.Contains(readIds, item.Id) {
				unreadIds = append(unreadIds, item.Id)
			}
		}

		params.Filter = append(params.Filter, base_model.FilterInfo{
			Field: sys_dao.SysAnnouncement.Columns().Id,
			Where: "in",
			Value: unreadIds,
		})

		//announcementList, _ = s.QueryAnnouncement(ctx, params, isExport)
		announcementList, _ = s.queryMyAnnouncementList(ctx, userId, unionMainId, params, isExport)
	} else if qType == 1 { // 已读
		params.Filter = append(params.Filter, base_model.FilterInfo{
			Field: sys_dao.SysAnnouncement.Columns().Id,
			Where: "in",
			Value: announcementIds,
		})

		//announcementList, _ = s.QueryAnnouncement(ctx, params, isExport)
		announcementList, _ = s.queryMyAnnouncementList(ctx, userId, unionMainId, params, isExport)
	} else if qType == 2 { // 全部
		//announcementList, _ = s.queryMyAnnouncementList(ctx, userId, unionMainId)
		announcementList, _ = s.queryMyAnnouncementList(ctx, userId, unionMainId, params, isExport)
	}

	for i, record := range announcementList.Records {
		res := announcementList.Records[i]

		if lo.Contains(readIds, record.Id) {
			res.ReadState = sys_enum.Announcement.FlagRead.Readed.Code()
		} else {
			res.ReadState = sys_enum.Announcement.FlagRead.UnRead.Code()
		}

		announcementList.Records[i] = res
	}

	return announcementList, nil
}

// QueryAnnouncementReadUserList 查询公告的已读用户｜列表
//func (s *sAnnouncement) QueryAnnouncementReadUserList(ctx context.Context, announcementId int64) (*sys_model.SysAnnouncementReadUserListRes, error) {
//
//	return ret, err
//}

// RevokedAnnouncement 撤销公告
func (s *sAnnouncement) RevokedAnnouncement(ctx context.Context, announcementId int64) (bool, error) {
	announcement, err := s.GetAnnouncementById(ctx, announcementId)
	if err != nil {
		return false, err
	}

	if announcement.State&sys_enum.Announcement.State.Revoked.Code() == sys_enum.Announcement.State.Revoked.Code() {
		return true, err
	}

	return s.SetAnnouncementState(ctx, announcementId, sys_enum.Announcement.State.Revoked.Code())
}

// SetAnnouncementState 设置公告状态
func (s *sAnnouncement) SetAnnouncementState(ctx context.Context, announcementId int64, state int) (bool, error) {

	count, err := sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{Id: announcementId}).Count()
	if err != nil || count == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "公告不存在", sys_dao.SysAnnouncement.Table())
	}

	affected, err := daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).OmitNilData().Data(
		sys_do.SysAnnouncement{
			State: state,
		}).Where(sys_dao.SysAnnouncement.Columns().Id, announcementId))
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "公告状态修改失败", sys_dao.SysAnnouncement.Table())
	}

	return affected > 0, nil
}
