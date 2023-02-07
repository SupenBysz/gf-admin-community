package sms

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type Test struct {
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error)
func (t *Test) VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
	return true, nil
}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (t *Test) SendSms(ctx context.Context, provider sys_model.SmsServiceProviderConfig, template sys_model.SmsTemplateConfig, req sys_model.SmsSendMessageReq) (*sys_model.SmsResponse, error) {
	return nil, nil
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (t *Test) ReceiveSms(ctx context.Context, req sys_model.SmsSendMessageReq) (bool, error) {
	return true, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func (t *Test) GetAppAvailableNumber(ctx context.Context, appNo string) int64 {
	return 0
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (t *Test) UpdateAppNumber(ctx context.Context, appNo string, fee uint64) (bool, error) {
	return false, nil
}

// RegisterTemplate 添加短信模版
func (t *Test) RegisterTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error) {
	return false, nil
}

// AuditTemplate 短信模版审核
func (t *Test) AuditTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error) {
	return true, nil
}

// GetByTemplateNo 根据模版编号查询模版信息
func (t *Test) GetByTemplateNo(ctx context.Context, templateNo string) (*sys_model.SmsTemplateConfig, error) {
	return nil, nil
}

// RegisterSign 添加短信签名
func (t *Test) RegisterSign(ctx context.Context, signInfo *sys_model.SmsSignConfig) (bool, error) {
	return false, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func (t *Test) GetSignBySignName(ctx context.Context, signName string) (*sys_model.SmsSignConfig, error) {
	return nil, nil
}
