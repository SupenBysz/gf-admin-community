package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysUser 鉴权
var SysUser = cSysUser{}

type cSysUser struct{}

// CreateUser 创建用户|信息
func (c *cSysUser) CreateUser(ctx context.Context, req *sys_api.CreateUserReq) (res api_v1.BoolRes, err error) {
	_, err = sys_service.SysUser().CreateUser(ctx, req.UserInnerRegister, sys_enum.User.State.Normal, sys_enum.User.Type.User)

	return err == nil, err
}

// QueryUserList 获取用户|列表
func (c *cSysUser) QueryUserList(ctx context.Context, req *sys_api.QueryUserListReq) (*sys_api.UserListRes, error) {
	data, err := sys_service.SysUser().QueryUserList(ctx, &req.SearchParams, false)
	if err != nil {
		return nil, err
	}
	return (*sys_api.UserListRes)(data), err
}

// SetUserRoleIds 设置用户角色
func (c *cSysRole) SetUserRoleIds(ctx context.Context, req *sys_api.SetUserRoleIdsReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysUser().SetUserRoleIds(ctx, req.RoleIds, req.UserId)

	return result == true, err
}

// SetUserPermissionIds 设置用户权限
func (c *cSysUser) SetUserPermissionIds(ctx context.Context, req *sys_api.SetUserPermissionIdsReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysUser().SetUserPermissionIds(ctx, req.Id, req.PermissionIds)
	return result == true, err
}

// GetUserPermissionIds 获取用户权限Ids
func (c *cSysUser) GetUserPermissionIds(ctx context.Context, req *sys_api.GetUserPermissionIdsReq) (*api_v1.Int64ArrRes, error) {
	result, err := sys_service.SysUser().GetUserPermissionIds(ctx, req.Id)
	return (*api_v1.Int64ArrRes)(&result), err
}

// SetUsername 设置用户登陆名
func (c *cSysUser) SetUsername(ctx context.Context, req *sys_api.SetUsernameByIdReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysUser().SetUsername(ctx, req.NewUsername)
	return result == true, err
}

// UpdateUserPassword 修改密码
func (c *cSysUser) UpdateUserPassword(ctx context.Context, req *sys_api.UpdateUserPasswordReq) (api_v1.BoolRes, error) {
	_, err := sys_service.SysUser().UpdateUserPassword(ctx, req.UpdateUserPassword)

	if err != nil {
		return false, err
	}
	return true, nil
}

// ResetUserPassword 重置用户密码
func (c *cSysUser) ResetUserPassword(ctx context.Context, req *sys_api.ResetUserPasswordReq) (res api_v1.BoolRes, err error) {
	_, err = sys_service.SysUser().ResetUserPassword(ctx, req.UserId, req.Password, req.ConfirmPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
