package sys_role

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type sSysRole struct {
}

func init() {
	sys_service.RegisterSysRole(New())
}

// New Auth 验证码管理服务
func New() *sSysRole {
	return &sSysRole{}
}

// QueryRoleList 获取角色列表
func (s *sSysRole) QueryRoleList(ctx context.Context, info sys_model.SearchParams) (*sys_model.RoleListRes, error) {
	result, err := daoctl.Query[sys_entity.SysRole](sys_dao.SysRole.Ctx(ctx), &info, false)

	return (*sys_model.RoleListRes)(result), err
}

// Create 创建角色信息
func (s *sSysRole) Create(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error) {
	info.Id = 0
	return s.Save(ctx, info)
}

// Update 更新角色信息
func (s *sSysRole) Update(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "ID参数错误"), "", sys_dao.SysRole.Table())
	}
	return s.Save(ctx, info)
}

// Save 新增或保存角色信息
func (s *sSysRole) Save(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error) {
	roleInfo := sys_entity.SysRole{
		Id:          info.Id,
		Name:        info.Name,
		Description: info.Description,
		UnionMainId: info.UnionMainId,
		IsSys:       info.IsSys,
		UpdatedAt:   gtime.Now(),
	}

	err := sys_dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if roleInfo.Id == 0 {
			roleInfo.Id = idgen.NextId()
			// count, err := sys_dao.SysRole.Ctx(ctx).WhereOr(sys_do.SysRole{Name: roleInfo.Name}).Count()

			count, err := sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Name: info.Name, UnionMainId: info.UnionMainId}).Count()
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "创建角色失败", sys_dao.SysRole.Table())
			}

			// 通过Union_main_id去判断
			if count > 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "角色名称在该域已经存在"), "", sys_dao.SysRole.Table())
			}

			roleInfo.CreatedAt = gtime.Now()

			_, err = sys_dao.SysRole.Ctx(ctx).Insert(roleInfo)

			if err != nil {
				return err
			}

			result, err := sys_service.Casbin().AddRoleForUserInDomain(gconv.String(roleInfo.Id), sys_consts.CasbinSuperRole, sys_consts.CasbinDomain)

			if !result || err != nil {
				return err
			}
		} else {
			_, err := sys_dao.SysRole.Ctx(ctx).OmitEmpty().Where(sys_do.SysRole{Id: roleInfo.Id}).Update(sys_do.SysRole{
				Name:        roleInfo.Name,
				Description: roleInfo.Description,
				UpdatedAt:   roleInfo.UpdatedAt,
			})
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "保存角色失败"), "", sys_dao.SysRole.Table())
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Id: roleInfo.Id}).Scan(&roleInfo)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询角色信息失败", sys_dao.SysRole.Table())
	}

	return &roleInfo, nil
}

// Delete 删除角色信息
func (s *sSysRole) Delete(ctx context.Context, roleId int64) (bool, error) {
	info := &sys_entity.SysRole{}
	err := sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Id: roleId}).Scan(&info)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除角色失败", sys_dao.SysRole.Table())
	}

	userIds, err := sys_service.Casbin().Enforcer().GetRoleManager().GetUsers(gconv.String(roleId), sys_consts.CasbinDomain)

	if len(userIds) > 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除角色失败，请先删除角色下的用户"), "", sys_dao.SysRole.Table())
	}

	if info.Id == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除角色ID不存在"), "", sys_dao.SysRole.Table())
	}

	err = sys_dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = sys_dao.SysRole.Ctx(ctx).Delete(sys_do.SysRole{Id: roleId})

		result, err := sys_service.Casbin().DeleteRoleForUserInDomain(gconv.String(info.Id), sys_consts.CasbinSuperRole, sys_consts.CasbinDomain)

		if !result || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "删除角色失败", sys_dao.SysRole.Table())
		}
		return nil
	})

	return err == nil, err
}

// SetRoleForUser 设置角色用户
func (s *sSysRole) SetRoleForUser(ctx context.Context, roleId int64, userId int64) (bool, error) {
	roleInfo := sys_entity.SysRole{}
	err := sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Id: roleId}).Scan(&roleInfo)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", sys_dao.SysRole.Table())
	}

	userInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysRole.Table())
	}

	return sys_service.Casbin().AddRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)
}

