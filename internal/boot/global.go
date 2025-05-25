package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

// InitGlobal 初始化公共对象
func InitGlobal() {
	for _, clientConfig := range g.Cfg().MustGet(context.Background(), "service.clientConfig").Array() {
		configItem := sys_model.ClientConfig{
			DefaultRegisterType:      0,
			DefaultUserState:         sys_enum.User.State.Normal.Code(),
			XClientToken:             "",
			AllowLoginUserTypeArr:    garray.NewSortedIntArray(),
			AllowRegister:            false,
			LoginRule:                garray.NewSortedIntArray(),
			RegisterRule:             garray.NewSortedIntArray(),
			EnableRegisterInviteCode: false,
			EnableSendCaptcha:        false,
			ApiPermissionWhitelist:   []string{},
		}

		err := gconv.Struct(clientConfig, &configItem)

		if err != nil {
			g.Log().Error(context.Background(), err)
		}

		sys_consts.Global.ClientConfig = append(sys_consts.Global.ClientConfig, configItem)
	}

	// 加载接口前缀
	sys_consts.Global.ApiPreFix = g.Cfg().MustGet(context.Background(), "service.apiPrefix").String()
	// 邀约码默认有效期天数
	sys_consts.Global.InviteCodeExpireDay = g.Cfg().MustGet(context.Background(), "service.inviteCodeExpireDay").Int()
	// 邀约码最大激活次数上限
	sys_consts.Global.InviteCodeMaxActivateNumber = g.Cfg().MustGet(context.Background(), "service.inviteCodeMaxActivateNumber").Int()

	// 加载ORM表缓存参数
	g.Cfg().MustGet(context.Background(), "service.ormCache", []interface{}{}).Structs(&sys_consts.Global.OrmCacheConf)
}
