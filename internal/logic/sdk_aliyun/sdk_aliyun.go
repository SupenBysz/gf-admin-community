package sdk_aliyun

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// 阿里云服务平台
type sSdkAliyun struct {
	AliyunSdkConfTokenList []*sys_model.AliyunSdkConfToken
	sysConfigName          string
	conf                   gdb.CacheOption
}

func init() {
	sys_service.RegisterSdkAliyun(New())
}

// New SdkBaidu 系统配置逻辑实现
func New() *sSdkAliyun {
	return &sSdkAliyun{
		AliyunSdkConfTokenList: make([]*sys_model.AliyunSdkConfToken, 0),
		sysConfigName:          "aliyun_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

// fetchALiYHunSdkConfToken 根据 identifier 获取阿里云API Token  （API获取方式）
func (s *sSdkAliyun) fetchAliyunSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.AliyunSdkConfToken, err error) {
	info, err := s.GetAliyunSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}

	// 准备请求参数
	var host = "http://nls-meta.cn-shanghai.aliyuncs.com/"

	// 请求头
	// Host 请求的服务器URL
	// User-Agent: curl/7.49.1 Accept: */*
	// Content-type: application/x-www-form-urlencoded

	//	请求参数
	param := g.Map{
		// 阿里云账号AccessKey ID
		"AccessKeyId": info.AESKey,
		// 签名算法版本：1.0
		"SignatureVersion": 1.0,
		// API名称：CreateToken
		"Action": "CreateToken",
		// 响应返回的类型：JSON
		"Format": "JSON",
		// 唯一随机数uuid，用于请求的防重放攻击，每次请求唯一，不能重复使用
		"SignatureNonce": "8d1e6a7a-f44e-40d5-aedb-fe4a1c80f434",
		// API版本：2019-02-28
		"Version": "2019-02-28",
		// 签名生成结果
		"Signature": "oT8A8RgvFE1tMD%2B3hDbGuoMQSi8%3D",
		// 签名生成算法
		"SignatureMethod": "HMAC-SHA1",
		// 服务所在地域
		"RegionId": "cn-shanghai",
		// 时间戳
		"Timestamp": gtime.Now().Timestamp(),
	}

	response, err := g.Client().Post(ctx, host, param)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取阿里云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 接受返回数据，json解析
	newTokenInfo := sys_model.AliyunAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取百度API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *sys_model.AliyunSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.AliyunSdkConfTokenList {
		if item.Identifier == identifier {
			result = &sys_model.AliyunSdkConfToken{
				AliyunSdkConf:     *info,
				AliyunAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &sys_model.AliyunSdkConfToken{
			AliyunSdkConf:     *info,
			AliyunAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	// 返回我们需要的token信息
	return result, nil
}

// GetAliyunSdkToken 通过SDK获取Token （SDK获取方式）
func (s *sSdkAliyun) GetAliyunSdkToken(ctx context.Context, tokenInfo sys_model.AliyunSdkConfToken, err error) {
	// client, err := sdk.NewClientWithAccessKey("cn-shanghai", "<yourAccessKey Id>", "<yourAccessKey Secret>")
	// if err != nil {
	// 	panic(err)
	// }
	// request := requests.NewCommonRequest()
	// request.Method = "POST"
	// request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	// request.ApiName = "CreateToken"
	// request.Version = "2019-02-28"
	// response, err := client.ProcessCommonRequest(request)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(response.GetHttpStatus())
	// fmt.Print(response.GetHttpContentString())
}

// GetAliyunSdkConfToken 根据 identifier 查询百度SDK应用配置和Token信息
func (s *sSdkAliyun) GetAliyunSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.AliyunSdkConfToken, err error) {
	for _, conf := range s.AliyunSdkConfTokenList {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}
	return s.fetchAliyunSdkConfToken(ctx, identifier)
}

// syncBaiduSdkConfTokenList 同步百度SDK应用配置信息Token列表缓存
func (s *sSdkAliyun) syncAliyunSdkConfTokenList(ctx context.Context) error {
	items, err := s.GetAliyunSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.AliyunSdkConfToken, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.AliyunSdkConfTokenList {
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.AliyunSdkConfTokenList = newTokenItems

	return nil
}

// GetAliyunSdkConfList 获取阿里云SDK应用配置列表
func (s *sSdkAliyun) GetAliyunSdkConfList(ctx context.Context) ([]*sys_model.AliyunSdkConf, error) {
	items := make([]*sys_model.AliyunSdkConf, 0)

	data := sys_entity.SysConfig{}

	err := sys_dao.SysConfig.Ctx(ctx).Where(sys_do.SysConfig{
		Name: s.sysConfigName,
	}).Scan(&data)

	if err != nil && err != sql.ErrNoRows {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("阿里云SDK配置信息获取失败"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if data.Value == "" {
		return items, nil
	}

	if nil == gjson.DecodeTo(data.Value, &items) {
		return items, nil
	}

	return items, nil
}

// GetAliyunSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkAliyun) GetAliyunSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.AliyunSdkConf, err error) {
	items, err := s.GetAliyunSdkConfList(ctx)

	if err != nil {
		return nil, err
	}

	// 循环所有配置，筛选出符合条件的配置
	for _, conf := range items {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询阿里云SDK应用配置信息失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveAliyunSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkAliyun) SaveAliyunSdkConf(ctx context.Context, info *sys_model.AliyunSdkConf, isCreate bool) (*sys_model.AliyunSdkConf, error) {
	items, _ := s.GetAliyunSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.AliyunSdkConf, 0)
	for _, conf := range items {
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
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("阿里云SDK配置信息保存失败，标识符错误"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)

	count, err := sys_dao.SysConfig.Ctx(ctx).Count(sys_do.SysConfig{
		Name: s.sysConfigName,
	})

	if count > 0 { // 已经存在，Save更新
		_, err = sys_dao.SysConfig.Ctx(ctx).Data(sys_do.SysConfig{Value: jsonString}).Where(sys_do.SysConfig{
			Name: s.sysConfigName,
		}).Update()
	} else { // 不存在，Insert添加
		_, err = sys_dao.SysConfig.Ctx(ctx).Insert(sys_do.SysConfig{
			Name:  s.sysConfigName,
			Value: jsonString,
		})
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "阿里云SDK配置信息保存失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(sys_dao.SysConfig.DB(), s.sysConfigName)

	// 同步token列表
	return info, nil
}

// DeleteAliyunSdkConf 删除百度SDK应用配置信息
func (s *sSdkAliyun) DeleteAliyunSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetAliyunSdkConfList(ctx)

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
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "要删除的阿里云SDK配置信息不存在", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)

	if sys_dao.SysConfig.Ctx(ctx).Where(sys_do.SysConfig{Name: s.sysConfigName}).Update(sys_do.SysConfig{Value: jsonString}); err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "阿里云SDK配置信息删除失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(sys_dao.SysConfig.DB(), s.sysConfigName)

	// 同步Token列表
	s.syncAliyunSdkConfTokenList(ctx)

	return true, nil
}

// 阿里云服务的具体应用实例
