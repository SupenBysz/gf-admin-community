// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICaptcha interface {
		// MakeCaptcha 创建验证码，直接输出验证码图片内容到HTTP Response.
		MakeCaptcha(ctx context.Context) error
		// VerifyAndClear 校验验证码，并清空缓存的验证码信息
		VerifyAndClear(_ *ghttp.Request, value string) bool
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
