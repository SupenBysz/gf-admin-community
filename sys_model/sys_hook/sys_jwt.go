package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type JwtHookFunc func(ctx context.Context, claims *sys_model.JwtCustomClaims) (*sys_model.JwtCustomClaims, error)

type JwtHookInfo sys_model.HookEventType[sys_enum.UserType, JwtHookFunc]
