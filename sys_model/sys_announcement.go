package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type SysAnnouncement struct {
	Title           string      `json:"title"         orm:"title"           description:"公告标题"`
	PublicAt        *gtime.Time `json:"publicAt"      orm:"public_at"       description:"公示时间，只有到了公示时间用户才可见"`
	Body            string      `json:"body"          orm:"body"            description:"公告正文"`
	UserTypeScope   int64       `json:"userTypeScope" orm:"user_type_scope" description:"受众用户类型：0则所有，复合类型 "`
	ExpireAt        *gtime.Time `json:"expireAt"      orm:"expire_at"       description:"过期时间，过期后前端用户不可见"`
	State           int         `json:"state"         orm:"state"           description:"状态：1草稿、2待发布、4已发布、8已过期、16已撤销"`
	ExtDataJson     string      `json:"extDataJson"   orm:"ext_data_json"   description:"扩展json数据"`
	CategoryId      int64       `json:"categoryId"    orm:"category_id"     description:"公告分类ID"`
	Priority        int         `json:"priority"      orm:"priority"        description:"优先级：1普通、2重要、3紧急"`
	IsPinned        int         `json:"isPinned"      orm:"is_pinned"       description:"是否置顶：0否、1是"`
	IsForceRead     int         `json:"isForceRead"   orm:"is_force_read"   description:"是否强制阅读：0否、1是"`
	Tags            string      `json:"tags"          orm:"tags"            description:"公告标签，多个用逗号分隔"`
	ConfirmRequired int         `json:"confirmRequired" orm:"confirm_required" description:"是否需要确认：0否、1是"`
}

type UpdateSysAnnouncement struct {
	Id              *int64      `json:"id"  v:"required#公告Id不能为空"`
	Title           *string     `json:"title"         orm:"title"           description:"公告标题"`
	Body            *string     `json:"body"          orm:"body"            description:"公告正文"`
	UserTypeScope   *int64      `json:"userTypeScope" orm:"user_type_scope" description:"受众用户类型：0则所有，复合类型"`
	State           int         `json:"state"         orm:"state"           description:"状态：1草稿、2待发布、4已发布、8已过期、16已撤销"`
	ExtDataJson     *string     `json:"extDataJson"   orm:"ext_data_json"   description:"扩展json数据"`
	PublicAt        *gtime.Time `json:"publicAt"  orm:"public_at"       description:"公示时间，只有到了公示时间用户才可见"`
	ExpireAt        *gtime.Time `json:"expireAt"  orm:"expire_at"       description:"过期时间，过期后前端用户不可见"`
	CategoryId      *int64      `json:"categoryId"    orm:"category_id"     description:"公告分类ID"`
	Priority        *int        `json:"priority"      orm:"priority"        description:"优先级：1普通、2重要、3紧急"`
	IsPinned        *int        `json:"isPinned"      orm:"is_pinned"       description:"是否置顶：0否、1是"`
	IsForceRead     *int        `json:"isForceRead"   orm:"is_force_read"   description:"是否强制阅读：0否、1是"`
	Tags            *string     `json:"tags"          orm:"tags"            description:"公告标签，多个用逗号分隔"`
	ConfirmRequired *int        `json:"confirmRequired" orm:"confirm_required" description:"是否需要确认：0否、1是"`
}

type SysAnnouncementRes struct {
	ReadState     int    `json:"readState" dc:"阅读状态：1未读，2已读"`
	CategoryName  string `json:"categoryName" dc:"分类名称"`
	ConfirmStatus int    `json:"confirmStatus" dc:"确认状态：0未确认，1已确认"`
	ConfirmCount  int    `json:"confirmCount" dc:"已确认人数"`
	sys_entity.SysAnnouncement
}

type SysAnnouncementListRes base_model.CollectRes[SysAnnouncementRes]

// 公告分类相关模型
type SysAnnouncementCategory struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name" v:"required#分类名称不能为空" dc:"分类名称"`
	Code        string `json:"code" v:"required#分类编码不能为空" dc:"分类编码"`
	UnionMainId int64  `json:"unionMainId" dc:"所属主体ID"`
	Sort        int    `json:"sort" dc:"排序"`
	Description string `json:"description" dc:"描述"`
}

type SysAnnouncementCategoryRes struct {
	AnnouncementCount int `json:"announcementCount" dc:"该分类下公告数量"`
	sys_entity.SysAnnouncementCategory
}

type SysAnnouncementCategoryListRes base_model.CollectRes[SysAnnouncementCategoryRes]

// 公告确认相关模型
type SysAnnouncementConfirm struct {
	Id             int64       `json:"id,omitempty"`
	UserId         int64       `json:"userId" v:"required#用户ID不能为空" dc:"用户ID"`
	AnnouncementId int64       `json:"announcementId" v:"required#公告ID不能为空" dc:"公告ID"`
	ConfirmAt      *gtime.Time `json:"confirmAt" dc:"确认时间"`
	ConfirmComment string      `json:"confirmComment" dc:"确认备注"`
}

type SysAnnouncementConfirmRes struct {
	UserName string `json:"userName" dc:"用户名称"`
	sys_entity.SysAnnouncementConfirm
}

type SysAnnouncementConfirmListRes base_model.CollectRes[SysAnnouncementConfirmRes]

// 公告统计信息
type SysAnnouncementStatistics struct {
	AnnouncementId int64   `json:"announcementId" dc:"公告ID"`
	ReadCount      int     `json:"readCount" dc:"已读人数"`
	UnreadCount    int     `json:"unreadCount" dc:"未读人数"`
	ConfirmCount   int     `json:"confirmCount" dc:"已确认人数"`
	TotalUserCount int     `json:"totalUserCount" dc:"总用户数"`
	ReadRate       float64 `json:"readRate" dc:"阅读率"`
	ConfirmRate    float64 `json:"confirmRate" dc:"确认率"`
}
