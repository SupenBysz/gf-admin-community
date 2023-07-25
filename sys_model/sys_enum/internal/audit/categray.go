package sys_enum_audit

import "github.com/kysion/base-library/utility/enum"

// CategoryEnum 审核业务类别：1个人资质审核、2主体资质审核、4数据审核
type CategoryEnum enum.IEnumCode[int]

type category struct {
	PersonLicenseAudit  CategoryEnum
	CompanyLicenseAudit CategoryEnum
	DataAudit           CategoryEnum
}

var Category = category{
	PersonLicenseAudit:  enum.New[CategoryEnum](1, "个人资质审核"),
	CompanyLicenseAudit: enum.New[CategoryEnum](2, "主体资质审核"),
	DataAudit:           enum.New[CategoryEnum](4, "数据审核"),
}

func (e category) New(code int, description string) CategoryEnum {
	if (code&Category.PersonLicenseAudit.Code()) == Category.PersonLicenseAudit.Code() ||
		(code&Category.CompanyLicenseAudit.Code()) == Category.CompanyLicenseAudit.Code() ||
		(code&Category.DataAudit.Code()) == Category.DataAudit.Code() {
		return enum.New[CategoryEnum](code, description)
	}
	panic("kyAudit.Category.New: error")
}
