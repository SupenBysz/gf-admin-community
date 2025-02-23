// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComment is the golang structure of table sys_comment for DAO operations like Where/Data.
type SysComment struct {
	g.Meta        `orm:"table:sys_comment, do:true"`
	Id            interface{} // ID
	UserId        interface{} // 用户ID
	UnionMainId   interface{} // 关联主体ID
	UnionMainType interface{} // 关联主体类型
	Body          interface{} // 图文评论
	MediaIds      interface{} // 媒体资源：图文、视频等
	ReplyId       interface{} // 评论回复信息ID，即关联父级评论ID
	Score         interface{} // 评分0-5，间隔0.1
	CreatedAt     *gtime.Time // 评论发表时间
	UnionId       interface{} // 关联业务ID
}
