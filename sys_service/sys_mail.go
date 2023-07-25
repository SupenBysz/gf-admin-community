// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type (
	ISysMails interface {
		// SendCaptcha 发送邮件验证码
		SendCaptcha(ctx context.Context, mailTo string, typeIdentifier int) (res bool, err error)
		// Verify 校验验证码
		Verify(ctx context.Context, email string, captcha string, typeIdentifier ...sys_enum.CaptchaType) (bool, error)
	}
)

var (
	localSysMails ISysMails
)

func SysMails() ISysMails {
	if localSysMails == nil {
		panic("implement not found for interface ISysMails, forgot register?")
	}
	return localSysMails
}

func RegisterSysMails(i ISysMails) {
	localSysMails = i
}
