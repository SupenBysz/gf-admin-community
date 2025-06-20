// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysInvite is the golang structure for table sys_invite.
type SysInvite struct {
	Id             int64       `json:"id"             orm:"id"              description:"ID"`
	UserId         int64       `json:"userId"         orm:"user_id"         description:"用户ID, 也就是邀约人ID"`
	Value          string      `json:"value"          orm:"value"           description:"邀约码背后的关联业务Json数据,"`
	ExpireAt       *gtime.Time `json:"expireAt"       orm:"expire_at"       description:"邀约码的过期失效"`
	ActivateNumber int         `json:"activateNumber" orm:"activate_number" description:"邀约码的激活次数限制，小于0，则无限制"`
	State          int         `json:"state"          orm:"state"           description:"状态： 0失效、1正常"`
	Type           int         `json:"type"           orm:"type"            description:"类型： 1注册、2加入团队、4加入角色 (复合类型)"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`
	Identifier     string      `json:"identifier"     orm:"identifier"      description:"标识符"`
	InviteCode     string      `json:"inviteCode"     orm:"invite_code"     description:"邀请码"`
}
