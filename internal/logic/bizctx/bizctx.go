package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sBizCtx struct {
	contextKey string
}

func init() {
	service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return &sBizCtx{
		contextKey: g.Cfg().MustGet(context.Background(), "service.HeaderContextKey").String(),
	}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(s.contextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(s.contextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sBizCtx) SetUser(ctx context.Context, claimsUser *model.JwtCustomClaims) {
	s.Get(ctx).ClaimsUser = claimsUser
}
