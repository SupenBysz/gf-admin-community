package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type MemberLevelHookFunc func(ctx context.Context, state sys_enum.MemberLevelEvent, info *sys_model.SysMemberLevelRes) error
type MemberLevelHookInfo struct {
	Key   sys_enum.MemberLevelEvent
	Value MemberLevelHookFunc
}
