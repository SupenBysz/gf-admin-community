// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysInvitePerson is the golang structure for table sys_invite_person.
type SysInvitePerson struct {
	Id                      int64       `json:"id"                      orm:"id"                        description:"ID"`
	FormUserId              int64       `json:"formUserId"              orm:"form_user_id"              description:"邀请人"`
	ByUserId                int64       `json:"byUserId"                orm:"by_user_id"                description:"被邀请人"`
	InviteCode              string      `json:"inviteCode"              orm:"invite_code"               description:"邀请码"`
	InviteAt                *gtime.Time `json:"inviteAt"                orm:"invite_at"                 description:"邀请时间"`
	UserIdentifierPrefix    string      `json:"userIdentifierPrefix"    orm:"user_identifier_prefix"    description:"用户标识符前缀"`
	CompanyIdentifierPrefix string      `json:"companyIdentifierPrefix" orm:"company_identifier_prefix" description:"公司标识符前缀"`
	InviteId                int64       `json:"inviteId"                orm:"invite_id"                 description:"邀请码ID，来自sys_invite 表"`
}
