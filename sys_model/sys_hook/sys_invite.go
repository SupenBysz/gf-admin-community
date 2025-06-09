package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

// func(上下文，匹配状态，邀约记录，注册表单)

type InviteRegisterHookFunc func(ctx context.Context, state sys_enum.InviteType, invite *sys_model.InviteRes, invitePerson *sys_entity.SysInvitePerson, registerInfo *sys_model.SysUser) (bool, error)
type InviteRegisterHookInfo struct {
	Key   sys_enum.InviteType
	Value InviteRegisterHookFunc
	//Category int `json:"category" dc:"业务类别"`
}

// InviteStateHookFunc 订阅邀约状态
type InviteStateHookFunc func(ctx context.Context, state sys_enum.InviteState, invite *sys_model.InviteRes) error
