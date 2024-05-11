package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysUtil 工具
var SysUtil = cSysUtil{}

type cSysUtil struct{}

// GetNerChEcom 中文分词
func (c *cSysUtil) GetNerChEcom(ctx context.Context, req *sys_api.GetNerChEcomReq) (sys_model.AliyunNlpDataRes, error) {
	ret, err := sys_service.SdkAliyun().GetWsCustomizedChGeneral(ctx, req.Text)
	return ret, err
}

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
