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

// AccessTokenByTencent 腾讯云API的AccessToken的响应参数。（AccessToken 是调用接口的认证Token）
type AccessTokenByTencent struct {
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	TransactionTime string `json:"transactionTime"`
	AccessToken     string `json:"access_token"`
	ExpireTime      string `json:"expire_time"`
	ExpireIn        string `json:"expire_in"`
}

// SignTicketRes 腾讯云API的Ticket的响应参数。（SIGN ticket 是合作方后台服务端业务请求生成签名鉴权参数之一，用于后台查询验证结果、调用其他业务服务等。）
type SignTicketRes struct {
	Code            string       `json:"code" dc:"code 不为 0 则表示获取失败，可以根据 code 和 msg 字段进行定位和调试。"`
	Msg             string       `json:"msg" `
	TransactionTime string       `json:"transactionTime" `
	Tickets         []SignTicket `json:"tickets" dc:"tickets 只有一个。"`
}
type SignTicket struct {
	Value      string `json:"value" dc:"ticket字符串值"`
	ExpireIn   string `json:"expire_in" dc:"expire_in 为 SIGN ticket 的最大生存时间，单位秒，合作伙伴在 判定有效期时以此为准。"`
	ExpireTime string `json:"expire_time" dc:"expire_time 为 SIGN ticket 失效的绝对时间，"`
}

// GetAdvFaceIdRes 腾讯云-合作方后台上传身份信息响应参数  ` POST https://kyc1.qcloud.com/api/server/getAdvFaceId?orderNo=xxx `
type GetAdvFaceIdRes struct {
	Code     string `json:"code" dc:"0：成功、非0：失败"`
	Msg      string `json:"msg" dc:"请求结果描述"`
	BizSeqNo string `json:"bizSeqNo" dc:"请求业务流水号"`
	Result   struct {
		BizSeqNo        string `json:"bizSeqNo" dc:"请求业务流水号"`
		TransactionTime string `json:"transactionTime" dc:"接口请求的时间"`
		OrderNo         string `json:"orderNo" dc:"订单编号"`
		FaceId          string `json:"faceId" dc:"此次刷脸用户标识"`
		OptimalDomain   string `json:"optimalDomain" dc:"启动 H5 人脸核身步骤中调用 login 接口使用的域名"`
		Success         bool   `json:"success" dc:""`
	} `json:"result"`
	TransactionTime string `json:"transactionTime" dc:"接口请求的时间"`
}

type StartAdvFaceAuthRes struct {
	OrderNo string `json:"orderNo" dc:"订单编号,查询认证结果时需要此参数"`
	FaceId  string `json:"faceId" dc:"此次刷脸用户标识"`
	Url     string `json:"url"  dc:"用于发起H5人脸核身流程的URL"`
}

// QueryFaceRecordRes 腾讯云-人脸核身结果查询响应参数
type QueryFaceRecordRes struct {
	Code     string `json:"code" dc:"0：表示身份验证成功且认证为同一人"`
	Msg      string `json:"msg" dc:"返回结果描述"`
	BizSeqNo string `json:"bizSeqNo" dc:"业务流水号"`
	Result   struct {
		OrderNo      string `json:"orderNo" dc:"订单编号"`
		LiveRate     string `json:"liveRate" dc:"活体检测得分"`
		Similarity   string `json:"similarity" dc:"人脸比对得分"`
		OccurredTime string `json:"occurredTime" dc:"进行刷脸的时间"`
		AppId        string `json:"appId" dc:"腾讯云控制台申请的 appid"`
		Photo        string `json:"photo" dc:"人脸核身时的照片，base64 位编码"`
		Video        string `json:"video" dc:"人脸核身时的视频，base64 位编码"`
		BizSeqNo     string `json:"bizSeqNo" dc:"业务流水号"`
		SdkVersion   string `json:"sdkVersion" dc:"人脸核身时的 sdk 版本号"`
		TrtcFlag     string `json:"trtcFlag" dc:"Trtc 渠道刷脸则标识Y"`
	} `json:"result" dc:"返回结果"`

	TransactionTime string `json:"transactionTime" dc:"请求接口的时间"`
}

/*


liveRate

String



similarity

String



occurredTime

String



photo

Base 64 string



video

Base 64 string



sdkVersion

String

人脸核身时的 sdk 版本号

trtcFlag

String

Trtc 渠道刷脸则标识"Y"

appId

String

腾讯云控制台申请的 appid


*/

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
