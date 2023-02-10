package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
)

// SysUser 鉴权
var SysUser = cSysUser{}

type cSysUser struct{}

// QueryUserList 获取用户|列表
func (c *cSysUser) QueryUserList(ctx context.Context, req *sys_api.QueryUserListReq) (*sys_model.SysUserListRes, error) {
	return funs.CheckPermission(ctx,
		func() (*sys_model.SysUserListRes, error) {
			sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
			return sys_service.SysUser().QueryUserList(
				c.makeMore(ctx),
				&req.SearchParams,
				sessionUser.UnionMainId,
				false,
			)
		},
		sys_enum.User.PermissionType.List,
	)
}

// SetUserPermissionIds 设置用户权限
func (c *cSysUser) SetUserPermissionIds(ctx context.Context, req *sys_api.SetUserPermissionIdsReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := sys_service.SysUser().SetUserPermissionIds(
				ctx,
				req.Id,
				req.PermissionIds,
			)
			return ret == true, err
		},
		sys_enum.User.PermissionType.SetPermission,
	)
}

// GetUserPermissionIds 获取用户权限Ids
func (c *cSysUser) GetUserPermissionIds(ctx context.Context, req *sys_api.GetUserPermissionIdsReq) (api_v1.Int64ArrRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.Int64ArrRes, error) {
			return sys_service.SysPermission().GetPermissionsByResource(
				ctx,
				req.Id,
			)
		},
		sys_enum.User.PermissionType.SetPermission,
	)
}

// GetUserDetail 查看详情
func (c *cSysUser) GetUserDetail(ctx context.Context, req *sys_api.GetUserDetailReq) (*sys_api.UserInfoRes, error) {
	return funs.CheckPermission(ctx,
		func() (*sys_api.UserInfoRes, error) {
			ret, err := sys_service.SysUser().GetUserDetail(c.makeMore(ctx), req.Id)
			return kconv.Struct(ret, &sys_api.UserInfoRes{}), err
		},
		sys_enum.User.PermissionType.ViewMoreDetail,
	)
}

// SetUserRoles 设置用户角色
func (c *cSysUser) SetUserRoles(ctx context.Context, req *sys_api.SetUserRolesReq) (res api_v1.BoolRes, err error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
			ret, err := sys_service.SysUser().SetUserRoles(
				ctx,
				req.UserId,
				req.RoleIds,
				sessionUser.UnionMainId,
			)
			return ret == true, err
		},
		sys_enum.Role.PermissionType.SetMember,
	)
}

// ResetUserPassword 重置用户密码
func (c *cSysUser) ResetUserPassword(ctx context.Context, req *sys_api.ResetUserPasswordReq) (res api_v1.BoolRes, err error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := sys_service.SysUser().ResetUserPassword(
				ctx,
				req.Id,
				req.Password,
				req.ConfirmPassword,
			)
			return ret == true, err
		},
		sys_enum.User.PermissionType.ResetPassword,
	)
}

// SetUserState 设置用户状态
func (c *cSysUser) SetUserState(ctx context.Context, req *sys_api.SetUserStateReq) (res api_v1.BoolRes, err error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := sys_service.SysUser().SetUserState(
				ctx, req.Id, sys_enum.User.State.New(req.State, ""),
			)
			return ret == true, err
		},
		sys_enum.User.PermissionType.SetState,
	)
}

// makeMore 是否订阅附加数据
func (c *cSysUser) makeMore(ctx context.Context) context.Context {
	return funs.AttrBuilder[sys_model.SysUser, *sys_entity.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
}
