// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysMemberLevel interface {
		// QueryMemberLevelList 获取会员等级列表
		QueryMemberLevelList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysMemberLevelListRes, error)
		// CreateMemberLevel 创建会员等级
		CreateMemberLevel(ctx context.Context, info *sys_model.SysMemberLevel, userId int64, unionMainId int64) (*sys_model.SysMemberLevelRes, error)
		// UpdateMemberLevel 更新会员等级
		UpdateMemberLevel(ctx context.Context, info *sys_model.UpdateSysMemberLevel, unionMainId int64) (*sys_model.SysMemberLevelRes, error)
		// DeleteMemberLevel 删除会员等级
		DeleteMemberLevel(ctx context.Context, id int64, unionMainId int64) (bool, error)
		// GetMemberLevelById 获取会员等级详情
		GetMemberLevelById(ctx context.Context, id int64) (*sys_model.SysMemberLevelRes, error)
		// QueryMemberLevelUserList 获取会员等级用户列表
		QueryMemberLevelUserList(ctx context.Context, memberLevelId int64) (*sys_model.SysMemberLevelUserListRes, error)
		// AddMemberLevelUser 添加会员等级用户
		AddMemberLevelUser(ctx context.Context, memberLevelId int64, userIds []int64) (bool, error)
		// DeleteMemberLevelUser 批量删除会员等级的用户
		DeleteMemberLevelUser(ctx context.Context, memberLevelId int64, userIds []int64) (bool, error)
	}
)

var (
	localSysMemberLevel ISysMemberLevel
)

func SysMemberLevel() ISysMemberLevel {
	if localSysMemberLevel == nil {
		panic("implement not found for interface ISysMemberLevel, forgot register?")
	}
	return localSysMemberLevel
}

func RegisterSysMemberLevel(i ISysMemberLevel) {
	localSysMemberLevel = i
}
