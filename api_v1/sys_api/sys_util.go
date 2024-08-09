package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*
	第三方能力：
	 - nlp 自然语言处理
	 - 实名核身鉴权
     - 活体人脸核身
     - 人脸核身
*/

// *********************************** 工具 *****************************************

type GetNerChEcomReq struct {
	g.Meta `path:"/getNerChEcom" method:"POST" tags:"工具/中文分词" summary:"中文分词"`
	Text   string `json:"text" dc:"需要分词的内容"`
}

// ************************************ 认证 *****************************************

type LivenessRecognitionReq struct {
	g.Meta       `path:"/livenessRecognition" method:"POST" tags:"认证/身份验证" summary:"活体人脸核身"`
	IdCard       string `json:"idCard" dc:"身份证号码" v:"required#身份证号码不能为空"`
	Name         string `json:"name" dc:"姓名" v:"required#姓名不能为空"`
	LivenessType string `json:"LivenessType" dc:"活体检测类型，取值：LIP/ACTION/SILENT。LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。" default:"SILENT"`
}

type DetectAuthReq struct {
	g.Meta    `path:"/detectAuth" method:"POST" tags:"认证/身份验证" summary:"实名核身鉴权"`
	IdCard    string `json:"idCard" dc:"身份证号码" v:"required#身份证号码不能为空"`
	Name      string `json:"name" dc:"姓名" v:"required#姓名不能为空"`
	ReturnUrl string `json:"returnUrl" dc:"认证结束后重定向的回调链接地址。最长长度1024位。" `
	//RuleId   string `json:"ruleId" dc:"规则Id" v:"required#规则Id不能为空"`

	//LivenessType string `json:"LivenessType" dc:"活体检测类型，取值：LIP/ACTION/SILENT。LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。" default:"SILENT"`
}

type GetDetectAuthResultReq struct {
	g.Meta   `path:"/getDetectAuthResult" method:"POST" tags:"认证/身份验证" summary:"获取实名核身鉴权结果"`
	BizToken string `json:"bizToken" dc:"人脸核身流程的标识，调用DetectAuth接口时生成。" v:"required#人脸核身的查询标识不能为空"`
	//RuleId   string `json:"ruleId" dc:"规则Id" v:"required#规则Id不能为空"`
}

type StartAdvFaceAuthReq struct {
	g.Meta      `path:"/faceAuth" method:"POST" tags:"认证/人脸核身" summary:"启动H5人脸核身"`
	IdCard      string `json:"idCard" dc:"身份证号码" v:"required#身份证号码不能为空"`
	Name        string `json:"name" dc:"姓名" v:"required#姓名不能为空"`
	CallbackUrl string `json:"callbackUrl" dc:"认证结束后重定向的回调链接地址。最长长度1024位。" `
}

type FaceAuthCallbackReq struct {
	g.Meta       `path:"/startAdvFaceAuth" method:"GET" tags:"认证/人脸核身" summary:"人脸核身回调地址"`
	Code         string `json:"code" dc:"人脸核身结果的返回码，0 表示人脸核身成功，其他错误码表示失败。" `
	OrderNo      string `json:"orderNo" dc:"本次人脸核身上送的订单号。" `
	H5faceId     string `json:"h5faceId" dc:"本次请求返回的唯一标识，此信息为本次人脸核身上传的信息。" `
	NewSignature string `json:"newSignature" dc:"对 URL 参数 App ID、orderNo 和 SIGN ticket、code 的签名。" `
	LiveRate     string `json:"liveRate" dc:"本次人脸核身的活体检测分数。" `
}

type QueryFaceRecordReq struct {
	g.Meta  `path:"/queryFaceRecord" method:"POST" tags:"认证/人脸核身" summary:"人脸核身结果查询"`
	OrderNo string `json:"orderNo" dc:"订单号" v:"required#人脸核身订单号不能为空"`
}
