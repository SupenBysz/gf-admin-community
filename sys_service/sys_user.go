// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysUser interface {
		// InstallHook 安装Hook
		InstallHook(event sys_enum.UserEvent, hookFunc sys_hook.UserHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除所有Hook
		CleanAllHook()
		// UpdateHeartbeatAt 更新用户心跳监测时间
		UpdateHeartbeatAt(ctx context.Context, heartbeatTimeout int) (api_v1.BoolRes, error)
		// QueryUserList 获取用户列表
		QueryUserList(ctx context.Context, info *base_model.SearchParams, unionMainId int64, isExport bool) (response *sys_model.SysUserListRes, err error)
		// SetUserRoleIds 设置用户角色
		SetUserRoleIds(ctx context.Context, roleIds []int64, userId int64) (bool, error)
		// CreateUser 创建用户
		CreateUser(ctx context.Context, info sys_model.UserInnerRegister, userState sys_enum.UserState, userType sys_enum.UserType, customId ...int64) (*sys_model.SysUser, error)
		// SetUserPermissions 设置用户权限
		SetUserPermissions(ctx context.Context, userId int64, permissionIds []int64) (bool, error)
		// GetSysUserByUsername 根据用户名获取用户
		GetSysUserByUsername(ctx context.Context, username string) (response *sys_model.SysUser, err error)
		// CheckPassword 检查密码是否正确
		CheckPassword(ctx context.Context, userId int64, password string) (bool, error)
		// HasSysUserByUsername 判断用户名是否存在
		HasSysUserByUsername(ctx context.Context, username string) bool
		// GetSysUserById 根据用户ID获取用户信息
		GetSysUserById(ctx context.Context, userId int64) (*sys_model.SysUser, error)
		MakeSession(ctx context.Context, userId int64)
		// SetUserPermissionIds 设置用户权限
		SetUserPermissionIds(ctx context.Context, userId int64, permissionIds []int64) (bool, error)
		// DeleteUser 删除用户信息，该方法一般由后端业务层内部调用
		DeleteUser(ctx context.Context, id int64) (bool, error)
		// SetUsername 修改自己的账号登陆名称
		SetUsername(ctx context.Context, newUsername string, userId int64) (bool, error)
		// SetUserState 设置用户状态
		SetUserState(ctx context.Context, userId int64, state sys_enum.UserType) (bool, error)
		// UpdateUserPassword 修改用户登录密码
		UpdateUserPassword(ctx context.Context, info sys_model.UpdateUserPassword, userId int64) (bool, error)
		// ResetUserPassword 重置用户密码 (超级管理员无需验证验证，XX商管理员重置员工密码无需验证)
		ResetUserPassword(ctx context.Context, userId int64, password string, confirmPassword string) (bool, error)
		// HasSysUserEmail 邮箱是否存在
		HasSysUserEmail(ctx context.Context, email string) bool
		// GetSysUserByEmail 根据邮箱获取用户信息
		GetSysUserByEmail(ctx context.Context, email string) (response *sys_model.SysUser, err error)
		// ResetUserEmail 重置用户邮箱
		ResetUserEmail(ctx context.Context, userId int64, email string) (bool, error)
		// SetUserRoles 设置用户角色
		SetUserRoles(ctx context.Context, userId int64, roleIds []int64, makeUserUnionMainId int64) (bool, error)
		// UpdateUserExDetail 更新用户扩展信息
		UpdateUserExDetail(ctx context.Context, user *sys_model.SysUser) (*sys_model.SysUser, error)
		// GetUserDetail 查看用户详情，含完整手机号
		GetUserDetail(ctx context.Context, userId int64) (*sys_model.SysUser, error)
		// GetUserListByMobileOrMail 根据手机号或者邮箱查询用户列表
		GetUserListByMobileOrMail(ctx context.Context, info string) (*sys_model.SysUserListRes, error)
		// SetUserMobile 设置用户手机号
		SetUserMobile(ctx context.Context, newMobile string, captcha string, password string, userId int64) (bool, error)
		// SetUserMail 设置用户邮箱
		SetUserMail(ctx context.Context, oldMail string, newMail string, captcha string, password string, userId int64) (bool, error)
		// Heartbeat 用户在线心跳
		Heartbeat(ctx context.Context, userId int64) (bool, error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
