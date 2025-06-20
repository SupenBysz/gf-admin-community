package sys_rules

import (
	"context"
	"strings"

	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/i18n"
	"github.com/SupenBysz/gf-admin-community/utility/invite_id"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/base_verify"
)

// CheckEnableSendCaptchaRule 检查是否允许发送验证码
func CheckEnableSendCaptchaRule(ctx context.Context) bool {
	clientConfig, _ := sys_consts.Global.GetClientConfig(ctx)

	if clientConfig == nil {
		return false
	}

	return clientConfig.EnableSendCaptcha
}

// CheckLoginRule 校验用户是否可以通过此方式登陆
func CheckLoginRule(ctx context.Context, loginIdentifier string) bool {
	clientConfig, _ := sys_consts.Global.GetClientConfig(ctx)

	if clientConfig == nil {
		return false
	}

	if base_verify.IsPhone(loginIdentifier) {
		return clientConfig.LoginRule.Contains(2)

	} else if base_verify.IsEmail(loginIdentifier) {
		return clientConfig.LoginRule.Contains(4)
	} else {
		return clientConfig.LoginRule.Contains(1)
	}
}

// CheckRegisterRule 校验用户是否可以通过此方式注册
func CheckRegisterRule(ctx context.Context, registerIdentifier string) bool {
	clientConfig, _ := sys_consts.Global.GetClientConfig(ctx)

	if clientConfig == nil {
		return false
	}

	if base_verify.IsPhone(registerIdentifier) {
		return clientConfig.RegisterRule.Contains(2)

	} else if base_verify.IsEmail(registerIdentifier) {
		return clientConfig.RegisterRule.Contains(4)
	} else {
		return clientConfig.RegisterRule.Contains(1)
	}
}

// CheckApiPermissionWhiteList 检查API权限白名单
func CheckApiPermissionWhiteList(ctx context.Context) bool {
	clientConfig, _ := sys_consts.Global.GetClientConfig(ctx)

	if clientConfig == nil {
		return false
	}

	apiRequestPath := ghttp.RequestFromCtx(ctx).URL.Path

	apiRequestPath = strings.TrimPrefix(apiRequestPath, sys_consts.Global.ApiPreFix)

	return base_funs.HasInSlice(clientConfig.ApiPermissionWhitelist, func(pattern string) bool {
		if pattern == "*" {
			return true
		}
		if strings.HasPrefix(pattern, "*") && strings.HasSuffix(pattern, "*") {
			middle := strings.Trim(pattern, "*")
			return strings.Contains(apiRequestPath, middle)
		}

		if strings.HasPrefix(pattern, "*") {
			suffix := strings.TrimPrefix(pattern, "*")
			return strings.HasSuffix(apiRequestPath, suffix)
		}

		if strings.HasSuffix(pattern, "*") {
			prefix := strings.TrimSuffix(pattern, "*")
			return strings.HasPrefix(apiRequestPath, prefix)
		}
		return apiRequestPath == pattern
	})
}

func CheckInviteCode(ctx context.Context, code string) (res *sys_model.InviteRes, err error) {
	clientConfig, _ := sys_consts.Global.GetClientConfig(ctx)

	// 判断是否填写邀约码
	if clientConfig != nil && clientConfig.EnableRegisterInviteCode && code == "" {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, i18n.T(ctx, "error_invite_code_required"))
	}

	inviteId := int64(0)

	// 只要填写了必需进行校验
	if code != "" {
		inviteId = invite_id.CodeToInviteId(code)
		if inviteId <= 0 {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, i18n.T(ctx, "error_invite_code_incorrect"))
		}
	}

	result := &sys_model.InviteRes{
		Code: code,
	}

	result.Id = inviteId

	return result, nil
}
