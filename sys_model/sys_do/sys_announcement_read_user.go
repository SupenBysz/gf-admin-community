// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncementReadUser is the golang structure of table sys_announcement_read_user for DAO operations like Where/Data.
type SysAnnouncementReadUser struct {
	g.Meta             `orm:"table:sys_announcement_read_user, do:true"`
	Id                 interface{} // ID
	UserId             interface{} // 用户ID
	ReadAnnouncementId interface{} // 用户已阅读的公告id
	ReadAt             *gtime.Time // 用户阅读时间
	ExtDataJson        interface{} // 扩展数据Json，由业务端决定用途
	FlagRead           interface{} // 标记已读，0未读，1已读，用户首次打开公告即标记已读，可手动标记未读，但read_at 数据不变，下次点开时更新阅读时间，并标记已读
}
