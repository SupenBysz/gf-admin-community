package sys_session

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sSysSession struct {
	contextKey string
}

func init() {
	sys_service.RegisterSysSession(New())
}

func New() *sSysSession {
	return &sSysSession{
		contextKey: g.Cfg().MustGet(context.Background(), "service.sessionContextKey", "gf-admin-community").String(),
	}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sSysSession) Init(sessionContext *sys_model.SessionContext, r *ghttp.Request, ctx ...*context.Context) {
	if r != nil {
		r.SetCtxVar(s.contextKey, sessionContext)
	} else if len(ctx) > 0 {
		*ctx[0] = context.WithValue(*ctx[0], s.contextKey, sessionContext)
	}
}

func (s *sSysSession) NewSessionCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, "hasCustom", "")

}

func (s *sSysSession) HasCustom(ctx context.Context) bool {
	return ctx.Value("hasCustom") != nil
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sSysSession) Get(ctx context.Context) *sys_model.SessionContext {
	value := ctx.Value(s.contextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*sys_model.SessionContext); ok {
		return localCtx
	}
	return nil
}

func (s *sSysSession) SetUserById(ctx *context.Context, userId int64) *sys_model.SessionContext {

	user, err := sys_service.SysUser().GetSysUserById(*ctx, userId)

	if err != nil {
		return nil
	}

	token, err := sys_service.Jwt().GenerateToken(*ctx, user)

	if token != nil && err == nil && s.HasCustom(*ctx) {
		claim := sys_service.Jwt().MakeSession(*ctx, token.Token)

		s.Init(&sys_model.SessionContext{
			JwtClaimsUser:     claim,
			Ipv4:              "",
			SessionErrorQueue: nil,
		}, nil, ctx)

		//s.SetUser(ctx, claim)
	}

	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sSysSession) SetUser(ctx context.Context, claimsUser *sys_model.JwtCustomClaims) {
	if claimsUser.Type == -1 {
		claimsUser.IsSuperAdmin = true
	}
	s.Get(ctx).JwtClaimsUser = claimsUser
}
