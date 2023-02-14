package sms

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
)

// SmsTest 企业信使
type SmsTest struct {
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error) （不属于接口，属于调用者校验）
//func (t *SmsTest) VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
//	return true, nil
//}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (t *SmsTest) SendSms(ctx context.Context, provider sys_model.SmsServiceProviderConfig, template sys_model.SmsTemplateConfig, req sys_model.SmsSendMessageReq) (*sys_model.SmsResponse, error) {
	return nil, nil
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (t *SmsTest) ReceiveSms(ctx context.Context, req sys_model.SmsSendMessageReq) (bool, error) {
	return true, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func (t *SmsTest) GetAppAvailableNumber(ctx context.Context, appNo string) (int, error) {
	return 0, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (t *SmsTest) UpdateAppNumber(ctx context.Context, appNo string, fee uint64) (bool, error) {
	return false, nil
}

// RegisterTemplate 添加短信模版
func (t *SmsTest) RegisterTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error) {
	return false, nil
}

// AuditTemplate 短信模版审核
func (t *SmsTest) AuditTemplate(ctx context.Context, id int64, info *sys_model.AuditInfo) (bool, error) {
	return true, nil
}

// GetByTemplateNo 根据模版编号查询模版信息
func (t *SmsTest) GetByTemplateNo(ctx context.Context, templateNo string) (*sys_model.SmsTemplateConfig, error) {
	return nil, nil
}

// RegisterSign 添加短信签名
func (t *SmsTest) RegisterSign(ctx context.Context, signInfo *sys_model.SmsSignConfig) (bool, error) {
	return false, nil
}

// AuditSign 审核短信签名
func (t *SmsTest) AuditSign(ctx context.Context, id int64, info *sys_model.AuditInfo) (bool, error) {
	return false, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func (t *SmsTest) GetSignBySignName(ctx context.Context, signName string) (*sys_model.SmsSignConfig, error) {
	return nil, nil
}

// CreateProvider 添加渠道商
func (t *SmsTest) CreateProvider(ctx context.Context, info *sys_model.SmsServiceProviderConfig) (bool, error) {
	return false, nil
}

// GetProviderByNo 根据No编号获取渠道商
func (t *SmsTest) GetProviderByNo(ctx context.Context, no string) (*sys_model.SmsServiceProviderConfig, error) {
	return nil, nil
}

// QueryProviderList 获取渠道商列表
func (t *SmsTest) QueryProviderList(ctx context.Context, search *sys_model.SearchParams, isExport bool) (*[]sys_model.SmsServiceProviderConfig, error) {
	return nil, nil
}
