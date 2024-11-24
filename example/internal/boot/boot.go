package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_controller"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/kysion/base-library/utility/en_crypto"
	"github.com/kysion/oss-library/oss_consts"
	"github.com/kysion/sms-library/sms_consts"
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
				// Permission 初始化
				_ = sys_service.SysPermission().ImportPermissionTree(ctx, sys_consts.Global.PermissionTree, nil)

				// 导入oss + sms权限树
				_ = sys_service.SysPermission().ImportPermissionTree(ctx, oss_consts.PermissionTree, nil)
				_ = sys_service.SysPermission().ImportPermissionTree(ctx, sms_consts.PermissionTree, nil)
			}

			// 业务端密码加密规则重写示例
			{
				sys_consts.Global.CryptoPasswordFunc = func(ctx context.Context, passwordStr string, user ...sys_entity.SysUser) (pwdEncode string) {
					// TODO 以下加密规则可替换
					slat := "kysion.com"
					if len(user) > 0 {
						slat = gconv.String(user[0].Id)
					}

					pwdHash, _ := en_crypto.PwdHash(passwordStr, slat)

					return pwdHash
				}
			}

			{
				// 初始化路由
				apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
				s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
					// 注册中间件
					group.Middleware(
						// sys_service.Middleware().Casbin,
						sys_service.Middleware().CTX,
						sys_service.Middleware().CORS,
						sys_service.Middleware().ResponseHandler,
					)

					// 匿名路由绑定
					group.Group("/", func(group *ghttp.RouterGroup) {
						// 鉴权：登录，注册，找回密码等
						group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.Auth) })
						// 图型验证码、短信验证码、地区
						group.Group("/common", func(group *ghttp.RouterGroup) {
							group.Bind(
								// 验证码（图形、短信、邮箱）
								sys_controller.Captcha,
								// 地区
								sys_controller.SysArea,
								// 公共：获取图片...
								sys_controller.Common,
								// 匿名访问文件
								sys_controller.SysFileAllowAnonymous,
							)
						})

					})

					// 权限路由绑定
					group.Group("/", func(group *ghttp.RouterGroup) {
						// 注册中间件
						group.Middleware(
							sys_service.Middleware().Auth,
							// sys_service.Middleware().CheckPermission,
						)

						// 用户鉴权 - 需要登陆
						group.Group("/auth", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.UserAuth) })
						// 文件上传
						group.Group("/common/file", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysFile) })
						// 应用配置
						group.Group("/system/config", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysConfig) })
						// 工具
						group.Group("/utils", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysUtil) })
						// 认证工具
						group.Group("/utils", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysAuthUtil) })
						// 系统配置
						group.Group("/system/settings", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysSettings) })
						// 系统前端配置
						group.Group("/system/frontSettings", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysFrontSettings) })
						// 用户
						group.Group("/user", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysUser) })
						// 角色
						group.Group("/role", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysRole) })
						// 权限
						group.Group("/permission", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysPermission) })
						// 组织架构
						group.Group("/organization", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysOrganization) })
						// 我的
						group.Group("/my", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysMy) })
						// 个人资质
						group.Group("/person_license", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysLicense) })
						// 个人资质审核
						group.Group("/person_audit", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysAudit) })
						// 菜单
						group.Group("/menu", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysMenu) })
						// 邀约
						group.Group("/invite", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysInvite) })
						// 行业类别
						group.Group("/industry", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysIndustry) })
						// 消息
						group.Group("/message", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysMessage) })
						// 公告
						group.Group("/announcement", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysAnnouncement) })
						// 会员等级
						group.Group("/memberLevel", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysMemberLevel) })
						// 类目管理
						group.Group("/category", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysCategory) })
						// UEditor 编辑器
						group.Group("/ueditor", func(group *ghttp.RouterGroup) { group.Bind(sys_controller.SysUEditor) })
					})
				})
			}

			s.Run()
			return nil
		},
	}
)
