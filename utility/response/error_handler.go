package response

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

// ErrorLevel 错误级别
type ErrorLevel int

const (
	ErrorLevelInfo ErrorLevel = iota
	ErrorLevelWarn
	ErrorLevelError
	ErrorLevelFatal
)

// ErrorContext 错误上下文信息
type ErrorContext struct {
	RequestID   string                 `json:"request_id"`
	UserID      int64                  `json:"user_id,omitempty"`
	Method      string                 `json:"method"`
	URL         string                 `json:"url"`
	UserAgent   string                 `json:"user_agent,omitempty"`
	IP          string                 `json:"ip"`
	Timestamp   time.Time              `json:"timestamp"`
	StackTrace  string                 `json:"stack_trace,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
}

// ErrorHandler 统一错误处理器
type ErrorHandler struct {
	logger *glog.Logger
}

// NewErrorHandler 创建错误处理器
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		logger: g.Log(),
	}
}

// HandleError 处理错误
func (eh *ErrorHandler) HandleError(ctx context.Context, err error, level ErrorLevel, extra ...map[string]interface{}) {
	if err == nil {
		return
	}

	// 构建错误上下文
	errorCtx := eh.buildErrorContext(ctx, extra...)
	
	// 记录错误日志
	eh.logError(ctx, err, level, errorCtx)
	
	// 可以在这里添加其他处理逻辑，如发送告警、上报监控等
}

// HandleHTTPError 处理HTTP请求错误
func (eh *ErrorHandler) HandleHTTPError(r *ghttp.Request, err error, code int, level ErrorLevel) {
	if err == nil {
		return
	}

	// 构建错误上下文
	errorCtx := eh.buildHTTPErrorContext(r)
	
	// 记录错误日志
	eh.logError(r.Context(), err, level, errorCtx)
	
	// 返回错误响应
	eh.sendErrorResponse(r, err, code)
}

// buildErrorContext 构建错误上下文
func (eh *ErrorHandler) buildErrorContext(ctx context.Context, extra ...map[string]interface{}) *ErrorContext {
	errorCtx := &ErrorContext{
		Timestamp: time.Now(),
		Extra:     make(map[string]interface{}),
	}

	// 从上下文中获取请求ID
	if requestID := ctx.Value("RequestId"); requestID != nil {
		errorCtx.RequestID = fmt.Sprintf("%v", requestID)
	}

	// 合并额外信息
	for _, e := range extra {
		for k, v := range e {
			errorCtx.Extra[k] = v
		}
	}

	return errorCtx
}

// buildHTTPErrorContext 构建HTTP错误上下文
func (eh *ErrorHandler) buildHTTPErrorContext(r *ghttp.Request) *ErrorContext {
	errorCtx := &ErrorContext{
		RequestID: r.GetCtxVar("RequestId").String(),
		Method:    r.Method,
		URL:       r.URL.String(),
		UserAgent: r.UserAgent(),
		IP:        r.GetClientIp(),
		Timestamp: time.Now(),
		Extra:     make(map[string]interface{}),
	}

	// 获取用户ID（如果存在）
	if userID := r.GetCtxVar("UserId"); !userID.IsNil() {
		errorCtx.UserID = userID.Int64()
	}

	return errorCtx
}

// logError 记录错误日志
func (eh *ErrorHandler) logError(ctx context.Context, err error, level ErrorLevel, errorCtx *ErrorContext) {
	// 获取堆栈信息
	if level >= ErrorLevelError {
		errorCtx.StackTrace = eh.getStackTrace()
	}

	logData := g.Map{
		"error":   err.Error(),
		"context": errorCtx,
	}

	switch level {
	case ErrorLevelInfo:
		eh.logger.Info(ctx, "Error occurred", logData)
	case ErrorLevelWarn:
		eh.logger.Warning(ctx, "Warning occurred", logData)
	case ErrorLevelError:
		eh.logger.Error(ctx, "Error occurred", logData)
	case ErrorLevelFatal:
		eh.logger.Fatal(ctx, "Fatal error occurred", logData)
	}
}

// sendErrorResponse 发送错误响应
func (eh *ErrorHandler) sendErrorResponse(r *ghttp.Request, err error, code int) {
	// 检查是否是gerror类型，如果是则提取错误码
	if gErr, ok := err.(*gerror.Error); ok {
		if gErr.Code().Code() != 0 {
			code = gErr.Code().Code()
		}
	}

	// 发送错误响应
	Error(r, code, err.Error())
}

// getStackTrace 获取堆栈跟踪信息
func (eh *ErrorHandler) getStackTrace() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stack := string(buf[:n])
	
	// 过滤掉不必要的堆栈信息
	lines := strings.Split(stack, "\n")
	var filteredLines []string
	skip := true
	
	for _, line := range lines {
		if strings.Contains(line, "error_handler.go") {
			skip = false
			continue
		}
		if !skip {
			filteredLines = append(filteredLines, line)
		}
	}
	
	return strings.Join(filteredLines, "\n")
}

// 全局错误处理器实例
var globalErrorHandler = NewErrorHandler()

// HandleError 全局错误处理函数
func HandleError(ctx context.Context, err error, level ErrorLevel, extra ...map[string]interface{}) {
	globalErrorHandler.HandleError(ctx, err, level, extra...)
}

// HandleHTTPError 全局HTTP错误处理函数
func HandleHTTPError(r *ghttp.Request, err error, code int, level ErrorLevel) {
	globalErrorHandler.HandleHTTPError(r, err, code, level)
}

// 便捷的错误处理函数

// LogInfo 记录信息级别错误
func LogInfo(ctx context.Context, err error, extra ...map[string]interface{}) {
	HandleError(ctx, err, ErrorLevelInfo, extra...)
}

// LogWarn 记录警告级别错误
func LogWarn(ctx context.Context, err error, extra ...map[string]interface{}) {
	HandleError(ctx, err, ErrorLevelWarn, extra...)
}

// LogError 记录错误级别错误
func LogError(ctx context.Context, err error, extra ...map[string]interface{}) {
	HandleError(ctx, err, ErrorLevelError, extra...)
}

// LogFatal 记录致命级别错误
func LogFatal(ctx context.Context, err error, extra ...map[string]interface{}) {
	HandleError(ctx, err, ErrorLevelFatal, extra...)
}

// HTTPError 处理HTTP错误（错误级别）
func HTTPError(r *ghttp.Request, err error, code int) {
	HandleHTTPError(r, err, code, ErrorLevelError)
}

// HTTPWarn 处理HTTP警告（警告级别）
func HTTPWarn(r *ghttp.Request, err error, code int) {
	HandleHTTPError(r, err, code, ErrorLevelWarn)
}