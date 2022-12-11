package sys_model

import "time"

// 华为云服务配置信息

type HuaweiSdkConf struct {
	Identifier  string `json:"identifier" v:"required#业务标识符参数错误" dc:"业务标识符，唯一，且不可修改"`
	Description string `json:"description" dc:"描述"`
	AppID       string `json:"appID" v:"required#请输入 AppID" dc:"AppID"`
	AESKey      string `json:"aesKey" v:"required#请输入 AES Key" dc:"AES Key"`
	APIKey      string `json:"apiKey" v:"required#请输入 API Key" dc:"API Key"`
	SecretKey   string `json:"secretKey" v:"required#请输入 Secret Key" dc:"Secret Key"`
}

// HuaweiAccessToken 华为云平台返回的Token
type HuaweiAccessToken struct {
	AccessToken string   `json:"access_token" dc:"token值"`
	TokenType   string   `json:"token_type" dc:"token类型"`
	ExpiresIn   int      `json:"expires_in" dc:"token有效期" `
	Scope       string   `json:"scope" dc:"该参数可忽略"`
	PlatUser    PlatUser `json:"plat_user"`
}

type UserExtendMap struct {
	ClientID string `json:"clientId" dc:"客户端秘钥"`
}

type PlatUser struct {
	UserID           int           `json:"userId" dc:"用户id"`
	UserNo           string        `json:"userNo" dc:"编号"`
	UserName         string        `json:"userName" dc:"用户名"`
	UserType         string        `json:"userType" dc:"用户类型"`
	Phone            interface{}   `json:"phone" dc:"电话"`
	TenantID         int           `json:"tenantId"`
	Status           string        `json:"status"`
	DeleteFlag       int           `json:"deleteFlag"`
	PwdUpdateDate    time.Time     `json:"pwdUpdateDate"`
	UserExtendMap    UserExtendMap `json:"userExtendMap"`
	Password         interface{}   `json:"password"`
	Salt             interface{}   `json:"salt"`
	CurrentLoginType string        `json:"currentLoginType"`
}

type HuaweiToken struct {
	ExpireTime string `json:"expire_time" dc:"Access Token的有效期"`
	Id         string `json:"id" dc:"请求分配的Token值"`
	UserId     string `json:"user_id" dc:"用户id"`
}

// HuaweiSdkConfToken 配置信息 + Token信息
type HuaweiSdkConfToken struct {
	HuaweiSdkConf
	HuaweiAccessToken
}

type HuaweiSdkConfList CollectRes[HuaweiSdkConf]
