package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// Captcha 图形验证码
var Captcha = cCaptcha{}

type cCaptcha struct{}

func (a *cCaptcha) Index(ctx context.Context, _ *sysapi.CaptchaIndexReq) (res *sysapi.CaptchaIndexRes, err error) {
	err = sys_service.Captcha().MakeCaptcha(ctx)
	return
}
