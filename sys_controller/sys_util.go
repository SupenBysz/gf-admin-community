package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
)

// SysUtil 工具
var SysUtil = cSysUtil{}

type cSysUtil struct{}

// *********************************** 工具 *****************************************

// GetNerChEcom 中文分词
func (c *cSysUtil) GetNerChEcom(ctx context.Context, req *sys_api.GetNerChEcomReq) (sys_model.AliyunNlpDataRes, error) {
	ret, err := sys_service.SdkAliyun().GetWsCustomizedChGeneral(ctx, req.Text)
	return ret, err
}

// GetAiSummary 文字分析总结
func (c *cSysUtil) GetAiSummary(ctx context.Context, req *sys_api.GetAiSummaryReq) (api_v1.StringRes, error) {
	ret, err := sys_service.SdkBaidu().GetAiSummary(ctx, req.Text, req.Identifier)

	return (api_v1.StringRes)(ret), err
}

// ************************************ 认证 *****************************************

// LivenessRecognition 身份验证/活体核验
func (c *cSysUtil) LivenessRecognition(ctx context.Context, req *sys_api.LivenessRecognitionReq) (api_v1.BoolRes, error) {
	sys_service.SdkTencent().LivenessRecognition(ctx, req.IdCard, req.Name, req.LivenessType)

	return false, nil
	//return ret, err
}

// DetectAuth 腾讯云-实名核身鉴权
func (c *cSysUtil) DetectAuth(ctx context.Context, req *sys_api.DetectAuthReq) (*sys_model.DetectAuthRes, error) {
	ret, err := sys_service.SdkTencent().DetectAuth(ctx, req.IdCard, req.Name, req.ReturnUrl)

	return ret, err
}

// GetDetectAuthResult 腾讯云-查询实名核身鉴权结果
func (c *cSysUtil) GetDetectAuthResult(ctx context.Context, req *sys_api.GetDetectAuthResultReq) (*sys_model.GetDetectAuthResultRes, error) {
	ret, err := sys_service.SdkTencent().GetDetectAuthResult(ctx, req.BizToken)

	return ret, err
}

// StartAdvFaceAuth  腾讯云-启动H5人脸核身
func (c *cSysUtil) StartAdvFaceAuth(ctx context.Context, req *sys_api.StartAdvFaceAuthReq) (*sys_model.StartAdvFaceAuthRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	orderId := idgen.NextId() // 注意：实际业务中，orderId需要管理，后续还需要通过此查询结果

	ret, err := sys_service.SdkTencent().GetAdvFaceIdAndAuth(ctx, user.Id, orderId, req.IdCard, req.Name, req.CallbackUrl)

	return ret, err
}

// FaceAuthCallback 人脸核身回调地址
func (c *cSysUtil) FaceAuthCallback(ctx context.Context, req *sys_api.FaceAuthCallbackReq) (api_v1.BoolRes, error) {
	res := req.Code == "0"

	// 注意：实际业务中，处理逻辑更加复杂，这里仅仅做了简单的结果处理返回

	return res == true, nil
}

// QueryFaceRecord 腾讯云-人脸核身结果查询
func (c *cSysUtil) QueryFaceRecord(ctx context.Context, req *sys_api.QueryFaceRecordReq) (*sys_model.QueryFaceRecordRes, error) {
	ret, err := sys_service.SdkTencent().QueryFaceRecord(ctx, req.OrderNo)

	return ret, err
}
