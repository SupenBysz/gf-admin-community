// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSmsLogs is the golang structure for table sys_sms_logs.
type SysSmsLogs struct {
	Id        float64     `json:"id"        orm:"id"         description:""`
	Type      string      `json:"type"      orm:"type"       description:"短信平台：qyxs：企业信使"`
	Context   string      `json:"context"   orm:"context"    description:"短信内容"`
	Mobile    string      `json:"mobile"    orm:"mobile"     description:"手机号"`
	State     string      `json:"state"     orm:"state"      description:"发送状态"`
	Result    string      `json:"result"    orm:"result"     description:"短信接口返回内容"`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户ID"`
	LicenseId int64       `json:"licenseId" orm:"license_id" description:"主体ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
