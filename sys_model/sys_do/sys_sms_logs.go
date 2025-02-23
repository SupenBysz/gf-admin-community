// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLogs is the golang structure of table sys_sms_logs for DAO operations like Where/Data.
type SysSmsLogs struct {
	g.Meta    `orm:"table:sys_sms_logs, do:true"`
	Id        interface{} //
	Type      interface{} // 短信平台：qyxs：企业信使
	Context   interface{} // 短信内容
	Mobile    interface{} // 手机号
	State     interface{} // 发送状态
	Result    interface{} // 短信接口返回内容
	UserId    interface{} // 用户ID
	LicenseId interface{} // 主体ID
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
