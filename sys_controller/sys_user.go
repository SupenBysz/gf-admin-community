package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
)

// SysUser 鉴权
var SysUser = cSysUser{}

type cSysUser struct{}

// CreateUser 创建用户|信息
func (c *cSysUser) CreateUser(ctx context.Context, req *sys_api.CreateUserReq) (res api_v1.BoolRes, err error) {
	return funs.ProxyFunc3[api_v1.BoolRes, sys_model.UserInnerRegister, sys_enum.UserState, sys_enum.UserType, *sys_model.SysUserRegisterRes](ctx,
		req.UserInnerRegister, sys_enum.User.State.Normal, sys_enum.User.Type.User,
		func(ctx context.Context, data sys_model.UserInnerRegister, data1 sys_enum.UserState, data2 sys_enum.UserType) (*sys_model.SysUserRegisterRes, error) {
			return sys_service.SysUser().CreateUser(ctx, data, data1, data2)
		}, false,
		sys_enum.User.PermissionType.Create,
	)
}

// QueryUserList 获取用户|列表
func (c *cSysUser) QueryUserList(ctx context.Context, req *sys_api.QueryUserListReq) (*sys_api.UserListRes, error) {
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.ProxyFunc3[*sys_api.UserListRes](ctx,
		&req.SearchParams, unionMainId, false,
		sys_service.SysUser().QueryUserList, nil,
		sys_enum.User.PermissionType.List,
	)
}

// SetUserRoleIds 设置用户角色
func (c *cSysRole) SetUserRoleIds(ctx context.Context, req *sys_api.SetUserRoleIdsReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc2[api_v1.BoolRes](
		ctx, req.RoleIds, req.UserId,
		sys_service.SysUser().SetUserRoleIds, false,
		sys_enum.User.PermissionType.SetUserRole,
	)
}

// SetUserPermissionIds 设置用户权限
func (c *cSysUser) SetUserPermissionIds(ctx context.Context, req *sys_api.SetUserPermissionIdsReq) (api_v1.BoolRes, error) {
	return funs.ProxyFunc2[api_v1.BoolRes](
		ctx, req.Id, req.PermissionIds,
		sys_service.SysUser().SetUserPermissionIds, false,
		sys_enum.User.PermissionType.SetPermission,
	)
}

// GetUserPermissionIds 获取用户权限Ids
func (c *cSysUser) GetUserPermissionIds(ctx context.Context, req *sys_api.GetUserPermissionIdsReq) (*api_v1.Int64ArrRes, error) {
	return funs.ProxyFunc1[*api_v1.Int64ArrRes](
		ctx, req.Id,
		sys_service.SysUser().GetUserPermissionIds, nil,
		sys_enum.User.PermissionType.SetPermission,
	)
}

// ResetUserPassword 重置用户密码
func (c *cSysUser) ResetUserPassword(ctx context.Context, req *sys_api.ResetUserPasswordReq) (res api_v1.BoolRes, err error) {
	// 获取当前登录用户
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	return funs.ProxyFunc4[api_v1.BoolRes](
		ctx, req.UserId, req.Password, req.ConfirmPassword, user.SysUser,
		sys_service.SysUser().ResetUserPassword, false,
		sys_enum.User.PermissionType.ResetPassword,
	)
}
