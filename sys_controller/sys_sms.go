package sys_controller

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/kysion/sms-library/api/sms_api"
	"github.com/kysion/sms-library/sms_controller"
	"github.com/kysion/sms-library/sms_global"
	"github.com/kysion/sms-library/sms_model"
)

// SysSms 短信
var SysSms = cSysSms{}

type cSysSms struct{}

func (c *cSysSms) SendCaptchaBySms(ctx context.Context, req *sys_api.SendCaptchaBySmsReq) (api_v1.BoolRes, error) {
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
