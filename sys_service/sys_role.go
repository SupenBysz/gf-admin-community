// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysRole interface {
		// InstallInviteRegisterHook 订阅邀约注册Hook
		InstallInviteRegisterHook(action sys_enum.RoleMemberChange, hookFunc sys_hook.RoleMemberChangeHookFunc)
		// QueryRoleList 获取角色列表
		QueryRoleList(ctx context.Context, info base_model.SearchParams, unionMainId int64) (*sys_model.RoleListRes, error)
		// GetRoleById 根据id获取角色
		GetRoleById(ctx context.Context, id int64) (*sys_entity.SysRole, error)
		// Create 创建角色信息
		Create(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error)
		// Update 更新角色信息
		Update(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error)
		// Save 新增或保存角色信息
		Save(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error)
		// Delete 删除角色信息
		Delete(ctx context.Context, roleId int64) (bool, error)
		// SetRoleMember 设置角色用户
		SetRoleMember(ctx context.Context, roleId int64, userIds []int64, makeUserUnionMainId int64) (bool, error)
		// RemoveRoleMember 移除角色中的用户
		RemoveRoleMember(ctx context.Context, roleId int64, userIds []int64) (bool, error)
		// GetRoleMemberIds 获取角色下的所有用户ID
		GetRoleMemberIds(ctx context.Context, roleId int64, makeUserUnionMainId int64) ([]int64, error)
		// GetRoleMemberList 获取角色下的所有用户
		GetRoleMemberList(ctx context.Context, roleId int64, makeUserUnionMainId int64) ([]*sys_model.SysUser, error)
		// GetRoleListByUserId 获取用户拥有的所有角色
		GetRoleListByUserId(ctx context.Context, userId int64) ([]*sys_entity.SysRole, error)
		// SetRolePermissions 设置角色权限
		SetRolePermissions(ctx context.Context, roleId int64, permissionIds []int64, makeUserUnionMainId int64) (bool, error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
