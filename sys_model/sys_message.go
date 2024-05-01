package sys_model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type SysMessage struct {
	Title     string  `json:"title"        description:"标题" v:"required#标题不能为空"`
	Summary   string  `json:"summary"      description:"摘要" `
	Content   string  `json:"content"      description:"内容" v:"required#内容不能为空"`
	Type      int     `json:"type"         description:"类型：1系统消息"  v:"required#请限定消息类型"`
	ToUserIds []int64 `json:"toUserIds"     description:"接收者UserIds，允许有多个接收者" v:"required#接收者id不能为空"`
	//ToUserType   int    `json:"toUserType"   description:"接收者类型" v:"required#接收者类型不能为空"`
	FromUserId     int64  `json:"fromUserId"   description:"发送者id"`
	FromUserType   int    `json:"fromUserType" description:"发送者用户类型，和UserType保持一致"`
	ExtJson        string `json:"extJson"      description:"拓展数据Json"`
	DataIdentifier string `json:"dataIdentifier" description:"数据标识"`

	//ToUserType   int         `json:"toUserType"   description:"接收者类型用户类型，和UserType保持一致"`
	//CreatedAt   *gtime.Time `json:"createdAt"    description:"创建时间"`
	//SendAt      *gtime.Time `json:"sendAt"       description:"发送时间"`
	//ReadUserIds string      `json:"readUserIds"  description:"已读UserIds"`
}

type UpdateSysMessage struct {
	Title     *string `json:"title"        description:"标题" `
	Summary   *string `json:"summary"      description:"摘要" `
	Content   *string `json:"content"      description:"内容" `
	Type      *int    `json:"type"         description:"类型：1系统消息"  `
	ToUserIds []int64 `json:"toUserId"     description:"接收者UserIds，允许有多个接收者" `
	//ToUserType   int    `json:"toUserType"   description:"接收者类型" v:"required#接收者类型不能为空"`
	ExtJson        string  `json:"extJson"      description:"拓展数据Json"`
	DataIdentifier *string `json:"dataIdentifier" description:"数据标识"`
}

type SysMessageRes struct {
	Id             int64       `json:"id"             description:"ID"`
	Title          string      `json:"title"          description:"标题"`
	Summary        string      `json:"summary"        description:"摘要"`
	Content        string      `json:"content"        description:"内容"`
	Type           int         `json:"type"           description:"消息类型: 1系统消息，支持自定义"`
	Link           string      `json:"link"           description:"跳转链接"`
	ToUserIds      string      `json:"toUserIds"      description:"接收者UserIds，允许有多个接收者"`
	ToUserType     int         `json:"toUserType"     description:"接收者类型用户类型，和UserType保持一致"`
	FromUserId     int64       `json:"fromUserId"     description:"发送者ID，为-1代表系统消息"`
	FromUserType   int         `json:"fromUserType"   description:"发送者类型"`
	SendAt         *gtime.Time `json:"sendAt"         description:"发送时间"`
	ExtJson        string      `json:"extJson"        description:"拓展数据Json"`
	ReadUserIds    string      `json:"readUserIds"    description:"已读用户UserIds"`
	DataIdentifier string      `json:"dataIdentifier" description:"关联的数据标识"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:""`
}

type SysMessageListRes base_model.CollectRes[SysMessageRes]
