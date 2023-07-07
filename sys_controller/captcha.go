package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// Captcha 图形验证码
var Captcha = cCaptcha{}

type cCaptcha struct{}

func (a *cCaptcha) Index(ctx context.Context, _ *sys_api.CaptchaIndexReq) (res *sys_api.CaptchaIndexRes, err error) {
	err = sys_service.Captcha().MakeCaptcha(ctx)
	return
}
}

// SendCaptchaByMail 发送邮箱验证码
func (c *cCaptcha) SendCaptchaByMail(ctx context.Context, req *sys_api.SendCaptchaByMailReq) (api_v1.BoolRes, error) {

	ret, err := sys_service.SysMails().SendCaptcha(ctx, req.Mail, req.CaptchaType) // TODO 加上类型

	return ret == true, err
}
