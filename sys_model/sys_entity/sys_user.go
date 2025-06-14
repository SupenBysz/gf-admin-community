// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id         int64       `json:"id"         orm:"id"          description:""`
	Username   string      `json:"username"   orm:"username"    description:"账号"`
	Password   string      `json:"password"   orm:"password"    description:"密码"`
	State      int         `json:"state"      orm:"state"       description:"状态：0未激活、1正常、-1封号、-2异常、-3已注销"`
	Type       int         `json:"type"       orm:"type"        description:"用户类型：0匿名、1用户、2微商、4商户、8广告主、16服务商、32运营中心、64后台"`
	Mobile     string      `json:"mobile"     orm:"mobile"      description:"手机号"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:""`
	Email      string      `json:"email"      orm:"email"       description:"邮箱"`
	InviteCode string      `json:"inviteCode" orm:"invite_code" description:"邀请码，代表通过这个邀请码注册"`
}
