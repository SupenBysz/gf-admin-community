// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPersonLicense is the golang structure for table sys_person_license.
type SysPersonLicense struct {
	Id               int64       `json:"id"               orm:"id"                 description:"ID"`
	IdcardFrontPath  string      `json:"idcardFrontPath"  orm:"idcard_front_path"  description:"身份证头像面照片"`
	IdcardBackPath   string      `json:"idcardBackPath"   orm:"idcard_back_path"   description:"身份证国徽面照片"`
	No               string      `json:"no"               orm:"no"                 description:"身份证号"`
	Gender           int         `json:"gender"           orm:"gender"             description:"性别"`
	Nation           string      `json:"nation"           orm:"nation"             description:"名族"`
	Name             string      `json:"name"             orm:"name"               description:"姓名"`
	Birthday         string      `json:"birthday"         orm:"birthday"           description:"出生日期"`
	Address          string      `json:"address"          orm:"address"            description:"家庭住址"`
	IssuingAuthorit  string      `json:"issuingAuthorit"  orm:"issuing_authorit"   description:"签发机关"`
	IssuingDate      string      `json:"issuingDate"      orm:"issuing_date"       description:"签发日期"`
	ExpriyDate       string      `json:"expriyDate"       orm:"expriy_date"        description:""`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:""`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:""`
	State            int         `json:"state"            orm:"state"              description:"状态：0失效、1正常"`
	AuthType         int         `json:"authType"         orm:"auth_type"          description:"认证类型:"`
	Remark           string      `json:"remark"           orm:"remark"             description:"备注信息"`
	LatestAuditLogId int64       `json:"latestAuditLogId" orm:"latest_audit_logId" description:"最新的审核记录id"`
	UserId           int64       `json:"userId"           orm:"user_id"            description:"关联的用户ID"`
}
