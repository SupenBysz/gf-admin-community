package sys_role

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sysapi"
	"github.com/SupenBysz/gf-admin-community/internal/consts"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/container/garray"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type sSysRole struct {
}

func init() {
	service.RegisterSysRole(New())
}

// New Auth 验证码管理服务
func New() *sSysRole {
	return &sSysRole{}
}

// QueryRoleList 获取角色列表
func (s *sSysRole) QueryRoleList(ctx context.Context, info model.SearchFilter) (*sysapi.RoleListRes, error) {
	result, err := daoctl.Query[entity.SysRole](dao.SysRole.Ctx(ctx), &info, false)

	return (*sysapi.RoleListRes)(result), err
}

// Create 创建角色信息
func (s *sSysRole) Create(ctx context.Context, info model.SysRole) (*entity.SysRole, error) {
	info.Id = 0
	return s.Save(ctx, info)
}

// Update 更新角色信息
func (s *sSysRole) Update(ctx context.Context, info model.SysRole) (*entity.SysRole, error) {
	if info.Id <= 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "ID参数错误"), "", dao.SysRole.Table())
	}
	return s.Save(ctx, info)
}

// Save 新增或保存角色信息
func (s *sSysRole) Save(ctx context.Context, info model.SysRole) (*entity.SysRole, error) {
	roleInfo := entity.SysRole{
		Id:          info.Id,
		Name:        info.Name,
		Description: info.Description,
		UpdatedAt:   gtime.Now(),
	}

	err := dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if roleInfo.Id == 0 {
			roleInfo.Id = idgen.NextId()
			count, err := dao.SysRole.Ctx(ctx).WhereOr(do.SysRole{Name: roleInfo.Name}).Count()

			if err != nil {
				return service.SysLogs().ErrorSimple(ctx, err, "创建角色失败", dao.SysRole.Table())
			}

			if count > 0 {
				return service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "角色名称已经存在"), "", dao.SysRole.Table())
			}

			roleInfo.CreatedAt = gtime.Now()

			_, err = dao.SysRole.Ctx(ctx).Insert(roleInfo)

			if err != nil {
				return err
			}

			result, err := service.Casbin().AddRoleForUserInDomain(gconv.String(roleInfo.Id), consts.CasbinSuperRole, consts.CasbinDomain)

			if !result || err != nil {
				return err
			}
		} else {
			_, err := dao.SysRole.Ctx(ctx).OmitEmpty().Where(do.SysRole{Id: roleInfo.Id}).Update(do.SysRole{
				Name:        roleInfo.Name,
				Description: roleInfo.Description,
				UpdatedAt:   roleInfo.UpdatedAt,
			})
			if err != nil {
				return service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "保存角色失败"), "", dao.SysRole.Table())
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = dao.SysRole.Ctx(ctx).Where(do.SysRole{Id: roleInfo.Id}).Scan(&roleInfo)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "查询角色信息失败", dao.SysRole.Table())
	}

	return &roleInfo, nil
}

// Delete 删除角色信息
func (s *sSysRole) Delete(ctx context.Context, roleId int64) (bool, error) {
	info := &entity.SysRole{}
	err := dao.SysRole.Ctx(ctx).Where(do.SysRole{Id: roleId}).Scan(&info)

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "删除角色失败", dao.SysRole.Table())
	}

	userIds, err := service.Casbin().Enforcer().GetRoleManager().GetUsers(gconv.String(roleId), consts.CasbinDomain)

	if len(userIds) > 0 {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除角色失败，请先删除角色下的用户"), "", dao.SysRole.Table())
	}

	if info.Id == 0 {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "删除角色ID不存在"), "", dao.SysRole.Table())
	}

	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.SysRole.Ctx(ctx).Delete(do.SysRole{Id: roleId})

		result, err := service.Casbin().DeleteRoleForUserInDomain(gconv.String(info.Id), consts.CasbinSuperRole, consts.CasbinDomain)

		if !result || err != nil {
			return service.SysLogs().ErrorSimple(ctx, err, "删除角色失败", dao.SysRole.Table())
		}
		return nil
	})

	return err == nil, err
}

