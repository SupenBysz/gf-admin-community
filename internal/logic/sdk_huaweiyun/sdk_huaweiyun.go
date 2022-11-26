package sdk_huaweiyun

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// 华为云服务平台

type sSdkHuaWeiYun struct {
	HuaWeiSdkConfTokenList []model.HuaWeiYunSdkConfToken
	CacheDuration          time.Duration
	sysConfigName          string
}

// New SdkBaidu 系统配置逻辑实现
func New() *sSdkHuaWeiYun {
	return &sSdkHuaWeiYun{
		HuaWeiSdkConfTokenList: make([]model.HuaWeiYunSdkConfToken, 0),
		CacheDuration:          time.Hour,
		sysConfigName:          "huawei_sdk_conf",
	}
}

func init() {
	service.RegisterSdkHuaWeiYun(New())
}

// fetchHuaWeiYUnSdkConfToken 根据 identifier 获取华为云API Token  （API获取方式）
func (s *sSdkHuaWeiYun) fetchHuaWeiYunSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.HuaWeiYunSdkConfToken, err error) {
	info, err := s.GetHuaWeiYunSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}
	client := g.Client()

	// URL
	//var host = "https://iam.myhuaweicloud.com/v3/auth/tokens?nocatalog=true"
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
		return nil, service.SysLogs().ErrorSimple(ctx, err, "获取华为云API Token 失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 接受返回数据，json解析
	newTokenInfo := model.HuaWeiYunAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "获取华为云API Token 失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *model.HuaWeiYunSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.HuaWeiSdkConfTokenList {
		if item.Identifier == identifier {
			result = &model.HuaWeiYunSdkConfToken{
				HuaWeiYunSdkConf:     *info,
				HuaWeiYunAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &model.HuaWeiYunSdkConfToken{
			HuaWeiYunSdkConf:     *info,
			HuaWeiYunAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	// 返回我们需要的token信息
	return result, nil
}

// GetHuaWeiYunSdkConfToken 根据 identifier 查询华为SDK应用配置和Token信息
func (s *sSdkHuaWeiYun) GetHuaWeiYunSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.HuaWeiYunSdkConfToken, err error) {
	for _, conf := range s.HuaWeiSdkConfTokenList {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}
	return s.fetchHuaWeiYunSdkConfToken(ctx, identifier)
}

// syncHuaWeiYunSdkConfTokenList 同步华为云SDK应用配置信息Token列表缓存
func (s *sSdkHuaWeiYun) syncHuaWeiYunSdkConfTokenList(ctx context.Context) error {
	items, err := s.GetHuaWeiYunSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]model.HuaWeiYunSdkConfToken, 0)
	for _, conf := range *items {
		for _, tokenInfo := range s.HuaWeiSdkConfTokenList {
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.HuaWeiSdkConfTokenList = newTokenItems

	return nil
}

// GetHuaWeiYunSdkConfList 获取华为云SDK应用配置列表
func (s *sSdkHuaWeiYun) GetHuaWeiYunSdkConfList(ctx context.Context) (*[]model.HuaWeiYunSdkConf, error) {
	items := make([]model.HuaWeiYunSdkConf, 0)

	data := entity.SysConfig{}

	err := dao.SysConfig.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.CacheDuration,
		Name:     s.sysConfigName,
		Force:    true,
	}).Where(do.SysConfig{
		Name: s.sysConfigName,
	}).Scan(&data)

	if err != nil && err != sql.ErrNoRows {
		return &items, service.SysLogs().ErrorSimple(ctx, gerror.New("华为云 SDK配置信息获取失败"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if data.Value == "" {
		return &items, nil
	}

	if nil == gjson.DecodeTo(data.Value, &items) {
		return &items, nil
	}

	return &items, nil
}

// GetHuaWeiYunSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkHuaWeiYun) GetHuaWeiYunSdkConf(ctx context.Context, identifier string) (tokenInfo *model.HuaWeiYunSdkConf, err error) {
	items, err := s.GetHuaWeiYunSdkConfList(ctx)

	if err != nil {
		return nil, err
	}

	//循环所有配置，筛选出符合条件的配置
	for _, conf := range *items {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}

	return nil, service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询华为云SDK应用配置信息失败", dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveHuaWeiYunSdkConf 保存华为SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkHuaWeiYun) SaveHuaWeiYunSdkConf(ctx context.Context, info model.HuaWeiYunSdkConf, isCreate bool) (*model.HuaWeiYunSdkConf, error) {
	items, _ := s.GetHuaWeiYunSdkConfList(ctx)

	isHas := false
	newItems := make([]model.HuaWeiYunSdkConf, 0)
	for _, conf := range *items {
		if conf.Identifier == info.Identifier { // 如果标识符相等，说明已经存在
			isHas = true
			newItems = append(newItems, info)
			continue
		}

		newItems = append(newItems, conf)
	}

	if !isHas {
		if isCreate {
			newItems = append(newItems, info)
		} else {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("华为云SDK配置信息保存失败，标识符错误"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)

	count, err := dao.SysConfig.Ctx(ctx).Count(do.SysConfig{
		Name: s.sysConfigName,
	})

	if count > 0 { // 已经存在，Save更新
		_, err = dao.SysConfig.Ctx(ctx).Data(do.SysConfig{Value: jsonString}).Where(do.SysConfig{
			Name: s.sysConfigName,
		}).Update()
	} else { // 不存在，Insert添加
		_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
			Name:  s.sysConfigName,
			Value: jsonString,
		})
	}

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "华为云SDK配置信息保存失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)

	// 同步token列表
	return &info, nil
}

// DeleteHuaWeiYunSdkConf 删除华为SDK应用配置信息
func (s *sSdkHuaWeiYun) DeleteHuaWeiYunSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetHuaWeiYunSdkConfList(ctx)

	isHas := false
	newItems := garray.New(false)
	for _, conf := range *items {
		if conf.Identifier == identifier {
			isHas = true
			continue
		}
		newItems.Append(conf)
	}

	if !isHas {
		return false, service.SysLogs().ErrorSimple(ctx, err, "要删除的华为云SDK配置信息不存在", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)

	if dao.SysConfig.Ctx(ctx).Where(do.SysConfig{Name: s.sysConfigName}).Update(do.SysConfig{Value: jsonString}); err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "华为云SDK配置信息删除失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)

	// 同步Token列表
	s.syncHuaWeiYunSdkConfTokenList(ctx)

	return true, nil
}

// 华为云服务的具体应用实例
