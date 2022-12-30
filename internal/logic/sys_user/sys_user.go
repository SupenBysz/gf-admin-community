package user

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/en_crypto"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/masker"
)

type hookInfo sys_model.KeyValueT[int64, sys_model.UserHookInfo]

type sSysUser struct {
	CacheDuration time.Duration
	CachePrefix   string
	hookArr       []hookInfo
}

func init() {
	sys_service.RegisterSysUser(New())
}

func New() *sSysUser {
	return &sSysUser{
		CacheDuration: time.Hour,
		CachePrefix:   sys_dao.SysUser.Table() + "_",
		hookArr:       make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sSysUser) InstallHook(event sys_enum.UserEvent, hookFunc sys_model.UserHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_model.UserHookInfo{Key: event, Value: hookFunc}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sSysUser) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除所有Hook
func (s *sSysUser) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// QueryUserList 获取用户列表
func (s *sSysUser) QueryUserList(ctx context.Context, info *sys_model.SearchParams, isExport bool) (response *sys_model.SysUserRes, err error) {
	if info != nil {
		newFields := make([]sys_model.FilterInfo, 0)

		for _, field := range info.Filter {
			if field.Field != sys_dao.SysUser.Columns().Type {
				newFields = append(newFields, field)
			}
		}
	}

	result, err := daoctl.Query[sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), info, isExport)

	newList := make([]sys_model.SysUser, 0)
	if result != nil && result.List != nil && len(*result.List) > 0 {
		for _, user := range *result.List {
			user.RoleNames = make([]string, 0)
			roleIds, err := sys_service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(user.Id), sys_consts.CasbinDomain)

			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}

			if len(roleIds) > 0 {
				roles, err := sys_service.SysRole().QueryRoleList(ctx, sys_model.SearchParams{
					Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
						Field:     sys_dao.SysRole.Columns().Id,
						Where:     "in",
						IsOrWhere: false,
						Value:     roleIds,
					}),
					Pagination: sys_model.Pagination{},
				}, sys_service.BizCtx().Get(ctx).ClaimsUser.UnionMainId)
				if err == nil && len(*roles.List) > 0 {
					for _, role := range *roles.List {
						user.RoleNames = append(user.RoleNames, role.Name)
					}
				}
			}
			user.Password = ""
			newList = append(newList, user)
		}
	}

	if newList != nil {
		result.List = &newList
	}

	return (*sys_model.SysUserRes)(result), err
}

// SetUserRoleIds 设置用户角色
func (s *sSysUser) SetUserRoleIds(ctx context.Context, roleIds []int64, userId int64) (bool, error) {
	for _, roleId := range roleIds {
		roleInfo := sys_entity.SysRole{}
		err := sys_dao.SysRole.Ctx(ctx).Where(sys_do.SysRole{Id: roleId}).Scan(&roleInfo)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", sys_dao.SysUser.Table())
		}

		userInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysUser.Table())
		}

		result, err := sys_service.Casbin().AddRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)

		if result == false || err != nil {
			return result, err
		}
	}

	return true, nil
}

// CreateUser 创建用户
func (s *sSysUser) CreateUser(ctx context.Context, info sys_model.UserInnerRegister, userState sys_enum.UserState, userType sys_enum.UserType, customId ...int64) (*sys_model.SysUserRegisterRes, error) {
	count, _ := sys_dao.SysUser.Ctx(ctx).Unscoped().Count(sys_dao.SysUser.Columns().Username, info.Username)
	if count > 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名已经存在"), "", sys_dao.SysUser.Table())
	}

	data := sys_entity.SysUser{
		Id:        idgen.NextId(),
		Username:  info.Username,
		Password:  info.Password,
		Mobile:    info.Mobile,
		State:     userState.Code(),
		Type:      userType.Code(),
		CreatedAt: gtime.Now(),
	}

	if len(customId) > 0 && customId[0] > 0 {
		data.Id = customId[0]
	}

	pwdHash, err := en_crypto.PwdHash(info.Password, gconv.String(data.Id))

	// 密码赋值
	data.Password = pwdHash

	result := sys_model.SysUserRegisterRes{
		UserInfo:     data,
		RoleInfoList: make([]sys_entity.SysRole, 0),
	}

	err = sys_dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 创建前
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.BeforeCreate.Code() == sys_enum.User.Event.BeforeCreate.Code() {
					hook.Value.Value(ctx, sys_enum.User.Event.BeforeCreate, data)
				}
			}
		})

		_, err := sys_dao.SysUser.Ctx(ctx).OmitNilData().Data(data).Insert()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "账号注册失败", sys_dao.SysUser.Table())
		}

		if len(info.RoleIds) > 0 {
			ret, err := s.SetUserRoleIds(ctx, info.RoleIds, data.Id)
			if ret != true || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "角色设置失败！"+err.Error(), sys_dao.SysUser.Table())
			}

			err = sys_dao.SysRole.Ctx(ctx).WhereIn(sys_dao.SysRole.Columns().Id, info.RoleIds).Scan(&result.RoleInfoList)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "查询角色信息失败！", sys_dao.SysUser.Table())
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	// 建后
	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.User.Event.AfterCreate.Code() == sys_enum.User.Event.AfterCreate.Code() {
				hook.Value.Value(ctx, sys_enum.User.Event.AfterCreate, data)
			}
		}
	})
	return &result, nil
}

