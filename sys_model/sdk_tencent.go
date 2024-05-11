package sys_model

import (
	"github.com/kysion/base-library/base_model"
)

// 阿里云服务配置信息

type TencentSdkConf struct {
	Identifier  string `json:"identifier" v:"required#业务标识符参数错误" dc:"业务标识符，唯一，且不可修改"`
	Description string `json:"description" dc:"描述"`
	AppID       string `json:"appID" v:"required#请输入 AppID" dc:"AppID"`      // APPID
	AESKey      string `json:"aesKey" v:"required#请输入 AES Key" dc:"AES Key"` // SecretId
	APIKey      string `json:"apiKey" v:"required#请输入 API Key" dc:"API Key"`
	SecretKey   string `json:"secretKey" v:"required#请输入 Secret Key" dc:"Secret Key"` // SecretKey

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

// 腾讯云API能力-----

// DetectAuthRes 腾讯云-实名身份核验接口 响应
type DetectAuthRes struct {
	// 用于发起核身流程的URL，仅微信H5场景使用。
	Url *string `json:"Url,omitnil,omitempty" name:"Url" dc:"用于发起核身流程的URL，仅微信H5场景使用。"`

	// 一次核身流程的标识，有效时间为7,200秒；
	// 完成核身后，可用该标识获取验证结果信息。
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken" dc:"一次核身流程的标识，有效时间为7,200秒；完成核身后，可用该标识获取验证结果信息。"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId" dc:"唯一请求 ID，由服务端生成，每次请求都会返回"`
}

// GetDetectAuthResultRes 腾讯云-获取实名身份核验接口接口 响应
// type GetDetectAuthResultRes struct {  }

// GetDetectAuthResultRes 腾讯云-获取实名身份核验接口接口 响应
type GetDetectAuthResultRes struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId" dc:"唯一请求 ID，由服务端生成，每次请求都会返回"`

	Text struct {
		ErrCode       int         `json:"ErrCode"`
		ErrMsg        string      `json:"ErrMsg"`
		IdCard        string      `json:"IdCard"`
		Name          string      `json:"Name"`
		OcrNation     interface{} `json:"OcrNation"`
		OcrAddress    interface{} `json:"OcrAddress"`
		OcrBirth      interface{} `json:"OcrBirth"`
		OcrAuthority  interface{} `json:"OcrAuthority"`
		OcrValidDate  interface{} `json:"OcrValidDate"`
		OcrName       string      `json:"OcrName"`
		OcrIdCard     string      `json:"OcrIdCard"`
		OcrGender     interface{} `json:"OcrGender"`
		LiveStatus    int         `json:"LiveStatus"`
		LiveMsg       string      `json:"LiveMsg"`
		Comparestatus int         `json:"Comparestatus"`
		Comparemsg    string      `json:"Comparemsg"`
		Sim           string      `json:"Sim"`
		Location      interface{} `json:"Location"`
		Extra         string      `json:"Extra"`
		Detail        struct {
			LivenessData []struct {
				ErrCode int    `json:"ErrCode"`
				ErrMsg  string `json:"ErrMsg"`
				ReqTime string `json:"ReqTime"`
				IdCard  string `json:"IdCard"`
				Name    string `json:"Name"`
			} `json:"LivenessData"`
		} `json:"Detail"`
	} `json:"Text"  dc:"文本类信息"`

	IdCardData struct {
		OcrFront interface{} `json:"OcrFront"`
		OcrBack  interface{} `json:"OcrBack"`
	} `json:"IdCardData" dc:"身份证照片信息。"`

	BestFrame struct {
		BestFrame string `json:"BestFrame"`
	} `json:"BestFrame" dc:"最佳帧信息Base64编码。注意：此字段可能返回 null，表示取不到有效值。"`
}

// type GetDetectAuthPlusResponseRes v20180301.GetDetectInfoEnhancedResponseParams

// GetDetectAuthPlusResponseRes 获取实名核身结果信息增强版
type GetDetectAuthPlusResponseRes struct {
	//Response struct {
	RequestId string `json:"RequestId" dc:"	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。"`

	Text struct {
		CompareLibType string `json:"CompareLibType"`
		Comparemsg     string `json:"Comparemsg"`
		Comparestatus  int    `json:"Comparestatus"`
		ErrCode        int    `json:"ErrCode"`
		ErrMsg         string `json:"ErrMsg"`
		Extra          string `json:"Extra"`
		IdCard         string `json:"IdCard"`
		IdInfoFrom     string `json:"IdInfoFrom"`
		LiveMsg        string `json:"LiveMsg"`
		LiveStatus     int    `json:"LiveStatus"`
		LivenessDetail []struct {
			CompareLibType string `json:"CompareLibType"`
			Comparemsg     string `json:"Comparemsg"`
			Comparestatus  int    `json:"Comparestatus"`
			Errcode        int    `json:"Errcode"`
			Errmsg         string `json:"Errmsg"`
			Idcard         string `json:"Idcard"`
			IsNeedCharge   bool   `json:"IsNeedCharge"`
			Livemsg        string `json:"Livemsg"`
			LivenessMode   int    `json:"LivenessMode"`
			Livestatus     int    `json:"Livestatus"`
			Name           string `json:"Name"`
			ReqTime        string `json:"ReqTime"`
			Seq            string `json:"Seq"`
			Sim            string `json:"Sim"`
		} `json:"LivenessDetail"`
		LivenessMode     int           `json:"LivenessMode"`
		Location         interface{}   `json:"Location"`
		Mobile           string        `json:"Mobile"`
		NFCBillingCounts int           `json:"NFCBillingCounts"`
		NFCRequestIds    []interface{} `json:"NFCRequestIds"`
		Name             string        `json:"Name"`
		OcrAddress       interface{}   `json:"OcrAddress"`
		OcrAuthority     interface{}   `json:"OcrAuthority"`
		OcrBirth         interface{}   `json:"OcrBirth"`
		OcrGender        interface{}   `json:"OcrGender"`
		OcrIdCard        string        `json:"OcrIdCard"`
		OcrName          string        `json:"OcrName"`
		OcrNation        interface{}   `json:"OcrNation"`
		OcrValidDate     interface{}   `json:"OcrValidDate"`
		PassNo           interface{}   `json:"PassNo"`
		Sim              string        `json:"Sim"`
		UseIDType        int           `json:"UseIDType"`
		VisaNum          interface{}   `json:"VisaNum"`
	} `json:"Text" dc:"文本类信息"`

	// TODO 如下不需要的字段先行注释了。。。。有需要解开！
	/*
		BestFrame struct {
			BestFrame  string      `json:"BestFrame"`
			BestFrames interface{} `json:"BestFrames"`
		} `json:"BestFrame" dc:"最佳帧信息。注意：此字段可能返回 null，表示取不到有效值。"`

		EncryptedBody string `json:"EncryptedBody"`
		Encryption    struct {
			Algorithm      string        `json:"Algorithm"`
			CiphertextBlob string        `json:"CiphertextBlob"`
			EncryptList    []interface{} `json:"EncryptList"`
			Iv             string        `json:"Iv"`
			TagList        []interface{} `json:"TagList"`
		} `json:"Encryption" dc:"敏感数据加密信息"`

		IdCardData struct {
			Avatar              interface{} `json:"Avatar"`
			BackWarnInfos       interface{} `json:"BackWarnInfos"`
			OcrBack             interface{} `json:"OcrBack"`
			OcrFront            interface{} `json:"OcrFront"`
			ProcessedBackImage  interface{} `json:"ProcessedBackImage"`
			ProcessedFrontImage interface{} `json:"ProcessedFrontImage"`
			WarnInfos           interface{} `json:"WarnInfos"`
		} `json:"IdCardData" dc:"身份证照片信息。"`

		IntentionActionResult interface{} `json:"IntentionActionResult" dc:"意愿核身点头确认模式的结果信息，若未使用该意愿核身功能，该字段返回值可以不处理。"`

		IntentionQuestionResult struct {
			AsrResult             []interface{} `json:"AsrResult"`
			Audios                []interface{} `json:"Audios"`
			FinalResultCode       interface{}   `json:"FinalResultCode"`
			FinalResultDetailCode interface{}   `json:"FinalResultDetailCode"`
			FinalResultMessage    interface{}   `json:"FinalResultMessage"`
			ResultCode            []interface{} `json:"ResultCode"`
			ScreenShot            []interface{} `json:"ScreenShot"`
			Video                 interface{}   `json:"Video"`
		} `json:"IntentionQuestionResult" dc:"意愿核身问答模式结果。若未使用该意愿核身功能，该字段返回值可以不处理。"`

		IntentionVerifyData struct {
			AsrResult                interface{} `json:"AsrResult"`
			ErrorCode                interface{} `json:"ErrorCode"`
			ErrorMessage             interface{} `json:"ErrorMessage"`
			IntentionVerifyBestFrame interface{} `json:"IntentionVerifyBestFrame"`
			IntentionVerifyVideo     interface{} `json:"IntentionVerifyVideo"`
		} `json:"IntentionVerifyData" dc:"意愿核身朗读模式结果信息。若未使用意愿核身功能，该字段返回值可以不处理。"`

		VideoData struct {
			LivenessVideo string      `json:"LivenessVideo"`
			BestFrames    interface{} `json:"BestFrames"`
		} `json:"VideoData" dc:"视频信息"`
	*/
	//} `json:"Response"`
}
