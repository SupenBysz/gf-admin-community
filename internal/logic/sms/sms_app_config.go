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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

// 短信应用配置管理
type sSmsAppConfig struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSmsAppConfig(NewAppConfig())
}

func NewAppConfig() *sSmsAppConfig {
	return &sSmsAppConfig{
		cachePrefix: sys_dao.SmsAppConfig.Table() + "_",
	}
}

// GetAppConfigByNo 根据应用编号获取AppConfig
func (s *sSmsAppConfig) GetAppConfigByNo(ctx context.Context, appNo string) (*sys_model.SmsAppConfig, error) {
	if appNo == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "应用编号不能为空", sys_dao.SmsAppConfig.Table())
	}

	data := sys_entity.SmsAppConfig{}

	err := sys_dao.SmsAppConfig.Ctx(ctx).Where(sys_do.SmsAppConfig{AppNo: appNo}).Scan(&data)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据应用编号获取应用信息失败", sys_dao.SmsAppConfig.Table())
	}

	res := kconv.Struct[*sys_model.SmsAppConfig](data, &sys_model.SmsAppConfig{})

	return res, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余短信数量)
func (*sSmsAppConfig) GetAppAvailableNumber(ctx context.Context, appNo string) (int, error) {
	if appNo == "" {
		return 0, sys_service.SysLogs().ErrorSimple(ctx, nil, "应用编号不能为空", sys_dao.SmsAppConfig.Table())
	}

	data := sys_entity.SmsAppConfig{}

	err := sys_dao.SmsAppConfig.Ctx(ctx).Where(sys_do.SmsAppConfig{AppNo: appNo}).Scan(&data)
	if err != nil {
		return 0, sys_service.SysLogs().ErrorSimple(ctx, err, "根据应用编号获取应用信息失败", sys_dao.SmsAppConfig.Table())
	}

	num := data.AvailableNumber - data.UseNumber

	return num, nil
}

// CreateAppNumber 创建应用 (上下文, 应用编号, 花费数量)
func (s *sSmsAppConfig) CreateAppNumber(ctx context.Context, config *sys_model.SmsAppConfig) (bool, error) {
	// 编号查重
	count, _ := sys_dao.SmsAppConfig.Ctx(ctx).Where(sys_do.SmsAppConfig{
		AppNo: config.AppNo,
	}).Count()

	if count > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "应用编号重复", sys_dao.SmsAppConfig.Table())
	}

	// 应用名称查重
	count1, _ := sys_dao.SmsAppConfig.Ctx(ctx).Where(sys_do.SmsAppConfig{
		AppName: config.AppName,
	}).Count()

	if count1 > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "应用名称重复", sys_dao.SmsAppConfig.Table())
	}

	// 生成id
	appConfig := sys_do.SmsAppConfig{}
	gconv.Struct(config, &appConfig)
	appConfig.Id = idgen.NextId()
	appConfig.Status = 1

	_, err := sys_dao.SmsAppConfig.Ctx(ctx).Insert(appConfig)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "短信应用创建失败", sys_dao.SmsAppConfig.Table())
	}

	return true, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (s *sSmsAppConfig) UpdateAppNumber(ctx context.Context, appNo string, fee uint64) (bool, error) {
	if appNo == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "应用编号不能为空", sys_dao.SmsAppConfig.Table())
	}

	// 获取原来的数量
	appConfig, err := s.GetAppConfigByNo(ctx, appNo)
	if err != nil {
		return false, err
	}
	newUseNum := appConfig.UseNumber + gconv.Int64(fee)
	newAvailableNum := appConfig.AvailableNumber - gconv.Int64(fee)

	affected, err := daoctl.UpdateWithError(sys_dao.SmsAppConfig.Ctx(ctx).
		Data(sys_do.SmsAppConfig{UseNumber: newUseNum, AvailableNumber: newAvailableNum}).
		Where(sys_do.SmsAppConfig{AppNo: appNo}))

	return affected > 0, nil
}
