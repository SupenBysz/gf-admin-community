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