// GetSysUserByUsername 根据用户名获取用户UID
func (s *sSysUser) GetSysUserByUsername(ctx context.Context, username string) (*sys_entity.SysUser, error) {
	data := &sys_entity.SysUser{}
	err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Username: username}).Scan(data)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败", sys_dao.SysUser.Table())
	}

	return data, nil
}

// HasSysUserByUsername 判断用户名是否存在
func (s *sSysUser) HasSysUserByUsername(ctx context.Context, username string) bool {
	count, _ := sys_dao.SysUser.Ctx(ctx).Count(sys_do.SysUser{Username: username})
	return count > 0
}

// GetSysUserById 根据用户ID获取用户信息
func (s *sSysUser) GetSysUserById(ctx context.Context, userId int64) (*sys_entity.SysUser, error) {
	data := &sys_entity.SysUser{}
	err := sys_dao.SysUser.Ctx(ctx).Scan(data, sys_do.SysUser{Id: userId})

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败", sys_dao.SysUser.Table())
	}

	data.Password = masker.MaskString(data.Password, masker.Password)
	return data, nil
}

// SetUserPermissionIds 设置用户权限
func (s *sSysUser) SetUserPermissionIds(ctx context.Context, userId int64, permissionIds []int64) (bool, error) {
	err := sys_dao.SysCasbin.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		{
			// 先清除roleId所有权限
			_, err := sys_service.Casbin().DeletePermissionsForUser(gconv.String(userId))

			if len(permissionIds) <= 0 {
				return err
			}
		}

		// 重新赋予roleId新的权限清单
		for _, item := range permissionIds {
			ret, err := sys_service.Casbin().Enforcer().AddPermissionForUser(gconv.String(userId), sys_consts.CasbinDomain, gconv.String(item), "allow")
			if err != nil || ret == false {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设置用户权限失败", sys_dao.SysUser.Table())
	}

	return true, nil
}

// GetUserPermissionIds 获取用户权限，返回权限Id数组
func (s *sSysUser) GetUserPermissionIds(ctx context.Context, userId int64) ([]int64, error) {
	result, err := sys_service.Casbin().Enforcer().GetImplicitPermissionsForUser(gconv.String(userId), sys_consts.CasbinDomain)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户权限查询失败"), "", sys_dao.SysUser.Table())
	}

	permissionIds := garray.NewFrom(g.Slice{})

	for _, items := range result {
		if len(items) >= 3 {
			permissionIds.Append(items[2])
		}
	}

	return gconv.Int64s(permissionIds.Unique().Slice()), nil
}

// SetUsername 修改自己的账号登陆名称
func (s *sSysUser) SetUsername(ctx context.Context, newUsername string, userId int64) (bool, error) {
	result, err := sys_dao.SysUser.Ctx(ctx).
		Where(sys_do.SysUser{Id: userId}).
		Update(sys_do.SysUser{Username: newUsername})

	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}

// UpdateUserPassword 修改用户登录密码
func (s *sSysUser) UpdateUserPassword(ctx context.Context, info sys_model.UpdateUserPassword, userId int64) (bool, error) {
	// 查询到用户信息
	sysUserInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户不存在")
	}

	// 判断输入的两次密码是否相同
	if info.Password != info.ConfirmPassword {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "两次输入的密码不一致，修改失败")
	}

	{
		// 传入用户输入的原始密码，进行hash，看是否和数据库中原始密码一致
		hash1, _ := en_crypto.PwdHash(info.OldPassword, gconv.String(sysUserInfo.Id))
		if sysUserInfo.Password != hash1 {
			return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "原密码输入错误，修改失败")
		}
	}

	{
		// 处理hook订阅
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.ChangePassword.Code() == sys_enum.User.Event.ChangePassword.Code() {
					// 调用hook
					err = hook.Value.Value(ctx, sys_enum.User.Event.ChangePassword, *sysUserInfo)
					if err != nil {
						break
					}
				}
			}
		})
		if err != nil {
			return false, err
		}
	}

	pwdHash, err := en_crypto.PwdHash(info.Password, gconv.String(sysUserInfo.Id))

	_, err = sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: sysUserInfo.Id}).Update(sys_do.SysUser{Password: pwdHash})

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码修改失败")
	}

	return true, nil
}

// ResetUserPassword 重置用户密码 (超级管理员无需验证验证，XX商管理员重置员工密码无需验证)
func (s *sSysUser) ResetUserPassword(ctx context.Context, userId int64, password string, confirmPassword string, userInfo sys_entity.SysUser) (bool, error) {
	// hook判断当前登录身份是否可以重置密码
	{
		var err error
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.ResetPassword.Code() == sys_enum.User.Event.ResetPassword.Code() {
					err = hook.Value.Value(ctx, sys_enum.User.Event.ResetPassword, userInfo)
					if err != nil {
						break
					}
				}
			}
		})

		if err != nil {
			return false, err
		}
	}

	// 生成密码，重置密码
	{
		if password != confirmPassword {
			return false, gerror.NewCode(gcode.CodeValidationFailed, "两次密码不一致，请重新输入")
		}
		// 取盐
		salt := gconv.String(userId)

		// 加密
		pwdHash, _ := en_crypto.PwdHash(password, salt)

		result, err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: userId}).Update(sys_do.SysUser{Password: pwdHash})

		// 受影响的行数
		count, _ := result.RowsAffected()

		if err != nil || count != 1 {
			return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "重置密码失败")
		}

	}

	return true, nil
}
