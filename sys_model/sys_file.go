package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileUploadInput 上传文件输入参数
type FileUploadInput struct {
	File       *ghttp.UploadFile `json:"sys_file"  dc:"请选择文件，以form-data方式提交"`    // 上传文件对象
	Name       string            `json:"name" dc:"自定义文件名"`                       // 自定义文件名称
	RandomName bool              `json:"randomName" dc:"是否随机文件名" default:"true"` // 是否随机命名文件
}

type FileInfo struct {
	sys_entity.SysFile
	ExpiresAt *gtime.Time
}

type UploadFile sys_entity.SysFile

// OCRIDCardFileUploadInput 上传身份证请求参数
type OCRIDCardFileUploadInput struct {
	FileUploadInput
	DetectRisk string `json:"detectRisk" default:"false" v:"in:true,false" dc:"是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)检测功能，默认不开启"`
	IDCardSide string `json:"idCardSide" default:"front" v:"in:front,back#参数错误|正反面参数错误" dc:"front身份证含照片的一面,back身份证带国徽的一面,自动检测身份证正反面，如果传参指定方向与图片相反，支持正常识别，返回参数image_status字段为 reversed_side "`
}

// IDCardWithOCR 身份证识别响应信息
type IDCardWithOCR struct {
	sys_entity.SysFile
	BaiduSdkOCRIDCard
	// OCRBusinessLicense OCRBusinessLicense `json:"orcBusinessLicense" dc:"营业执照识别的信息"`
}

// OCRBusinessLicense 上传新版营业执照
type OCRBusinessLicense struct {
	FileUploadInput
}

// BusinessLicenseWithOCR 营业执照识别响应信息
type BusinessLicenseWithOCR struct {
	sys_entity.SysFile
	BusinessLicenseOCR
}

// BankCardWithOCRInput 上传银行卡请求参数
type BankCardWithOCRInput struct {
	FileUploadInput
}

// BankCardWithOCR 银行卡识别响应信息
type BankCardWithOCR struct {
	// 上传文件返回数据
	sys_entity.SysFile

	// SDK返回的识别数据
	BaiduSdkOCRBankCard
}

// PictureWithOCRInput 上传图片
type PictureWithOCRInput struct {
	FileUploadInput
	ImageType uint64 `json:"imageType" dc:"图片类型：0静态图片,1GIF动态图片（仅对首帧进行审核）" default:"0" v:"in:0,1#请选择正确的数据范围"`
}

// PictureWithOCR 图片审核响应信息
type PictureWithOCR struct {
	sys_entity.SysFile
	Conclusion     string            `json:"conclusion" dc:"审核结果，可取值描述：合规、不合规、疑似、审核失败"`
	ConclusionType int               `json:"conclusionType" dc:"审核结果类型，可取值1、2、3、4，分别代表1：合规，2：不合规，3：疑似，4：审核失败："`
	Data           []DescriptionData `json:"data" dc:"不合规/疑似/命中白名单项详细信息。响应成功并且conclusion为疑似或不合规或命中白名单时才返回，响应失败或conclusion为合规且未命中白名单时不返回。"`
}

// DescriptionData 图片不合规描述
type DescriptionData struct {
	Type           int    `json:"type"`
	SubType        int    `json:"subType"`
	Conclusion     string `json:"conclusion" dc:"审核结果，可取值描述：合规、不合规、疑似、审核失败"`
	ConclusionType int    `json:"conclusionType" dc:"审核结果类型，可取值1、2、3、4，分别代表1：合规，2：不合规，3：疑似，4：审核失败："`
	Msg            string `json:"msg"`
}
