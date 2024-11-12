package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type SysAnnouncementReadUser struct {
	Id                 int64       `json:"id"                 orm:"id"                   description:"ID"`
	UserId             int64       `json:"userId"             orm:"user_id"              description:"用户ID"`
	ReadAnnouncementId string      `json:"readAnnouncementId" orm:"read_announcement_id" description:"用户已阅读的公告id"`
	ReadAt             *gtime.Time `json:"readAt"             orm:"read_at"              description:"用户阅读时间"`
	ExtDataJson        string      `json:"extDataJson"        orm:"ext_data_json"        description:"扩展数据Json，由业务端决定用途"`
	FlagRead           int         `json:"flagRead"           orm:"flag_read"            description:"标记已读，0未读，1已读，用户首次打开公告即标记已读，可手动标记未读，但read_at 数据不变，下次点开时更新阅读时间，并标记已读"`
}
type SysAnnouncementReadUserRes sys_entity.SysAnnouncementReadUser

type SysAnnouncementReadUserListRes base_model.CollectRes[SysAnnouncementReadUserRes]