// RemoveRoleForUser 移除角色中的用户
func (s *sSysRole) RemoveRoleForUser(ctx context.Context, roleId int64, userId int64) (bool, error) {
	roleInfo := sys_entity.SysRole{}
	err := sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Id: roleId}).Scan(&roleInfo)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", sys_dao.SysRole.Table())
	}

	userInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysRole.Table())
	}

	return sys_service.Casbin().DeleteRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)
}

// GetRoleUsers 获取角色下的所有用户
func (s *sSysRole) GetRoleUsers(ctx context.Context, roleId int64) (*[]sys_model.SysUser, error) {
	roleInfo := sys_entity.SysRole{}
	err := sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Id: roleId}).Scan(&roleInfo)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", sys_dao.SysRole.Table())
	}

	if roleInfo.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "查询角色ID信息不存在")
	}

	userIds, err := sys_service.Casbin().Enforcer().GetRoleManager().GetUsers(gconv.String(roleId), sys_consts.CasbinDomain)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysRole.Table())
	}

	userInfoArr := make([]sys_model.SysUser, 0)

	err = sys_dao.SysUser.Ctx(ctx).WhereIn(sys_dao.SysUser.Columns().Id, userIds).Scan(&userInfoArr)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询用户信息失败", sys_dao.SysRole.Table())
	}

	result := make([]sys_model.SysUser, 0)
	// 移除密码信息
	for _, user := range userInfoArr {
		user.Password = ""
		user.RoleNames = make([]string, 0)

		roles, err := sys_service.SysRole().GetUserRoleList(ctx, user.Id)
		if err == nil && len(*roles) > 0 {
			for _, role := range *roles {
				user.RoleNames = append(user.RoleNames, role.Name)
			}
		}
		result = append(result, user)
	}

	return &result, nil
}

// GetUserRoleList 获取用户拥有的所有角色
func (s *sSysRole) GetUserRoleList(ctx context.Context, userId int64) (*[]sys_entity.SysRole, error) {
	userInfo := sys_entity.SysUser{}

	err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: userId}).Scan(&userInfo)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysRole.Table())
	}

	if userInfo.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "查询用户ID信息不存在"), "", sys_dao.SysRole.Table())
	}

	roleIds, err := sys_service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(userId), sys_consts.CasbinDomain)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysRole.Table())
	}

	roleInfoArr := make([]sys_entity.SysRole, 0)

	err = sys_dao.SysRole.Ctx(ctx).WhereIn(sys_dao.SysRole.Columns().Id, roleIds).Scan(&roleInfoArr)

	return &roleInfoArr, nil
}

// SetRolePermissions 设置角色权限
func (s *sSysRole) SetRolePermissions(ctx context.Context, roleId int64, permissionIds []int64) (bool, error) {

	err := sys_dao.SysCasbin.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		{
			// 先清除roleId所有权限
			_, err := sys_service.Casbin().DeletePermissionsForUser(gconv.String(roleId))

			if len(permissionIds) <= 0 {
				return err
			}
		}

		// 重新赋予roleId新的权限清单
		for _, item := range permissionIds {
			ret, err := sys_service.Casbin().Enforcer().AddPermissionForUser(gconv.String(roleId), sys_consts.CasbinDomain, gconv.String(item), "allow")
			if err != nil || ret == false {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设置角色权限失败", sys_dao.SysRole.Table())
	}

	return true, nil
}

// GetRolePermissions 获取角色权限Ids，返回权限Id数组
func (s *sSysRole) GetRolePermissions(ctx context.Context, roleId int64) ([]int64, error) {
	result, err := sys_service.Casbin().Enforcer().GetImplicitPermissionsForUser(gconv.String(roleId), sys_consts.CasbinDomain)
	if err != nil {
		return make([]int64, 0), sys_service.SysLogs().ErrorSimple(ctx, err, "角色权限查询失败", sys_dao.SysRole.Table())
	}

	permissionIds := garray.NewFrom(g.Slice{})

	for _, items := range result {
		if len(items) >= 3 {
			permissionIds.Append(items[2])
		}
	}

	return gconv.Int64s(permissionIds.Unique().Slice()), nil
}
