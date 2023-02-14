package sms

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信签名管理
type sSmsSignConfig struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSmsSignConfig(NewSmsSignConfig())
}

func NewSmsSignConfig() *sSmsSignConfig {
	return &sSmsSignConfig{
		cachePrefix: sys_dao.SmsSignConfig.Table() + "_",
	}
}

// CreateSign 添加短信签名
func (s *sSmsSignConfig) CreateSign(ctx context.Context, info *sys_model.SmsSignConfig) (bool, error) {
	// 根据渠道商编号判断渠道商是否存在
	providerData, err := sys_service.SmsServiceProviderConfig().GetProviderByNo(ctx, info.ProviderNo)

	if err != nil || providerData == nil {
		return false, err
	}

	// 添加短信签名
	data := kconv.Struct(info, &sys_do.SmsSignConfig{})

	data.Id = idgen.NextId()
	// 未审核的短信签名是禁用状态
	data.Status = 0

	_, err = sys_dao.SmsSignConfig.Ctx(ctx).OmitNilData().Insert(data)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信签名添加失败", sys_dao.SmsSignConfig.Table())
	}

	return true, nil
}

// AuditSign 审核短信签名, 将短信签名Status状态改变为1
func (s *sSmsSignConfig) AuditSign(ctx context.Context, id int64, info *sys_model.AuditInfo) (bool, error) {
	// 判断审核行为，只能是审核通过或者不通过 -1不通过 1通过
	if info.State != sys_enum.Sms.Action.Reject.Code() && info.State != sys_enum.Sms.Action.Approve.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", sys_dao.SmsSignConfig.Table())
	}

	// 审核不通过需要有原因
	if info.State == sys_enum.Sms.Action.Reject.Code() && info.ReplyMsg == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", sys_dao.SmsSignConfig.Table())
	}

	// 代表已审过的
	if info.State > sys_enum.Sms.Action.WaitReview.Code() {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", sys_dao.SmsSignConfig.Table())
	}

	// 判断签名是否存在
	sign, err := s.GetSignById(ctx, id)
	if err != nil || sign == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信签名不存在", sys_dao.SmsSignConfig.Table())
	}

	// 改变状态为正常代表审核成功
	_, err = sys_dao.SmsSignConfig.Ctx(ctx).OmitNilData().Data(sys_do.SmsSignConfig{
		AuditUserId:   info.AuditUserId,
		AuditReplyMsg: info.ReplyMsg,
		AuditAt:       gtime.Now(),
		Status:        info.State,
	}).Where(sys_do.SmsSignConfig{
		Id: id,
	}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "短信签名审核失败", sys_dao.SmsSignConfig.Table())
	}

	return true, nil
}

// GetSignBySignName 根据签名名称查找签名数据
func (s *sSmsSignConfig) GetSignBySignName(ctx context.Context, signName string) (res *sys_model.SmsSignConfig, err error) {
	err = sys_dao.SmsSignConfig.Ctx(ctx).Where(sys_do.SmsSignConfig{
		SignName: signName,
	}).Scan(&res)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该签名不存在", sys_dao.SmsSignConfig.Table())
	}
	return res, nil
}

// GetSignById 根据id查找签名数据
func (s *sSmsSignConfig) GetSignById(ctx context.Context, id int64) (*sys_model.SmsSignConfig, error) {
	if id == 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "签名id不能为空", sys_dao.SmsSignConfig.Table())
	}

	result, err := daoctl.GetByIdWithError[*sys_model.SmsSignConfig](sys_dao.SmsSignConfig.Ctx(ctx), id)

	return *result, err
}
