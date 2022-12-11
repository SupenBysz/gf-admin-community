package sys_captcha

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/mojocn/base64Captcha"
)

type sCaptcha struct{}

var (
	captchaStore  = base64Captcha.DefaultMemStore
	captchaDriver = newDriver()
)

func init() {
	sys_service.RegisterCaptcha(New())
}

// New Captcha 验证码管理服务
func New() *sCaptcha {
	return &sCaptcha{}
}

func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:     44,
		Width:      126,
		NoiseCount: 9,
		// ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		// ShowLineOptions: base64Captcha.OptionShowSineLine,
		Length: 4,
		Source: "1234567890",
		Fonts:  []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

// MakeCaptcha 创建验证码，直接输出验证码图片内容到HTTP Response.
func (s *sCaptcha) MakeCaptcha(ctx context.Context) error {
	var (
		request = g.RequestFromCtx(ctx)
		captcha = base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	)

	_, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	item, _ := captcha.Driver.DrawCaptcha(content)
	request.Response.Header().Add("captchaId", gconv.String(idgen.NextId()))
	captcha.Store.Set(gmd5.MustEncryptString(answer), answer)
	_, err := item.WriteTo(request.Response.Writer)
	return err
}

// VerifyAndClear 校验验证码，并清空缓存的验证码信息
func (s *sCaptcha) VerifyAndClear(_ *ghttp.Request, value string) bool {
	// 开发模式不校验验证码
	if gmode.IsDevelop() {
		return true
	}
	return captchaStore.Verify(gmd5.MustEncryptString(value), value, true)
}
