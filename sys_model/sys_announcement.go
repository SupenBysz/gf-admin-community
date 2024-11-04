package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type SysAnnouncement struct {
	Title string `json:"title"         orm:"title"           description:"公告标题"`
	//UnionMainId   int64       `json:"unionMainId"   orm:"union_main_id"   description:"发布主体，0值则代表平台发布的公告"`
	PublicAt      *gtime.Time `json:"publicAt"      orm:"public_at"       description:"公示时间，只有到了公示时间用户才可见"`
	Body          string      `json:"body"          orm:"body"            description:"公告正文"`
	UserTypeScope int64       `json:"userTypeScope" orm:"user_type_scope" description:"受众用户类型：0则所有，复合类型 "`
	ExpireAt      *gtime.Time `json:"expireAt"      orm:"expire_at"       description:"过期时间，过期后前端用户不可见"`
	State         int         `json:"state"         orm:"state"           description:"状态：1草稿、2待发布、4已发布、8已过期、16已撤销"`
	ExtDataJson   string      `json:"extDataJson"   orm:"ext_data_json"   description:"扩展json数据"`
	//CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:""`
	//UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:""`
	//CreatedBy     int64       `json:"createdBy"     orm:"created_by"      description:"创建用户"`
	//UpdatedBy   int64       `json:"updatedBy"     orm:"updated_by"      description:"最后修改用户"`
	//DeletedAt   *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:""`
	//DeletedBy   int64       `json:"deletedBy"     orm:"deleted_by"      description:""`
}

type UpdateSysAnnouncement struct {
	Id            int64  `json:"id"  v:"required#公告Id不能为空"`
	Title         string `json:"title"         orm:"title"           description:"公告标题"`
	Body          string `json:"body"          orm:"body"            description:"公告正文"`
	UserTypeScope int64  `json:"userTypeScope" orm:"user_type_scope" description:"受众用户类型：0则所有，复合类型"`
	ExtDataJson   string `json:"extDataJson"   orm:"ext_data_json"   description:"扩展json数据"`

	PublicAt *gtime.Time `json:"publicAt"      orm:"public_at"       description:"公示时间，只有到了公示时间用户才可见"`
	ExpireAt *gtime.Time `json:"expireAt"      orm:"expire_at"       description:"过期时间，过期后前端用户不可见"`

	//UserTypeScope int64  `json:"userTypeScope" orm:"user_type_scope" description:"受众用户类型：0则所有，复合类型，注意：已达到公示时间的公告不能修改"`
	//PublicAt *gtime.Time `json:"publicAt"      orm:"public_at"       description:"公示时间，只有到了公示时间用户才可见，注意：已达到公示时间的公告不能修改公示时间"`
}

type SysAnnouncementRes sys_entity.SysAnnouncement

type SysAnnouncementListRes base_model.CollectRes[SysAnnouncementRes]
