package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
)

// SysAuthUtil 认证工具
var SysAuthUtil = cSysAuthUtil{}

type cSysAuthUtil struct{}

// ************************************ 认证 *****************************************

// LivenessRecognition 身份验证/活体核验
func (c *cSysAuthUtil) LivenessRecognition(ctx context.Context, req *sys_api.LivenessRecognitionReq) (api_v1.BoolRes, error) {
	sys_service.SdkTencent().LivenessRecognition(ctx, req.IdCard, req.Name, req.LivenessType)

	return false, nil
	//return ret, err
}

// DetectAuth 腾讯云-实名核身鉴权
func (c *cSysAuthUtil) DetectAuth(ctx context.Context, req *sys_api.DetectAuthReq) (*sys_model.DetectAuthRes, error) {
	ret, err := sys_service.SdkTencent().DetectAuth(ctx, req.IdCard, req.Name, req.ReturnUrl)

	return ret, err
}

// GetDetectAuthResult 腾讯云-查询实名核身鉴权结果
func (c *cSysAuthUtil) GetDetectAuthResult(ctx context.Context, req *sys_api.GetDetectAuthResultReq) (*sys_model.GetDetectAuthResultRes, error) {
	ret, err := sys_service.SdkTencent().GetDetectAuthResult(ctx, req.BizToken)

	return ret, err
}

// StartAdvFaceAuth  腾讯云-启动H5人脸核身
func (c *cSysAuthUtil) StartAdvFaceAuth(ctx context.Context, req *sys_api.StartAdvFaceAuthReq) (*sys_model.StartAdvFaceAuthRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	orderId := idgen.NextId() // 注意：实际业务中，orderId需要管理，后续还需要通过此查询结果

	ret, err := sys_service.SdkTencent().GetAdvFaceIdAndAuth(ctx, user.Id, orderId, req.IdCard, req.Name, req.CallbackUrl)

	return ret, err
}

// FaceAuthCallback 人脸核身回调地址
func (c *cSysAuthUtil) FaceAuthCallback(ctx context.Context, req *sys_api.FaceAuthCallbackReq) (api_v1.BoolRes, error) {
	res := req.Code == "0"

	// 注意：实际业务中，处理逻辑更加复杂，这里仅仅做了简单的结果处理返回

	return res == true, nil
}

// QueryFaceRecord 腾讯云-人脸核身结果查询
func (c *cSysAuthUtil) QueryFaceRecord(ctx context.Context, req *sys_api.QueryFaceRecordReq) (*sys_model.QueryFaceRecordRes, error) {
	ret, err := sys_service.SdkTencent().QueryFaceRecord(ctx, req.OrderNo)

	return ret, err
}
