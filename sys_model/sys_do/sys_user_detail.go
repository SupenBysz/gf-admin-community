// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserDetail is the golang structure of table sys_user_detail for DAO operations like Where/Data.
type SysUserDetail struct {
	g.Meta          `orm:"table:sys_user_detail, do:true"`
	Id              interface{} // ID，保持与USERID一致
	Realname        interface{} // 姓名
	UnionMainName   interface{} // 关联主体名称
	LastLoginIp     interface{} // 最后登录IP
	LastLoginArea   interface{} // 最后登录地区
	LastLoginAt     *gtime.Time // 最后登录时间
	LastHeartbeatAt *gtime.Time // 最后在线时间
}
