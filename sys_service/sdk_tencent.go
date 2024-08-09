// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type (
	ISdkTencent interface {
		// GetAccessToken 获取 AccessToken
		GetAccessToken(ctx context.Context, wBAppId, wBAppSecret, version string) (accessToken string, err error)
		// GetApiTicket 获取API Ticket
		GetApiTicket(ctx context.Context, wBAppId, accessToken, version string) (ticket string, err error)
		// GetAdvFaceIdAndAuth 合作方后台上传身份信息 (idCard, name)获取FaceId + 拼接H5人脸核身认证URL
		// func (s *sSdkTencent) getAdvFaceIdAndAuth(ctx context.Context, userId int64, orderId int64, idCard, name string, callbackUrl string) (*sys_model.GetAdvFaceIdRes, error) { // userId 就是sysUserID、orderId 就是sysAuditId
		GetAdvFaceIdAndAuth(ctx context.Context, userId int64, orderId int64, idCard, name string, callbackUrl string) (*sys_model.StartAdvFaceAuthRes, error)
		// QueryFaceRecord 腾讯云-人脸核身结果查询
		QueryFaceRecord(ctx context.Context, orderNo string) (*sys_model.QueryFaceRecordRes, error)
		// LivenessRecognition_Http 人脸核身(HTTP接入模式)
		LivenessRecognition_Http(ctx context.Context, action, secretId, secretKey string)
		// GetTencentSdkConfList 获取腾讯云SDK应用配置列表
		GetTencentSdkConfList(ctx context.Context) ([]*sys_model.TencentSdkConf, error)
		// GetTencentSdkConf 根据identifier标识获取SDK配置信息
		GetTencentSdkConf(ctx context.Context, identifier string) (tokenInfo *sys_model.TencentSdkConf, err error)
		// SaveTencentSdkConf 保存SDK应用配信息, isCreate判断是更新还是新建
		SaveTencentSdkConf(ctx context.Context, info *sys_model.TencentSdkConf, isCreate bool) (*sys_model.TencentSdkConf, error)
		// DeleteTencentSdkConf 删除腾讯云SDK应用配置信息
		DeleteTencentSdkConf(ctx context.Context, identifier string) (bool, error)
		// LivenessRecognition 人脸核身(SDK接入模式)
		LivenessRecognition(ctx context.Context, idCard, name, livenessType string)
		// DetectAuth 腾讯云-实名核身鉴权
		DetectAuth(ctx context.Context, idCard, name, returnUrl string) (*sys_model.DetectAuthRes, error)
		// GetDetectAuthResult 获取腾讯云-实名核身鉴权结果
		GetDetectAuthResult(ctx context.Context, bizToken string, ruleId ...string) (*sys_model.GetDetectAuthResultRes, error)
		// GetDetectAuthPlusResponse 获取腾讯云-实名核身鉴权增强版结果 （v3.0接口）
		GetDetectAuthPlusResponse(ctx context.Context, bizToken, ruleId string) (*sys_model.GetDetectAuthPlusResponseRes, error)
	}
)

var (
	localSdkTencent ISdkTencent
)

func SdkTencent() ISdkTencent {
	if localSdkTencent == nil {
		panic("implement not found for interface ISdkTencent, forgot register?")
	}
	return localSdkTencent
}

func RegisterSdkTencent(i ISdkTencent) {
	localSdkTencent = i
}
