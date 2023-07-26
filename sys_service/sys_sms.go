// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/kysion/base-library/base_model/base_enum"
)

type (
	ISysSms interface {
		// Verify 校验验证码
		Verify(ctx context.Context, mobile string, captcha string, typeIdentifier ...base_enum.CaptchaType) (bool, error)
	}
)

var (
	localSysSms ISysSms
)

func SysSms() ISysSms {
	if localSysSms == nil {
		panic("implement not found for interface ISysSms, forgot register?")
	}
	return localSysSms
}

func RegisterSysSms(i ISysSms) {
	localSysSms = i
}
