package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
)

// SysRole 角色
var SysRole = cSysRole{}

type cSysRole struct{}

// GetRoleList 获取角色|列表
func (c *cSysRole) GetRoleList(ctx context.Context, req *sys_api.QueryRoleListReq) (*sys_model.RoleListRes, error) {
	return sys_service.SysRole().QueryRoleList(ctx, req.SearchParams)
}

// CreateRoleInfo 新增或保存角色|信息
func (c *cSysRole) CreateRoleInfo(ctx context.Context, req *sys_api.CreateRoleInfoReq) (*sys_api.RoleInfoRes, error) {
	result, err := sys_service.SysRole().Create(ctx, req.SysRole)

	return (*sys_api.RoleInfoRes)(result), err
}

// UpdateRoleInfo 更新角色信息
func (c *cSysRole) UpdateRoleInfo(ctx context.Context, req *sys_api.UpdateRoleInfoReq) (*sys_api.RoleInfoRes, error) {
	result, err := sys_service.SysRole().Update(ctx, req.SysRole)

	return (*sys_api.RoleInfoRes)(result), err
}

// DeleteRoleInfo 删除角色
func (c *cSysRole) DeleteRoleInfo(ctx context.Context, req *sys_api.DeleteRoleInfoReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysRole().Delete(ctx, req.Id)

	return result == true, err
}

// SetRoleForUser 设置角色用户
func (c *cSysRole) SetRoleForUser(ctx context.Context, req *sys_api.SetRoleForUserReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysRole().SetRoleForUser(ctx, req.RoleId, req.UserId)

	return result == true, err
}

// RemoveRoleForUser 移除用户所拥有的角色
func (c *cSysRole) RemoveRoleForUser(ctx context.Context, req *sys_api.RemoveRoleForUserReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysRole().RemoveRoleForUser(ctx, req.RoleId, req.UserId)

	return result == true, err
}

// GetRoleUserList 获取角色下的所有用户|列表
func (c *cSysRole) GetRoleUserList(ctx context.Context, req *sys_api.GetRoleUsersReq) (*sys_api.UserListRes, error) {
	data, err := sys_service.SysRole().GetRoleUsers(ctx, req.RoleId)

	if err != nil {
		return nil, err
	}

	count := len(*data)

	return &sys_api.UserListRes{
		List: data,
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: count,
			},
			PageTotal: 1,
			Total:     gconv.Int64(count),
		},
	}, err
}

// GetUserRoleList 获取用户拥有的所有角色列表
func (c *cSysRole) GetUserRoleList(ctx context.Context, req *sys_api.GetUserRolesReq) (*sys_model.RoleListRes, error) {
	data, err := sys_service.SysRole().GetUserRoleList(ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	count := len(*data)

	return &sys_model.RoleListRes{
		List: data,
		PaginationRes: sys_model.PaginationRes{
			Pagination: sys_model.Pagination{
				Page:     1,
				PageSize: count,
			},
			PageTotal: 1,
			Total:     gconv.Int64(count),
		},
	}, err
}

// SetRolePermissions 设置角色权限
func (c *cSysRole) SetRolePermissions(ctx context.Context, req *sys_api.SetRolePermissionsReq) (api_v1.BoolRes, error) {
	result, err := sys_service.SysRole().SetRolePermissions(ctx, req.Id, req.PermissionIds)
	return result == true, err
}

// GetRolePermissionIds 获取角色权限Ids
func (c *cSysRole) GetRolePermissionIds(ctx context.Context, req *sys_api.GetRolePermissionsReq) (*api_v1.Int64ArrRes, error) {
	result, err := sys_service.SysRole().GetRolePermissions(ctx, req.Id)
	return (*api_v1.Int64ArrRes)(&result), err
}
