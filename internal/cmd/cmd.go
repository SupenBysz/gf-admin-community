package cmd

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	sysController "github.com/SupenBysz/gf-admin-community/controller"
	"github.com/SupenBysz/gf-admin-community/internal/consts"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/validator"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gmode"

	"github.com/yitter/idgenerator-go/idgen"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "Kysion Gf Admin Community",
		Usage: "Kysion Gf Admin Community",
		Brief: "Start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			{
				// OpenApi自定义信息
				oai.Info.Title = `API Reference`
				oai.Config.CommonResponse = api_v1.JsonRes{}
				oai.Config.CommonResponseDataField = `Data`
			}

			{
				// 静态目录设置
				uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
				if uploadPath == "" {
					g.Log().Fatal(ctx, "文件上传配置路径不能为空!")
				}
				if !gfile.Exists(uploadPath) {
					_ = gfile.Mkdir(uploadPath)
				}
				// 上传目录添加至静态资源
				s.AddStaticPath("/upload", uploadPath)
			}

			{
				// HOOK, 开发阶段禁止浏览器缓存,方便调试
				if gmode.IsDevelop() {
					s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
						r.Response.Header().Set("Cache-Control", "no-store")
					})
				}
			}

			{
				// 用户默认类型：0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心、-1超级管理员；
				// 独立调用创建用户、查询用户信息等相关接口时强制过滤类型
				consts.Global.DefaultRegisterType = g.Cfg().MustGet(ctx, "service.userDefaultType", 0).Int()
				// 加载不允许登录的用户类型
				consts.Global.NotAllowLoginUserTypeArr = garray.NewSortedIntArrayFrom(g.Cfg().MustGet(ctx, "service.allowLoginUserType", "[-1]").Ints())
				// 去重
				consts.Global.NotAllowLoginUserTypeArr.Unique()
			}

			{
				serviceWorkerId := g.Cfg().MustGet(ctx, "service.idGeneratorWorkerId").Uint16()
				if serviceWorkerId < 1 || serviceWorkerId > 63 {
					g.Log().Fatal(ctx, "service.serviceWorkerId 取值范围只能是 1 ~ 63")
					return nil
				}

				// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
				var options = idgen.NewIdGeneratorOptions(serviceWorkerId)
				options.WorkerIdBitLength = 10 // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
				// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。
				// 保存参数（必须的操作，否则以上设置都不能生效）：
				idgen.SetIdGenerator(options)
			}

			{
				// 注册自定义参数校验规则
				validator.RegisterServicePhone()
			}

			{
				// CASBIN 初始化
				sys_service.Casbin().Enforcer()
			}

			{
				// 初始化路由
				apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
				s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
					// 注册中间件
					group.Middleware(
						sys_service.Middleware().Casbin,
						sys_service.Middleware().CTX,
						sys_service.Middleware().ResponseHandler,
					)

					// 匿名路由绑定
					group.Group("/", func(group *ghttp.RouterGroup) {
						// 鉴权：登录，注册，找回密码等
						group.Group("/sys_auth", func(group *ghttp.RouterGroup) { group.Bind(sysController.Auth) })
						// 图型验证码、短信验证码、地区
						group.Group("/common", func(group *ghttp.RouterGroup) {
							group.Bind(
								// 图型验证码
								sysController.Captcha,
								// 短信验证码
								sysController.SysSms,
								// 地区
								sysController.SysArea,
							)
						})
					})

					// 权限路由绑定
					group.Group("/", func(group *ghttp.RouterGroup) {
						// 注册中间件
						group.Middleware(
							sys_service.Middleware().Auth,
						)

						// 文件上传
						group.Group("/common/sys_file", func(group *ghttp.RouterGroup) { group.Bind(sysController.SysFile) })
						// 系统配置
						group.Group("/system/config", func(group *ghttp.RouterGroup) { group.Bind(sysController.SysConfig) })
						// 用户
						group.Group("/user", func(group *ghttp.RouterGroup) { group.Bind(sysController.SysUser) })
						// 角色
						group.Group("/role", func(group *ghttp.RouterGroup) { group.Bind(sysController.SysRole) })
						// 权限
						group.Group("/permission", func(group *ghttp.RouterGroup) { group.Bind(sysController.SysPermission) })
						// 组织架构
						group.Group("/organization", func(group *ghttp.RouterGroup) { group.Bind(sysController.SysOrganization) })
					})
				})
			}

			s.Run()
			return nil
		},
	}
)
