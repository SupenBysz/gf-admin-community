// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncement is the golang structure for table sys_announcement.
type SysAnnouncement struct {
	Id              int64       `json:"id"              orm:"id"               description:""`
	Title           string      `json:"title"           orm:"title"            description:"公告标题"`
	UnionMainId     int64       `json:"unionMainId"     orm:"union_main_id"    description:"发布主体，0值则代表平台发布的公告"`
	PublicAt        *gtime.Time `json:"publicAt"        orm:"public_at"        description:"公示时间，只有到了公示时间用户才可见"`
	Body            string      `json:"body"            orm:"body"             description:"公告正文"`
	UserTypeScope   int         `json:"userTypeScope"   orm:"user_type_scope"  description:"受众用户类型：0则所有，复合类型"`
	ExpireAt        *gtime.Time `json:"expireAt"        orm:"expire_at"        description:"过期时间，过期后前端用户不可见"`
	State           int         `json:"state"           orm:"state"            description:"状态：1草稿、2待发布、4已发布、8已过期、16已撤销"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:""`
	CreatedBy       int64       `json:"createdBy"       orm:"created_by"       description:"创建用户"`
	UpdatedBy       int64       `json:"updatedBy"       orm:"updated_by"       description:"最后修改用户"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:""`
	DeletedBy       int64       `json:"deletedBy"       orm:"deleted_by"       description:""`
	ExtDataJson     string      `json:"extDataJson"     orm:"ext_data_json"    description:"扩展json数据"`
	CategoryId      int64       `json:"categoryId"      orm:"category_id"      description:"公告分类ID"`
	Priority        int         `json:"priority"        orm:"priority"         description:"优先级：1普通、2重要、3紧急"`
	IsPinned        int         `json:"isPinned"        orm:"is_pinned"        description:"是否置顶：0否、1是"`
	IsForceRead     int         `json:"isForceRead"     orm:"is_force_read"    description:"是否强制阅读：0否、1是"`
	Tags            string      `json:"tags"            orm:"tags"             description:"公告标签，多个用逗号分隔"`
	ReadCount       int64       `json:"readCount"       orm:"read_count"       description:"阅读次数"`
	ConfirmRequired int         `json:"confirmRequired" orm:"confirm_required" description:"是否需要确认：0否、1是"`
}
