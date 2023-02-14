package sms

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

// 渠道商管理
type sSmsServiceProviderConfig struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSmsServiceProviderConfig(NewSmsServiceProviderConfig())
}

func NewSmsServiceProviderConfig() *sSmsServiceProviderConfig {
	return &sSmsServiceProviderConfig{
		cachePrefix: sys_dao.SmsServiceProviderConfig.Table() + "_",
	}
}

// CreateProvider 添加渠道商
func (t *sSmsServiceProviderConfig) CreateProvider(ctx context.Context, info *sys_model.SmsServiceProviderConfig) (bool, error) {
	model := sys_dao.SmsServiceProviderConfig.Ctx(ctx)
	// 判断渠道商是否已经存在
	count, _ := model.Where(sys_do.SmsServiceProviderConfig{
		ProviderNo: info.ProviderNo,
	}).Count()

	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信渠道商已存在", sys_dao.SmsServiceProviderConfig.Table())
	}

	// 插入渠道商配置信息
	data := sys_do.SmsServiceProviderConfig{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.CreatedAt = gtime.Now()

	_, err := model.OmitNilData().Insert(data)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信渠道商添加失败", sys_dao.SmsServiceProviderConfig.Table())
	}

	return true, nil
}

// GetProviderByNo 根据No编号获取渠道商
func (t *sSmsServiceProviderConfig) GetProviderByNo(ctx context.Context, no string) (*sys_model.SmsServiceProviderConfig, error) {
	if no == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "渠道商编号不能为空", sys_dao.SmsServiceProviderConfig.Table())
	}

	data := sys_entity.SmsServiceProviderConfig{}

	err := sys_dao.SmsServiceProviderConfig.Ctx(ctx).Where(sys_do.SmsServiceProviderConfig{ProviderNo: no}).Scan(&data)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据应用编号获取渠道商信息失败", sys_dao.SmsServiceProviderConfig.Table())
	}

	res := kconv.Struct[*sys_model.SmsServiceProviderConfig](data, &sys_model.SmsServiceProviderConfig{})

	return res, nil
}

// QueryProviderList 获取渠道商列表
func (t *sSmsServiceProviderConfig) QueryProviderList(ctx context.Context, search *sys_model.SearchParams, isExport bool) (*sys_model.ServiceProviderConfigListRes, error) {
	result, err := daoctl.Query[*sys_model.SmsServiceProviderConfig](sys_dao.SmsServiceProviderConfig.Ctx(ctx), search, isExport)

	return (*sys_model.ServiceProviderConfigListRes)(result), err
}
