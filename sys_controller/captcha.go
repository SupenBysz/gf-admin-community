package sys_controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/enum"
	"github.com/kysion/sms-library/api/sms_api"
	"github.com/kysion/sms-library/sms_controller"
	"github.com/kysion/sms-library/sms_global"
	"github.com/kysion/sms-library/sms_model"
	"time"
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
	// 如下只能满足一个验证码用一次
	//smsType := base_enum.Captcha.Type.New(req.CaptchaType, "")
	//g.DB().GetCache().Set(ctx, smsType.Description()+"_"+req.Mobile, "666666", time.Minute*5)

	// TODO 测试代码 666666
	// 一个验证码支持多种业务场景的，那验证码类型就传入复合类型的进来，如：1登录 8找回密码/重置密码，
	captchaTypes := enum.GetTypes[int, base_enum.CaptchaType](req.CaptchaType, base_enum.Captcha.Type)
	cacheTimeLen := 5 * len(captchaTypes)

	for _, value := range captchaTypes {
		// 存储缓存：key = 业务场景 + 邮箱号   register_1212121@163.com  login_1212121@163.com
		cacheKey := value.Description() + "_" + req.Mobile

		// 方式1：保持验证码到缓存
		err := g.DB().GetCache().Set(ctx, cacheKey, "666666", time.Minute*time.Duration(int64(cacheTimeLen)))
		// 方式2：保持验证码到缓存
		//_, err = g.Redis().Set(ctx, cacheKey, code)
		//if err == nil {
		//_, err = g.Redis().Do(ctx, "EXPIRE", cacheKey, time.Minute*time.Duration(int64(cacheTimeLen)))
		//}
		if err != nil {
			return false, errors.New("验证码缓存失败")
		}
	}

	onSendCaptcha, err := g.Cfg().Get(ctx, "service.onSendCaptcha", false)

	if onSendCaptcha == nil || !onSendCaptcha.Bool() {
		return true, nil
	}

	// TODO 如下是正式代码
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

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "短信发送失败", "Sms")
	}

	return res.SmsSendStatus[0].Code == "OK", err
}

// SendCaptchaByMail 发送邮箱验证码
func (c *cCaptcha) SendCaptchaByMail(ctx context.Context, req *sys_api.SendCaptchaByMailReq) (api_v1.BoolRes, error) {
	// 如下只能满足一个验证码用一次
	//smsType := base_enum.Captcha.Type.New(req.CaptchaType, "")
	//g.DB().GetCache().Set(ctx, smsType.Description()+"_"+req.Mail, "666666", time.Minute*5)
	//
	//return true, nil

	// TODO 测试代码 666666
	// 一个验证码支持多种业务场景的，那验证码类型就传入复合类型的进来，如：1登录 8找回密码/重置密码，
	captchaTypes := enum.GetTypes[int, base_enum.CaptchaType](req.CaptchaType, base_enum.Captcha.Type)
	cacheTimeLen := 5 * len(captchaTypes)

	for _, value := range captchaTypes {
		// 存储缓存：key = 业务场景 + 邮箱号   register_1212121@163.com  login_1212121@163.com
		cacheKey := value.Description() + "_" + req.Mail

		// 方式1：保持验证码到缓存
		err := g.DB().GetCache().Set(ctx, cacheKey, "666666", time.Minute*time.Duration(int64(cacheTimeLen)))
		// 方式2：保持验证码到缓存
		//_, err = g.Redis().Set(ctx, cacheKey, code)
		//if err == nil {
		//_, err = g.Redis().Do(ctx, "EXPIRE", cacheKey, time.Minute*time.Duration(int64(cacheTimeLen)))
		//}
		if err != nil {
			return false, errors.New("验证码缓存失败")
		}
	}
	return true, nil

	// TODO 正式代码
	//ret, err := sys_service.SysMails().SendCaptcha(ctx, req.Mail, req.CaptchaType) // TODO 加上类型
	//return ret == true, err
}
