package sdk_tencent

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/json"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	faceid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 腾讯云服务平台

type sSdkTencent struct {
	TencentSdkConfTokenList []*sys_model.TencentSdkConfToken
	sysConfigName           string
	conf                    gdb.CacheOption
}

// New SdkBaidu 系统配置逻辑实现
func New() sys_service.ISdkTencent {
	return &sSdkTencent{
		TencentSdkConfTokenList: make([]*sys_model.TencentSdkConfToken, 0),
		sysConfigName:           "tencent_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

func init() {
	sys_service.RegisterSdkTencent(New())
}

// fetchTencentSdkConfToken 根据 identifier 获取腾讯云API Token  （API获取方式）
func (s *sSdkTencent) fetchTencentSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConfToken, err error) {

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
	// header["Authorization"] = ""
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取腾讯云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 接受返回数据，json解析
	newTokenInfo := sys_model.TencentAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取腾讯云API Token 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *sys_model.TencentSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.TencentSdkConfTokenList {
		if item.Identifier == identifier {
			result = &sys_model.TencentSdkConfToken{
				TencentSdkConf:     *info,
				TencentAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &sys_model.TencentSdkConfToken{
			TencentSdkConf:     *info,
			TencentAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	// 返回我们需要的token信息
	return result, nil
}

// GetTencentSdkConfList 获取腾讯云SDK应用配置列表
func (s *sSdkTencent) GetTencentSdkConfList(ctx context.Context) ([]*sys_model.TencentSdkConf, error) {
	items := make([]*sys_model.TencentSdkConf, 0)
	config, err := sys_service.SysConfig().GetByName(ctx, s.sysConfigName)
	if err != nil {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("腾讯云SDK配置信息获取失败"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if config.Value == "" {
		return items, nil
	}

	_ = gjson.DecodeTo(config.Value, &items)

	return items, nil
}

// GetTencentSdkConf 根据identifier标识获取SDK配置信息
func (s *sSdkTencent) GetTencentSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConf, err error) {
	items, err := s.GetTencentSdkConfList(ctx)
	if err != nil {
		return nil, err
	}

	// 循环所有配置，筛选出符合条件的配置
	for _, conf := range items {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询腾讯云SDK应用配置信息失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveTencentSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
func (s *sSdkTencent) SaveTencentSdkConf(ctx context.Context, info *sys_model.TencentSdkConf, isCreate bool) (*sys_model.TencentSdkConf, error) {
	oldItems, _ := s.GetTencentSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.TencentSdkConf, 0)
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
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("腾讯云SDK配置信息保存失败，标识符错误"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	// 序列化后进行保存至数据库
	jsonString := gjson.MustEncodeString(newItems)
	_, err := sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "腾讯云SDK配置信息保存失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步腾讯云SDK应用配置缓存列表
	s.syncTencentSdkConfList(ctx)

	return info, nil
}

// syncTencentSdkConfList 同步腾讯云SDK应用配置信息列表缓存  （代码中要是用到了s.TencentSdkConfList缓存变量的话，一定需要在CUD操作后调用此方法更新缓存变量）
func (s *sSdkTencent) syncTencentSdkConfList(ctx context.Context) error {
	items, err := s.GetTencentSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.TencentSdkConfToken, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.TencentSdkConfTokenList { // tokenList
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.TencentSdkConfTokenList = newTokenItems

	return nil
}

// DeleteTencentSdkConf 删除腾讯云SDK应用配置信息
func (s *sSdkTencent) DeleteTencentSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetTencentSdkConfList(ctx)

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
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "要删除的腾讯云SDK配置信息不存在", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err = sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "腾讯云SDK配置信息删除失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 同步Token列表
	s.syncTencentSdkConfList(ctx)

	return true, nil
}

// 腾讯云服务的具体应用实例

// LivenessRecognition 人脸核身(SDK接入模式)
func (s *sSdkTencent) LivenessRecognition(ctx context.Context, idCard, name, livenessType string) {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取

	config, err := s.GetTencentSdkConf(ctx, "tencent_config")

	credential := common.NewCredential(
		config.AESKey,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := faceid.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := faceid.NewLivenessRecognitionRequest()
	request.IdCard = common.StringPtr(idCard)
	request.Name = common.StringPtr(name)
	request.LivenessType = common.StringPtr(livenessType)

	// 返回的resp是一个LivenessRecognitionResponse的实例，与请求对象对应
	response, err := client.LivenessRecognition(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

}

// DetectAuth 腾讯云-实名核身鉴权
func (s *sSdkTencent) DetectAuth(ctx context.Context, idCard, name, returnUrl string) (*sys_model.DetectAuthRes, error) {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	config, err := s.GetTencentSdkConf(ctx, "tencent_config")

	credential := common.NewCredential(
		config.AESKey,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := faceid.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := faceid.NewDetectAuthRequest()

	request.RuleId = common.StringPtr("1") // TODO 现在先硬编码写死了，后续RuleId应该做数据库管理
	request.IdCard = common.StringPtr(idCard)
	request.Name = common.StringPtr(name)
	request.RedirectUrl = common.StringPtr(returnUrl)

	// 返回的resp是一个DetectAuthResponse的实例，与请求对象对应
	response, err := client.DetectAuth(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "An API error has returned: "+err.Error(), s.sysConfigName)
	}
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "腾讯云-实名核身鉴权请求失败", s.sysConfigName)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

	res := sys_model.DetectAuthRes{}
	if response.Response != nil {
		_ = gjson.DecodeTo(response.Response, &res)

	}

	return &res, err

	/*
		响应示例：
			{"Response":{"Url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx1f7125112b74db52\u0026redirect_uri=https%3A%2F%2Fopen.faceid.qq.com%2Fv1%2Fapi%2FgetCode%3FbizRedirect%3Dhttps%253A%252F%252Ffaceid.qq.com%252Fapi%252Fauth%252FgetOpenidAndSaveToken%253Ftoken%253D43DCB8C3-D330-429C-AF46-DB73BA9EE794\u0026response_type=code\u0026scope=snsapi_base\u0026state=\u0026component_appid=wx9802ee81e68d6dee#wechat_redirect","BizToken":"43DCB8C3-D330-429C-AF46-DB73BA9EE794","RequestId":"95be8b7f-f6f8-4735-895f-ee1c3d1f4ab9"}}
	*/
}

// GetDetectAuthResult 获取腾讯云-实名核身鉴权结果
func (s *sSdkTencent) GetDetectAuthResult(ctx context.Context, bizToken string, ruleId ...string) (*sys_model.GetDetectAuthResultRes, error) {
	config, err := s.GetTencentSdkConf(ctx, "tencent_config")
	credential := common.NewCredential(
		config.AESKey,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := faceid.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := faceid.NewGetDetectInfoRequest()

	request.BizToken = common.StringPtr(bizToken)
	id := "1"
	request.RuleId = common.StringPtr(id) // TODO 现在先硬编码写死了，后续RuleId应该做数据库管理
	//request.RuleId = common.StringPtr(ruleId)

	// 返回的resp是一个GetDetectInfoResponse的实例，与请求对象对应
	response, err := client.GetDetectInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "An API error has returned: "+err.Error(), s.sysConfigName)
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	//fmt.Printf("%s", response.ToJsonString())

	var tempInfo struct {
		Response struct {
			RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId" `

			DetectInfo *string `json:"DetectInfo,omitnil,omitempty" name:"DetectInfo" `
		} `json:"Response"`
	}

	res := sys_model.GetDetectAuthResultRes{}
	if response.Response != nil {
		//gjson.DecodeTo(response.Response.DetectInfo, &res)

		//jsonString := response.ToJsonString()
		//gjson.DecodeTo(jsonString, &res)

		t := kconv.Struct(response.ToJsonString(), &tempInfo)
		jsonStr, _ := json.UnescapeJSON(*t.Response.DetectInfo)
		gjson.DecodeTo(jsonStr, &res)
		res.RequestId = response.Response.RequestId
	}

	return &res, err

	/*
		 响应示例：
			{
		  "Text": {
		    "ErrCode": 0,
		    "ErrMsg": "成功",
		    "IdCard": "****",
		    "Name": "林菲菲",
		    "OcrNation": null,
		    "OcrAddress": null,
		    "OcrBirth": null,
		    "OcrAuthority": null,
		    "OcrValidDate": null,
		    "OcrName": "",
		    "OcrIdCard": "",
		    "OcrGender": null,
		    "LiveStatus": 0,
		    "LiveMsg": "成功",
		    "Comparestatus": 0,
		    "Comparemsg": "成功",
		    "Sim": "96.42",
		    "Location": null,
		    "Extra": "",
		    "Detail": {
		      "LivenessData": [
		        {
		          "ErrCode": 0,
		          "ErrMsg": "成功",
		          "ReqTime": "1715329563289",
		          "IdCard": "*****",
		          "Name": "林菲菲"
		        }
		      ]
		    }
		  },
		  "IdCardData": {
		    "OcrFront": null,
		    "OcrBack": null
		  },
		  "BestFrame": {
		    "BestFrame": ""
		  }
		}
	*/
}

// GetDetectAuthPlusResponse 获取腾讯云-实名核身鉴权增强版结果 （v3.0接口）
func (s *sSdkTencent) GetDetectAuthPlusResponse(ctx context.Context, bizToken, ruleId string) (*sys_model.GetDetectAuthPlusResponseRes, error) {
	config, err := s.GetTencentSdkConf(ctx, "tencent_config")

	credential := common.NewCredential(
		config.AESKey,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := faceid.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := faceid.NewGetDetectInfoEnhancedRequest()

	request.BizToken = common.StringPtr(bizToken)
	request.RuleId = common.StringPtr(ruleId)

	// 返回的resp是一个GetDetectInfoEnhancedResponse的实例，与请求对象对应
	response, err := client.GetDetectInfoEnhanced(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "An API error has returned: "+err.Error(), s.sysConfigName)
	}
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "腾讯云-获取实名核身鉴权增强版结果请求失败", s.sysConfigName)
	}

	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

	res := sys_model.GetDetectAuthPlusResponseRes{}
	if response.Response != nil {
		_ = gjson.DecodeTo(response.Response, &res)
	}

	return &res, err

	/*
	 响应示例：

	*/
}

// LivenessRecognition_Http 人脸核身(HTTP接入模式)
func (s *sSdkTencent) LivenessRecognition_Http(ctx context.Context, action, secretId, secretKey string) { // LivenessRecognition 活体核验
	service := "faceid"
	version := "2018-03-01"
	//action := "LivenessRecognition"
	region := ""
	token := ""
	host := "faceid.tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"
	var timestamp = time.Now().Unix()

	// ************* 步骤 1：拼接规范请求串 *************
	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	contentType := "application/json; charset=utf-8"
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\nx-tc-action:%s\n",
		contentType, host, strings.ToLower(action))
	signedHeaders := "content-type;host;x-tc-action"
	payload := "{}"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	log.Println(canonicalRequest)

	// ************* 步骤 2：拼接待签名字符串 *************
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)
	log.Println(string2sign)

	// ************* 步骤 3：计算签名 *************
	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	log.Println(signature)

	// ************* 步骤 4：拼接 Authorization *************
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature)
	log.Println(authorization)

	// ************* 步骤 5：构造并发起请求 *************
	url := "https://" + host
	httpRequest, _ := http.NewRequest("POST", url, strings.NewReader(payload))
	httpRequest.Header = map[string][]string{
		"Host":           {host},
		"X-TC-Action":    {action},
		"X-TC-Version":   {version},
		"X-TC-Timestamp": {strconv.FormatInt(timestamp, 10)},
		"Content-Type":   {contentType},
		"Authorization":  {authorization},
	}
	if region != "" {
		httpRequest.Header["X-TC-Region"] = []string{region}
	}
	if token != "" {
		httpRequest.Header["X-TC-Token"] = []string{token}
	}
	httpClient := http.Client{}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(body.String())
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
