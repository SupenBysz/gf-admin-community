// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMessage is the golang structure for table sys_message.
type SysMessage struct {
	Id             int64       `json:"id"             orm:"id"              description:"ID"`
	Title          string      `json:"title"          orm:"title"           description:"标题"`
	Summary        string      `json:"summary"        orm:"summary"         description:"摘要"`
	Content        string      `json:"content"        orm:"content"         description:"内容"`
	Type           int         `json:"type"           orm:"type"            description:"消息类型"`
	Link           string      `json:"link"           orm:"link"            description:"跳转链接"`
	ToUserIds      string      `json:"toUserIds"      orm:"to_user_ids"     description:"接收者UserIds，允许有多个接收者"`
	ToUserType     int         `json:"toUserType"     orm:"to_user_type"    description:"接收者类型用户类型，和UserType保持一致"`
	FromUserId     int64       `json:"fromUserId"     orm:"from_user_id"    description:"发送者ID，为-1代表系统消息"`
	FromUserType   int         `json:"fromUserType"   orm:"from_user_type"  description:"发送者类型"`
	SendAt         *gtime.Time `json:"sendAt"         orm:"send_at"         description:"发送时间"`
	ExtJson        string      `json:"extJson"        orm:"ext_json"        description:"拓展数据Json"`
	ReadUserIds    string      `json:"readUserIds"    orm:"read_user_ids"   description:"已读用户UserIds"`
	DataIdentifier string      `json:"dataIdentifier" orm:"data_identifier" description:"关联的数据标识"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""`
	SceneDesc      string      `json:"sceneDesc"      orm:"scene_desc"      description:"场景描述"`
	SceneType      int         `json:"sceneType"      orm:"scene_type"      description:"场景类型【业务层自定义】例如：1活动即将开始提醒、2活动开始提醒、3活动即将结束提醒、4活动结束提醒、5活动获奖提醒、6券即将生效提醒、7券的生效提醒、8券的失效提醒、9券即将失效提醒、10券核销提醒、8192系统通知、"`
}
