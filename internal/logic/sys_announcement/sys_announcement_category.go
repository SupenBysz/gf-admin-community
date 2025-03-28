package sys_announcement

import (
	"context"
	"fmt"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

// GetCategoryById 根据ID获取公告分类
func (s *sAnnouncement) GetCategoryById(ctx context.Context, id int64) (*sys_model.SysAnnouncementCategoryRes, error) {
	result, err := daoctl.GetByIdWithError[sys_model.SysAnnouncementCategoryRes](sys_dao.SysAnnouncementCategory.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_query_by_id_failed}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 获取该分类下的公告数量
	count, _ := sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{CategoryId: id}).Count()
	result.AnnouncementCount = count

	return result, nil
}

// CreateCategory 创建公告分类
func (s *sAnnouncement) CreateCategory(ctx context.Context, info *sys_model.SysAnnouncementCategory, unionMainId int64, userId int64) (*sys_model.SysAnnouncementCategoryRes, error) {
	// 检查编码是否已存在
	count, err := sys_dao.SysAnnouncementCategory.Ctx(ctx).
		Where(sys_do.SysAnnouncementCategory{Code: info.Code, UnionMainId: unionMainId}).
		Count()
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_check_code_failed}", sys_dao.SysAnnouncementCategory.Table())
	}
	if count > 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_category_code_exists}", sys_dao.SysAnnouncementCategory.Table())
	}

	data := kconv.Struct(info, &sys_do.SysAnnouncementCategory{})
	data.Id = idgen.NextId()
	data.UnionMainId = unionMainId
	data.CreatedBy = userId
	data.CreatedAt = gtime.Now()

	affected, err := daoctl.InsertWithError(sys_dao.SysAnnouncementCategory.Ctx(ctx).OmitNilData().Data(data))
	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_add_failed}", sys_dao.SysAnnouncementCategory.Table())
	}

	return s.GetCategoryById(ctx, gconv.Int64(data.Id))
}

// UpdateCategory 更新公告分类
func (s *sAnnouncement) UpdateCategory(ctx context.Context, info *sys_model.SysAnnouncementCategory, unionMainId int64, userId int64) (*sys_model.SysAnnouncementCategoryRes, error) {
	if info.Id == 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_category_id_required}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 检查分类是否存在
	category, err := s.GetCategoryById(ctx, info.Id)
	if err != nil || category == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 检查主体是否匹配
	if category.UnionMainId != unionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_cross_subject_modification_forbidden}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 检查编码是否已存在（非当前记录）
	if info.Code != category.Code {
		count, _ := sys_dao.SysAnnouncementCategory.Ctx(ctx).
			Where(sys_do.SysAnnouncementCategory{Code: info.Code, UnionMainId: unionMainId}).
			WhereNot(sys_dao.SysAnnouncementCategory.Columns().Id, info.Id).
			Count()
		if count > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_category_code_exists}", sys_dao.SysAnnouncementCategory.Table())
		}
	}

	data := kconv.Struct(info, &sys_do.SysAnnouncementCategory{})
	data.UpdatedBy = userId
	data.UpdatedAt = gtime.Now()

	// ID不能更新
	id := data.Id
	data.Id = 0

	affected, err := daoctl.UpdateWithError(sys_dao.SysAnnouncementCategory.Ctx(ctx).
		Where(sys_do.SysAnnouncementCategory{Id: id}).
		OmitNilData().Data(data))

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_update_failed}", sys_dao.SysAnnouncementCategory.Table())
	}

	return s.GetCategoryById(ctx, gconv.Int64(id))
}

// DeleteCategory 删除公告分类
func (s *sAnnouncement) DeleteCategory(ctx context.Context, id int64, unionMainId int64) (bool, error) {
	// 检查分类是否存在
	category, err := s.GetCategoryById(ctx, id)
	if err != nil || category == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 检查主体是否匹配
	if category.UnionMainId != unionMainId {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_cross_subject_modification_forbidden}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 检查分类是否被使用
	isUsed, count, err := s.CheckCategoryUsage(ctx, id)
	if err != nil {
		return false, err
	}

	if isUsed {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, fmt.Sprintf("{#error_announcement_category_in_use}:%d", count), sys_dao.SysAnnouncementCategory.Table())
	}

	affected, err := daoctl.DeleteWithError(sys_dao.SysAnnouncementCategory.Ctx(ctx).
		Where(sys_do.SysAnnouncementCategory{Id: id, UnionMainId: unionMainId}))

	if affected == 0 || err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_delete_failed}", sys_dao.SysAnnouncementCategory.Table())
	}

	return true, nil
}

// QueryCategory 查询公告分类列表
func (s *sAnnouncement) QueryCategory(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysAnnouncementCategoryListRes, error) {
	result, err := daoctl.Query[sys_model.SysAnnouncementCategoryRes](
		sys_dao.SysAnnouncementCategory.Ctx(ctx).OrderDesc(sys_dao.SysAnnouncementCategory.Columns().Sort),
		params,
		isExport,
	)

	if err != nil {
		return &sys_model.SysAnnouncementCategoryListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_category_list_query_failed}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 获取每个分类下的公告数量
	for i, item := range result.Records {
		count, _ := sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{CategoryId: item.Id}).Count()
		result.Records[i].AnnouncementCount = count
	}

	return (*sys_model.SysAnnouncementCategoryListRes)(result), nil
}

// CheckCategoryUsage 检查分类是否被使用
func (s *sAnnouncement) CheckCategoryUsage(ctx context.Context, categoryId int64) (bool, int, error) {
	count, err := sys_dao.SysAnnouncement.Ctx(ctx).Where(sys_do.SysAnnouncement{CategoryId: categoryId}).Count()
	if err != nil {
		return false, 0, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_check_category_usage_failed}", sys_dao.SysAnnouncementCategory.Table())
	}

	return count > 0, gconv.Int(count), nil
}
