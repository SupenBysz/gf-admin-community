package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/service"
)

// SysRole 角色
var SysRole = cSysRole{}

type cSysRole struct{}

// GetRoleList 获取角色|列表
func (c *cSysRole) GetRoleList(ctx context.Context, req *sysapi.QueryRoleListReq) (*sysapi.RoleListRes, error) {
	return service.SysRole().QueryRoleList(ctx, req.SearchParams)
}

// CreateeRoleInfo 新增或保存角色|信息
func (c *cSysRole) CreateeRoleInfo(ctx context.Context, req *sysapi.CreateRoleInfoReq) (*sysapi.RoleInfoRes, error) {
	result, err := service.SysRole().Create(ctx, req.SysRole)

	return (*sysapi.RoleInfoRes)(result), err
}

// UpdateRoleInfo 更新角色信息
func (c *cSysRole) UpdateRoleInfo(ctx context.Context, req *sysapi.UpdateRoleInfoReq) (*sysapi.RoleInfoRes, error) {
	result, err := service.SysRole().Update(ctx, req.SysRole)

	return (*sysapi.RoleInfoRes)(result), err
}

// DeleteRoleInfo 删除角色
func (c *cSysRole) DeleteRoleInfo(ctx context.Context, req *sysapi.DeleteRoleInfoReq) (api_v1.BoolRes, error) {
	result, err := service.SysRole().Delete(ctx, req.Id)

	return result == true, err
}

// SetRoleForUser 设置角色用户
func (c *cSysRole) SetRoleForUser(ctx context.Context, req *sysapi.SetRoleForUserReq) (api_v1.BoolRes, error) {
	result, err := service.SysRole().SetRoleForUser(ctx, req.RoleId, req.UserId)

	return result == true, err
}

// RemoveRoleForUser 移除用户所拥有的角色
func (c *cSysRole) RemoveRoleForUser(ctx context.Context, req *sysapi.RemoveRoleForUserReq) (api_v1.BoolRes, error) {
	result, err := service.SysRole().RemoveRoleForUser(ctx, req.RoleId, req.UserId)

	return result == true, err
}

// GetRoleUserList 获取角色下的所有用户|列表
func (c *cSysRole) GetRoleUserList(ctx context.Context, req *sysapi.GetRoleUsersReq) (*sysapi.UserListRes, error) {
	data, err := service.SysRole().GetRoleUsers(ctx, req.RoleId)

	if err != nil {
		return nil, err
	}

	count := len(*data)

	return &sysapi.UserListRes{
		List: data,
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: count,
			},
			PageTotal: 1,
		},
	}, err
}

// GetUserRoleList 获取用户拥有的所有角色列表
func (c *cSysRole) GetUserRoleList(ctx context.Context, req *sysapi.GetUserRolesReq) (*sysapi.RoleListRes, error) {
	data, err := service.SysRole().GetUserRoleList(ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	count := len(*data)

	return &sysapi.RoleListRes{
		List: data,
		PaginationRes: model.PaginationRes{
			Pagination: model.Pagination{
				Page:     1,
				PageSize: count,
			},
			PageTotal: 1,
		},
	}, err
}

// SetRolePermissions 设置角色权限
func (c *cSysUser) SetRolePermissions(ctx context.Context, req *sysapi.SetRolePermissionsReq) (api_v1.BoolRes, error) {
	result, err := service.SysRole().SetRolePermissions(ctx, req.Id, req.PermissionIds)
	return result == true, err
}

// GetRolePermissionIds 获取角色权限Ids
func (c *cSysUser) GetRolePermissionIds(ctx context.Context, req *sysapi.GetRolePermissionsReq) (*api_v1.Int64ArrRes, error) {
	result, err := service.SysRole().GetRolePermissions(ctx, req.Id)
	return (*api_v1.Int64ArrRes)(&result), err
}
