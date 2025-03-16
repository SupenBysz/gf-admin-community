package sdk_ctyun

import (
	"context"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 天翼云服务平台

type sSdkCtyun struct {
	CtyunSdkConfTokenList []*sys_model.CtyunSdkConfInfo
	sysConfigName         string
	conf                  gdb.CacheOption
}

// New SdkBaidu 系统配置逻辑实现
func New() sys_service.ISdkCtyun {
	return &sSdkCtyun{
		CtyunSdkConfTokenList: make([]*sys_model.CtyunSdkConfInfo, 0),
		sysConfigName:         "ctyun_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

func init() {
	sys_service.RegisterSdkCtyun(New())
}

// 我估计没有token的认证，后期可能需要添加一个签名函数

// GetCtyunSdkConfList 获取天翼云SDK应用配置列表
func (s *sSdkCtyun) GetCtyunSdkConfList(ctx context.Context) ([]*sys_model.CtyunSdkConf, error) {
	items := make([]*sys_model.CtyunSdkConf, 0)
	config, err := sys_service.SysConfig().GetByName(ctx, s.sysConfigName)
	if err != nil {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_ctyun_sdk_config_fetch_failed"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if config.Value == "" {
		return items, nil
	}

	_ = gjson.DecodeTo(config.Value, &items)

	return items, nil
}

// GetCtyunSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkCtyun) GetCtyunSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.CtyunSdkConf, err error) {
	items, err := s.GetCtyunSdkConfList(ctx)
	if err != nil {
		return nil, err
	}

	// 循环所有配置，筛选出符合条件的配置
	for _, conf := range items {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_ctyun_sdk_app_config_query_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveCtyunSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkCtyun) SaveCtyunSdkConf(ctx context.Context, info *sys_model.CtyunSdkConf, isCreate bool) (*sys_model.CtyunSdkConf, error) {
	oldItems, _ := s.GetCtyunSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.CtyunSdkConf, 0)
	for _, conf := range oldItems {
		if conf.Identifier == info.Identifier { // 如果标识符相等，说明已经存在， 将最新的追加到新的容器中
			isHas = true
			newItems = append(newItems, info)
			continue
		}

		newItems = append(newItems, conf) // 将旧的Item追加到新的容器中
	}

	if !isHas { // 不存在
		if isCreate { // 创建 --- 追加info （原有的 + 最新的Info）
			newItems = append(newItems, info)
		} else { // 更新 --- 不存在此配置，那么就提示错误
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_ctyun_sdk_config_save_failed_identifier_error"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)
	_, err := sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_ctyun_sdk_config_save_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步天翼云SDK应用配置缓存列表
	s.syncCtyunSdkConfList(ctx)

	return info, nil
}

// syncCtyunSdkConfList 同步天翼云SDK应用配置信息列表缓存  （代码中要是用到了s.CtyunSdkConfList缓存变量的话，一定需要在CUD操作后调用此方法更新缓存变量）
func (s *sSdkCtyun) syncCtyunSdkConfList(ctx context.Context) error {
	items, err := s.GetCtyunSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.CtyunSdkConfInfo, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.CtyunSdkConfTokenList { // tokenList
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.CtyunSdkConfTokenList = newTokenItems

	return nil
}

// DeleteCtyunSdkConf 删除天翼云SDK应用配置信息
func (s *sSdkCtyun) DeleteCtyunSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetCtyunSdkConfList(ctx)

	isHas := false
	newItems := garray.New(false)
	for _, conf := range items {
		if conf.Identifier == identifier {
			isHas = true
			continue
		}
		newItems.Append(conf)
	}

	if !isHas {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_ctyun_sdk_config_delete_not_exist", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err = sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_ctyun_sdk_config_delete_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步Token列表
	s.syncCtyunSdkConfList(ctx)

	return true, nil
}

// 天翼云服务的具体应用实例
