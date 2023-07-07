package sys_controller

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/sms-library/api/sms_api"
	"github.com/kysion/sms-library/sms_controller"
	"github.com/kysion/sms-library/sms_global"
	"github.com/kysion/sms-library/sms_model"
)

// Captcha 验证码
var Captcha = cCaptcha{}

type cCaptcha struct{}

// Index 获取默认的图形验证码
func (c *cCaptcha) Index(ctx context.Context, _ *sys_api.CaptchaIndexReq) (res *sys_api.CaptchaIndexRes, err error) {
	err = sys_service.Captcha().MakeCaptcha(ctx)
	return
}

// SendCaptchaBySms 发送短信验证码
func (c *cCaptcha) SendCaptchaBySms(ctx context.Context, req *sys_api.SendCaptchaBySmsReq) (api_v1.BoolRes, error) {
	//smsType := sys_enum.Sms.CaptchaType.New(req.CaptchaType, "")
	//g.DB().GetCache().Set(ctx, smsType.Description()+"_"+req.Mobile, "666666", time.Minute*5)

	//return true, nil

	sendReq := sms_api.SendSmsReq{
		CaptchaType: req.CaptchaType,
		SmsSendMessageReq: sms_model.SmsSendMessageReq{
			Phones:      []string{req.Mobile},
			CaptchaType: req.CaptchaType,
		},
	}

	fmt.Println(sendReq)
	modules := sms_global.Global.Modules
	res, err := sms_controller.Sms(modules).SendSms(ctx, &sendReq)

	return res.SmsSendStatus[0].Code == "OK", err
}

// SendCaptchaByMail 发送邮箱验证码
func (c *cCaptcha) SendCaptchaByMail(ctx context.Context, req *sys_api.SendCaptchaByMailReq) (api_v1.BoolRes, error) {

	ret, err := sys_service.SysMails().SendCaptcha(ctx, req.Mail, req.CaptchaType) // TODO 加上类型

	return ret == true, err
}
