package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// SysMy 公共我的模块
var SysMy = cSysMy{}

type cSysMy struct{}

// SetUsername 设置用户登陆名
func (c *cSysMy) SetUsername(ctx context.Context, req *sys_api.SetUsernameByIdReq) (api_v1.BoolRes, error) {
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

	result, err := sys_service.SysUser().SetUsername(ctx, req.NewUsername, userId)
	return result == true, err
}

// UpdateUserPassword 修改密码
func (c *cSysMy) UpdateUserPassword(ctx context.Context, req *sys_api.UpdateUserPasswordReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.User.PermissionType.ChangePassword); has != true {
		return false, err
	}

	// 获取到当前登录用户名称
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err := sys_service.SysUser().UpdateUserPassword(ctx, req.UpdateUserPassword, user.Id)

	if err != nil {
		return false, err
	}
	return true, nil
}
