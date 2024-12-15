package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

var SysComment = cSysComment{}

type cSysComment struct{}

// GetCommentById 根据ID查下评论
func (c *cSysComment) GetCommentById(ctx context.Context, req *sys_api.GetCommentByIdReq) (*sys_model.SysCommentRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Comment.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	return sys_service.SysComment().GetCommentById(ctx, req.Id, true)
}

// CreateComment 新建评论
func (c *cSysComment) CreateComment(ctx context.Context, req *sys_api.CreateCommentReq) (*sys_model.SysCommentRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Comment.PermissionType.Create); has != true {
		return nil, err
	}

	return sys_service.SysComment().CreateComment(ctx, &req.SysComment, true)
}

// DeleteComment 删除评论
func (c *cSysComment) DeleteComment(ctx context.Context, req *sys_api.DeleteCommentReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Comment.PermissionType.Delete); has != true {
		return false, err
	}

	return sys_service.SysComment().DeleteComment(ctx, req.Id)
}

// QueryComment 查询评论
func (c *cSysComment) QueryComment(ctx context.Context, req *sys_api.QueryCommentReq) (*sys_model.SysCommentListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Comment.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	return sys_service.SysComment().QueryComment(ctx, &req.SearchParams)
}
