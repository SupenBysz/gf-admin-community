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
	"github.com/gogf/gf/v2/text/gstr"
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
func (s *sAnnouncement) checkPublic(ctx context.Context, result *sys_model.SysAnnouncementRes) *sys_model.SysAnnouncementRes {
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
			_ = sys_service.SysLogs().ErrorSimple(ctx, err, "error_announcement_status_update_failed", sys_dao.SysAnnouncement.Table())
		}
	}

	// 已过期 ---- 发布时间小于现在时间，并且状态小于已过期，则修改状态为已过期
	if result.ExpireAt != nil && gtime.Now().After(result.ExpireAt) && result.State < sys_enum.Announcement.State.Expired.Code() { // 已过期
		_, err := daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{Id: result.Id}), sys_do.SysAnnouncement{State: sys_enum.Announcement.State.Expired.Code()})
		if err != nil {
			_ = sys_service.SysLogs().ErrorSimple(ctx, err, "error_announcement_status_update_failed", sys_dao.SysAnnouncement.Table())
		}
	}

	// 已移除 （需要手动触发）

	return result
}

// GetAnnouncementById 根据id查询公告｜信息
func (s *sAnnouncement) GetAnnouncementById(ctx context.Context, id int64, userId ...int64) (*sys_model.SysAnnouncementRes, error) {

	result, err := daoctl.GetByIdWithError[sys_model.SysAnnouncementRes](sys_dao.SysAnnouncement.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_query_by_id_failed}", sys_dao.SysAnnouncement.Table())
	}

	// 获取分类名称
	if result.CategoryId > 0 {
		category, _ := s.GetCategoryById(ctx, result.CategoryId)
		if category != nil {
			result.CategoryName = category.Name
		}
	}

	// 获取确认人数
	confirmCount, _ := sys_dao.SysAnnouncementConfirm.Ctx(ctx).Where(sys_do.SysAnnouncementConfirm{AnnouncementId: id}).Count()
	result.ConfirmCount = gconv.Int(confirmCount)

	// TODO 增加公告的已读用户记录
	if len(userId) > 0 && userId[0] != 0 {
		// 获取确认状态
		isConfirmed, _ := s.IsAnnouncementConfirmed(ctx, id, userId[0])
		if isConfirmed {
			result.ConfirmStatus = 1
		}

		newCtx := context.Background()
		go func(ctx context.Context) {
			for _, uId := range userId {
				_, err = s.MarkRead(ctx, id, uId)
			}
		}(newCtx)
	}

	return s.checkPublic(ctx, result), nil
}

// CreateAnnouncement 添加公告｜信息
func (s *sAnnouncement) CreateAnnouncement(ctx context.Context, info *sys_model.SysAnnouncement, unionMainId, userId int64) (*sys_model.SysAnnouncementRes, error) {
	data := kconv.Struct(info, &sys_do.SysAnnouncement{})

	// 检查分类是否存在
	if info.CategoryId > 0 {
		category, err := s.GetCategoryById(ctx, info.CategoryId)
		if err != nil || category == nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
		}
	}

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

		// 设置默认值
		if info.Priority == 0 {
			data.Priority = 1 // 默认普通优先级
		}

		// 设置阅读次数初始值
		data.ReadCount = 0

		data.CreatedAt = gtime.Now()
		affected, err := daoctl.InsertWithError(sys_dao.SysAnnouncement.Ctx(ctx).OmitNilData().Data(data))

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_add_failed}", sys_dao.SysAnnouncement.Table())
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_cross_subject_modification_forbidden}", sys_dao.SysAnnouncement.Table())
	}

	// 如果有分类ID，检查分类是否存在
	if info.CategoryId != nil && *info.CategoryId > 0 {
		category, err := s.GetCategoryById(ctx, *info.CategoryId)
		if err != nil || category == nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
		}
	}

	/*
		判断是否可以编辑：
			不能编辑：
				- 公告的状态不是草稿
			能编辑：
				- 公告未发布则允许编辑

	*/
	if announcement.State == sys_enum.Announcement.State.Published.Code() || announcement.State == sys_enum.Announcement.State.Expired.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_public_expired_announcement_edit_forbidden}", sys_dao.SysAnnouncement.Table())
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
			return sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_update_failed}", sys_dao.SysAnnouncement.Table())
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
			return sys_service.SysLogs().ErrorSimple(ctx, err, "error_announcement_delete_failed", sys_dao.SysAnnouncement.Table())
		}

		// 2、设置删除用户
		_, err = daoctl.UpdateWithError(sys_dao.SysAnnouncement.Ctx(ctx).Where(
			sys_do.SysAnnouncement{
				Id:          id,
				UnionMainId: unionMainId,
			},
		).OmitNilData().Data(sys_do.SysAnnouncement{DeletedAt: gtime.Now(), DeletedBy: userId}))
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "error_announcement_delete_user_set_failed", sys_dao.SysAnnouncement.Table())
		}

		// 3、删除公告关联的已读用户记录
		affected, err := daoctl.DeleteWithError(sys_dao.SysAnnouncementReadUser.Ctx(ctx).Where(sys_do.SysAnnouncementReadUser{ReadAnnouncementId: id}))
		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "error_announcement_delete_read_user_failed", sys_dao.SysAnnouncementReadUser.Table())
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
	filter := make([]base_model.FilterInfo, 0)

	m := sys_dao.SysAnnouncement.Ctx(ctx)

	if params != nil {
		for _, info := range params.Filter {
			if gstr.ToUpper(info.Field) == gstr.ToUpper(sys_dao.SysAnnouncement.Columns().ExpireAt) {
				m = m.Where(m.Builder().WhereGTE(info.Field, info.Value).WhereOrNull(sys_dao.SysAnnouncement.Columns().ExpireAt))
			} else {
				filter = append(filter, info)
			}
		}
		params.Filter = filter
	}

	// 添加排序：优先置顶，然后按优先级，最后按发布时间
	m = m.Order(sys_dao.SysAnnouncement.Columns().IsPinned + " DESC, " +
		sys_dao.SysAnnouncement.Columns().Priority + " DESC, " +
		sys_dao.SysAnnouncement.Columns().PublicAt + " DESC")

	res, err := daoctl.Query[sys_model.SysAnnouncementRes](m, params, isExport)

	if err != nil {
		return &sys_model.SysAnnouncementListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_list_query_failed}", sys_dao.SysAnnouncement.Table())
	}

	for i, record := range res.Records {
		s.checkPublic(ctx, &record)

		// 获取分类名称
		if record.CategoryId > 0 {
			category, _ := s.GetCategoryById(ctx, record.CategoryId)
			if category != nil {
				res.Records[i].CategoryName = category.Name
			}
		}

		// 获取确认人数
		confirmCount, _ := sys_dao.SysAnnouncementConfirm.Ctx(ctx).Where(sys_do.SysAnnouncementConfirm{AnnouncementId: record.Id}).Count()
		res.Records[i].ConfirmCount = gconv.Int(confirmCount)
	}

	return (*sys_model.SysAnnouncementListRes)(res), nil
}
