package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	kyUser2 "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/user"
)

// SysUser 鉴权
var SysUser = cSysUser{}

type cSysUser struct{}

// CreateUser 创建用户|信息
func (c *cSysUser) CreateUser(ctx context.Context, req *sysapi.CreateUserReq) (res api_v1.BoolRes, err error) {
	_, err = sys_service.SysUser().CreateUser(ctx, req.UserInnerRegister, kyUser2.State.Normal, kyUser2.Type.User)

	return err == nil, err
}

// QueryUserList 获取用户|列表
func (c *cSysUser) QueryUserList(ctx context.Context, req *sysapi.QueryUserListReq) (*sysapi.UserListRes, error) {
	data, err := sys_service.SysUser().QueryUserList(ctx, &req.SearchParams, false)
	if err != nil {
		return nil, err
	}
	return (*sysapi.UserListRes)(data), err
}

// SetUserRoleIds 设置用户角色
func (c *cSysRole) SetUserRoleIds(ctx context.Context, req *sysapi.SetUserRoleIdsReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysUser().SetUserRoleIds(ctx, req.RoleIds, req.UserId)

	return result == true, err
}

// SetUserPermissionIds 设置用户权限
func (c *cSysUser) SetUserPermissionIds(ctx context.Context, req *sysapi.SetUserPermissionIdsReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysUser().SetUserPermissionIds(ctx, req.Id, req.PermissionIds)
	return result == true, err
}

// GetUserPermissionIds 获取用户权限Ids
func (c *cSysUser) GetUserPermissionIds(ctx context.Context, req *sysapi.GetUserPermissionIdsReq) (*api_v1.Int64ArrRes, error) {
	result, err := sys_service.SysUser().GetUserPermissionIds(ctx, req.Id)
	return (*api_v1.Int64ArrRes)(&result), err
}

// SetUsername 设置用户登陆名
func (c *cSysUser) SetUsername(ctx context.Context, req *sysapi.SetUsernameByIdReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysUser().SetUsername(ctx, req.NewUsername)
	return result == true, err
}

// UpdateUserPassword 修改密码
func (c *cSysUser) UpdateUserPassword(ctx context.Context, req *sysapi.UpdateUserPasswordReq) (api_v1.BoolRes, error) {
	_, err := sys_service.SysUser().UpdateUserPassword(ctx, req.UpdateUserPassword)

	if err != nil {
		return false, err
	}
	return true, nil
}
