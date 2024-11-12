package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	v1 "github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_model"
)

type cSysAnnouncement struct{}

var SysAnnouncement cSysAnnouncement

// GetAnnouncementById 根据id查询公告｜信息
func (c *cSysAnnouncement) GetAnnouncementById(ctx context.Context, req *v1.GetAnnouncementByIdReq) (*sys_model.SysAnnouncementRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().GetAnnouncementById(ctx, req.Id, user.UnionMainId)

	return ret, err
}

// CreateAnnouncement 添加公告｜信息
func (c *cSysAnnouncement) CreateAnnouncement(ctx context.Context, req *v1.CreateAnnouncementReq) (*sys_model.SysAnnouncementRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().CreateAnnouncement(ctx, &req.SysAnnouncement, user.UnionMainId, user.Id)

	return ret, err
}

// UpdateAnnouncement 编辑公告｜信息
func (c *cSysAnnouncement) UpdateAnnouncement(ctx context.Context, req *v1.UpdateAnnouncementReq) (*sys_model.SysAnnouncementRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().UpdateAnnouncement(ctx, &req.UpdateSysAnnouncement, user.UnionMainId, user.Id)

	return ret, err
}

// DeleteAnnouncement 删除公告
func (c *cSysAnnouncement) DeleteAnnouncement(ctx context.Context, req *v1.DeleteAnnouncementReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().DeleteAnnouncement(ctx, req.Id, user.UnionMainId, user.Id)

	return ret == true, err
}

// QueryAnnouncement 查询公告｜列表
func (c *cSysAnnouncement) QueryAnnouncement(ctx context.Context, req *v1.QueryAnnouncementReq) (*sys_model.SysAnnouncementListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	req.SearchParams.Filter = append(req.SearchParams.Filter, base_model.FilterInfo{
		Field: sys_dao.SysAnnouncement.Columns().UnionMainId,
		Where: "in",
		Value: []int64{user.UnionMainId},
	})

	ret, err := sys_service.Announcement().QueryAnnouncement(ctx, &req.SearchParams, req.IsExport)

	return ret, err
}
