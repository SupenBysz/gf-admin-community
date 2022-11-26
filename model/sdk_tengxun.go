package model

// 阿里云服务配置信息

type TengxunSdkConf struct {
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

// TengxunAccessToken 腾讯平台返回的Token
type TengxunAccessToken struct {
	TengxunToken TengxunToken `json:"response" dc:"token信息"`
}

type TengxunToken struct {
	ExpireTime string `json:"expire_time" dc:"Access Token的有效期"`
	RequestId  string `json:"request_id" dc:"唯一请求 ID" `
	Token      string `json:"token" dc:"请求分配的Token值"`
}

// TengxunSdkConfToken 配置信息 + Token信息
type TengxunSdkConfToken struct {
	TengxunSdkConf
	TengxunAccessToken
}

type TengxunSdkConfList CollectRes[TengxunSdkConf]

// 腾讯云服务应用列表
