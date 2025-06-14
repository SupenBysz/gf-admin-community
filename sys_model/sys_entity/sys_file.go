// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFile is the golang structure for table sys_file.
type SysFile struct {
	Id             int64       `json:"id"             orm:"id"              description:"自增ID"`
	Name           string      `json:"name"           orm:"name"            description:"文件名称"`
	Src            string      `json:"src"            orm:"src"             description:"存储路径"`
	Url            string      `json:"url"            orm:"url"             description:"URL地址"`
	Ext            string      `json:"ext"            orm:"ext"             description:"扩展名"`
	Size           int64       `json:"size"           orm:"size"            description:"文件大小"`
	Category       string      `json:"category"       orm:"category"        description:"文件分类"`
	UserId         int64       `json:"userId"         orm:"user_id"         description:"用户ID"`
	UnionMainId    int64       `json:"unionMainId"    orm:"union_main_id"   description:"关联主体ID"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`
	LocalPath      string      `json:"localPath"      orm:"local_path"      description:"本地路径"`
	AllowAnonymous int         `json:"allowAnonymous" orm:"allow_anonymous" description:"是否允许匿名访问"`
}
