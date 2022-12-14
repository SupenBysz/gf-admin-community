package cmd

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	sysController "github.com/SupenBysz/gf-admin-community/sys_controller"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gmode"

	_ "github.com/SupenBysz/gf-admin-community/internal/boot"
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