// SetRoleForUser 设置角色用户
func (s *sSysRole) SetRoleForUser(ctx context.Context, roleId int64, userId int64) (bool, error) {
	roleInfo := entity.SysRole{}
	err := dao.SysRole.Ctx(ctx).Where(do.SysRole{Id: roleId}).Scan(&roleInfo)

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", dao.SysRole.Table())
	}

	userInfo, err := service.SysUser().GetSysUserById(ctx, userId)

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", dao.SysRole.Table())
	}

	return service.Casbin().AddRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), consts.CasbinDomain)
}

// RemoveRoleForUser 移除角色中的用户
func (s *sSysRole) RemoveRoleForUser(ctx context.Context, roleId int64, userId int64) (bool, error) {
	roleInfo := entity.SysRole{}
	err := dao.SysRole.Ctx(ctx).Where(do.SysRole{Id: roleId}).Scan(&roleInfo)

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", dao.SysRole.Table())
	}

	userInfo, err := service.SysUser().GetSysUserById(ctx, userId)

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", dao.SysRole.Table())
	}

	return service.Casbin().DeleteRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), consts.CasbinDomain)
}

// GetRoleUsers 获取角色下的所有用户
func (s *sSysRole) GetRoleUsers(ctx context.Context, roleId int64) (*[]model.SysUser, error) {
	roleInfo := entity.SysRole{}
	err := dao.SysRole.Ctx(ctx).Where(do.SysRole{Id: roleId}).Scan(&roleInfo)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", dao.SysRole.Table())
	}

	if roleInfo.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "查询角色ID信息不存在")
	}

	userIds, err := service.Casbin().Enforcer().GetRoleManager().GetUsers(gconv.String(roleId), consts.CasbinDomain)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", dao.SysRole.Table())
	}

	userInfoArr := make([]model.SysUser, 0)

	err = dao.SysUser.Ctx(ctx).WhereIn(dao.SysUser.Columns().Id, userIds).Scan(&userInfoArr)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "查询用户信息失败", dao.SysRole.Table())
	}

	result := make([]model.SysUser, 0)
	// 移除密码信息
	for _, user := range userInfoArr {
		user.Password = ""
		user.RoleNames = make([]string, 0)

		roles, err := service.SysRole().GetUserRoleList(ctx, user.Id)
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
func (s *sSysRole) GetUserRoleList(ctx context.Context, userId int64) (*[]entity.SysRole, error) {
	userInfo := entity.SysUser{}

	err := dao.SysUser.Ctx(ctx).Where(do.SysUser{Id: userId}).Scan(&userInfo)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", dao.SysRole.Table())
	}

	if userInfo.Id <= 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "查询用户ID信息不存在"), "", dao.SysRole.Table())
	}

	roleIds, err := service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(userId), consts.CasbinDomain)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", dao.SysRole.Table())
	}

	roleInfoArr := make([]entity.SysRole, 0)

	err = dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, roleIds).Scan(&roleInfoArr)

	return &roleInfoArr, nil
}

// SetRolePermissions 设置角色权限
func (s *sSysRole) SetRolePermissions(ctx context.Context, roleId int64, permissionIds []int64) (bool, error) {

	err := dao.SysCasbin.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		{
			// 先清除roleId所有权限
			_, err := service.Casbin().DeletePermissionsForUser(gconv.String(roleId))

			if len(permissionIds) <= 0 {
				return err
			}
		}

		// 重新赋予roleId新的权限清单
		for _, item := range permissionIds {
			ret, err := service.Casbin().Enforcer().AddPermissionForUser(gconv.String(roleId), consts.CasbinDomain, gconv.String(item), "allow")
			if err != nil || ret == false {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "设置角色权限失败", dao.SysRole.Table())
	}

	return true, nil
}

// GetRolePermissions 获取角色权限Ids，返回权限Id数组
func (s *sSysRole) GetRolePermissions(ctx context.Context, roleId int64) ([]int64, error) {
	result, err := service.Casbin().Enforcer().GetImplicitPermissionsForUser(gconv.String(roleId), consts.CasbinDomain)
	if err != nil {
		return make([]int64, 0), service.SysLogs().ErrorSimple(ctx, err, "角色权限查询失败", dao.SysRole.Table())
	}

	permissionIds := garray.NewFrom(g.Slice{})

	for _, items := range result {
		if len(items) >= 3 {
			permissionIds.Append(items[2])
		}
	}

	return gconv.Int64s(permissionIds.Unique().Slice()), nil
}
