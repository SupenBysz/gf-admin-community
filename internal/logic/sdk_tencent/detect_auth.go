package sdk_tencent

import (
	"context"
	"fmt"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/kysion/base-library/utility/json"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	faceid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
)

// 腾讯云服务的具体应用实例

/*
	实名核身鉴权 （人脸核身SasS服务接口）
*/

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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_tencent_auth_request_failed", s.sysConfigName)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

	res := sys_model.DetectAuthRes{}
	if response.Response != nil {
		//_ = gjson.DecodeTo(response.Response, &res)

		kconv.Struct(response.Response, &res)
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_tencent_auth_enhanced_result_failed", s.sysConfigName)
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
