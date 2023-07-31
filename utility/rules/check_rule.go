package rules

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/invite_id"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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

func CheckInviteCode(ctx context.Context, code string) (res *sys_model.InviteRes, err error) {
	// 判断是否填写邀约码
	if sys_consts.Global.RegisterIsNeedInviteCode && code == "" {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "系统要求注册必需填写邀约码！")
	}
	// 只要填写了必需进行校验
	if code != "" {
		id := invite_id.CodeToInviteId(code)
		res, err = sys_service.SysInvite().GetInviteById(ctx, id)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "填写的邀约码错误，请检查！")
		}
	}

	return res, err
}
