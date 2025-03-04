// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLogs is the golang structure for table sys_logs.
type SysLogs struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户UID"`
	Error     string      `json:"error"     orm:"error"      description:"错误信息"`
	Category  string      `json:"category"  orm:"category"   description:"分类"`
	Level     int         `json:"level"     orm:"level"      description:"等级"`
	Content   string      `json:"content"   orm:"content"    description:"日志内容"`
	Context   string      `json:"context"   orm:"context"    description:"上下文数据"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
