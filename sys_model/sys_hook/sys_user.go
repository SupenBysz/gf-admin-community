package sys_hook

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type UserHookFunc sys_model.HookFuncRes[sys_enum.UserEvent, sys_model.SysUser]
type UserHookInfo sys_model.HookEventType[sys_enum.UserEvent, UserHookFunc]
