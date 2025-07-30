package response

import (
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/i18n"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// 配置项
var (
	DefaultTimeFormat = "Y-m-d H:i:s"
	timeFormat        = DefaultTimeFormat
)

// ResponseBuilder 响应构建器
type ResponseBuilder struct {
	code    int
	message string
	data    interface{}
	i18nKey string
}

// ResponseOption 响应选项函数
type ResponseOption func(*ResponseBuilder)

// WithCode 设置响应码
func WithCode(code int) ResponseOption {
	return func(rb *ResponseBuilder) {
		rb.code = code
	}
}

// WithMessage 设置响应消息
func WithMessage(message string) ResponseOption {
	return func(rb *ResponseBuilder) {
		rb.message = message
	}
}

// WithI18nKey 设置国际化键
func WithI18nKey(key string) ResponseOption {
	return func(rb *ResponseBuilder) {
		rb.i18nKey = key
	}
}

// WithData 设置响应数据
func WithData(data interface{}) ResponseOption {
	return func(rb *ResponseBuilder) {
		rb.data = data
	}
}

// SetTimeFormat 设置时间格式
func SetTimeFormat(format string) {
	timeFormat = format
}

// buildResponse 构建响应对象
func buildResponse(r *ghttp.Request, options ...ResponseOption) *api_v1.JsonRes {
	builder := &ResponseBuilder{
		code: 0,
		data: g.Map{},
	}
	
	for _, option := range options {
		option(builder)
	}
	
	// 处理国际化
	message := builder.message
	if builder.i18nKey != "" {
		message = i18n.T(r.Context(), builder.i18nKey)
	} else if builder.message != "" {
		message = i18n.T(r.Context(), builder.message)
	}
	
	return &api_v1.JsonRes{
		Code:    builder.code,
		Message: message,
		Data:    builder.data,
		Time:    gtime.Now().Format(timeFormat),
	}
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
		
		// 使用类型断言检查是否已经是JsonRes类型
		if jsonRes, ok := data[0].(api_v1.JsonRes); ok {
			r.Response.WriteJson(jsonRes)
			return
		}
		
		// 检查是否是JsonRes指针
		if jsonRes, ok := data[0].(*api_v1.JsonRes); ok {
			r.Response.WriteJson(*jsonRes)
			return
		}
	} else {
		responseData = g.Map{}
	}

	// 使用构建器模式创建响应
	response := buildResponse(r, 
		WithCode(code),
		WithMessage(message),
		WithData(responseData),
	)
	
	r.Response.WriteJson(response)
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.ExitAll()
}

// JsonRedirect 返回标准JSON数据引导客户端跳转。
func JsonRedirect(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
		
		// 使用类型断言检查是否已经是JsonRes类型
		if jsonRes, ok := data[0].(api_v1.JsonRes); ok {
			r.Response.WriteJson(jsonRes)
			return
		}
		
		// 检查是否是JsonRes指针
		if jsonRes, ok := data[0].(*api_v1.JsonRes); ok {
			r.Response.WriteJson(*jsonRes)
			return
		}
	} else {
		responseData = g.Map{}
	}

	// 使用构建器模式创建响应
	response := buildResponse(r,
		WithCode(code),
		WithMessage(message),
		WithData(responseData),
	)
	
	r.Response.WriteJson(response)
}

// JsonRedirectExit 返回标准JSON数据引导客户端跳转，并退出当前HTTP执行函数。
func JsonRedirectExit(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	JsonRedirect(r, code, message, redirect, data...)
	r.Exit()
}

// ============ 便捷响应方法 ============

// Success 成功响应
func Success(r *ghttp.Request, data interface{}, message ...string) {
	options := []ResponseOption{
		WithCode(0),
		WithData(data),
	}
	if len(message) > 0 && message[0] != "" {
		options = append(options, WithMessage(message[0]))
	}
	response := buildResponse(r, options...)
	r.Response.WriteJson(response)
}

// SuccessExit 成功响应并退出
func SuccessExit(r *ghttp.Request, data interface{}, message ...string) {
	Success(r, data, message...)
	r.ExitAll()
}

// Error 错误响应
func Error(r *ghttp.Request, code int, message string) {
	response := buildResponse(r, WithCode(code), WithMessage(message))
	r.Response.WriteJson(response)
}

// ErrorExit 错误响应并退出
func ErrorExit(r *ghttp.Request, code int, message string) {
	Error(r, code, message)
	r.ExitAll()
}

// I18nResponse 带国际化键的响应
func I18nResponse(r *ghttp.Request, code int, i18nKey string, data interface{}) {
	response := buildResponse(r, 
		WithCode(code), 
		WithI18nKey(i18nKey), 
		WithData(data),
	)
	r.Response.WriteJson(response)
}

// I18nResponseExit 带国际化键的响应并退出
func I18nResponseExit(r *ghttp.Request, code int, i18nKey string, data interface{}) {
	I18nResponse(r, code, i18nKey, data)
	r.ExitAll()
}

// JsonWithError 带错误处理的响应
func JsonWithError(r *ghttp.Request, err error, successData interface{}, successMessage ...string) {
	if err != nil {
		Error(r, 1, err.Error())
		return
	}
	Success(r, successData, successMessage...)
}

// JsonWithErrorExit 带错误处理的响应并退出
func JsonWithErrorExit(r *ghttp.Request, err error, successData interface{}, successMessage ...string) {
	if err != nil {
		ErrorExit(r, 1, err.Error())
		return
	}
	SuccessExit(r, successData, successMessage...)
}

// JsonWithValidation 带验证的数据响应
func JsonWithValidation(r *ghttp.Request, data interface{}, validator func(interface{}) error, options ...ResponseOption) {
	if validator != nil {
		if err := validator(data); err != nil {
			Error(r, 1, err.Error())
			return
		}
	}
	
	// 默认添加数据
	allOptions := append([]ResponseOption{WithData(data)}, options...)
	response := buildResponse(r, allOptions...)
	r.Response.WriteJson(response)
}

// JsonBuilder 使用构建器模式的响应
func JsonBuilder(r *ghttp.Request, options ...ResponseOption) {
	response := buildResponse(r, options...)
	r.Response.WriteJson(response)
}

// JsonBuilderExit 使用构建器模式的响应并退出
func JsonBuilderExit(r *ghttp.Request, options ...ResponseOption) {
	JsonBuilder(r, options...)
	r.ExitAll()
}
