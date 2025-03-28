// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncementConfirm is the golang structure for table sys_announcement_confirm.
type SysAnnouncementConfirm struct {
	Id             int64       `json:"id"             orm:"id"              description:""`
	UserId         int64       `json:"userId"         orm:"user_id"         description:"用户ID"`
	AnnouncementId int64       `json:"announcementId" orm:"announcement_id" description:"公告ID"`
	ConfirmAt      *gtime.Time `json:"confirmAt"      orm:"confirm_at"      description:"确认时间"`
	ConfirmComment string      `json:"confirmComment" orm:"confirm_comment" description:"确认备注"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"      description:""`
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""`
	DeletedBy      int64       `json:"deletedBy"      orm:"deleted_by"      description:""`
}
