// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysMemberLevelUser is the golang structure of table sys_member_level_user for DAO operations like Where/Data.
type SysMemberLevelUser struct {
	g.Meta           `orm:"table:sys_member_level_user, do:true"`
	Id               interface{} // ID
	UserId           interface{} // 用户ID
	ExtMemberLevelId interface{} // 会员级别
	UnionMainId      interface{} // 保留字段
}
