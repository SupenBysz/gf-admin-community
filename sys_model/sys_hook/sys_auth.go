package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type AuthHookFunc func(ctx context.Context, actionType sys_enum.AuthActionType, info *sys_model.SysUser) error
type AuthHookInfo struct {
	Key      sys_enum.AuthActionType
	Value    AuthHookFunc
}
