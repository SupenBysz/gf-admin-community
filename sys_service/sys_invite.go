// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysInvite interface {
		InstallInviteStateHook(actionType sys_enum.InviteState, hookFunc sys_hook.InviteStateHookFunc)
		// GetInviteById 根据id获取邀约
		GetInviteById(ctx context.Context, id int64) (*sys_model.InviteRes, error)
		// QueryInviteList 查询邀约｜列表
		QueryInviteList(ctx context.Context, filter *base_model.SearchParams) (*sys_model.InviteListRes, error)
		// CreateInvite 创建邀约信息
		CreateInvite(ctx context.Context, info *sys_model.Invite) (*sys_model.InviteRes, error)
		// DeleteInvite 删除邀约信息
		DeleteInvite(ctx context.Context, inviteId int64) (bool, error)
		// SetInviteState 修改邀约信息状态
		SetInviteState(ctx context.Context, id int64, state int) (bool, error)
		// SetInviteNumber 修改邀约剩余次数
		SetInviteNumber(ctx context.Context, id int64, num int, isAdd bool) (res bool, err error)
	}
)

var (
	localSysInvite ISysInvite
)

func SysInvite() ISysInvite {
	if localSysInvite == nil {
		panic("implement not found for interface ISysInvite, forgot register?")
	}
	return localSysInvite
}

func RegisterSysInvite(i ISysInvite) {
	localSysInvite = i
}
