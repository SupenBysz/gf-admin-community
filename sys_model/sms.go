package sys_model

import "github.com/gogf/gf/v2/os/gtime"

// SmsSendMessageReq 发送短信请求对象
type SmsSendMessageReq struct {
	BusinessNo string   `json:"businessNo"  dc:"服务编号"`
	TemplateNo string   `json:"templateNo"  dc:"模版编号"`
	Phones     []string `json:"phones"  dc:"手机号集合"`
	Params     []string `json:"params"  dc:"参数列表"`
}

// SmsResponse 短信请求响应体
type SmsResponse struct {
	SmsSendStatus []SmsSendStatus `json:"sms_send_status" dc:"响应集合"`
	RequestId     string          `json:"request_id" dc:"请求ID"`
}

type SmsSendStatus struct {
	Fee     uint64 `json:"fee" dc:"计费条数"`
	Message string `json:"message" dc:"响应信息"`
	Code    string `json:"code" dc:"响应消息"`
	Phone   string `json:"phone" dc:"手机号码"`
}

// SmsServiceProviderConfig 渠道商来源
type SmsServiceProviderConfig struct {
	ProviderName    string `json:"providerName" `
	ProviderNo      string `json:"providerNo" dc:"渠道商编号"`
	AccessKeyId     string `json:"accessKeyId" dc:"身份标识"`
	AccessKeySecret string `json:"accessKeySecret" dc:"身份认证密钥"`
	Endpoint        string `json:"endpoint" dc:"调用域名"`
	SdkAppId        string `json:"sdkAppId" dc:"应用ID"`
	Region          string `json:"region" dc:"地域列表"`
	Remark          string `json:"remark" dc:"备注"`
	Status          int    `json:"status" dc:"状态: 0禁用 1正常"`
}

// SmsTemplateConfig 短信模版
type SmsTemplateConfig struct {
	SignName             string      `json:"signName" dc:"签名名称"`
	TemplateNo           string      `json:"templateNo" dc:"模版编号"`
	TemplateName         string      `json:"templateName" dc:"模版名称"`
	TemplateContent      string      `json:"templateContent" dc:"模版内容"`
	ThirdPartyTemplateNo string      `json:"thirdPartyTemplateNo" dc:"第三方模版编号"`
	ProviderNo           string      `json:"providerNo" dc:"渠道商编号"`
	Remark               string      `json:"remark" dc:"备注"`
	Status               int         `json:"status"        dc:"状态: -1不通过 0待审核 1正常"`
	AuditUserId          int64       `json:"auditUserId"   dc:"审核者UserID"`
	AuditReplyMsg        string      `json:"auditReplyMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt              *gtime.Time `json:"auditAt"       dc:"审核时间"`
}

// SmsSignConfig 短信签名
type SmsSignConfig struct {
	SignName      string      `json:"signName" dc:"签名名称"`
	ProviderNo    string      `json:"providerNo" dc:"渠道商编号"`
	ProviderName  string      `json:"providerName" dc:"渠道商名称"`
	Remark        string      `json:"remark" dc:"备注"`
	Status        int         `json:"status"        dc:"状态: -1不通过 0待审核 1正常"`
	AuditUserId   int64       `json:"auditUserId"   dc:"审核者UserID"`
	AuditReplyMsg string      `json:"auditReplyMsg" dc:"审核回复，仅审核不通过时才有值"`
	AuditAt       *gtime.Time `json:"auditAt"       dc:"审核时间"`
}

// SmsBusinessConfig 短信业务
type SmsBusinessConfig struct {
	AppNo        string `json:"appNo" dc:"应用ID"`
	BusinessName string `json:"businessName" dc:"业务名称"`
	BusinessNo   string `json:"businessNo" dc:"业务编号"`
	TemplateNo   string `json:"templateNo" dc:"模版编号"`
	BusinessDesc string `json:"businessDesc" dc:"业务说明"`
	Remark       string `json:"remark" dc:"备注"`
	Status       int    `json:"status" dc:"状态: 0禁用 1正常"`
}

// SmsAppConfig 短信应用
type SmsAppConfig struct {
	AppNo           string `json:"appNo" dc:"应用编号"`
	AppName         string `json:"appName" dc:"应用名称"`
	AvailableNumber int64  `json:"availableNumber" dc:"可用数量"`
	CurrentLimiting int64  `json:"currentLimiting" dc:"限流数量"`
	UseNumber       int64  `json:"useNumber" dc:"已用数量"`
	Remark          string `json:"remark" dc:"备注"`
	Status          int    `json:"status" dc:"状态: 0禁用 1正常"`
}

// SmsSendLog 短信发送日志
type SmsSendLog struct {
	AppNo       string `json:"appNo" dc:"应用ID"`
	BusinessNo  string `json:"businessNo" dc:"业务编号"`
	Status      int    `json:"status" dc:"状态"`
	Fee         string `json:"fee" dc:"计价条数"`
	PhoneNumber string `json:"phoneNumber" dc:"发送手机号"`
	Message     string `json:"message" dc:"接口响应消息"`
	Code        string `json:"code" dc:"接口响应状态码"`
	Content     string `json:"content" dc:"发送内容"`
	Remark      string `json:"remark" dc:"备注"`
}

// AuditInfo 审核
type AuditInfo struct {
	State       int    `json:"state" dc:"审核状态"`
	AuditUserId int64  `json:"auditUserId" dc:"审核者UserId"`
	ReplyMsg    string `json:"replyMsg" dc:"审核失败时必填的原因回复"`
}

type ServiceProviderConfigListRes CollectRes[*SmsServiceProviderConfig]
