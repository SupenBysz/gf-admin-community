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
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

/*
注意：后续公告需要审核，再加入审核状态的字段， audit_state
*/

type sAnnouncement struct {
	//AnnouncementHook base_hook.BaseHook[sys_enum.AnnouncementType, sys_hook.AnnouncementTypeHookFunc]
}

func NewAnnouncement() sys_service.IAnnouncement {
	return &sAnnouncement{}
}

func init() {
	sys_service.RegisterAnnouncement(NewAnnouncement())
}

// checkPublic 检查公告是否已发布
func (s *sAnnouncement) checkPublic(ctx context.Context, result *sys_model.SysAnnouncementRes) {
	/*
		发布时间：19:20
		过期时间：19:30
	*/
	/*
		现在时间：19:07
		state = 待发布

		现在时间：19:21
		state = 已发布

		现在时间：19:31
		state = 已过期
	*/

	// 草稿 ---- 默认添加就是草稿

	// 待发布 ---- 审核通过就是待发布 （后续审核功能接入后支持）

	// 已发布 ---- 发布时间大于现在时间，并且状态小于已发布，则修改状态为已发布
	if result.PublicAt != nil && gtime.Now().After(result.PublicAt) && result.State < sys_enum.Announcement.State.Published.Code() { // 已发布
		_, err := daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{Id: result.Id}), sys_do.SysAnnouncement{State: sys_enum.Announcement.State.Published.Code()})
		if err != nil {
			_ = sys_service.SysLogs().ErrorSimple(ctx, err, "公告状态修改为已发布失败", sys_dao.SysAnnouncement.Table())
		}
	}

	// 已过期 ---- 发布时间小于现在时间，并且状态小于已过期，则修改状态为已过期
	if result.ExpireAt != nil && gtime.Now().After(result.ExpireAt) && result.State < sys_enum.Announcement.State.Expired.Code() { // 已过期
		_, err := daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{Id: result.Id}), sys_do.SysAnnouncement{State: sys_enum.Announcement.State.Expired.Code()})
		if err != nil {
			_ = sys_service.SysLogs().ErrorSimple(ctx, err, "公告状态修改为已发布失败", sys_dao.SysAnnouncement.Table())
		}
	}

	// 已移除 （需要手动触发）

}

// GetAnnouncementById 根据id查询公告｜信息
func (s *sAnnouncement) GetAnnouncementById(ctx context.Context, id int64, userId ...int64) (*sys_model.SysAnnouncementRes, error) {

	result, err := daoctl.GetByIdWithError[sys_entity.SysAnnouncement](sys_dao.SysAnnouncement.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询公告失败", sys_dao.SysAnnouncement.Table())
	}

	// TODO 增加公告的已读用户记录
	if len(userId) > 0 && userId[0] != 0 {
		newCtx := context.Background()
		go func(ctx context.Context) {
			for _, uId := range userId {
				_, err = s.MarkRead(ctx, id, uId)
			}
		}(newCtx)
	}

	s.checkPublic(ctx, (*sys_model.SysAnnouncementRes)(result))

	return (*sys_model.SysAnnouncementRes)(result), nil
}

// CreateAnnouncement 添加公告｜信息
func (s *sAnnouncement) CreateAnnouncement(ctx context.Context, info *sys_model.SysAnnouncement, unionMainId, userId int64) (*sys_model.SysAnnouncementRes, error) {
	data := kconv.Struct(info, &sys_do.SysAnnouncement{})

	err := sys_dao.SysAnnouncement.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		data.Id = idgen.NextId()
		data.UnionMainId = unionMainId
		data.CreatedBy = userId
		if info.ExtDataJson == "" {
			data.ExtDataJson = nil
		}

		if info.State == 0 { // 没填就是草稿
			data.State = sys_enum.Announcement.State.Draft.Code()
		}

		data.CreatedAt = gtime.Now()
		affected, err := daoctl.InsertWithError(sys_dao.SysAnnouncement.Ctx(ctx).OmitNilData().Data(data))

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "添加公告失败", sys_dao.SysAnnouncement.Table())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetAnnouncementById(ctx, gconv.Int64(data.Id))
}

