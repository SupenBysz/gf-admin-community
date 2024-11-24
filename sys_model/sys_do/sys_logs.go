// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLogs is the golang structure of table sys_logs for DAO operations like Where/Data.
type SysLogs struct {
	g.Meta    `orm:"table:sys_logs, do:true"`
	Id        interface{} // ID
	UserId    interface{} // 用户UID
	Error     interface{} // 错误信息
	Category  interface{} // 分类
	Level     interface{} // 等级
	Content   interface{} // 日志内容
	Context   interface{} // 上下文数据
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
