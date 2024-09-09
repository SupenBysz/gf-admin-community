package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type PersonLicense struct {
	Id               int64  `json:"id"              description:"ID"`
	IdcardFrontPath  string `json:"idcardFrontPath"  description:"身份证头像面照片"`
	IdcardBackPath   string `json:"idcardBackPath"   description:"身份证国徽面照片"`
	No               string `json:"no"              description:"身份证号"`
	Gender           int    `json:"gender"          description:"性别"`
	Nation           string `json:"nation"          description:"名族"`
	Name             string `json:"name"            description:"姓名"`
	Birthday         string `json:"birthday"        description:"出生日期"`
	Address          string `json:"address"         description:"家庭住址"`
	IssuingAuthorit  string `json:"issuingAuthorit" description:"签发机关"`
	IssuingDate      string `json:"issuingDate"     description:"签发日期"`
	ExpriyDate       string `json:"expriyDate"      description:""`
	Remark           string `json:"remark"          description:"备注信息"`
	LatestAuditLogId int64  `json:"latestAuditLogId" description:"最新的审核记录id"`

	State    int    `json:"state"           description:"状态：0失效、1正常" v:"in:0,1#状态错误"`
	AuthType int    `json:"authType"        description:"认证类型:"`
	Summary  string `json:"summary"              description:"概述"`
}

type PersonLicenseRes sys_entity.SysPersonLicense
type PersonLicenseListRes base_model.CollectRes[sys_entity.SysPersonLicense]

type AuditPersonLicense struct {
	UnionMainId int64  `json:"unionMainId"             dc:"资质审核关联的业务主体ID"` // 个人资质的unionMainId, 没有则为0
	LicenseId   int64  `json:"licenseId"             dc:"资质ID"`
	UserId      int64  `json:"userId" dc:"上传资质的userId"` // 个人资质存在待上传的问题， 所以userID代表上传者
	OwnerUserId int64  `json:"ownerUserId" dc:"资质的所属userId" `
	Summary     string `json:"summary"               description:"概述"`
	PersonLicense
}
