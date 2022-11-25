package sdk_baidu

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type sSdkBaidu struct {
	BaiduSdkConfTokenList []model.BaiduSdkConfToken
	CacheDuration         time.Duration
	sysConfigName         string
}

func init() {
	service.RegisterSdkBaidu(New())
}

// New SdkBaidu 系统配置逻辑实现
func New() *sSdkBaidu {
	return &sSdkBaidu{
		BaiduSdkConfTokenList: make([]model.BaiduSdkConfToken, 0),
		CacheDuration:         time.Hour,
		sysConfigName:         "baidu_sdk_conf",
	}
}

// fetchBaiduSdkConfToken 根据 identifier 获取百度API Token
func (s *sSdkBaidu) fetchBaiduSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.BaiduSdkConfToken, err error) {
	info, err := s.GetBaiduSdkConf(ctx, identifier)
	if err != nil {
		return nil, err
	}

	var host = "https://aip.baidubce.com/oauth/2.0/token"

	param := g.Map{
		"grant_type":    "client_credentials",
		"client_id":     info.APIKey,
		"client_secret": info.SecretKey,
	}

	response, err := g.Client().Post(ctx, host, param)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "获取百度API Token 失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	newTokenInfo := model.BaiduSdkConfAccessToken{}

	err = gjson.DecodeTo(response.ReadAllString(), &newTokenInfo)
	if nil != err {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "获取百度API Token 失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	var result *model.BaiduSdkConfToken = nil
	newItems := garray.New()
	for _, item := range s.BaiduSdkConfTokenList {
		if item.Identifier == identifier {
			result = &model.BaiduSdkConfToken{
				BaiduSdkConf:            *info,
				BaiduSdkConfAccessToken: newTokenInfo,
			}
			newItems.Append(*result)
			continue
		}

		newItems.Append(*result)
	}

	if result == nil {
		result = &model.BaiduSdkConfToken{
			BaiduSdkConf:            *info,
			BaiduSdkConfAccessToken: newTokenInfo,
		}
		newItems.Append(*result)
	}

	return result, nil
}

// syncBaiduSdkConfTokenList 同步百度SDK应用配置信息Token列表缓存
func (s *sSdkBaidu) syncBaiduSdkConfTokenList(ctx context.Context) error {
	items, err := s.GetBaiduSdkConfList(ctx)
	if err != nil {
		return err
	}

	newTokenItems := make([]model.BaiduSdkConfToken, 0)
	for _, conf := range *items {
		for _, tokenInfo := range s.BaiduSdkConfTokenList {
			if tokenInfo.Identifier == conf.Identifier {
				newTokenItems = append(newTokenItems, tokenInfo)
			}
		}
	}

	s.BaiduSdkConfTokenList = newTokenItems

	return nil
}

// GetBaiduSdkConfToken 根据 identifier 查询百度SDK应用配置和Token信息
func (s *sSdkBaidu) GetBaiduSdkConfToken(ctx context.Context, identifier string) (tokenInfo *model.BaiduSdkConfToken, err error) {
	for _, conf := range s.BaiduSdkConfTokenList {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}
	return s.fetchBaiduSdkConfToken(ctx, identifier)
}

// GetBaiduSdkConfList 获取百度SDK应用配置列表
func (s *sSdkBaidu) GetBaiduSdkConfList(ctx context.Context) (*[]model.BaiduSdkConf, error) {
	items := make([]model.BaiduSdkConf, 0)
	data := entity.SysConfig{}
	err := dao.SysConfig.Ctx(ctx).
		Cache(gdb.CacheOption{
			Duration: s.CacheDuration,
			Name:     s.sysConfigName,
			Force:    true,
		}).
		Where(do.SysConfig{Name: s.sysConfigName}).Scan(&data)
	if err != nil && err != sql.ErrNoRows {
		return &items, service.SysLogs().ErrorSimple(ctx, gerror.New("百度SDK配置信息获取失败"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	if data.Value == "" {
		return &items, nil
	}

	if nil == gjson.DecodeTo(data.Value, &items) {
		return &items, nil
	}

	return &items, nil
}

// GetBaiduSdkConf 根据 identifier 查询百度SDK应用配置信息
func (s *sSdkBaidu) GetBaiduSdkConf(ctx context.Context, identifier string) (*model.BaiduSdkConf, error) {
	items, err := s.GetBaiduSdkConfList(ctx)
	if err != nil {
		return nil, err
	}

	for _, conf := range *items {
		if conf.Identifier == identifier {
			return &conf, nil
		}
	}
	return nil, service.SysLogs().ErrorSimple(ctx, err, "根据 identifier 查询百度SDK应用配置信息失败", dao.SysConfig.Table()+":"+s.sysConfigName)
}

// SaveBaiduSdkConf 保存百度SDK应用配信息
func (s *sSdkBaidu) SaveBaiduSdkConf(ctx context.Context, info model.BaiduSdkConf, isCreate bool) (*model.BaiduSdkConf, error) {
	items, _ := s.GetBaiduSdkConfList(ctx)

	isHas := false
	newItems := make([]model.BaiduSdkConf, 0)
	for _, conf := range *items {
		if conf.Identifier == info.Identifier {
			isHas = true
			newItems = append(newItems, info)
			continue
		}

		newItems = append(newItems, conf)
	}
	if !isHas {
		if isCreate {
			newItems = append(newItems, info)
		} else {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("百度SDK配置信息保存失败，标识符错误"), "", dao.SysConfig.Table()+":"+s.sysConfigName)
		}
	}

	jsonString := gjson.MustEncodeString(newItems)

	count, err := dao.SysConfig.Ctx(ctx).Count(do.SysConfig{Name: s.sysConfigName})
	if count > 0 {
		_, err = dao.SysConfig.Ctx(ctx).Where(do.SysConfig{Name: s.sysConfigName}).Save(do.SysConfig{Value: jsonString})
	} else {
		_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{Name: s.sysConfigName, Value: jsonString})
	}

	if nil != err {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "百度SDK配置信息保存失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除列表缓存
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)
	// 同步 Token 列表
	s.syncBaiduSdkConfTokenList(ctx)
	return &info, nil
}

// DeleteBaiduSdkConf 删除百度SDK应用配置信息
func (s *sSdkBaidu) DeleteBaiduSdkConf(ctx context.Context, identifier string) (bool, error) {
	items, err := s.GetBaiduSdkConfList(ctx)

	isHas := false
	newItems := garray.New(false)
	for _, conf := range *items {
		if conf.Identifier == identifier {
			isHas = true
			continue
		}

		newItems.Append(conf)
	}
	if !isHas {
		return false, service.SysLogs().ErrorSimple(ctx, err, "要删除的百度SDK配置信息不存在", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	jsonString := gjson.MustEncodeString(newItems)
	if _, err = dao.SysConfig.Ctx(ctx).Where(do.SysConfig{Name: s.sysConfigName}).Save(do.SysConfig{Value: jsonString}); err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "百度SDK配置信息删除失败", dao.SysConfig.Table()+":"+s.sysConfigName)
	}

	// 移除列表缓存
	daoctl.RemoveQueryCache(dao.SysConfig.DB(), s.sysConfigName)
	// 同步 Token 列表
	s.syncBaiduSdkConfTokenList(ctx)
	return true, nil
}

// OCRBankCard OCR识别银行卡
func (s *sSdkBaidu) OCRBankCard(ctx context.Context, imageBase64 string) (*model.OCRBankCard, error) {
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
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("识别证照信息失败"), "", dao.SysFile.Table())
	}

	// json解析
	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("解析证照识别数据失败"), "", dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("证照识别失败"), "", dao.SysFile.Table())
	}

	// 组装返回数据
	ret := model.OCRBankCard{
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
func (s *sSdkBaidu) OCRIDCard(ctx context.Context, imageBase64 string, detectRisk string, idCardSide string) (*model.BaiduSdkOCRIDCard, error) {

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
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("识别证照信息失败"), "", dao.SysFile.Table())
	}

	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("解析证照识别数据失败"), "", dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("证照识别失败"), "", dao.SysFile.Table())
	}

	ret := model.BaiduSdkOCRIDCard{}

	if idCardSide == "front" {
		ret.OCRIDCardA = &model.BaiduSdkOCRIDCardA{
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
		ret.OCRIDCardB = &model.BaiduSdkOCRIDCardB{
			ExpiryDate:       jsonObj.Get("words_result.失效日期.words", "").String(),
			IssuingAuthority: jsonObj.Get("words_result.签发机关.words", "").String(),
			IssuingDate:      jsonObj.Get("words_result.签发日期.words", "").String(),
		}
	}

	return &ret, nil
}

// OCRBusinessLicense OCR识别营业执照
func (s *sSdkBaidu) OCRBusinessLicense(ctx context.Context, imageBase64 string) (*model.BusinessLicenseOCR, error) {
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
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("识别证照信息失败"), "", dao.SysFile.Table())
	}

	jsonObj, err := gjson.DecodeToJson(response.ReadAll())

	if err != nil || jsonObj.IsNil() {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("解析证照识别数据失败"), "", dao.SysFile.Table())
	}

	if jsonObj.Get("error_code").Int() != 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("证照识别失败"), "", dao.SysFile.Table())
	}

	ret := model.BusinessLicenseOCR{
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
