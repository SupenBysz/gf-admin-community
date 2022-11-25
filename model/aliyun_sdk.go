package model

// 阿里云服务配置信息

type AliyunSdkConf struct {
	Identifier  string `json:"identifier" v:"required#业务标识符参数错误" dc:"业务标识符，唯一，且不可修改"`
	Description string `json:"description" dc:"描述"`
	AppID       string `json:"appID" v:"required#请输入 AppID" dc:"AppID"`
	AESKey      string `json:"aesKey" v:"required#请输入 AES Key" dc:"AES Key"`
	APIKey      string `json:"apiKey" v:"required#请输入 API Key" dc:"API Key"`
	SecretKey   string `json:"secretKey" v:"required#请输入 Secret Key" dc:"Secret Key"`
}

// ALiYunAccessToken 阿里云平台返回的Token
type ALiYunAccessToken struct {
	NlsRequestId string      `json:"NlsRequestId" dc:"该参数可忽略"`
	RequestId    string      `json:"RequestId" dc:"该参数忽略,请求ID"`
	AliYunToken  ALiYunToken `json:"Token" dc:"token信息"`
}

type ALiYunToken struct {
	ExpireTime string `json:"expire_time" dc:"Access Token的有效期"`
	Id         string `json:"id" dc:"请求分配的Token值"`
	UserId     string `json:"user_id" dc:"用户id"`
}

type AliyunSdkConfToken struct {
	AliyunSdkConf
	ALiYunAccessToken
}

type AliyunSdkConfList CollectRes[AliyunSdkConf]

// 阿里云服务应用列表
