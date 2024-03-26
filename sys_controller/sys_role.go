package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
)

// SysRole 角色
var SysRole = cSysRole{}

type cSysRole struct{}

// GetRoleList 获取角色|列表
func (c *cSysRole) GetRoleList(ctx context.Context, req *sys_api.QueryRoleListReq) (*sys_model.RoleListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.List); has != true {
		return nil, err
	}

	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return sys_service.SysRole().QueryRoleList(ctx, req.SearchParams, unionMainId)
}

// CreateRoleInfo 新增或保存角色|信息
func (c *cSysRole) CreateRoleInfo(ctx context.Context, req *sys_api.CreateRoleInfoReq) (*sys_api.RoleInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.Create); has != true {
		return nil, err
	}

	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId
	req.UnionMainId = unionMainId

	result, err := sys_service.SysRole().Create(ctx, req.SysRole)

	return (*sys_api.RoleInfoRes)(result), err
}

// UpdateRoleInfo 更新角色信息
func (c *cSysRole) UpdateRoleInfo(ctx context.Context, req *sys_api.UpdateRoleInfoReq) (*sys_api.RoleInfoRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.Update); has != true {
		return nil, err
	}

	result, err := sys_service.SysRole().Update(ctx, req.SysRole)

	return (*sys_api.RoleInfoRes)(result), err
}

// DeleteRoleInfo 删除角色
func (c *cSysRole) DeleteRoleInfo(ctx context.Context, req *sys_api.DeleteRoleInfoReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.Delete); has != true {
		return false, err
	}

	result, err := sys_service.SysRole().Delete(ctx, req.Id)

	return result == true, err
}

// SetRoleMember 设置角色用户
func (c *cSysRole) SetRoleMember(ctx context.Context, req *sys_api.SetRoleMemberReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.SetMember); has != true {
		return false, err
	}

	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	result, err := sys_service.SysRole().SetRoleMember(ctx, req.RoleId, req.UserIds, unionMainId)

	return result == true, err
}

// RemoveRoleMember 移除用户所拥有的角色
func (c *cSysRole) RemoveRoleMember(ctx context.Context, req *sys_api.RemoveRoleMemberReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.SetMember); has != true {
		return false, err
	}

	result, err := sys_service.SysRole().RemoveRoleMember(ctx, req.RoleId, req.UserIds)

	return result == true, err
}

// GetRoleMemberIds 获取角色下的所有用户Ids|列表
func (c *cSysRole) GetRoleMemberIds(ctx context.Context, req *sys_api.GetRoleMemberIdsReq) (api_v1.Int64ArrRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.ViewMember); has != true {
		return nil, err
	}
	// 获取当前登录用户的UnionMainId
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return sys_service.SysRole().GetRoleMemberIds(ctx, req.RoleId, unionMainId)
}

// GetRoleMemberList 获取角色下的所有用户|列表
func (c *cSysRole) GetRoleMemberList(ctx context.Context, req *sys_api.GetRoleMemberReq) (*sys_api.UserListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.ViewMember); has != true {
		return nil, err
	}

	// 获取当前登录用户的UnionMainId
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	data, err := sys_service.SysRole().GetRoleMemberList(ctx, req.RoleId, unionMainId)

	if err != nil {
		return nil, err
	}

	count := len(data)

	return &sys_api.UserListRes{
		Records: data,
		PaginationRes: base_model.PaginationRes{
			Pagination: base_model.Pagination{
				PageNum:  1,
				PageSize: count,
			},
			PageTotal: 1,
			Total:     gconv.Int64(count),
		},
	}, err
}

// GetRoleByUserIdList 获取用户ID获取所有关联角色
func (c *cSysRole) GetRoleByUserIdList(ctx context.Context, req *sys_api.GetRoleByUserIdListReq) (*sys_model.RoleListRes, error) {
	data, err := sys_service.SysRole().GetRoleListByUserId(ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	count := len(data)

	return &sys_model.RoleListRes{
		Records: data,
		PaginationRes: base_model.PaginationRes{
			Pagination: base_model.Pagination{
				PageNum:  1,
				PageSize: count,
			},
			PageTotal: 1,
			Total:     gconv.Int64(count),
		},
	}, err
}

// SetRolePermissions 设置角色权限
func (c *cSysRole) SetRolePermissions(ctx context.Context, req *sys_api.SetRolePermissionsReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Role.PermissionType.SetPermission); has != true {
		return false, err
	}

	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	result, err := sys_service.SysRole().SetRolePermissions(ctx, req.Id, req.PermissionIds, unionMainId)
	return result == true, err
}

// GetRolePermissionIds 获取角色权限Ids
func (c *cSysRole) GetRolePermissionIds(ctx context.Context, req *sys_api.GetRolePermissionsIdsReq) (*api_v1.Int64ArrRes, error) {
	result, err := sys_service.SysPermission().GetPermissionsByResource(ctx, gconv.String(req.Id))
	return (*api_v1.Int64ArrRes)(&result), err
}
