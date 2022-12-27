package sys_model

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type BizCtxHookFunc func(ctx context.Context, userInfo sys_entity.SysUser) (int64, error)

type BizCtxHookInfo HookEventType[sys_enum.UserType, BizCtxHookFunc]