// UpdateAnnouncement 编辑公告｜信息
func (s *sAnnouncement) UpdateAnnouncement(ctx context.Context, info *sys_model.UpdateSysAnnouncement, unionMainId int64, userId int64) (*sys_model.SysAnnouncementRes, error) {
	var err error

	// 判断公告是否存在
	announcement, err := s.GetAnnouncementById(ctx, *info.Id)
	if err != nil || announcement == nil {
		return nil, err
	}

	// 判断是否是本主体公告
	if unionMainId != announcement.UnionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "禁止跨主体修改公告信息", sys_dao.SysAnnouncement.Table())
	}

	/*
		判断是否可以编辑：
			不能编辑：
				- 公告的状态不是草稿
			能编辑：
				- 公告未发布则允许编辑

	*/
	if announcement.State == sys_enum.Announcement.State.Published.Code() || announcement.State == sys_enum.Announcement.State.Expired.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "公示中和已过期的公告禁止编辑", sys_dao.SysAnnouncement.Table())
	}

	//if gtime.Now().After(announcement.PublicAt) && announcement.State != sys_enum.Announcement.State.Draft.Code() {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "只有未发送的公告支持编辑", sys_dao.SysAnnouncement.Table())
	//}

	data := kconv.Struct(info, &sys_do.SysAnnouncement{})
	data.Id = nil
	data.UpdatedBy = userId
	data.UpdatedAt = gtime.Now()

	if *info.ExtDataJson == "" {
		data.ExtDataJson = nil
	}

	err = sys_dao.SysAnnouncement.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		affected, err := daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(
			sys_do.SysAnnouncement{
				Id: info.Id,
			},
		).OmitNilData().Data(&data))

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "公告修改失败", sys_dao.SysAnnouncement.Table())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetAnnouncementById(ctx, *info.Id)
}

// DeleteAnnouncement 删除公告
func (s *sAnnouncement) DeleteAnnouncement(ctx context.Context, id int64, unionMainId, userId int64) (bool, error) {
	// 判断公告是否存在
	announcement, err := s.GetAnnouncementById(ctx, id)
	if err != nil || announcement == nil {
		return false, err
	}
	err = sys_dao.SysAnnouncement.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 1、删除
		_, err = daoctl.DeleteWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{Id: id, UnionMainId: unionMainId}))
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "公告删除失败", sys_dao.SysAnnouncement.Table())
		}

		// 2、设置删除用户
		_, err = daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(
			sys_do.SysAnnouncement{
				Id:          id,
				UnionMainId: unionMainId,
			},
		).OmitNilData().Data(sys_do.SysAnnouncement{DeletedAt: gtime.Now(), DeletedBy: userId}))
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "设置删除公告用户失败", sys_dao.SysAnnouncement.Table())
		}

		// 3、删除公告关联的已读用户记录
		affected, err := daoctl.DeleteWithError(sys_dao.SysAnnouncementReadUser.Ctx(ctx).Where(sys_do.SysAnnouncementReadUser{ReadAnnouncementId: id}))
		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "删除公告关联的已读用户记录失败", sys_dao.SysAnnouncementReadUser.Table())
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return err == nil, err
}

// QueryAnnouncement 查询公告｜列表
func (s *sAnnouncement) QueryAnnouncement(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementListRes, error) {
	res, err := daoctl.Query[sys_model.SysAnnouncementRes](sys_dao.SysAnnouncement.Ctx(ctx), params, isExport)

	if err != nil {
		return &sys_model.SysAnnouncementListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "公告列表查询失败", sys_dao.SysAnnouncement.Table())
	}

	for _, record := range res.Records {
		s.checkPublic(ctx, &record)
	}

	return (*sys_model.SysAnnouncementListRes)(res), nil
}
