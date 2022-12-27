package sys_bizctx

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/gogf/gf/v2/net/ghttp"
)

type hookInfo sys_model.KeyValueT[int64, sys_model.BizCtxHookInfo]

type sBizCtx struct {
	contextKey string
	hookArr    []hookInfo
}

func init() {
	sys_service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return &sBizCtx{
		contextKey: g.Cfg().MustGet(context.Background(), "service.bizCtxContextKey", "gf-admin-bizctx").String(),
	}
}

// InstallHook 安装Hook
func (s *sBizCtx) InstallHook(userType sys_enum.UserType, hookFunc sys_model.BizCtxHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_model.BizCtxHookInfo{Key: userType, Value: hookFunc}}

	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sBizCtx) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除所有Hook
func (s *sBizCtx) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *sys_model.Context) {
	r.SetCtxVar(s.contextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sBizCtx) Get(ctx context.Context) *sys_model.Context {
	value := ctx.Value(s.contextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*sys_model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sBizCtx) SetUser(ctx context.Context, claimsUser *sys_model.JwtCustomClaims) {
	s.Get(ctx).ClaimsUser = claimsUser
}

// GetUnionMainId 当前登录用户的主体id获取
func (s *sBizCtx) GetUnionMainId(ctx context.Context) (int64, error) {
	userId := s.Get(ctx).ClaimsUser.Id

	userInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

	var userUnionMainId int64

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&userInfo.Type == userInfo.Type {
				userUnionMainId, err = hook.Value.Value(ctx, *userInfo)
				if err != nil {
					break
				}
			}
		}
	})

	if err != nil {
		return 0, err
	}
	return userUnionMainId, nil
}
