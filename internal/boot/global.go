package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
)

// InitGlobal 初始化公共对象
func InitGlobal() {
	// 默认类型：0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心、-1超级管理员；
	// 独立调用创建用户、查询用户信息等相关接口时强制过滤类型
	userDefaultType := g.Cfg().MustGet(context.Background(), "service.userDefaultType", 0)
	sys_consts.Global.UserDefaultType = sys_enum.User.Type.New(userDefaultType.Int(), "")
	// 新增用户默认状态：0未激活，1正常，-1封号，-2异常，-3已注销
	sys_consts.Global.UserDefaultState = sys_enum.User.State.New(g.Cfg().MustGet(context.Background(), "service.userDefaultState", 0).Int(), "")
	// 加载不允许登录的用户类型，并去重
	sys_consts.Global.NotAllowLoginUserTypeArr = garray.NewSortedIntArrayFrom(g.Cfg().MustGet(context.Background(), "service.notAllowLoginUserType", "[-1]").Ints()).SetUnique(true)
	// 加载允许登录的用户类型，并去重 (如果NotAllowLoginUserTypeArr包含allowLoginUserType中的用户类型，那么前者优先级高于后者, 默认值为UserDefaultType)
	sys_consts.Global.AllowLoginUserTypeArr = garray.NewSortedIntArrayFrom(g.Cfg().MustGet(context.Background(), "service.allowLoginUserType", "["+userDefaultType.String()+"]").Ints()).SetUnique(true)
	// 加载接口前缀
	sys_consts.Global.ApiPreFix = g.Cfg().MustGet(context.Background(), "service.apiPrefix").String()
	// 注册是否需要邀约码
	sys_consts.Global.RegisterIsNeedInviteCode = g.Cfg().MustGet(context.Background(), "service.registerIsNeedInviteCode").Bool()
	// 加载ORM表缓存参数
	g.Cfg().MustGet(context.Background(), "service.ormCache", []interface{}{}).Structs(&sys_consts.Global.OrmCacheConf)
}
