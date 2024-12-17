package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	v1 "github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// MarkRead 标记已读｜公告
func (c *cSysAnnouncement) MarkRead(ctx context.Context, req *v1.MarkReadReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().MarkRead(ctx, req.Id, user.Id)

	return ret == true, err
}

// MarkUnRead 标记未读｜公告
func (c *cSysAnnouncement) MarkUnRead(ctx context.Context, req *v1.MarkUnReadReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().MarkUnRead(ctx, req.Id, user.UnionMainId)

	return ret == true, err
}

// HasUnReadAnnouncement 获取未读公告数量
func (c *cSysAnnouncement) HasUnReadAnnouncement(ctx context.Context, req *v1.HasUnReadAnnouncementReq) (api_v1.IntRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().HasUnReadAnnouncement(ctx, user.Id, user.UnionMainId)

	return (api_v1.IntRes)(ret), err
}

// QueryAnnouncementListByUser 查询用户的公告｜列表
func (c *cSysAnnouncement) QueryAnnouncementListByUser(ctx context.Context, req *v1.QueryAnnouncementListByUserReq) (*sys_model.SysAnnouncementListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Announcement().QueryAnnouncementListByUser(ctx, user.Id, user.UnionMainId, req.Type, &req.SearchParams, req.IsExport)

	return ret, err
}

// QueryAnnouncementReadUserList 查询公告的已读用户｜列表
//func (c *cSysAnnouncement) QueryAnnouncementReadUserList(ctx context.Context, req *v1.QueryAnnouncementReadUserListReq) (*sys_model.SysAnnouncementListRes, error) {
//	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
//
//	ret, err := sys_service.Announcement().QueryAnnouncementReadUserList(ctx, req.SearchParams, req.IsExport, user.UnionMainId)
//
//	return ret, err
//}
