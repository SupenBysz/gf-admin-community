package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/utility/env"
	"github.com/SupenBysz/gf-admin-community/utility/validator"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	env.LoadEnv()

	initCustomRules()
	initGlobal()
	initIdGenerator()
	initLogLevelToDatabase()
}

// initCustomRules 注册自定义参数校验规则
func initCustomRules() {
	// 注册电话验证规则
	validator.RegisterServicePhone()
}

// initGlobal 初始化公共对象
func initGlobal() {
	// 用户默认类型：0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心、-1超级管理员；
	// 独立调用创建用户、查询用户信息等相关接口时强制过滤类型
	sys_consts.Global.DefaultRegisterType = g.Cfg().MustGet(context.Background(), "service.userDefaultType", 0).Int()
	// 加载不允许登录的用户类型，并去重
	sys_consts.Global.NotAllowLoginUserTypeArr = garray.NewSortedIntArrayFrom(g.Cfg().MustGet(context.Background(), "service.notAllowLoginUserType", "[-1]").Ints()).SetUnique(true)
}

// initIdGenerator 初始化ID生成器
func initIdGenerator() {
	serviceWorkerId := g.Cfg().MustGet(context.Background(), "service.idGeneratorWorkerId").Uint16()
	if serviceWorkerId < 1 || serviceWorkerId > 63 {
		g.Log().Fatal(context.Background(), "service.serviceWorkerId 取值范围只能是 1 ~ 63")
		return
	}

	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(serviceWorkerId)
	options.WorkerIdBitLength = 10 // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
	// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。
	// 保存参数（必须的操作，否则以上设置都不能生效）：
	idgen.SetIdGenerator(options)
}

// logLevelToDatabase加载日志写数据库的配置
func initLogLevelToDatabase() {
	LogLevelToDatabaseArr := garray.NewSortedStrArrayFrom(g.Cfg().MustGet(context.Background(), "service.logLevelToDatabase", "[\"ALL\"]").Strings()).SetUnique(true)

	LogLevelToDatabaseArr.Iterator(func(_ int, value string) bool {
		switch gstr.ToUpper(value) {
		case "ALL":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_ALL)
			return false
		case "ERROR":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_ERRO)
		case "INFO":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_INFO)
		case "WARN":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_WARN)
		}
		return true
	})
}
