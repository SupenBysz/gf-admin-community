package sdk_aliyun

import (
	"context"
	"fmt"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/kysion/base-library/utility/kconv"
)

/*
  阿里云服务平台
*/

type sSdkAliyun struct {
	AliyunSdkConfTokenList []*sys_model.AliyunSdkConfToken
	sysConfigName          string
	conf                   gdb.CacheOption
}

func init() {
	sys_service.RegisterSdkAliyun(New())
}

func New() sys_service.ISdkAliyun {
	return &sSdkAliyun{
		AliyunSdkConfTokenList: make([]*sys_model.AliyunSdkConfToken, 0),
		sysConfigName:          "aliyun_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
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

// GetAliyunSdkConfList 获取阿里云SDK应用配置列表
func (s *sSdkAliyun) GetAliyunSdkConfList(ctx context.Context) ([]*sys_model.AliyunSdkConf, error) {
	items := make([]*sys_model.AliyunSdkConf, 0)
	config, err := sys_service.SysConfig().GetByName(ctx, s.sysConfigName)
	if err != nil {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_aliyun_sdk_config_fetch_failed"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if config.Value == "" {
		return items, nil
	}

	_ = gjson.DecodeTo(config.Value, &items)

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

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_aliyun_sdk_app_config_query_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveAliyunSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkAliyun) SaveAliyunSdkConf(ctx context.Context, info *sys_model.AliyunSdkConf, isCreate bool) (*sys_model.AliyunSdkConf, error) {
	oldItems, _ := s.GetAliyunSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.AliyunSdkConf, 0)
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
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_aliyun_sdk_config_save_failed_identifier_error"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)
	_, err := sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_aliyun_sdk_config_save_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步阿里云SDK应用配置缓存列表
	s.syncAliyunSdkConfList(ctx)

	return info, nil
}

// syncAliyunSdkConfList 同步阿里云SDK应用配置信息列表缓存  （代码中要是用到了s.AliyunSdkConfList缓存变量的话，一定需要在CUD操作后调用此方法更新缓存变量）
func (s *sSdkAliyun) syncAliyunSdkConfList(ctx context.Context) error {
	items, err := s.GetAliyunSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.AliyunSdkConfToken, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.AliyunSdkConfTokenList { // tokenList
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.AliyunSdkConfTokenList = newTokenItems

	return nil
}

// DeleteAliyunSdkConf 删除阿里云SDK应用配置信息
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
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_aliyun_sdk_config_delete_not_exist", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err = sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_aliyun_sdk_config_delete_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步Token列表
	s.syncAliyunSdkConfList(ctx)

	return true, nil
}

//  阿里云服务的具体应用实例

// GetWsCustomizedChGeneral 中文分词
func (s *sSdkAliyun) GetWsCustomizedChGeneral(ctx context.Context, text string) (sys_model.AliyunNlpDataRes, error) {
	/**
	 * 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	 * 此处以把AccessKey和AccessKeySecret保存在环境变量为例说明。您也可以根据业务需要，保存到配置文件里。
	 * 强烈建议不要把AccessKey和AccessKeySecret保存到代码里，会存在密钥泄漏风险
	 */

	config, err := s.GetAliyunSdkConf(ctx, "nlp")

	AccessKeyId := config.AESKey
	AccessKeySecret := config.SecretKey
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Domain = "alinlp.cn-hangzhou.aliyuncs.com"
	request.Version = "2020-06-29"
	// 因为是RPC接口，因此需指定ApiName(Action)
	request.ApiName = "GetWsCustomizedChGeneral" // ApiName 就是文档中的请求参数 Action
	request.QueryParams["ServiceCode"] = "alinlp"
	request.QueryParams["Text"] = text
	request.QueryParams["TokenizerId"] = "GENERAL_CHN"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetHttpContentString())
	var data struct {
		RequestId string `json:"RequestId" dc:"唯一请求id，排查问题的依据"`
		Data      string `json:"data"`
	}

	res := sys_model.AliyunNlpDataRes{}
	_ = gjson.DecodeTo(response.GetHttpContentString(), &data)

	kconv.Struct(data, &res)

	return res, nil
}
