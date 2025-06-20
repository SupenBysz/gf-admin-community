// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysInvite is the golang structure of table sys_invite for DAO operations like Where/Data.
type SysInvite struct {
	g.Meta         `orm:"table:sys_invite, do:true"`
	Id             interface{} // ID
	UserId         interface{} // 用户ID, 也就是邀约人ID
	Value          interface{} // 邀约码背后的关联业务Json数据,
	ExpireAt       *gtime.Time // 邀约码的过期失效
	ActivateNumber interface{} // 邀约码的激活次数限制，小于0，则无限制
	State          interface{} // 状态： 0失效、1正常
	Type           interface{} // 类型： 1注册、2加入团队、4加入角色 (复合类型)
	CreatedAt      *gtime.Time //
	Identifier     interface{} // 标识符
	InviteCode     interface{} // 邀请码
}
