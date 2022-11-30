package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	userState "github.com/SupenBysz/gf-admin-community/model/enum/user_state"
	userType "github.com/SupenBysz/gf-admin-community/model/enum/user_type"
	"github.com/SupenBysz/gf-admin-community/service"
)

// SysUser 鉴权
var SysUser = cSysUser{}

type cSysUser struct{}

// CreateUser 创建用户|信息
func (c *cSysUser) CreateUser(ctx context.Context, req *sysapi.CreateUserReq) (res api_v1.BoolRes, err error) {
	_, err = service.SysUser().CreateUser(ctx, req.UserInnerRegister, userState.Normal, userType.User)

	return err == nil, err
}

// QueryUserList 获取用户|列表
func (c *cSysUser) QueryUserList(ctx context.Context, req *sysapi.QueryUserListReq) (*sysapi.UserListRes, error) {
	data, err := service.SysUser().QueryUserList(ctx, &req.SearchParams, false)
	if err != nil {
		return nil, err
	}
	return (*sysapi.UserListRes)(data), err
}

// SetUserRoleIds 设置用户角色
func (c *cSysRole) SetUserRoleIds(ctx context.Context, req *sysapi.SetUserRoleIdsReq) (api_v1.BoolRes, error) {
	result, err := service.SysUser().SetUserRoleIds(ctx, req.RoleIds, req.UserId)

	return result == true, err
}

// SetUserPermissionIds 设置用户权限
func (c *cSysUser) SetUserPermissionIds(ctx context.Context, req *sysapi.SetUserPermissionIdsReq) (api_v1.BoolRes, error) {
	result, err := service.SysUser().SetUserPermissionIds(ctx, req.Id, req.PermissionIds)
	return result == true, err
}

// GetUserPermissionIds 获取用户权限Ids
func (c *cSysUser) GetUserPermissionIds(ctx context.Context, req *sysapi.GetUserPermissionIdsReq) (*api_v1.Int64ArrRes, error) {
	result, err := service.SysUser().GetUserPermissionIds(ctx, req.Id)
	return (*api_v1.Int64ArrRes)(&result), err
}

// SetUsername 设置用户登陆名
func (c *cSysUser) SetUsername(ctx context.Context, req *sysapi.SetUsernameByIdReq) (api_v1.BoolRes, error) {
	result, err := service.SysUser().SetUsername(ctx, req.NewUsername)
	return result == true, err
}
