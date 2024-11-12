// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserDetail is the golang structure for table sys_user_detail.
type SysUserDetail struct {
	Id              int64       `json:"id"              orm:"id"                description:"ID，保持与USERID一致"`
	Realname        string      `json:"realname"        orm:"realname"          description:"姓名"`
	UnionMainName   string      `json:"unionMainName"   orm:"union_main_name"   description:"关联主体名称"`
	LastLoginIp     string      `json:"lastLoginIp"     orm:"last_login_ip"     description:"最后登录IP"`
	LastLoginArea   string      `json:"lastLoginArea"   orm:"last_login_area"   description:"最后登录地区"`
	LastLoginAt     *gtime.Time `json:"lastLoginAt"     orm:"last_login_at"     description:"最后登录时间"`
	LastHeartbeatAt *gtime.Time `json:"lastHeartbeatAt" orm:"last_heartbeat_at" description:"最后在线时间"`
}
