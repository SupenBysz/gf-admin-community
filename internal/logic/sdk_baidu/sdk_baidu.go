package sdk_baidu

import (
	"context"
	"time"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/kconv"
)

type sSdkBaidu struct {
	// 内存缓存
	BaiduSdkConfTokenList []*sys_model.BaiduSdkConfToken
	// 百度SDK的配置名称
	sysConfigName string
	// 缓存配置
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSdkBaidu(New())
}

// New SdkBaidu 系统配置逻辑实现
func New() sys_service.ISdkBaidu {
	return &sSdkBaidu{
		BaiduSdkConfTokenList: make([]*sys_model.BaiduSdkConfToken, 0),
		sysConfigName:         "baidu_sdk_conf",
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

// fetchBaiduSdkConfToken_bak 根据 identifier 获取百度API Token
func (s *sSdkBaidu) fetchBaiduSdkConfToken_bak(ctx context.Context, identifier string) (tokenInfo *sys_model.BaiduSdkConfToken, err error) {
	info, err := s.GetBaiduSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}

	// 1、获取百度API访问的Token
	var host = "https://aip.baidubce.com/oauth/2.0/token"
	param := g.Map{
		"grant_type":    "client_credentials",
		"client_id":     info.APIKey,
		"client_secret": info.SecretKey,
	}
	response, err := g.Client().Post(ctx, host, param)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_api_token_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}
	newTokenInfo := sys_model.BaiduSdkConfAccessToken{}
	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_api_token_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 2、组装返回数据
	var result *sys_model.BaiduSdkConfToken = nil
	newItems := garray.New() // 这个变量有何作用？？？  ----> 想要替换 s.BaiduSdkConfTokenList

	for _, item := range s.BaiduSdkConfTokenList {
		if item.Identifier == identifier { // xxAPI能力
			result = &sys_model.BaiduSdkConfToken{ // 给result赋值
				BaiduSdkConf:            *info,        // 百度SDK某个API的应用配置
				BaiduSdkConfAccessToken: newTokenInfo, // 百度SDK的API访问认证的Token
			}
			newItems.Append(*result) //  追加
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &sys_model.BaiduSdkConfToken{
			BaiduSdkConf:            *info,
			BaiduSdkConfAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	return result, nil
}

// fetchBaiduSdkConfToken 根据 identifier 获取百度API Token
func (s *sSdkBaidu) fetchBaiduSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.BaiduSdkConfToken, err error) {
	// 1、获取API的应用配置
	info, err := s.GetBaiduSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}

	// 2、获取百度API访问的Token
	var host = "https://aip.baidubce.com/oauth/2.0/token"
	param := g.Map{
		"grant_type":    "client_credentials",
		"client_id":     info.APIKey,
		"client_secret": info.SecretKey,
	}
	response, err := g.Client().Post(ctx, host, param)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_api_token_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}
	newTokenInfo := sys_model.BaiduSdkConfAccessToken{}
	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_api_token_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 3、更新Token缓存列表
	var result *sys_model.BaiduSdkConfToken = nil
	newTokenItems := make([]*sys_model.BaiduSdkConfToken, 0)
	for _, item := range s.BaiduSdkConfTokenList {
		if item.Identifier == identifier { // xxAPI能力   -----  更新
			result = &sys_model.BaiduSdkConfToken{ // 给result赋值
				BaiduSdkConf:            *info,        // 百度SDK某个API的应用配置
				BaiduSdkConfAccessToken: newTokenInfo, // 百度SDK的API访问认证的Token
			}

			newTokenItems = append(newTokenItems, result)
			continue
		}

		newTokenItems = append(newTokenItems, item)
	}

	// 4、组装返回数据 & 更新最新的缓存
	if result == nil {
		result = &sys_model.BaiduSdkConfToken{
			BaiduSdkConf:            *info,
			BaiduSdkConfAccessToken: newTokenInfo,
		}

		// 最新的Token，原先的缓存列表都不存在的  --- 加进去
		newTokenItems = append(newTokenItems, result)
	}

	s.BaiduSdkConfTokenList = newTokenItems

	return result, nil
}

