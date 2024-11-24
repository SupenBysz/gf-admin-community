// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPersonLicense is the golang structure of table sys_person_license for DAO operations like Where/Data.
type SysPersonLicense struct {
	g.Meta           `orm:"table:sys_person_license, do:true"`
	Id               interface{} // ID
	IdcardFrontPath  interface{} // 身份证头像面照片
	IdcardBackPath   interface{} // 身份证国徽面照片
	No               interface{} // 身份证号
	Gender           interface{} // 性别
	Nation           interface{} // 名族
	Name             interface{} // 姓名
	Birthday         interface{} // 出生日期
	Address          interface{} // 家庭住址
	IssuingAuthorit  interface{} // 签发机关
	IssuingDate      interface{} // 签发日期
	ExpriyDate       interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	DeletedAt        *gtime.Time //
	State            interface{} // 状态：0失效、1正常
	AuthType         interface{} // 认证类型:
	Remark           interface{} // 备注信息
	LatestAuditLogid interface{} // 最新的审核记录id
	UserId           interface{} // 关联的用户ID
}
