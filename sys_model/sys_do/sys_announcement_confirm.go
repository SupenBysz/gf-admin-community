// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncementConfirm is the golang structure of table sys_announcement_confirm for DAO operations like Where/Data.
type SysAnnouncementConfirm struct {
	g.Meta         `orm:"table:sys_announcement_confirm, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户ID
	AnnouncementId interface{} // 公告ID
	ConfirmAt      *gtime.Time // 确认时间
	ConfirmComment interface{} // 确认备注
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	CreatedBy      interface{} //
	UpdatedBy      interface{} //
	DeletedAt      *gtime.Time //
	DeletedBy      interface{} //
}
