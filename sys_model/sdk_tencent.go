package sys_model

import "github.com/kysion/base-library/base_model"

// 阿里云服务配置信息

type TencentSdkConf struct {
	Identifier  string `json:"identifier" v:"required#业务标识符参数错误" dc:"业务标识符，唯一，且不可修改"`
	Description string `json:"description" dc:"描述"`
	AppID       string `json:"appID" v:"required#请输入 AppID" dc:"AppID"`
	AESKey      string `json:"aesKey" v:"required#请输入 AES Key" dc:"AES Key"`
	APIKey      string `json:"apiKey" v:"required#请输入 API Key" dc:"API Key"`
	SecretKey   string `json:"secretKey" v:"required#请输入 Secret Key" dc:"Secret Key"`

	Active  string `json:"active" v:"required#公共参数" dc:"代表应用接口的值"`
	Version string `json:"version" v:"required#公共参数" dc:"接口版本"`
	Region  string `json:"region" dc:"代表请求的地域，公共参数"`
}

// TencentAccessToken 腾讯平台返回的Token
type TencentAccessToken struct {
	TencentToken TencentToken `json:"response" dc:"token信息"`
}

type TencentToken struct {
	ExpireTime string `json:"expire_time" dc:"Access Token的有效期"`
	RequestId  string `json:"request_id" dc:"唯一请求 ID" `
	Token      string `json:"token" dc:"请求分配的Token值"`
}

// TencentSdkConfToken 配置信息 + Token信息
type TencentSdkConfToken struct {
	TencentSdkConf
	TencentAccessToken
}

type TencentSdkConfList base_model.CollectRes[*TencentSdkConf]

// 腾讯云服务应用列表

type TencentOSS struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	BaseURL    string `json:"baseURL"`
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
	PathPrefix string `json:"pathPrefix"`
}
