// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysComment interface {
		// GetCommentById 根据id获取评论
		GetCommentById(ctx context.Context, id int64, makeUrl bool) (*sys_model.SysCommentRes, error)
		// CreateComment 创建评论
		CreateComment(ctx context.Context, info *sys_model.SysComment, makeUrl bool) (*sys_model.SysCommentRes, error)
		// DeleteComment 删除评论
		DeleteComment(ctx context.Context, id int64) (api_v1.BoolRes, error)
		// QueryComment 查询评论
		QueryComment(ctx context.Context, search *base_model.SearchParams) (*sys_model.SysCommentListRes, error)
	}
)

var (
	localSysComment ISysComment
)

func SysComment() ISysComment {
	if localSysComment == nil {
		panic("implement not found for interface ISysComment, forgot register?")
	}
	return localSysComment
}

func RegisterSysComment(i ISysComment) {
	localSysComment = i
}
