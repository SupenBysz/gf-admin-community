package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type CasbinHookFunc func(ctx context.Context, info sys_entity.SysUser) (bool, error)

type CasbinHookInfo sys_model.HookEventType[sys_enum.UserType, CasbinHookFunc]
