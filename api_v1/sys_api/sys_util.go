package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*
	第三方能力：
	 - nlp 自然语言处理
*/

type GetNerChEcomReq struct {
	g.Meta `path:"/getNerChEcom" method:"POST" tags:"工具/中文分词" summary:"中文分词"`
	Text   string `json:"text" dc:"需要分词的内容"`
}

type LivenessRecognitionReq struct {
	g.Meta       `path:"/livenessRecognition" method:"POST" tags:"工具/身份验证" summary:"活体人脸核身"`
	IdCard       string `json:"idCard" dc:"身份证号码" v:"required#身份证号码不能为空"`
	Name         string `json:"name" dc:"姓名" v:"required#姓名不能为空"`
	LivenessType string `json:"LivenessType" dc:"活体检测类型，取值：LIP/ACTION/SILENT。LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。" default:"SILENT"`
}

type DetectAuthReq struct {
	g.Meta    `path:"/detectAuth" method:"POST" tags:"工具/身份验证" summary:"实名核身鉴权"`
	IdCard    string `json:"idCard" dc:"身份证号码" v:"required#身份证号码不能为空"`
	Name      string `json:"name" dc:"姓名" v:"required#姓名不能为空"`
	ReturnUrl string `json:"returnUrl" dc:"认证结束后重定向的回调链接地址。最长长度1024位。" `
	//RuleId   string `json:"ruleId" dc:"规则Id" v:"required#规则Id不能为空"`

	//LivenessType string `json:"LivenessType" dc:"活体检测类型，取值：LIP/ACTION/SILENT。LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。" default:"SILENT"`
}

type GetDetectAuthResultReq struct {
	g.Meta   `path:"/getDetectAuthResult" method:"POST" tags:"工具/身份验证" summary:"获取实名核身鉴权结果"`
	BizToken string `json:"bizToken" dc:"人脸核身流程的标识，调用DetectAuth接口时生成。" v:"required#人脸核身的查询标识不能为空"`
	//RuleId   string `json:"ruleId" dc:"规则Id" v:"required#规则Id不能为空"`
}
