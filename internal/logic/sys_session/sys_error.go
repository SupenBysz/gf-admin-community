package sys_session

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type sSessionError struct {
}

func init() {
	sys_service.RegisterSessionError(NewSessionError())
}

func NewSessionError() *sSessionError {
	return &sSessionError{}
}

// Append 追加错误至错误队列
func (s *sSessionError) Append(ctx context.Context, error error) error {
	bizctx := sys_service.BizCtx().Get(ctx)
	if bizctx == nil {
		return error
	}
	bizctx.SessionErrorQueue.Append(error)
	return error
}

// HasError 错误队列中检索指定错误
func (s *sSessionError) HasError(ctx context.Context, err error) (response bool) {
	response = false
	s.Iterator(ctx, func(_ int, v error) bool {
		if v != nil && v.Error() == err.Error() {
			response = true
			return false
		}
		return true
	})
	return response
}

// Iterator 获取错误信息队列
func (s *sSessionError) Iterator(ctx context.Context, f func(k int, err error) bool) {
	sys_service.BizCtx().Get(ctx).SessionErrorQueue.Iterator(func(key int, v interface{}) bool {
		return f(key, v.(error))
	})
}
