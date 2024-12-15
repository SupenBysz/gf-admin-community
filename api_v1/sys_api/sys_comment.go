package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetCommentByIdReq struct {
	g.Meta `path:"/getCommentById" method:"post" summary:"获取评论｜信息" tags:"评论管理"`
	Id     int64 `json:"id" dc:"评论ID" v:"required#请输入评论信息id"`
}

type CreateCommentReq struct {
	g.Meta `path:"/createComment" method:"post" summary:"创建获取评论｜信息" tags:"评论管理"`
	sys_model.SysComment
}

type DeleteCommentReq struct {
	g.Meta `path:"/deleteComment" method:"post" summary:"删除获取评论｜信息" tags:"评论管理"`
	Id     int64 `json:"id" dc:"评论ID" v:"required#请输入评论信息id"`
}

type QueryCommentReq struct {
	g.Meta `path:"/queryComment" method:"post" summary:"查询获取评论｜列表" tags:"评论管理"`
	base_model.SearchParams
}
