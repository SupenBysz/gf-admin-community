package sms

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
)

// 校验验证码

// 发送短信

// 接收短信

// 查询短信

// 短信签名管理

// 短信模版管理

// 渠道商管理

type SMSer interface {
	// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error)
	VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error)

	// SendSms 发送短信 (渠道商, 短信模版,请求内容)
	SendSms(ctx context.Context, provider sys_model.SmsServiceProviderConfig, template sys_model.SmsTemplateConfig, req sys_model.SmsSendMessageReq) (*sys_model.SmsResponse, error)

	// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
	ReceiveSms(ctx context.Context, req sys_model.SmsSendMessageReq) (bool, error)

	// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
	GetAppAvailableNumber(ctx context.Context, appNo string) int64

	// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
	UpdateAppNumber(ctx context.Context, appNo string, fee uint64) (bool, error)

	// RegisterTemplate 添加短信模版
	RegisterTemplate(ctx context.Context, signInfo *sys_model.SmsTemplateConfig) (bool, error)

	// AuditTemplate 短信模版审核
	AuditTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error)

	// GetByTemplateNo 根据模版编号查询模版信息
	GetByTemplateNo(ctx context.Context, templateNo string) (*sys_model.SmsTemplateConfig, error)

	// RegisterSign 添加短信签名
	RegisterSign(ctx context.Context, signInfo *sys_model.SmsSignConfig) (bool, error)

	// GetSignBySignName 根据签名名称查找签名数据
	GetSignBySignName(ctx context.Context, signName string) (*sys_model.SmsSignConfig, error)
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error)
func VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
	return true, nil
}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func SendSms(ctx context.Context, provider sys_model.SmsServiceProviderConfig, template sys_model.SmsTemplateConfig, req sys_model.SmsSendMessageReq) (*sys_model.SmsResponse, error) {
	return nil, nil
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func ReceiveSms(ctx context.Context, req sys_model.SmsSendMessageReq) (bool, error) {
	return true, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func GetAppAvailableNumber(ctx context.Context, appNo string) int64 {
	return 0
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func UpdateAppNumber(ctx context.Context, appNo string, fee uint64) (bool, error) {
	return false, nil
}

// RegisterTemplate 添加短信模版
func RegisterTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error) {
	return false, nil
}

// AuditTemplate 短信模版审核
func AuditTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error) {
	return true, nil
}

// GetByTemplateNo 根据模版编号查询模版信息
func GetByTemplateNo(ctx context.Context, templateNo string) (*sys_model.SmsTemplateConfig, error) {
	return nil, nil
}

// RegisterSign 添加短信签名
func RegisterSign(ctx context.Context, signInfo *sys_model.SmsSignConfig) (bool, error) {
	return false, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func GetSignBySignName(ctx context.Context, signName string) (*sys_model.SmsSignConfig, error) {
	return nil, nil
}
