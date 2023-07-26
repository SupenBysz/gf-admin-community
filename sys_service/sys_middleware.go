// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Auth 通讯鉴权
		Auth(r *ghttp.Request)
		// CTX 自定义上下文对象
		CTX(r *ghttp.Request)
		// CORS 允许接口跨域请求
		CORS(r *ghttp.Request)
		// ResponseHandler 响应函数
		ResponseHandler(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
