package sms

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信模板管理
type sSmsTemplateConfig struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSmsTemplateConfig(NewSmsTemplateConfig())
}

func NewSmsTemplateConfig() *sSmsTemplateConfig {
	return &sSmsTemplateConfig{
		cachePrefix: sys_dao.SmsTemplateConfig.Table() + "_",
	}
}

// CreateTemplate 添加短信模版
func (s *sSmsTemplateConfig) CreateTemplate(ctx context.Context, info *sys_model.SmsTemplateConfig) (bool, error) {
	// 根据渠道商编号判断渠道商是否存在
	providerData, err := sys_service.SmsServiceProviderConfig().GetProviderByNo(ctx, info.ProviderNo)

	if err != nil || providerData == nil {
		return false, err
	}

	// 判断签名是否存在
	signInfo, err := sys_service.SmsSignConfig().GetSignBySignName(ctx, info.SignName)
	if err != nil || signInfo == nil {
		return false, err
	}

	// 添加短信模版
	data := kconv.Struct(info, &sys_do.SmsTemplateConfig{})

	data.Id = idgen.NextId()
	// 未审核的短信模版是禁用状态
	data.Status = 0

	_, err = sys_dao.SmsTemplateConfig.Ctx(ctx).OmitNilData().Insert(data)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信模板添加失败", sys_dao.SmsTemplateConfig.Table())
	}

	return true, nil
}

// AuditTemplate 短信模版审核
func (s *sSmsTemplateConfig) AuditTemplate(ctx context.Context, id int64, info *sys_model.AuditInfo) (bool, error) {
	// 判断审核行为，只能是审核通过或者不通过 -1不通过 1通过
	if info.State != sys_enum.Sms.Action.Reject.Code() && info.State != sys_enum.Sms.Action.Approve.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", sys_dao.SmsTemplateConfig.Table())
	}

	// 审核不通过需要有原因
	if info.State == sys_enum.Sms.Action.Reject.Code() && info.ReplyMsg == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", sys_dao.SmsTemplateConfig.Table())
	}

	// 代表已审过的
	if info.State > sys_enum.Sms.Action.WaitReview.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", sys_dao.SmsTemplateConfig.Table())
	}

	// 判断签名是否存在
	sign, err := sys_service.SmsSignConfig().GetSignById(ctx, id)
	if err != nil || sign == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信签名不存在", sys_dao.SmsSignConfig.Table())
	}

	// 改变状态为正常代表审核成功
	_, err = sys_dao.SmsTemplateConfig.Ctx(ctx).OmitNilData().Data(sys_do.SmsTemplateConfig{
		AuditUserId:   info.AuditUserId,
		AuditReplyMsg: info.ReplyMsg,
		AuditAt:       gtime.Now(),
		Status:        info.State,
	}).Where(sys_do.SmsTemplateConfig{
		Id: id,
	}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "短信签名审核失败", sys_dao.SmsTemplateConfig.Table())
	}

	return true, nil
}

// GetByTemplateNo 根据模版编号查询模版信息
func (s *sSmsTemplateConfig) GetByTemplateNo(ctx context.Context, templateNo string) (*sys_model.SmsTemplateConfig, error) {
	if templateNo == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "模版编号不能为空", sys_dao.SmsTemplateConfig.Table())
	}

	data := sys_entity.SmsTemplateConfig{}

	err := sys_dao.SmsTemplateConfig.Ctx(ctx).Where(sys_do.SmsTemplateConfig{ProviderNo: templateNo}).Scan(&data)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据模版编号获取模版信息失败", sys_dao.SmsTemplateConfig.Table())
	}

	result := kconv.Struct[*sys_model.SmsTemplateConfig](data, &sys_model.SmsTemplateConfig{})

	return result, nil
}
