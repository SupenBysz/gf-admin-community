package sys_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	v1 "github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
)

type cSysAnnouncement struct{}

var SysAnnouncement cSysAnnouncement

// GetAnnouncementById 根据id查询公告｜信息
func (c *cSysAnnouncement) GetAnnouncementById(ctx context.Context, req *v1.GetAnnouncementByIdReq) (*sys_model.SysAnnouncementRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().GetAnnouncementById(ctx, req.Id, user.UnionMainId)

	return ret, err
}

// CreateAnnouncement 添加公告｜信息
func (c *cSysAnnouncement) CreateAnnouncement(ctx context.Context, req *v1.CreateAnnouncementReq) (*sys_model.SysAnnouncementRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.Create); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 验证用户是否有权限设置特定类型用户可见
	// 如果用户不是超级管理员或高级管理角色，限制他们只能设置与自己相同或更低权限的用户可见
	// 由于用户类型和用户权限可能使用不同的数据格式，这里做简单判断
	// 超级管理员判断可以使用user.Id==1的简单判断
	if user.Id != 1 { // 非超级管理员
		// 用户类型使用位运算存储，通常较大值表示更高权限
		// 确保用户不能设置比自己权限更高的用户范围
		userType := gconv.Int64(user.Type)
		if req.SysAnnouncement.UserTypeScope > userType {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_no_permission_set_user_type}", sys_dao.SysAnnouncement.Table())
		}
	}

	ret, err := sys_service.Announcement().CreateAnnouncement(ctx, &req.SysAnnouncement, user.UnionMainId, user.Id)

	return ret, err
}

// UpdateAnnouncement 编辑公告｜信息
func (c *cSysAnnouncement) UpdateAnnouncement(ctx context.Context, req *v1.UpdateAnnouncementReq) (*sys_model.SysAnnouncementRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.Update); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 首先获取原公告信息
	announcement, err := sys_service.Announcement().GetAnnouncementById(ctx, req.Id, user.UnionMainId)
	if err != nil || announcement == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "{#error_announcement_not_exists}", sys_dao.SysAnnouncement.Table())
	}

	// 验证用户是否有权限修改此公告的用户类型范围
	// 超级管理员判断可以使用user.Id==1的简单判断
	if user.Id != 1 { // 非超级管理员
		// 如果用户类型范围发生了变化
		if req.UpdateSysAnnouncement.UserTypeScope != nil {
			userType := gconv.Int64(user.Type)
			announcementScopeType := gconv.Int64(announcement.UserTypeScope)

			// 检查是否修改了用户类型范围
			if *req.UpdateSysAnnouncement.UserTypeScope != announcementScopeType {
				// 非超级管理员只能修改为自己类型及以下类型的用户范围
				if *req.UpdateSysAnnouncement.UserTypeScope > userType {
					return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_no_permission_set_user_type}", sys_dao.SysAnnouncement.Table())
				}
			}
		}
	}

	ret, err := sys_service.Announcement().UpdateAnnouncement(ctx, &req.UpdateSysAnnouncement, user.UnionMainId, user.Id)

	return ret, err
}

// DeleteAnnouncement 删除公告
func (c *cSysAnnouncement) DeleteAnnouncement(ctx context.Context, req *v1.DeleteAnnouncementReq) (api_v1.BoolRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.Delete); has != true {
		return false, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().DeleteAnnouncement(ctx, req.Id, user.UnionMainId, user.Id)

	return ret == true, err
}

// QueryAnnouncement 查询公告｜列表
func (c *cSysAnnouncement) QueryAnnouncement(ctx context.Context, req *v1.QueryAnnouncementReq) (*sys_model.SysAnnouncementListRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.List); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	req.SearchParams.Filter = append(req.SearchParams.Filter, base_model.FilterInfo{
		Field: sys_dao.SysAnnouncement.Columns().UnionMainId,
		Where: "in",
		Value: []int64{user.UnionMainId},
	})

	ret, err := sys_service.Announcement().QueryAnnouncement(ctx, &req.SearchParams, req.IsExport)

	return ret, err
}

// 新增控制器方法 - 公告分类相关

// GetCategoryById 根据ID获取公告分类
func (c *cSysAnnouncement) GetCategoryById(ctx context.Context, req *v1.GetAnnouncementCategoryByIdReq) (*sys_model.SysAnnouncementCategoryRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.CategoryManage); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().GetCategoryById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 验证分类所属主体
	if ret.UnionMainId != 0 && ret.UnionMainId != user.UnionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
	}

	return ret, nil
}

// CreateCategory 创建公告分类
func (c *cSysAnnouncement) CreateCategory(ctx context.Context, req *v1.CreateAnnouncementCategoryReq) (*sys_model.SysAnnouncementCategoryRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.CategoryManage); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	category := &sys_model.SysAnnouncementCategory{
		Name:        req.Name,
		Code:        req.Code,
		UnionMainId: user.UnionMainId,
		Sort:        req.Sort,
		Description: req.Description,
	}

	ret, err := sys_service.Announcement().CreateCategory(ctx, category, user.UnionMainId, user.Id)

	return ret, err
}

