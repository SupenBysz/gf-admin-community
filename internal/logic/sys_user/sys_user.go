package user

import (
	"context"
	kyUser "github.com/SupenBysz/gf-admin-community/model/enum/user"
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

	"github.com/SupenBysz/gf-admin-community/internal/consts"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/masker"
)

type hookInfo model.KeyValueT[int64, model.UserHookInfo]

type sSysUser struct {
	CacheDuration time.Duration
	CachePrefix   string
	hookArr       []hookInfo
}

func init() {
	service.RegisterSysUser(New())
}

func New() *sSysUser {
	return &sSysUser{
		CacheDuration: time.Hour,
		CachePrefix:   dao.SysUser.Table() + "_",
		hookArr:       make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sSysUser) InstallHook(event kyUser.EventEnum, hookFunc model.UserHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: model.UserHookInfo{Key: event, Value: hookFunc}}
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
func (s *sSysUser) QueryUserList(ctx context.Context, info *model.SearchParams, isExport bool) (response *model.SysUserRes, err error) {
	if info != nil {
		newFields := make([]model.FilterInfo, 0)

		newFields = append(newFields, model.FilterInfo{
			Field: dao.SysUser.Columns().Type, // type
			Where: "=",
			Value: consts.Global.UserDefaultType,
		})

		for _, field := range info.Filter {
			if field.Field != dao.SysUser.Columns().Type {
				newFields = append(newFields, field)
			}
		}
	}

	result, err := daoctl.Query[model.SysUser](dao.SysUser.Ctx(ctx), info, isExport)

	newList := make([]model.SysUser, 0)
	if result != nil && result.List != nil && len(*result.List) > 0 {
		for _, user := range *result.List {
			user.RoleNames = make([]string, 0)
			roles, err := service.SysRole().GetUserRoleList(ctx, user.Id)
			if err == nil && len(*roles) > 0 {
				for _, role := range *roles {
					user.RoleNames = append(user.RoleNames, role.Name)
				}
			}
			user.Password = ""
			newList = append(newList, user)
		}
	}

	if newList != nil {
		result.List = &newList
	}

	return (*model.SysUserRes)(result), err
}

// SetUserRoleIds 设置用户角色
func (s *sSysUser) SetUserRoleIds(ctx context.Context, roleIds []int64, userId int64) (bool, error) {
	for _, roleId := range roleIds {
		roleInfo := entity.SysRole{}
		err := dao.SysRole.Ctx(ctx).Where(do.SysRole{Id: roleId}).Scan(&roleInfo)

		if err != nil {
			return false, service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", dao.SysUser.Table())
		}

		userInfo, err := service.SysUser().GetSysUserById(ctx, userId)

		if err != nil {
			return false, service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", dao.SysUser.Table())
		}

		result, err := service.Casbin().AddRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), consts.CasbinDomain)

		if result == false || err != nil {
			return result, err
		}
	}

	return true, nil
}

// CreateUser 创建用户
func (s *sSysUser) CreateUser(ctx context.Context, info model.UserInnerRegister, userState kyUser.StateEnum, userType kyUser.TypeEnum, customId ...int64) (*model.SysUserRegisterRes, error) {
	count, _ := dao.SysUser.Ctx(ctx).Unscoped().Count(dao.SysUser.Columns().Username, info.Username)
	if count > 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名已经存在"), "", dao.SysUser.Table())
	}

	data := entity.SysUser{
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

	result := model.SysUserRegisterRes{
		UserInfo:     data,
		RoleInfoList: make([]entity.SysRole, 0),
	}

	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 创建前
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&kyUser.Event.BeforeCreate.Code() == kyUser.Event.BeforeCreate.Code() {
					hook.Value.Value(ctx, kyUser.Event.BeforeCreate, data)
				}
			}
		})

		_, err := dao.SysUser.Ctx(ctx).OmitNilData().Data(data).Insert()

		if err != nil {
			return service.SysLogs().ErrorSimple(ctx, err, "账号注册失败", dao.SysUser.Table())
		}

		if len(info.RoleIds) > 0 {
			ret, err := s.SetUserRoleIds(ctx, info.RoleIds, data.Id)
			if ret != true || err != nil {
				return service.SysLogs().ErrorSimple(ctx, err, "角色设置失败！"+err.Error(), dao.SysUser.Table())
			}

			err = dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, info.RoleIds).Scan(&result.RoleInfoList)
			if err != nil {
				return service.SysLogs().ErrorSimple(ctx, err, "查询角色信息失败！", dao.SysUser.Table())
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
			if hook.Value.Key.Code()&kyUser.Event.AfterCreate.Code() == kyUser.Event.AfterCreate.Code() {
				hook.Value.Value(ctx, kyUser.Event.AfterCreate, data)
			}
		}
	})
	return &result, nil
}

