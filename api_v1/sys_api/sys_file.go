package sys_api

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
type UploadFileInfoRes sys_model.FileInfo

type BankCardWithOCRRes sys_model.BankCardWithOCR

type GetFileByIdReq struct {
	g.Meta `path:"/getFileById" method:"get" summary:"通过id获取文件" tags:"工具"`
	Id     int64 `json:"id" v:"required#资源ID错误" dc:"文件资源ID"`
}

type UploadFileRes sys_model.UploadFile

type GetFileReq struct {
	g.Meta `path:"/getFile" method:"get" summary:"获取图片" tags:"工具"`
	Sign   string `json:"sign" dc:"签名数据，组成部分：(srcBase64 + srcMd5 + fileId) ==> md5加密"`
	Path   string `json:"path" dc:"文件Src的base64编码后的数据"`
	Id     int64  `json:"id" dc:"文件id"`
	CId    int64  `json:"cid" dc:"缓存id"`
}
