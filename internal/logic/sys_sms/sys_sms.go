package sys_sms

import (
	"context"
	"fmt"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/kysion/base-library/base_model/base_enum"
)

type sSysSms struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSysSms(New())
}

func New() sys_service.ISysSms {
	return &sSysSms{
		cachePrefix: sys_dao.SysSmsLogs.Table() + "_",
	}
}

// Verify 校验验证码
func (s *sSysSms) Verify(ctx context.Context, mobile string, captcha string, typeIdentifier ...base_enum.CaptchaType) (bool, error) {
	if mobile == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_sms_mobile_empty", "Sms")
	}
	if captcha == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_sms_captcha_empty", "Sms")
	}

	if  gmode.IsDevelop() {
		return true, nil
	}

	key := ""
	if len(typeIdentifier) > 0 {
		key = typeIdentifier[0].Description() + "_" + mobile
	} else {
		key = mobile
	}

	code, err := g.DB().GetCache().Get(ctx, key)

	fmt.Println("验证码：", code.String())

	if err != nil || code.String() != captcha {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_sms_captcha_incorrect", "Sms")
	}

	// 成功、清除该缓存
	_, _ = g.DB().GetCache().Remove(ctx, key)

	// 此验证码类型是复用类型
	//if (typeIdentifier[0].Code() & base_enum.Captcha.Type.ForgotUserNameAndPassword.Code()) == base_enum.Captcha.Type.ForgotUserNameAndPassword.Code() {
	//	cacheKey := base_enum.Captcha.Type.SetPassword.Description() + "_" + mobile
	//
	//	// 重新保持验证码到缓存
	//	_, err := g.Redis().Set(ctx, cacheKey, code)
	//	if err != nil {
	//		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "再次设置忘记密码验证码至缓存失败", "Sms")
	//	}
	//	// 设置验证码缓存时间
	//	_, _ = g.Redis().Do(ctx, "EXPIRE", cacheKey, time.Minute*5)
	//}

	return true, nil
}
