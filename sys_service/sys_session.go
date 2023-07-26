// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISessionError interface {
		// Append 追加错误至错误队列
		Append(ctx context.Context, error error) error
		// HasError 错误队列中检索指定错误
		HasError(ctx context.Context, err error) (response bool)
		// Iterator 获取错误信息队列
		Iterator(ctx context.Context, f func(k int, err error) bool)
	}
	ISysSession interface {
		// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
		Init(sessionContext *sys_model.SessionContext, r *ghttp.Request, ctx ...*context.Context)
		NewSessionCtx(ctx context.Context) context.Context
		HasCustom(ctx context.Context) bool
		// Get 获得上下文变量，如果没有设置，那么返回nil
		Get(ctx context.Context) *sys_model.SessionContext
		SetUserById(ctx *context.Context, userId int64) *sys_model.SessionContext
		// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetUser(ctx context.Context, claimsUser *sys_model.JwtCustomClaims)
	}
)

var (
	localSessionError ISessionError
	localSysSession   ISysSession
)

func SessionError() ISessionError {
	if localSessionError == nil {
		panic("implement not found for interface ISessionError, forgot register?")
	}
	return localSessionError
}

func RegisterSessionError(i ISessionError) {
	localSessionError = i
}

func SysSession() ISysSession {
	if localSysSession == nil {
		panic("implement not found for interface ISysSession, forgot register?")
	}
	return localSysSession
}

func RegisterSysSession(i ISysSession) {
	localSysSession = i
}
