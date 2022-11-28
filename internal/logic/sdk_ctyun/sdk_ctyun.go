package sdk_ctyun

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
	"time"
)

// 天翼云服务平台

type sSdkCtyun struct {
	CtyunSdkConfTokenList []model.CtyunSdkConfInfo
	CacheDuration         time.Duration
	sysConfigName         string
}

// New SdkBaidu 系统配置逻辑实现
func New() *sSdkCtyun {
	return &sSdkCtyun{
		CtyunSdkConfTokenList: make([]model.CtyunSdkConfInfo, 0),
		CacheDuration:         time.Hour,
		sysConfigName:         "ctyun_sdk_conf",
	}
}

func init() {
	service.RegisterSdkCtyun(New())
}

// 我估计没有token的认证，后期可能需要添加一个签名函数

// syncCtyunSdkConfTokenList 同步天翼云SDK应用配置信息Token列表缓存
func (s *sSdkCtyun) syncCtyunSdkConfTokenList(ctx context.Context) error {
	items, err := s.GetCtyunSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]model.CtyunSdkConfInfo, 0)
	for _, conf := range *items {
		for _, tokenInfo := range s.CtyunSdkConfTokenList {
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.CtyunSdkConfTokenList = newTokenItems

	return nil
}

// GetCtyunSdkConfList 获取天翼云SDK应用配置列表
func (s *sSdkCtyun) GetCtyunSdkConfList(ctx context.Context) (*[]model.CtyunSdkConf, error) {
	items := make([]model.CtyunSdkConf, 0)

	data := entity.SysConfig{}

	err := dao.SysConfig.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.CacheDuration,
		Name:     s.sysConfigName,
		Force:    true,
	}).Where(do.SysConfig{
		Name: s.sysConfigName,
	}).Scan(&data)

	if err != nil && err != sql.ErrNoRows {
		return &items, service.SysLogs().ErrorSimple(ctx, gerror.New("天翼云 SDK配置信息获取失败"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if data.Value == "" {
		return &items, nil
	}

	if nil == gjson.DecodeTo(data.Value, &items) {
		return &items, nil
	}

	return &items, nil
}

// GetCtyunSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkCtyun) GetCtyunSdkConf(ctx context.Context, identifier string) (tokenInfo *model.CtyunSdkConf, err error) {
	items, err := s.GetCtyunSdkConfList(ctx)

	if err != nil {
		return nil, err
	}

	//循环所有配置，筛选出符合条件的配置
	for _, conf := range *items {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}

	return nil, service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询天翼云SDK应用配置信息失败", dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveCtyunSdkConf 保存天翼SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkCtyun) SaveCtyunSdkConf(ctx context.Context, info model.CtyunSdkConf, isCreate bool) (*model.CtyunSdkConf, error) {
	items, _ := s.GetCtyunSdkConfList(ctx)

	isHas := false
	newItems := make([]model.CtyunSdkConf, 0)
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
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("天翼云SDK配置信息保存失败，标识符错误"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
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
		return nil, service.SysLogs().ErrorSimple(ctx, err, "天翼云SDK配置信息保存失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)

	// 同步token列表
	return &info, nil
}

// DeleteCtyunSdkConf 删除天翼SDK应用配置信息
func (s *sSdkCtyun) DeleteCtyunSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetCtyunSdkConfList(ctx)

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
		return false, service.SysLogs().ErrorSimple(ctx, err, "要删除的天翼云SDK配置信息不存在", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)

	if dao.SysConfig.Ctx(ctx).Where(do.SysConfig{Name: s.sysConfigName}).Update(do.SysConfig{Value: jsonString}); err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "天翼云SDK配置信息删除失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除缓存列表
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)

	// 同步Token列表
	s.syncCtyunSdkConfTokenList(ctx)

	return true, nil
}

// 天翼云服务的具体应用实例
