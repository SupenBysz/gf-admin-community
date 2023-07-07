package rules

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/rule"
)

// CheckLoginRule 校验用户是否可以通过此方式登陆
func CheckLoginRule(ctx context.Context, loginIdentifier string) bool {
	arr := g.Cfg().MustGet(ctx, "service.loginRule").Array()

	if rule.IsPhone(loginIdentifier) {
		for _, v := range arr {
			if gconv.Int(v) == 2 {
				return true
			}
		}

	} else if rule.IsEmail(loginIdentifier) {
		for _, v := range arr {
			if gconv.Int(v) == 4 {
				return true
			}
		}
	} else {
		for _, v := range arr {
			if gconv.Int(v) == 1 {
				return true
			}
		}
	}

	return false
}

// CheckRegisterRule 校验用户是否可以通过此方式注册
func CheckRegisterRule(ctx context.Context, registerIdentifier string) bool {
	arr := g.Cfg().MustGet(ctx, "service.registerRule").Array()

	if rule.IsPhone(registerIdentifier) {
		for _, v := range arr {
			if gconv.Int(v) == 2 {
				return true
			}
		}

	} else if rule.IsEmail(registerIdentifier) {
		for _, v := range arr {
			if gconv.Int(v) == 4 {
				return true
			}
		}
	} else {
		for _, v := range arr {
			if gconv.Int(v) == 1 {
				return true
			}
		}
	}

	return false
}
