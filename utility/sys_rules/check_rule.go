package sys_rules

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/invite_id"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/kysion/base-library/utility/base_verify"
)

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

	if base_verify.IsPhone(registerIdentifier) {
		return clientConfig.RegisterRule.Contains(2)

	} else if base_verify.IsEmail(registerIdentifier) {
		return clientConfig.RegisterRule.Contains(4)
	} else {
		return clientConfig.RegisterRule.Contains(1)
	}
}

func CheckInviteCode(ctx context.Context, code string) (res *sys_model.InviteRes, err error) {
	clientConfig, _ := sys_consts.Global.GetClientConfig(ctx)

	// 判断是否填写邀约码
	if clientConfig != nil && clientConfig.EnableRegisterInviteCode == true && code == "" {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "邀约码不能为空")
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
