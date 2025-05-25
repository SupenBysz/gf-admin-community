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
		// 通过标识符获取邀请信息
		GetInviteByIdentifier(ctx context.Context, identifier string) (*sys_model.InviteRes, error)
		// CreateInvite 创建邀约信息
		CreateInvite(ctx context.Context, info *sys_model.Invite) (*sys_model.InviteRes, error)
		// DeleteInvite 删除邀约信息
		DeleteInvite(ctx context.Context, inviteId int64) (bool, error)
		// SetInviteState 修改邀约信息状态
		SetInviteState(ctx context.Context, id int64, state int) (bool, error)
		// SetInviteNumber 修改邀约剩余次数
		SetInviteNumber(ctx context.Context, id int64, num int, isAdd bool) (res bool, err error)
		// GetInvitePersonById 获取被邀请信息
		GetInvitePersonById(ctx context.Context, id int64) (*sys_model.InvitePersonRes, error)
		// GetInvitePersonByUserId 获取被邀请信息
		GetInvitePersonByUserId(ctx context.Context, userId int64) (*sys_model.InvitePersonRes, error)
		// QueryInvitePersonList 获取邀请列表
		QueryInvitePersonList(ctx context.Context, inviteUserId int64) (*sys_model.InvitePersonListRes, error)
		// CreateInvitePerson 创建被邀请信息
		CreateInvitePerson(ctx context.Context, info *sys_model.InvitePersonInfo) (*sys_model.InvitePersonRes, error)
		// CountRegisterInvitePersonByInviteCode 统计邀请码邀请的人数
		CountRegisterInvitePersonByInviteCode(ctx context.Context, inviteCode string) (int, error)
		// CountRegisterInvitePersonByInviteId 统计邀请码邀请的人数
		CountRegisterInvitePersonByInviteId(ctx context.Context, inviteId int64) (int, error)
		// CountRegisterInvitePersonByFormUserId 统计邀请人邀请的人数
		C(ctx context.Context, formUserId int64) (int, error)
		// IsInviteCodeOverLimit 判断邀请码是否使用上限
		IsInviteCodeOverLimit(ctx context.Context, inviteCode string) (bool, error)
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
