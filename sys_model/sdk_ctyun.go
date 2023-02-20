package sys_model

import "github.com/kysion/base-library/base_model"

// CtyunSdkConf天翼云服务配置信息
type CtyunSdkConf struct {
	Identifier  string `json:"identifier" v:"required#业务标识符参数错误" dc:"业务标识符，唯一，且不可修改"`
	Description string `json:"description" dc:"描述"`
	AppID       string `json:"appID" v:"required#请输入 AppID" dc:"应用AppID"`
	AppKey      string `json:"appKey" v:"required#请输入 API Key" dc:"应用API Key"`
	AppSecret   string `json:"appSecret" v:"required#请输入 Secret Key" dc:"应用AppSecret"`
	AccessKey   string `json:"accessKey" v:"required#请输入 AccessKey" dc:"账户的AES Key"`
	SecurityKey string `json:"security_key" v:"required#请输入 SecurityKey" dc:"账户的 SecurityKey"`
}

// CtyunSdkConfToken 配置信息 + ...
type CtyunSdkConfInfo struct {
	CtyunSdkConf
}

// CtyunSdkConfList天翼云应用配置列表
type CtyunSdkConfList base_model.CollectRes[*CtyunSdkConf]

// 因为我看平台的认证方式好像不是token认证，是通过AK/SK进行认证的，所以我暂时没写token信息

// 具体的应用实例
