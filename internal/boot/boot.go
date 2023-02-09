package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/downloader"
	"github.com/SupenBysz/gf-admin-community/utility/env"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/SupenBysz/gf-admin-community/utility/validator"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
)

func init() {
	env.LoadEnv()

	Ip2region()
	InitCustomRules()
	InitGlobal()
	InitRedisCache()
	InitIdGenerator()
	InitLogLevelToDatabase()
	InitPermission()
}

// InitCustomRules 注册自定义参数校验规则
func InitCustomRules() {
	// 注册电话验证规则
	validator.RegisterServicePhone()
}

func Ip2region() {
	ip2regionPath := g.Cfg().MustGet(
		context.Background(),
		"service.ip2region.xdbPath",
		"./resources/assets/static/ip2resion.xdb",
	).String()

	if ip2regionPath == "" || gfile.Size(ip2regionPath) <= 0 {
		log.Println("开始下载IP信息库资源")
		d := downloader.NewDownloader(
			"https://ghproxy.com/https://github.com/lionsoul2014/ip2region/raw/master/data/ip2region.xdb",
			gfile.Basename(ip2regionPath),
			gfile.Abs(gfile.Dir(ip2regionPath)),
			10,
		)
		if err := d.Download(); err != nil {
			panic("ip2region 获取失败")
		}
	}
	if gfile.Size(ip2regionPath) <= 0 {
		panic("ip2region 校验失败")
	}

	cBuff, err := xdb.LoadContentFromFile(ip2regionPath)
	if err != nil {
		panic("ip2region 初始化失败")
	}
	sys_consts.Global.Searcher, _ = xdb.NewWithBuffer(cBuff)
}

// InitGlobal 初始化公共对象
func InitGlobal() {
	// 默认类型：0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心、-1超级管理员；
	// 独立调用创建用户、查询用户信息等相关接口时强制过滤类型
	sys_consts.Global.UserDefaultType = sys_enum.User.Type.New(g.Cfg().MustGet(context.Background(), "service.userDefaultType", 0).Int(), "")
	// 新增用户默认状态：0未激活，1正常，-1封号，-2异常，-3已注销
	sys_consts.Global.UserDefaultState = sys_enum.User.State.New(g.Cfg().MustGet(context.Background(), "service.userDefaultState", 0).Int(), "")
	// 加载不允许登录的用户类型，并去重
	sys_consts.Global.NotAllowLoginUserTypeArr = garray.NewSortedIntArrayFrom(g.Cfg().MustGet(context.Background(), "service.notAllowLoginUserType", "[-1]").Ints()).SetUnique(true)
	// 加载接口前缀
	sys_consts.Global.ApiPreFix = g.Cfg().MustGet(context.Background(), "service.apiPrefix").String()
	// 加载ORM表缓存参数
	g.Cfg().MustGet(context.Background(), "service.ormCache", []interface{}{}).Structs(&sys_consts.Global.OrmCacheConf)
}

