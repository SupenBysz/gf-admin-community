package sdk_huawei

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// 华为云服务平台

type sSdkHuawei struct {
	HuaweiSdkConfTokenList []*sys_model.HuaweiSdkConfToken
	sysConfigName          string
	conf                   gdb.CacheOption
}

// New SdkBaidu 系统配置逻辑实现
func New() sys_service.ISdkHuawei {
	return &sSdkHuawei{
		HuaweiSdkConfTokenList: make([]*sys_model.HuaweiSdkConfToken, 0),
		sysConfigName:          "huawei_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

func init() {
	sys_service.RegisterSdkHuawei(New())
}

// fetchHuaweiSdkConfToken 根据 identifier 获取华为云API Token  （API获取方式）
func (s *sSdkHuawei) fetchHuaweiSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.HuaweiSdkConfToken, err error) {
	info, err := s.GetHuaweiSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}
	client := g.Client()

	// URL
	// var host = "https://iam.myhuaweicloud.com/v3/auth/tokens?nocatalog=true"
	var host = "https://isdp+域名/oauth2/oauth/rest_token"

	// 请求头
	// Host 请求的服务器URL
	// Content-type: application/json
	header := make(map[string]string)

	header["Content-type"] = "application/json"

	client.Header(header)

	// 请求数据，是一个复杂的结构体
	param := g.Map{
		// 客户端id (固定值isdp-saas-openapi)
		"client_id": "isdp-saas-openapi",
		// 客户端秘钥 (固定值isdp-saas-openapi)
		"client_secret": info.SecretKey,
		// 授权方式(固定值)
		"grant_type": "password",
		// 订阅的应用ID
		"username": info.AppID,
		// 应用实例令牌
		"password": info.APIKey,
	}

	response, err := client.Post(ctx, host, param)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取华为云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 接受返回数据，json解析
	newTokenInfo := sys_model.HuaweiAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取华为云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *sys_model.HuaweiSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.HuaweiSdkConfTokenList {
		if item.Identifier == identifier {
			result = &sys_model.HuaweiSdkConfToken{
				HuaweiSdkConf:     *info,
				HuaweiAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &sys_model.HuaweiSdkConfToken{
			HuaweiSdkConf:     *info,
			HuaweiAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	// 返回我们需要的token信息
	return result, nil
}

// GetHuaweiSdkConfList 获取阿里云SDK应用配置列表
func (s *sSdkHuawei) GetHuaweiSdkConfList(ctx context.Context) ([]*sys_model.HuaweiSdkConf, error) {
	items := make([]*sys_model.HuaweiSdkConf, 0)
	config, err := sys_service.SysConfig().GetByName(ctx, s.sysConfigName)
	if err != nil {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("华为SDK配置信息获取失败"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if config.Value == "" {
		return items, nil
	}

	_ = gjson.DecodeTo(config.Value, &items)

	return items, nil
}

// GetHuaweiSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkHuawei) GetHuaweiSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.HuaweiSdkConf, err error) {
	items, err := s.GetHuaweiSdkConfList(ctx)
	if err != nil {
		return nil, err
	}

	// 循环所有配置，筛选出符合条件的配置
	for _, conf := range items {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询华为云SDK应用配置信息失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveHuaweiSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkHuawei) SaveHuaweiSdkConf(ctx context.Context, info *sys_model.HuaweiSdkConf, isCreate bool) (*sys_model.HuaweiSdkConf, error) {
	oldItems, _ := s.GetHuaweiSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.HuaweiSdkConf, 0)
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
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("华为云SDK配置信息保存失败，标识符错误"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)
	_, err := sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "华为云SDK配置信息保存失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步华为云SDK应用配置缓存列表
	s.syncHuaweiSdkConfList(ctx)

	return info, nil
}

// syncHuaweiSdkConfList 同步华为SDK应用配置信息列表缓存  （代码中要是用到了s.HuaweiSdkConfList缓存变量的话，一定需要在CUD操作后调用此方法更新缓存变量）
func (s *sSdkHuawei) syncHuaweiSdkConfList(ctx context.Context) error {
	items, err := s.GetHuaweiSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.HuaweiSdkConfToken, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.HuaweiSdkConfTokenList { // tokenList
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.HuaweiSdkConfTokenList = newTokenItems

	return nil
}

// DeleteHuaweiSdkConf 删除华为SDK应用配置信息
func (s *sSdkHuawei) DeleteHuaweiSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetHuaweiSdkConfList(ctx)

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
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "要删除的华为云SDK配置信息不存在", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err = sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "华为云SDK配置信息删除失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步Token列表
	s.syncHuaweiSdkConfList(ctx)

	return true, nil
}

// 华为云服务的具体应用实例

// 华为云服务的具体应用实例
