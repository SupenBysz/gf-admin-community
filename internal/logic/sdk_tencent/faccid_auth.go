package sdk_tencent

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"net/url"
	"sort"
	"strings"
)

/*
	人脸核身：
	http接入方式，无SDK
*/

/*
	人脸核身：https://cloud.tencent.com/document/product/1007/61073
		接入方式： `H5（移动端浏览器）`
		应用场景：`H5(非微信原生)`
	对接步骤：
		1、App 调用 H5 兼容性配置 (ios, android, Harmony)
		2、合作方后台上传身份信息 (idCard, name) 封装方法：GetAdvFaceId()
		3、启动 H5 人脸核身
		4、人脸核身结果返回及跳转   `callbackUrl`
		5、人脸核身结果查询
*/

// DetectAuthByFaceId 腾讯云-人脸核身API （接入方式 `H5（移动端浏览器）`、应用场景`H5(非微信原生)`）

// GetAdvFaceIdAndAuth 合作方后台上传身份信息 (idCard, name)获取FaceId + 拼接H5人脸核身认证URL
func (s *sSdkTencent) GetAdvFaceIdAndAuth(ctx context.Context, userId int64, orderId int64, idCard, name string, callbackUrl string) (*sys_model.StartAdvFaceAuthRes, error) { // userId 就是sysUserID、orderId 就是sysAuditId

	getAdvFaceIdRes := sys_model.GetAdvFaceIdRes{}
	startAdvFaceAuthRes := sys_model.StartAdvFaceAuthRes{}
	authUrl := ""

	/*
		合作方后台上传身份信息
		所需步骤：
			1、准备参数、生成签名
			2、请求腾讯云接口，`合作方后台上传身份信息`，拿到响应参数
			3、返回结果
	*/

	/*
		启动 H5 人脸核身
		所需步骤：
			1、准备参数、生成签名
			2、拼接认证 URL
			3、返回认证 URL
	*/

	// ************* 步骤 1: 准备参数、生成签名 *************

	// 获取配置
	//config, err := s.GetTencentSdkConf(ctx, "facere_auth") // 正式环境
	config, err := s.GetTencentSdkConf(ctx, "facere_auth_test") // 测试环境
	if err != nil {
		return nil, err
	}

	/*
		所需参数：
			- appId： 获取 WBappid 指引在人脸核身控制台内申请 `https://cloud.tencent.com/document/product/1007/61073`
			- userId： 用户唯一标识 userId自定义即可
			- version： 参数值为：1.0.0
			- ticket：合作伙伴服务端获取的 ticket，注意是 SIGN 类型 `https://cloud.tencent.com/document/product/1007/57613`
			- nonce：必须是32位随机数，自定义即可
	*/

	wBAppId := config.AppID
	wBAppSecret := config.SecretKey
	accessToken := ""
	ticket := ""
	nonce := ""
	version := "1.0.0"

	// 1.1 获取 AccessToken
	accessToken, err = s.GetAccessToken(ctx, wBAppId, wBAppSecret, version)
	if err != nil {
		return nil, err
	}

	// 1.2 获取Ticket
	ticket, err = s.GetApiTicket(ctx, wBAppId, accessToken, version)
	if err != nil {
		return nil, err
	}

	// 1.3生成 nonce
	{
		nonce = uuid.New().String() // db25ff2b-aae1-4f95-9a3f-4ebe565fa4f1
		nonce = gstr.Replace(nonce, "-", "")
	}

	// 1.4 将 nonce、userId、appId 连同 ticket、version 共五个参数的值进行字典序排序。
	strs := []string{nonce, gconv.String(userId), wBAppId, ticket, version}
	sortStrings(strs)

	//1.5  将排序后的所有参数字符串拼接成一个字符串。
	strJoin := strings.Join(strs, "")
	fmt.Println(strJoin)

	//1.6 将排序后的字符串进行 SHA1 编码，编码后的40位字符串作为签名（sign）。
	sign := sha1Encode(strJoin)

	// ************* 步骤 2：请求腾讯云接口，`合作方后台上传身份信息`，拿到响应参数 *************
	{
		// POST https://kyc1.qcloud.com/api/server/getAdvFaceId?orderNo=xxx
		url := fmt.Sprintf("https://kyc1.qcloud.com/api/server/getAdvFaceId?orderNo=%s", gconv.String(orderId))
		info := g.Map{
			"appId":   wBAppId,
			"orderNo": gconv.String(orderId),
			"name":    name,
			"idNo":    idCard,
			"userId":  gconv.String(userId),
			"version": version,
			"sign":    sign,
			"nonce":   nonce,
		}
		reqData, _ := gjson.Encode(info)
		response := g.Client().PostContent(ctx, url, reqData)
		// {"code":"0","msg":"请求成功","bizSeqNo":"24072620001184433114441312514267","result":{"bizSeqNo":"24072620001184433114441312514267","transactionTime":"20240726144413","oriCode":"0","orderNo":"9164541471817797","faceId":"tx11a42739e56889725d5be8ab492d9d","optimalDomain":"kyc1.qcloud.com","success":true},"transactionTime":"20240726144413"}
		_ = gjson.DecodeTo(response, &getAdvFaceIdRes)
		if getAdvFaceIdRes.Code != "0" || getAdvFaceIdRes.Result.FaceId == "" {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "合作方后台上传身份信息,获取FaceId 失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}

	}

	// ************* 步骤 3：启动 H5 人脸核身 *************
	{
		// 1.4 将 appId、userId、orderNo、version、faceId 连同 ticket、nonce 共7个参数的值进行字典序排序。 TODO 需要确认ticket nonce 是否需要和上传的接口保持一致
		strs2 := []string{wBAppId, gconv.String(userId), getAdvFaceIdRes.Result.OrderNo, version, getAdvFaceIdRes.Result.FaceId, ticket, nonce}
		sortStrings(strs2)

		//1.5  将排序后的所有参数字符串拼接成一个字符串。
		strJoin2 := strings.Join(strs2, "")
		fmt.Println(strJoin2)

		//1.6 将排序后的字符串进行 SHA1 编码，编码后的40位字符串作为签名（sign）。
		sign2 := sha1Encode(strJoin2)

		if "" == getAdvFaceIdRes.Result.OptimalDomain {
			getAdvFaceIdRes.Result.OptimalDomain = "kyc1.qcloud.com"
		}

		callbackUrlEncoded := url.QueryEscape(callbackUrl)

		//// GET https://{optimalDomain}/api/web/login
		//requestUrl := fmt.Sprintf("https://%s/api/web/login", getAdvFaceIdRes.Result.OptimalDomain)
		//info := g.Map{
		//	"appId":        wBAppId,
		//	"version":      version,
		//	"nonce":        nonce,
		//	"orderNo":      getAdvFaceIdRes.Result.OrderNo,
		//	"faceId":       getAdvFaceIdRes.Result.FaceId,
		//	"url":          callbackUrlEncoded, // H5 人脸核身完成后回调的第三方 URL，需要第三方提供完整 URL 且做 URL Encode。完整 URL Encode 示例：原 URL：https://cloud.tencent.com ---> Encode 后：https%3a%2f%2fcloud.tencent.com
		//	"resultType":   "1",                // 是否显示结果页面，参数值为“1”时直接跳转到 url 回调地址，null 或其他值跳转提供的结果页面
		//	"userId":       gconv.String(userId),
		//	"sign":         sign,
		//	"from":         "browser", // browser：表示在浏览器启动刷脸；App：表示在 App 里启动刷脸，默认值为 App
		//	"redirectType": "1",       // 跳转模式，参数值为“1”时，刷脸页面使用 replace 方式跳转，不在浏览器 history 中留下记录；不传或其他值则正常跳转
		//}
		//reqData, _ := gjson.Encode(info)

		authUrl = fmt.Sprintf("https://%s/api/web/login?appId=%s&version=%s&nonce=%s&orderNo=%s&faceId=%s&url=%s&from=%s&userId=%s&sign=%s&redirectType=%s",
			getAdvFaceIdRes.Result.OptimalDomain, wBAppId, version, nonce, getAdvFaceIdRes.Result.OrderNo, getAdvFaceIdRes.Result.FaceId, callbackUrlEncoded, "browser", gconv.String(userId), sign2, "1")
	}

	startAdvFaceAuthRes.OrderNo = getAdvFaceIdRes.Result.OrderNo
	startAdvFaceAuthRes.FaceId = getAdvFaceIdRes.Result.FaceId
	startAdvFaceAuthRes.Url = authUrl

	// ************* 步骤 4：返回认证结果，URL等 *************
	return &startAdvFaceAuthRes, nil
}

// QueryFaceRecord 腾讯云-人脸核身结果查询
func (s *sSdkTencent) QueryFaceRecord(ctx context.Context, orderNo string) (*sys_model.QueryFaceRecordRes, error) {
	// POST https://kyc1.qcloud.com/api/v2/base/queryfacerecord?orderNo=xxx

	ret := sys_model.QueryFaceRecordRes{}

	// ************* 步骤 1: 准备参数、生成签名 *************
	//config, err := s.GetTencentSdkConf(ctx, "facere_auth")
	config, err := s.GetTencentSdkConf(ctx, "facere_auth_test")
	if err != nil {
		return nil, err
	}

	wBAppId := config.AppID
	wBAppSecret := config.SecretKey
	accessToken := ""
	ticket := ""
	nonce := ""
	version := "1.0.0"

	// 1.1 获取 AccessToken
	accessToken, err = s.GetAccessToken(ctx, wBAppId, wBAppSecret, version)
	if err != nil {
		return nil, err
	}

	// 1.2 获取Ticket
	ticket, err = s.GetApiTicket(ctx, wBAppId, accessToken, version)
	if err != nil {
		return nil, err
	}

	// 1.3生成 nonce
	{
		nonce = uuid.New().String() // db25ff2b-aae1-4f95-9a3f-4ebe565fa4f1
		nonce = gstr.Replace(nonce, "-", "")
	}

	// 1.4 将  appId、orderNo、version、ticket、nonce 共5个参数的值进行字典序排序。
	strs := []string{wBAppId, orderNo, version, ticket, nonce}
	sortStrings(strs)

	//1.5  将排序后的所有参数字符串拼接成一个字符串。
	strJoin := strings.Join(strs, "")
	fmt.Println(strJoin)

	//1.6 将排序后的字符串进行 SHA1 编码，编码后的40位字符串作为签名（sign）。
	sign := sha1Encode(strJoin)

	// ************* 步骤 2: 请求腾讯云接口，`查询核身结果` *************
	url := fmt.Sprintf("https://kyc1.qcloud.com/api/v2/base/queryfacerecord?orderNo=%s", orderNo)
	info := g.Map{
		"appId":   wBAppId,
		"version": version,
		"nonce":   nonce,
		"orderNo": orderNo,
		"sign":    sign,
		//"getFile": 1, // 是否需要获取人脸识别的视频和文件，值为1则返回视频和照片、值为2则返回照片、值为3则返回视频；其他则不返回
	}
	reqData, _ := gjson.Encode(info)
	response := g.Client().PostContent(ctx, url, reqData)

	// ************* 步骤 3：解析响应参数 *************
	_ = gjson.DecodeTo(response, &ret)
	//if ret.Code != "0" || ret.Result.OrderNo == "" {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取人脸核身结果查询失败", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
	//}

	return &ret, err
}

// sortStrings 字符串字典排序
func sortStrings(strs []string) {
	sort.Strings(strs)
}

// sha1Encode 字符串 SHA1 编码
func sha1Encode(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	sha1Hash := h.Sum(nil)
	return hex.EncodeToString(sha1Hash)
}
