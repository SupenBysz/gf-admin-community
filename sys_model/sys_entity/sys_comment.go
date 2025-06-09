// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComment is the golang structure for table sys_comment.
type SysComment struct {
	Id            int64       `json:"id"            orm:"id"              description:"ID"`
	UserId        int64       `json:"userId"        orm:"user_id"         description:"用户ID"`
	UnionMainId   int64       `json:"unionMainId"   orm:"union_main_id"   description:"关联主体ID"`
	UnionMainType int         `json:"unionMainType" orm:"union_main_type" description:"关联主体类型"`
	Body          string      `json:"body"          orm:"body"            description:"图文评论"`
	MediaIds      string      `json:"mediaIds"      orm:"media_ids"       description:"媒体资源：图文、视频等"`
	ReplyId       int64       `json:"replyId"       orm:"reply_id"        description:"评论回复信息ID，即关联父级评论ID"`
	Score         int         `json:"score"         orm:"score"           description:"评分0-5，间隔0.1"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"评论发表时间"`
	UnionId       int64       `json:"unionId"       orm:"union_id"        description:"关联业务ID"`
}
