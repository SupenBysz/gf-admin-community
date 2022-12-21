package sys_model

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type ReqCasbin struct {
	UserId    int64  `p:"userId"`
	Domain    string `p:"domain"`
	Interface string `p:"i"`
	Action    string `p:"a"`
}

type CasbinHookFunc func(ctx context.Context, info sys_entity.SysUser) (bool, error)

type CasbinHookInfo HookEventType[sys_enum.UserType, CasbinHookFunc]
