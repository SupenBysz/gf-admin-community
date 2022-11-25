package sysapi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/SupenBysz/gf-admin-community/model"
)

type UploadReq struct {
	g.Meta `path:"/upload" method:"post" summary:"文件上传" tags:"工具"`
	model.FileUploadInput
}

type UploadIDCardWithOCRReq struct {
	g.Meta `path:"/uploadIDCardWithOCR" method:"post" summary:"上传身份证" tags:"工具"`
	model.OCRIDCardFileUploadInput
}

type UploadBusinessLicenseWithOCRReq struct {
	g.Meta `path:"/uploadBusinessLicenseWithOCR" method:"post" summary:"上传营业执照" tags:"工具"`
	model.OCRBusinessLicense
}

type UploadBankCardWithOCRReq struct {
	g.Meta `path:"/uploadBankCardWithOCR" method:"post" summary:"上传银行卡" tags:"工具"`
	model.BankCardWithOCRInput
}

type UploadBusinessLicenseWithOCRRes model.BusinessLicenseWithOCR

type IDCardWithOCRRes model.IDCardWithOCR
type UploadRes model.FileUploadOutput

type BankCardWithOCRRes model.BankCardWithOCR
