package sdk_tencent

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	faceid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
	活体人脸核身 （人脸核身PasS服务接口）
*/

// LivenessRecognition 人脸核身(SDK接入模式)
func (s *sSdkTencent) LivenessRecognition(ctx context.Context, idCard, name, livenessType string) {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取

	config, err := s.GetTencentSdkConf(ctx, "tencent_config")

	credential := common.NewCredential(
		config.AESKey,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := faceid.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := faceid.NewLivenessRecognitionRequest()
	request.IdCard = common.StringPtr(idCard)
	request.Name = common.StringPtr(name)
	request.LivenessType = common.StringPtr(livenessType)

	// 返回的resp是一个LivenessRecognitionResponse的实例，与请求对象对应
	response, err := client.LivenessRecognition(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

}

// LivenessRecognition_Http 人脸核身(HTTP接入模式)
func (s *sSdkTencent) LivenessRecognition_Http(ctx context.Context, action, secretId, secretKey string) { // LivenessRecognition 活体核验
	service := "faceid"
	version := "2018-03-01"
	//action := "LivenessRecognition"
	region := ""
	token := ""
	host := "faceid.tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"
	var timestamp = time.Now().Unix()

	// ************* 步骤 1：拼接规范请求串 *************
	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	contentType := "application/json; charset=utf-8"
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\nx-tc-action:%s\n",
		contentType, host, strings.ToLower(action))
	signedHeaders := "content-type;host;x-tc-action"
	payload := "{}"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	log.Println(canonicalRequest)

	// ************* 步骤 2：拼接待签名字符串 *************
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)
	log.Println(string2sign)

	// ************* 步骤 3：计算签名 *************
	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	log.Println(signature)

	// ************* 步骤 4：拼接 Authorization *************
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature)
	log.Println(authorization)

	// ************* 步骤 5：构造并发起请求 *************
	url := "https://" + host
	httpRequest, _ := http.NewRequest("POST", url, strings.NewReader(payload))
	httpRequest.Header = map[string][]string{
		"Host":           {host},
		"X-TC-Action":    {action},
		"X-TC-Version":   {version},
		"X-TC-Timestamp": {strconv.FormatInt(timestamp, 10)},
		"Content-Type":   {contentType},
		"Authorization":  {authorization},
	}
	if region != "" {
		httpRequest.Header["X-TC-Region"] = []string{region}
	}
	if token != "" {
		httpRequest.Header["X-TC-Token"] = []string{token}
	}
	httpClient := http.Client{}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(body.String())
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