// GetSysUserByUsername 根据用户名获取用户UID
func (s *sSysUser) GetSysUserByUsername(ctx context.Context, username string) (*entity.SysUser, error) {
	data := &entity.SysUser{}
	err := dao.SysUser.Ctx(ctx).Where(do.SysUser{Username: username}).Scan(data)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败", dao.SysUser.Table())
	}

	return data, nil
}

// HasSysUserByUsername 判断用户名是否存在
func (s *sSysUser) HasSysUserByUsername(ctx context.Context, username string) bool {
	count, _ := dao.SysUser.Ctx(ctx).Count(do.SysUser{Username: username})
	return count > 0
}

// GetSysUserById 根据用户ID获取用户信息
func (s *sSysUser) GetSysUserById(ctx context.Context, userId int64) (*entity.SysUser, error) {
	data := &entity.SysUser{}
	err := dao.SysUser.Ctx(ctx).Scan(data, do.SysUser{Id: userId})

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败", dao.SysUser.Table())
	}

	data.Password = masker.MaskString(data.Password, masker.Password)
	return data, nil
}

// SetUserPermissionIds 设置用户权限
func (s *sSysUser) SetUserPermissionIds(ctx context.Context, userId int64, permissionIds []int64) (bool, error) {
	err := dao.SysCasbin.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		{
			// 先清除roleId所有权限
			_, err := service.Casbin().DeletePermissionsForUser(gconv.String(userId))

			if len(permissionIds) <= 0 {
				return err
			}
		}

		// 重新赋予roleId新的权限清单
		for _, item := range permissionIds {
			ret, err := service.Casbin().Enforcer().AddPermissionForUser(gconv.String(userId), consts.CasbinDomain, gconv.String(item), "allow")
			if err != nil || ret == false {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "设置用户权限失败", dao.SysUser.Table())
	}

	return true, nil
}

// GetUserPermissionIds 获取用户权限，返回权限Id数组
func (s *sSysUser) GetUserPermissionIds(ctx context.Context, userId int64) ([]int64, error) {
	result, err := service.Casbin().Enforcer().GetImplicitPermissionsForUser(gconv.String(userId), consts.CasbinDomain)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户权限查询失败"), "", dao.SysUser.Table())
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
func (s *sSysUser) SetUsername(ctx context.Context, newUsername string) (bool, error) {
	userId := service.BizCtx().Get(ctx).ClaimsUser.Id
	result, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{Id: userId}).Update(do.SysUser{Username: newUsername})

	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}

// UpdateUserPassword 修改用户登录密码
func (s *sSysUser) UpdateUserPassword(ctx context.Context, info model.UpdateUserPassword) (bool, error) {
	// 获取到当前登录用户
	loginUserName := service.BizCtx().Get(ctx).ClaimsUser.Username

	if loginUserName == "" {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "需要先登录后才能修改密码")
	}

	// 查询到用户信息
	sysUserInfo, err := service.SysUser().GetSysUserByUsername(ctx, loginUserName)

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "当前登录用户不存在")
	}

	// 判断输入的两次密码是否相同
	if info.Password != info.ConfirmPassword {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "两次输入的密码不一致，修改失败")
	}

	// 传入用户输入的原始密码，进行hash，看是否和数据库中原始密码一致
	hash1, _ := en_crypto.PwdHash(info.OldPassword, gconv.String(sysUserInfo.Id))
	if sysUserInfo.Password != hash1 {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "原密码输入错误，修改失败")
	}

	pwdHash, err := en_crypto.PwdHash(info.Password, gconv.String(sysUserInfo.Id))

	_, err = dao.SysUser.Ctx(ctx).Where(do.SysUser{Id: sysUserInfo.Id}).Update(do.SysUser{Password: pwdHash})

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码修改失败")
	}

	return true, nil
}