// UpdateCategory 更新公告分类
func (c *cSysAnnouncement) UpdateCategory(ctx context.Context, req *v1.UpdateAnnouncementCategoryReq) (*sys_model.SysAnnouncementCategoryRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.CategoryManage); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 先获取原始分类信息
	category, err := sys_service.Announcement().GetCategoryById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 验证分类所属主体
	if category.UnionMainId != 0 && category.UnionMainId != user.UnionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
	}

	// 更新分类信息
	updateData := &sys_model.SysAnnouncementCategory{
		Id:          req.Id,
		Name:        req.Name,
		Code:        req.Code,
		Sort:        req.Sort,
		Description: req.Description,
	}

	ret, err := sys_service.Announcement().UpdateCategory(ctx, updateData, user.UnionMainId, user.Id)

	return ret, err
}

// DeleteCategory 删除公告分类
func (c *cSysAnnouncement) DeleteCategory(ctx context.Context, req *v1.DeleteAnnouncementCategoryReq) (api_v1.BoolRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.CategoryManage); has != true {
		return false, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 先获取原始分类信息
	category, err := sys_service.Announcement().GetCategoryById(ctx, req.Id)
	if err != nil {
		return false, err
	}

	// 验证分类所属主体
	if category.UnionMainId != 0 && category.UnionMainId != user.UnionMainId {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "{#error_announcement_category_not_exists}", sys_dao.SysAnnouncementCategory.Table())
	}

	ret, err := sys_service.Announcement().DeleteCategory(ctx, req.Id, user.UnionMainId)

	return (api_v1.BoolRes)(ret), err
}

// QueryCategory 查询公告分类列表
func (c *cSysAnnouncement) QueryCategory(ctx context.Context, req *v1.QueryAnnouncementCategoryReq) (*sys_model.SysAnnouncementCategoryListRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.CategoryManage); has != true {
		return nil, err
	}

	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 添加主体过滤
	req.SearchParams.Filter = append(req.SearchParams.Filter, base_model.FilterInfo{
		Field: sys_dao.SysAnnouncementCategory.Columns().UnionMainId,
		Where: "in",
		Value: []int64{0, user.UnionMainId}, // 包含平台公共分类和当前主体的分类
	})

	ret, err := sys_service.Announcement().QueryCategory(ctx, &req.SearchParams, req.IsExport)

	return ret, err
}

// 新增控制器方法 - 公告确认相关

// ConfirmAnnouncement 确认公告
func (c *cSysAnnouncement) ConfirmAnnouncement(ctx context.Context, req *v1.ConfirmAnnouncementReq) (api_v1.BoolRes, error) {
	// 用户接口，不需要权限检查
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().ConfirmAnnouncement(ctx, req.Id, user.Id, req.Comment)

	return (api_v1.BoolRes)(ret), err
}

// GetAnnouncementConfirmInfo 获取公告确认信息
func (c *cSysAnnouncement) GetAnnouncementConfirmInfo(ctx context.Context, req *v1.GetAnnouncementConfirmInfoReq) (*sys_model.SysAnnouncementConfirmRes, error) {
	// 用户接口，不需要权限检查
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().GetAnnouncementConfirmInfo(ctx, req.Id, user.Id)

	return ret, err
}

// QueryAnnouncementConfirms 查询公告确认记录
func (c *cSysAnnouncement) QueryAnnouncementConfirms(ctx context.Context, req *v1.QueryAnnouncementConfirmsReq) (*sys_model.SysAnnouncementConfirmListRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.ConfirmList); has != true {
		return nil, err
	}

	ret, err := sys_service.Announcement().QueryAnnouncementConfirms(ctx, req.Id, &req.SearchParams, req.IsExport)

	return ret, err
}

// 新增控制器方法 - 公告统计相关

// GetAnnouncementStatistics 获取公告统计信息
func (c *cSysAnnouncement) GetAnnouncementStatistics(ctx context.Context, req *v1.GetAnnouncementStatisticsReq) (*sys_model.SysAnnouncementStatistics, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.StatisticsView); has != true {
		return nil, err
	}

	ret, err := sys_service.Announcement().GetAnnouncementStatistics(ctx, req.Id)

	return ret, err
}

// 新增控制器方法 - 批量操作相关

// BatchMarkRead 批量标记公告为已读
func (c *cSysAnnouncement) BatchMarkRead(ctx context.Context, req *v1.BatchMarkReadReq) (api_v1.BoolRes, error) {
	// 用户接口，不需要权限检查
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().BatchMarkRead(ctx, req.Ids, user.Id)

	return (api_v1.BoolRes)(ret), err
}

// 新增控制器方法 - 撤销公告

// RevokedAnnouncement 撤销公告
func (c *cSysAnnouncement) RevokedAnnouncement(ctx context.Context, req *v1.RevokedAnnouncementReq) (api_v1.BoolRes, error) {
	// 权限检查
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Announcement.PermissionType.Revoke); has != true {
		return false, err
	}

	ret, err := sys_service.Announcement().RevokedAnnouncement(ctx, req.Id)

	return (api_v1.BoolRes)(ret), err
}