// InitIdGenerator 初始化ID生成器
func InitIdGenerator() {
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

// InitLogLevelToDatabase 加载日志写数据库的配置
func InitLogLevelToDatabase() {
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

func InitPermission() []*permission.SysPermissionTree {
	sys_consts.Global.PermissionTree = []*permission.SysPermissionTree{
		// 用户管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5947106208184773,
				Name:       "用户管理",
				Identifier: "User",
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看用户，查看某个用户登录账户
				sys_enum.User.PermissionType.ViewDetail,
				// 查看更多详情，含完整手机号
				sys_enum.User.PermissionType.ViewMoreDetail,
				// 用户列表，查看所有用户
				sys_enum.User.PermissionType.List,
				// 重置密码，重置某个用户的登录密码
				sys_enum.User.PermissionType.ResetPassword,
				// 设置状态，设置某个用户的状态
				sys_enum.User.PermissionType.SetState,
				// 修改密码，修改自己的登录密码
				sys_enum.User.PermissionType.ChangePassword,
				// "创建用户，创建一个新用户"
				// sys_enum.User.PermissionType.Create,
				// 修改用户名称，修改用户登录账户名称信息
				sys_enum.User.PermissionType.SetUsername,
				// 设置用户角色，设置某一个用户的角色
				sys_enum.User.PermissionType.SetUserRole,
				// 设置用户权限，设置某一个用户的权限
				sys_enum.User.PermissionType.SetPermission,
			},
		},
		// 组织架构权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948649344204869,
				Name:       "组织架构",
				Identifier: "Organization",
				Type:       1,
				IsShow:     0,
			},
			Children: []*permission.SysPermissionTree{
				// 查看，查看某个组织架构
				sys_enum.Organization.PermissionType.ViewDetail,
				// 查看列表，查看所有组织架构列表
				sys_enum.Organization.PermissionType.List,
				// 更新，更新某个组织架构
				sys_enum.Organization.PermissionType.Update,
				// 删除，删除某个组织架构
				sys_enum.Organization.PermissionType.Delete,
				// 创建，创建组织架构
				sys_enum.Organization.PermissionType.Create,
			},
		},
		// 角色管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5948684761759818,
				Name:       "角色管理",
				Identifier: "Role",
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看角色，查看某个角色
				sys_enum.Role.PermissionType.ViewDetail,
				// 角色列表，查看所有角色
				sys_enum.Role.PermissionType.List,
				// 更新角色信息，更新某个角色信息
				sys_enum.Role.PermissionType.Update,
				// 删除角色，删除某个角色
				sys_enum.Role.PermissionType.Delete,
				// 创建角色，创建一个新角色
				sys_enum.Role.PermissionType.Create,
				// 设置角色成员，增加或移除角色成员
				sys_enum.Role.PermissionType.SetMember,
				// 设置角色权限，设置某个角色的权限
				sys_enum.Role.PermissionType.SetPermission,
			},
		},
		// 权限管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5950408166668741,
				Name:       "权限管理",
				Identifier: "Permission",
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看权限，查看某个权限
				sys_enum.Permissions.PermissionType.ViewDetail,
				// 权限列表，查看所有权限
				sys_enum.Permissions.PermissionType.List,
				// 更新权限，更新某个权限
				sys_enum.Permissions.PermissionType.Update,
				// 删除权限，删除某个权限
				sys_enum.Permissions.PermissionType.Delete,
				// 创建权限，创建权限
				sys_enum.Permissions.PermissionType.Create,
			},
		},
	}

	// 添加资质和审核权限树
	licensePermission := initAuditAndLicensePermission()
	sys_consts.Global.PermissionTree = append(sys_consts.Global.PermissionTree, licensePermission...)

	return sys_consts.Global.PermissionTree
}

func initAuditAndLicensePermission() []*permission.SysPermissionTree {
	result := []*permission.SysPermissionTree{

		// 资质权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5953153121845333,
				Name:       "资质",
				Identifier: "License",
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看资质信息，查看某条资质信息
				sys_enum.License.PermissionType.ViewDetail,
				// 资质列表，查看所有资质信息
				sys_enum.License.PermissionType.List,
				// 更新资质审核信息，更新某条资质审核信息
				sys_enum.License.PermissionType.Update,
				// 创建资质，创建资质信息
				sys_enum.License.PermissionType.Create,
				// 设置资质状态，设置某资质认证状态
				sys_enum.License.PermissionType.SetState,
			},
		},
		// 审核管理权限树
		{
			SysPermission: &sys_entity.SysPermission{
				Id:         5953151699124300,
				Name:       "审核管理",
				Identifier: "Audit",
				Type:       1,
				IsShow:     1,
			},
			Children: []*permission.SysPermissionTree{
				// 查看审核信息，查看某条资质审核信息
				sys_enum.Audit.PermissionType.ViewDetail,
				// 资质审核列表，查看所有资质审核
				sys_enum.Audit.PermissionType.List,
				// 更新资质审核信息，更新某条资质审核信息
				sys_enum.Audit.PermissionType.Update,
			},
		},
	}
	return result
}

func InitRedisCache() {
	// 获取配置文件addr对象
	addr, _ := g.Cfg().Get(context.Background(), "redis.default.address")

	conf, _ := gredis.GetConfig("default")

	// 设置服务端口和ip
	conf.Address = addr.String()
	// 不同的表分配不同的redis数据库
	conf.Db = 1

	// 没配置redis ip+端口,配置信息也为空，那么使用内存缓存
	if addr.String() == "" || conf.Address == "" {
		g.DB().GetCache().SetAdapter(gcache.New())
		return
	}
	// 根据redis配置创建Redis
	redis, _ := gredis.New(conf)
	// 全局设置Redis适配器
	g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
}
