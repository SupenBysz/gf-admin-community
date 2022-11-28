package sdk_tencent

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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

// 腾讯云服务平台

type sSdkTencent struct {
	TencentSdkConfTokenList []model.TencentSdkConfToken
	CacheDuration           time.Duration
	sysConfigName           string
}

// New SdkBaidu 系统配置逻辑实现
func New() *sSdkTencent {
	return &sSdkTencent{
		TencentSdkConfTokenList: make([]model.TencentSdkConfToken, 0),
		CacheDuration:           time.Hour,
		sysConfigName:           "tencent_sdk_conf",
	}
}

func init() {
	service.RegisterSdkTencent(New())
}

// fetchTencentSdkConfToken 根据 identifier 获取腾讯云API Token  （API获取方式）
func (s *sSdkTencent) fetchTencentSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.TencentSdkConfToken, err error) {

	info, err := s.GetTencentSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}
	client := g.Client()

	// URL 请求的服务器URL
	var host = "https://rkp.tencentcloudapi.com"

	// 请求头
	header := make(map[string]string)

	header["X-TC-Action"] = "GetToken"
	header["Content-type"] = "application/json"
	header["X-TC-Region"] = ""
	header["X-TC-Timestamp"] = gtime.Now().TimestampStr()
	header["X-TC-Version"] = info.Version
	//header["Authorization"] = ""
	header["X-TC-Language"] = "zh-CN"

	client.Header(header)

	// 请求数据，
	param := g.Map{
		// 业务ID
		"BusinessId": gconv.Int64(info.AppID),
		// 业务子场景
		"Scene": 0,
		// 业务侧账号体系下的用户ID (不是必填)
		"BusinessUserId": info.AESKey,
		// 用户侧的IP (不是必填)
		"AppClientIp": info.AppID,
		// 过期时间 (不是必填)
		"ExpireTime": info.APIKey,
	}

	response, err := client.Post(ctx, host, param)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "获取腾讯云API Token 失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 接受返回数据，json解析
	newTokenInfo := model.TencentAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "获取腾讯云API Token 失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *model.TencentSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.TencentSdkConfTokenList {
		if item.Identifier == identifier {
			result = &model.TencentSdkConfToken{
				TencentSdkConf:     *info,
				TencentAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &model.TencentSdkConfToken{
			TencentSdkConf:     *info,
			TencentAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	// 返回我们需要的token信息
	return result, nil
}

// GetTencentSdkConfToken 根据 identifier 查询腾讯SDK应用配置和Token信息
func (s *sSdkTencent) GetTencentSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.TencentSdkConfToken, err error) {
	for _, conf := range s.TencentSdkConfTokenList {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}
	return s.fetchTencentSdkConfToken(ctx, identifier)
}

// syncTencentSdkConfTokenList 同步腾讯云SDK应用配置信息Token列表缓存
func (s *sSdkTencent) syncTencentSdkConfTokenList(ctx context.Context) error {
	items, err := s.GetTencentSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]model.TencentSdkConfToken, 0)
	for _, conf := range *items {
		for _, tokenInfo := range s.TencentSdkConfTokenList {
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.TencentSdkConfTokenList = newTokenItems

	return nil
}

// GetTencentSdkConfList 获取腾讯云SDK应用配置列表
func (s *sSdkTencent) GetTencentSdkConfList(ctx context.Context) (*[]model.TencentSdkConf, error) {
	items := make([]model.TencentSdkConf, 0)

	data := entity.SysConfig{}

	err := dao.SysConfig.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.CacheDuration,
		Name:     s.sysConfigName,
		Force:    true,
	}).Where(do.SysConfig{
		Name: s.sysConfigName,
	}).Scan(&data)

	if err != nil && err != sql.ErrNoRows {
		return &items, service.SysLogs().ErrorSimple(ctx, gerror.New("腾讯云 SDK配置信息获取失败"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if data.Value == "" {
		return &items, nil
	}

	if nil == gjson.DecodeTo(data.Value, &items) {
		return &items, nil
	}

	return &items, nil
}

// GetTencentSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkTencent) GetTencentSdkConf(ctx context.Context, identifier string) (tokenInfo *model.TencentSdkConf, err error) {
	items, err := s.GetTencentSdkConfList(ctx)

	if err != nil {
		return nil, err
	}

	//循环所有配置，筛选出符合条件的配置
	for _, conf := range *items {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}

	return nil, service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询腾讯云SDK应用配置信息失败", dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveTencentSdkConf 保存腾讯SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkTencent) SaveTencentSdkConf(ctx context.Context, info model.TencentSdkConf, isCreate bool) (*model.TencentSdkConf, error) {
	items, _ := s.GetTencentSdkConfList(ctx)

	isHas := false
	newItems := make([]model.TencentSdkConf, 0)
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
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("腾讯云SDK配置信息保存失败，标识符错误"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
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
		return nil, service.SysLogs().ErrorSimple(ctx, err, "腾讯云SDK配置信息保存失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)

	// 同步token列表
	return &info, nil
}

// DeleteTencentSdkConf 删除腾讯SDK应用配置信息
func (s *sSdkTencent) DeleteTencentSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetTencentSdkConfList(ctx)

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
		return false, service.SysLogs().ErrorSimple(ctx, err, "要删除的腾讯云SDK配置信息不存在", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)

	if dao.SysConfig.Ctx(ctx).Where(do.SysConfig{Name: s.sysConfigName}).Update(do.SysConfig{Value: jsonString}); err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "腾讯云SDK配置信息删除失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)

	// 同步Token列表
	s.syncTencentSdkConfTokenList(ctx)

	return true, nil
}

// 腾讯云服务的具体应用实例
