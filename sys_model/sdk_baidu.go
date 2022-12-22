package sys_model

type BaiduSdkConf struct {
	Identifier  string `json:"identifier" v:"required#业务标识符参数错误" dc:"业务标识符，唯一，且不可修改"`
	Description string `json:"description" dc:"描述"`
	AppID       string `json:"appID" v:"required#请输入 AppID" dc:"AppID"`
	AESKey      string `json:"aesKey" v:"required#请输入 AES Key" dc:"AES Key"`
	APIKey      string `json:"apiKey" v:"required#请输入 API Key" dc:"API Key"`
	SecretKey   string `json:"secretKey" v:"required#请输入 Secret Key" dc:"Secret Key"`
}

type BaiduSdkConfAccessToken struct {
	AccessToken   string `json:"access_token" dc:"获取的Access Token"`
	refreshToken  string `json:"refresh_token" dc:"该参数忽略"`
	ExpiresIn     int64  `json:"expires_in" dc:"Access Token的有效期(秒为单位，有效期30天)"`
	Scope         string `json:"scope" dc:"该参数忽略"`
	SessionKey    string `json:"session_key" dc:"该参数忽略"`
	SessionSecret string `json:"session_secret" dc:"该参数忽略"`
}

type BaiduSdkOCRIDCardA struct {
	Direction      int    `json:"direction" dc:"图像方向：-1 为定义,0正向,1逆时针90度,2逆时针180度,3逆时针270度"`
	ImageStateText string `json:"imageStateText" dc:"状态：normal识别正常,reversed_side身份证正反面颠倒,non_idcard上传的图片中不包含身份证,blurred身份证模糊,other_type_card,其他类型证照,over_exposure身份证关键字段反光或过曝,over_dark身份证欠曝（亮度过低）,unknown未知状态"`
	RiskType       string `json:"riskType" dc:"风险类型：normal正常身份证,copy复印件,temporary临时身份证,screen翻拍,unknown其他未知情况"`
	Address        string `json:"address" dc:"地址信息"`
	IDCardNumber   string `json:"idCardNumber" dc:"身份证号"`
	Birthday       string `json:"birthday" dc:"出生日期"`
	Realname       string `json:"realname" dc:"姓名"`
	Gender         string `json:"gender" dc:"性别"`
	Nation         string `json:"nation" dc:"民族"`
}

type BaiduSdkOCRIDCardB struct {
	ExpiryDate       string `json:"expiryDate" dc:"失效日期"`
	IssuingAuthority string `json:"issuingAuthority" dc:"签发机关"`
	IssuingDate      string `json:"issuingDate" dc:"签发日期"`
}

type BaiduSdkOCRIDCard struct {
	OCRIDCardA *BaiduSdkOCRIDCardA `json:"ocrIDCardA" dc:"身份证头像面识别的信息"`
	OCRIDCardB *BaiduSdkOCRIDCardB `json:"ocrIDCardB" dc:"身份证国徽面识别的信息"`
}

type BusinessLicenseOCR struct {
	Direction         int    `json:"direction" dc:"图像方向：-1 为定义,0正向,1逆时针90度,2逆时针180度,3逆时针270度"`
	RiskType          string `json:"riskType" dc:"风险类型：normal正常营业执照,copy复印件,screen翻拍,scan扫描,unknown其他未知情况"`
	CreditCode        string `json:"creditCode" dc:"社会信用代码"`
	CombiningForm     string `json:"combiningForm" dc:"组成形式"`
	BusinessScope     string `json:"businessScope" dc:"经营范围"`
	EstablishmentDate string `json:"establishmentDate" dc:"成立日期"`
	LegalPerson       string `json:"legalPerson" dc:"法人"`
	RegisteredCapital string `json:"registeredCapital" dc:"注册资本"`
	CertificateNumber string `json:"certificateNumber" dc:"证件编号"`
	RegisteredAddress string `json:"registeredAddress" dc:"注册地址"`
	CompanyName       string `json:"companyName" dc:"主体名称"`
	ExpirationDate    string `json:"expirationDate" dc:"有效期"`
	ApprovalDate      string `json:"approvalDate" dc:"核准日期"`
	RegistrationDate  string `json:"RegistrationDate" dc:"核准日期"`
}

type OCRBankCard struct {
	Direction      int    `json:"direction" dc:"图像方向：-1 为定义,0正向,1逆时针90度,2逆时针180度,3逆时针270度"`
	BankCardNumber string `json:"bankCardNumber" dc:"银行卡号"`
	ValidDate      string `json:"validDate" dc:"有效期"`
	BankCardType   int    `json:"bankCardType" dc:"银行卡类型"`
	BankName       string `json:"bankName" dc:"银行名字"`
	HolderName     string `json:"holderName" dc:"持卡人名字"`
}

type BaiduSdkOCRBankCard struct {
	OCRBankCard
}

type BaiduSdkConfToken struct {
	BaiduSdkConf
	BaiduSdkConfAccessToken
}

type BaiduSdkConfList CollectRes[BaiduSdkConf]
