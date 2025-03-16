package response

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/i18n"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}

	// 尝试国际化消息
	ctx := r.Context()
	translatedMessage := message
	if message != "" {
		// 尝试翻译消息
		translatedMessage = i18n.T(ctx, message)
	}

	if reflect.ValueOf(data) == reflect.ValueOf(api_v1.JsonRes{}) {
		r.Response.WriteJson(data)
	} else {
		if len(data) > 0 && strings.Contains(reflect.ValueOf(data[0]).String(), "JsonRes") && len(data) > 0 && reflect.TypeOf(data[0]).Name() == reflect.TypeOf([]api_v1.JsonRes{}).Name() {
			jsonRes := (*api_v1.JsonRes)(unsafe.Pointer(&data[0]))

			r.Response.WriteJson(api_v1.JsonRes{
				Code:    code,
				Message: translatedMessage,
				Data:    gconv.Map(jsonRes.Data)["Data"],
				Time:    gtime.Now().Format("Y-m-d H:i:s"),
			})
			return
		}

		r.Response.WriteJson(api_v1.JsonRes{
			Code:    code,
			Message: translatedMessage,
			Data:    responseData,
			Time:    gtime.Now().Format("Y-m-d H:i:s"),
		})
	}
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.ExitAll()
}

// JsonRedirect 返回标准JSON数据引导客户端跳转。
func JsonRedirect(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	// 尝试国际化消息
	ctx := r.Context()
	translatedMessage := message
	if message != "" {
		// 尝试翻译消息
		translatedMessage = i18n.T(ctx, message)
	}

	if reflect.ValueOf(data) == reflect.ValueOf(api_v1.JsonRes{}) {
		r.Response.WriteJson(data)
	} else {
		r.Response.WriteJson(api_v1.JsonRes{
			Code:    code,
			Message: translatedMessage,
			Data:    responseData,
			Time:    gtime.Now().Format("Y-m-d H:i:s"),
		})
	}
}

// JsonRedirectExit 返回标准JSON数据引导客户端跳转，并退出当前HTTP执行函数。
func JsonRedirectExit(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	JsonRedirect(r, code, message, redirect, data...)
	r.Exit()
}