// syncBaiduSdkTokenConfList 同步百度SDK应用配置信息列表缓存  （代码中要是用到了s.BaiduSdkConfTokenList缓存变量的话，一定需要在CUD操作后调用此方法更新缓存变量）
func (s *sSdkBaidu) syncBaiduSdkTokenConfList(ctx context.Context) error {
	items, err := s.GetBaiduSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]*sys_model.BaiduSdkConfToken, 0)
	for _, conf := range items {
		for _, tokenInfo := range s.BaiduSdkConfTokenList { // tokenList
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.BaiduSdkConfTokenList = newTokenItems

	return nil
}

// GetBaiduSdkConfToken 根据 identifier 查询百度SDK应用配置和Token信息
func (s *sSdkBaidu) GetBaiduSdkConfToken(ctx context.Context, identifier string) (tokenInfo *sys_model.BaiduSdkConfToken, err error) {
	// 迭代内存缓存
	for _, conf := range s.BaiduSdkConfTokenList {
		if conf.Identifier == identifier { // 缓存已存在
			return conf, nil
		}
	}
	return s.fetchBaiduSdkConfToken(ctx, identifier)
}

// GetBaiduSdkConfList 获取百度SDK应用配置列表
func (s *sSdkBaidu) GetBaiduSdkConfList(ctx context.Context) ([]*sys_model.BaiduSdkConf, error) {
	items := make([]*sys_model.BaiduSdkConf, 0)
	config, err := sys_service.SysConfig().GetByName(ctx, s.sysConfigName)
	if err != nil {
		return items, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_sdk_config_fetch_failed"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if config.Value == "" {
		return items, nil
	}

	//_ = gjson.DecodeTo(config.Value, &items)
	// 反序列化失败
	if nil != gjson.DecodeTo(config.Value, &items) {
		return items, nil
	}

	return items, nil
}

// GetBaiduSdkConf 根据 identifier 查询百度SDK应用配置信息
func (s *sSdkBaidu) GetBaiduSdkConf(ctx context.Context, identifier string) (*sys_model.BaiduSdkConf, error) {
	items, err := s.GetBaiduSdkConfList(ctx)
	if err != nil {
		return nil, err
	}

	for _, conf := range items {
		if conf.Identifier == identifier {
			return conf, nil
		}
	}
	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_sdk_app_config_query_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveBaiduSdkConf 保存百度SDK应用配信息
func (s *sSdkBaidu) SaveBaiduSdkConf(ctx context.Context, info *sys_model.BaiduSdkConf, isCreate bool) (*sys_model.BaiduSdkConf, error) {
	items, _ := s.GetBaiduSdkConfList(ctx)

	isHas := false
	newItems := make([]*sys_model.BaiduSdkConf, 0)
	for _, conf := range items {
		if conf.Identifier == info.Identifier { // 如果标识符相等，说明已经存在，将最新的追加到新的容器中 （更新已存在的数据）
			isHas = true
			newItems = append(newItems, info)
			continue
		}

		newItems = append(newItems, conf) // 将旧的Item追加到新的容器中 （保持原有的）
	}

	if !isHas { // 不存在
		if isCreate { // 创建 --- 追加info （原有的 + 最新的Info）
			newItems = append(newItems, info)
		} else { // 更新 --- 不存在此配置，那么就提示错误
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_sdk_config_save_failed_identifier_error"), "", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err := sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})
	if nil != err {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_sdk_config_save_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除列表缓存 -- 不需要，Dao对象会自动删除
	//daoctl.RemoveQueryCache(sys_dao.SysConfig.DB(), s.sysConfigName)

	// 同步百度SDK应用配置信息列表缓存
	s.syncBaiduSdkTokenConfList(ctx)

	return info, nil
}

// DeleteBaiduSdkConf 删除百度SDK应用配置信息
func (s *sSdkBaidu) DeleteBaiduSdkConf(ctx context.Context, identifier string) (bool, error) {
	// 查询已存在的百度SDK列表
	items, err := s.GetBaiduSdkConfList(ctx)

	isHas := false
	newItems := garray.New(false) // 初始化切片
	for _, conf := range items {
		if conf.Identifier == identifier { // 已存在  ----  删除操作：不加入到newItems 新容器中
			isHas = true
			continue
		}

		newItems.Append(conf) // 不需要删除的 ----  删除操作：保留到newItems 新容器中
	}
	if !isHas {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_sdk_config_delete_not_exist", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	_, err = sys_service.SysConfig().SaveConfig(ctx, &sys_model.SysConfig{
		Name:  s.sysConfigName,
		Value: jsonString,
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_baidu_sdk_config_delete_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除列表缓存 （无需，Dao对象自带该能力）
	//daoctl.RemoveQueryCache(sys_dao.SysConfig.DB(), s.sysConfigName)
	// 同步 内存缓存 列表
	s.syncBaiduSdkTokenConfList(ctx)
	return true, nil
}

// OCRBankCard OCR识别银行卡
func (s *sSdkBaidu) OCRBankCard(ctx context.Context, imageBase64 string) (*sys_model.OCRBankCard, error) {
	// 请求参数
	param := g.Map{
		"image": imageBase64,
	}

	// 获取图像应用的请求Token
	tokenInfo, err := s.GetBaiduSdkConfToken(ctx, "certificate_orc")

	if err != nil {
		return nil, err
	}

	// 模拟客户端发起请求
	response, err := g.Client().Header(g.MapStrStr{"Content-Type": "application/x-www-form-urlencoded"}).
		Post(ctx, "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard?access_token="+tokenInfo.AccessToken, param)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_license_data_parsing_failed"), "", sys_dao.SysFile.Table())
	}

	// json解析
	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_license_recognition_failed"), "", sys_dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_license_info_recognition_failed"), "", sys_dao.SysFile.Table())
	}

	// 组装返回数据
	ret := sys_model.OCRBankCard{
		Direction:      jsonObj.Get("direction", -1).Int(),
		BankCardNumber: jsonObj.Get("result.bank_card_number", "").String(),
		ValidDate:      jsonObj.Get("result.valid_date", "").String(),
		BankCardType:   jsonObj.Get("result.bank_card_type", -1).Int(),
		BankName:       jsonObj.Get("result.bank_name", "").String(),
		HolderName:     jsonObj.Get("result.holder_name", "").String(),
	}

	return &ret, nil

}

// OCRIDCard OCR识别身份证
func (s *sSdkBaidu) OCRIDCard(ctx context.Context, imageBase64 string, detectRisk string, idCardSide string) (*sys_model.BaiduSdkOCRIDCard, error) {

	param := g.Map{
		"image":            imageBase64,
		"id_card_side":     idCardSide,
		"detect_risk":      detectRisk,
		"detect_direction": "true",
	}

	tokenInfo, err := s.GetBaiduSdkConfToken(ctx, "certificate_orc")
	if err != nil {
		return nil, err
	}

	response, err := g.Client().Header(g.MapStrStr{"Content-Type": "application/x-www-form-urlencoded"}).
		Post(ctx, "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard?access_token="+tokenInfo.AccessToken, param)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	ret := sys_model.BaiduSdkOCRIDCard{}

	if idCardSide == "front" {
		ret.OCRIDCardA = &sys_model.BaiduSdkOCRIDCardA{
			Direction:      jsonObj.Get("direction", -1).Int(),
			ImageStateText: jsonObj.Get("image_status", "unknown").String(),
			RiskType:       jsonObj.Get("risk_type", "unknown").String(),
			Address:        jsonObj.Get("words_result.住址.words", "").String(),
			IDCardNumber:   jsonObj.Get("words_result.公民身份号码.words", "").String(),
			Birthday:       jsonObj.Get("words_result.出生.words", "").String(),
			Realname:       jsonObj.Get("words_result.姓名.words", "").String(),
			Gender:         jsonObj.Get("words_result.性别.words", "").String(),
			Nation:         jsonObj.Get("words_result.民族.words", "").String(),
		}
	} else if idCardSide == "back" {
		ret.OCRIDCardB = &sys_model.BaiduSdkOCRIDCardB{
			ExpiryDate:       jsonObj.Get("words_result.失效日期.words", "").String(),
			IssuingAuthority: jsonObj.Get("words_result.签发机关.words", "").String(),
			IssuingDate:      jsonObj.Get("words_result.签发日期.words", "").String(),
		}
	}

	return &ret, nil
}

// OCRBusinessLicense OCR识别营业执照
func (s *sSdkBaidu) OCRBusinessLicense(ctx context.Context, imageBase64 string) (*sys_model.BusinessLicenseOCR, error) {
	param := g.Map{
		"image":     imageBase64,
		"risk_warn": "true",
	}

	tokenInfo, err := s.GetBaiduSdkConfToken(ctx, "certificate_orc")
	if err != nil {
		return nil, err
	}

	response, err := g.Client().Header(g.MapStrStr{"Content-Type": "application/x-www-form-urlencoded"}).
		Post(ctx, "https://aip.baidubce.com/rest/2.0/ocr/v1/business_license?access_token="+tokenInfo.AccessToken, param)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	ret := sys_model.BusinessLicenseOCR{
		Direction:         jsonObj.Get("direction", -1).Int(),
		RiskType:          jsonObj.Get("risk_type", "unknown").String(),
		CreditCode:        jsonObj.Get("words_result.社会信用代码.words", "").String(),
		CombiningForm:     jsonObj.Get("words_result.组成形式.words", "").String(),
		BusinessScope:     jsonObj.Get("words_result.经营范围.words", "").String(),
		EstablishmentDate: jsonObj.Get("words_result.成立日期.words", "").String(),
		LegalPerson:       jsonObj.Get("words_result.法人.words", "").String(),
		RegisteredCapital: jsonObj.Get("words_result.注册资本.words", "").String(),
		CertificateNumber: jsonObj.Get("words_result.证件编号.words", "").String(),
		RegisteredAddress: jsonObj.Get("words_result.地址.words", "").String(),
		CompanyName:       jsonObj.Get("words_result.单位名称.words", "").String(),
		ExpirationDate:    jsonObj.Get("words_result.有效期.words", "").String(),
		ApprovalDate:      jsonObj.Get("words_result.核准日期.words", "").String(),
		RegistrationDate:  jsonObj.Get("words_result.类型.words", "").String(),
	}

	return &ret, nil
}

// AuditPicture 审核图片
func (s *sSdkBaidu) AuditPicture(ctx context.Context, imageBase64 string, imageType uint64) (*sys_model.PictureWithOCR, error) {
	param := g.Map{
		"image":   imageBase64,
		"imgType": imageType,
	}

	tokenInfo, err := s.GetBaiduSdkConfToken(ctx, "audit_picture")
	if err != nil {
		return nil, err
	}

	response, err := g.Client().Header(g.MapStrStr{"Content-Type": "application/x-www-form-urlencoded"}).
		Post(ctx, "https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined?access_token="+tokenInfo.AccessToken, param)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("error_baidu_image_review_failed"), "", sys_dao.SysFile.Table())
	}

	ret := sys_model.PictureWithOCR{Data: make([]sys_model.DescriptionData, 0)}

	if jsonObj.Get("conclusionType").Int() == 1 {
		ret.Conclusion = jsonObj.Get("conclusion").String()
		ret.ConclusionType = jsonObj.Get("conclusionType", "").Int()
	}

	if jsonObj.Get("conclusionType").Int() == 2 {
		ret.Conclusion = jsonObj.Get("conclusion").String()
		ret.ConclusionType = jsonObj.Get("conclusionType", "").Int()
		data := sys_model.DescriptionData{}
		for _, item := range jsonObj.Get("data").Array() {
			record := kconv.Struct(item, &sys_model.DescriptionData{})
			data.Type = record.Type
			data.SubType = record.SubType
			data.Conclusion = record.Conclusion
			data.ConclusionType = record.ConclusionType
			data.Msg = record.Msg
		}
		ret.Data = append(ret.Data, data)
	}

	return &ret, nil
}
