package sysapi

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type UploadReq struct {
	g.Meta `path:"/upload" method:"post" summary:"文件上传" tags:"工具"`
	sys_model.FileUploadInput
}

type UploadIDCardWithOCRReq struct {
	g.Meta `path:"/uploadIDCardWithOCR" method:"post" summary:"上传身份证" tags:"工具"`
	sys_model.OCRIDCardFileUploadInput
}

type UploadBusinessLicenseWithOCRReq struct {
	g.Meta `path:"/uploadBusinessLicenseWithOCR" method:"post" summary:"上传营业执照" tags:"工具"`
	sys_model.OCRBusinessLicense
}

type UploadBankCardWithOCRReq struct {
	g.Meta `path:"/uploadBankCardWithOCR" method:"post" summary:"上传银行卡" tags:"工具"`
	sys_model.BankCardWithOCRInput
}

type UploadBusinessLicenseWithOCRRes sys_model.BusinessLicenseWithOCR

type IDCardWithOCRRes sys_model.IDCardWithOCR
type UploadRes sys_model.FileUploadOutput

type BankCardWithOCRRes sys_model.BankCardWithOCR
