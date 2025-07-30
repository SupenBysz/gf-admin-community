package sys_model

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/kysion/base-library/utility/base_funs"
)

type ClientConfig struct {
	DefaultRegisterType      int                    `json:"defaultRegisterType"     dc:"默认注册类型"`
	DefaultUserState         int                    `json:"defaultUserState"     dc:"默认用户状态"`
	XClientToken             string                 `json:"identifier"     dc:"客户端token"`
	AllowLoginUserTypeArr    *garray.SortedIntArray `json:"allowLoginUserType"     dc:"允许登录的用户类型"`
	AllowRegister            bool                   `json:"allowRegister"     dc:"是否允许注册"`
	LoginRule                *garray.SortedIntArray `json:"loginRule"     dc:"登录规则"`
	EnableInviteRegister     bool                   `json:"enableInviteRegister" dc:"是否开启邀请注册功能"`
	RegisterRule             *garray.SortedIntArray `json:"registerRule"     dc:"注册规则"`
	EnableRegisterInviteCode bool                   `json:"enableRegisterInviteCode"     dc:"注册是否需要邀请码"`
	EnableSendCaptcha        bool                   `json:"enableSendCaptcha"     dc:"开启验证码发送"`
	ApiPermissionWhitelist   []string               `json:"apiPermissionWhitelist" dc:"接口权限白名单，填写API路径（不含前缀接口前缀[apiPrefix中定义的迁移]），支持*为通配符，* 表示所有接口都不需要权限验证，*xxx 表示所有以xxx结尾的接口都不需要权限验证, xxx* 表示所有以xxx开头的接口都不需要权限验证"`
	// 用户类型白名单，用户类型在 model/enum/user/type.go 中定义白名单中的用户类型将允许访问用户类型对应的额员工信息和公司信息
	UserTypeWhitelist []int `json:"user_type_whitelist" dc:"用户类型白名单"`
}

func (c ClientConfig) HasUserTypeWhitelist(userType int) bool {
	return base_funs.Contains(c.UserTypeWhitelist, userType)
}
