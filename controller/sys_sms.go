package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
)

// SysSms 短信
var SysSms = cSysSms{}

type cSysSms struct{}

func (c *cSysSms) SendCaptchaBySms(_ context.Context, _ *sys_api.SendCaptchaBySmsReq) (api_v1.BoolRes, error) {
	// 暂定总是成功，后续短信模块完成后这里再继续完善
	return true, nil
}
