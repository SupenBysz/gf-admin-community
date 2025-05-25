// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysInvitePerson is the golang structure of table sys_invite_person for DAO operations like Where/Data.
type SysInvitePerson struct {
	g.Meta                  `orm:"table:sys_invite_person, do:true"`
	Id                      interface{} // ID
	FormUserId              interface{} // 邀请人
	ByUserId                interface{} // 被邀请人
	InviteCode              interface{} // 邀请码
	InviteAt                *gtime.Time // 邀请时间
	UserIdentifierPrefix    interface{} // 用户标识符前缀
	CompanyIdentifierPrefix interface{} // 公司标识符前缀
	InviteId                interface{} // 邀请码ID，来自sys_invite 表
}
