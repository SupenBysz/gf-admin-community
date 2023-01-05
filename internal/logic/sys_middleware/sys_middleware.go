package sys_middleware

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/response"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

func init() {
	sys_service.RegisterMiddleware(New())
}

// New MiddlewareMiddleware 权限控制
func New() *sMiddleware {
	return &sMiddleware{}
}

// Auth 通讯鉴权
func (s *sMiddleware) Auth(r *ghttp.Request) {
	sys_service.Jwt().Middleware(r)
}

// CTX 自定义上下文对象
func (s *sMiddleware) CTX(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customSessionCtx := &sys_model.SessionContext{
		JwtClaimsUser:     &sys_model.JwtCustomClaims{},
		SessionErrorQueue: garray.NewArray(),
		Ipv4:              r.RemoteAddr,
	}

	sys_service.SysSession().Init(r, customSessionCtx)

	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customSessionCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// CORS 允许接口跨域请求
func (s *sMiddleware) CORS(r *ghttp.Request) {
	// corsOptions := r.Response.DefaultCORSOptions()
	// corsOptions.AllowOrigin = "*"
	// corsOptions.AllowCredentials = "true"
	// corsOptions.AllowHeaders = "Content-Type,Access-Token"
	// corsOptions.AllowDomain = []string{"127.0.0.1:3000", "0.0.0.0:3000"}
	// corsOptions.ExposeHeaders = "*"
	// corsOptions.AllowMethods = "POST,GET,OPTIONS,DELETE"
	// r.Response.CORS(corsOptions)
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// ResponseHandler 响应函数
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)

	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}

		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, code.Code(), "", res)
	}
}

func (s *sMiddleware) Casbin(r *ghttp.Request) {
	sys_service.Casbin().Middleware(r)
}
