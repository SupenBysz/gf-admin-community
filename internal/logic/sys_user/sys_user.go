package user

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/utility/en_crypto"
	"github.com/SupenBysz/gf-admin-community/utility/kconv"
	"github.com/SupenBysz/gf-admin-community/utility/kmap"
	"github.com/SupenBysz/gf-admin-community/utility/masker"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"math"
	"sort"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
)

type hookInfo sys_model.KeyValueT[int64, sys_model.UserHookInfo]

type sSysUser struct {
	hookArr       []hookInfo
	mapInt64Items *kmap.HashMap[int64, *sys_model.SysUser]
}

func init() {
	sys_service.RegisterSysUser(New())
}

func New() *sSysUser {
	return &sSysUser{
		hookArr:       make([]hookInfo, 0),
		mapInt64Items: kmap.New[int64, *sys_model.SysUser](),
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

func (s *sSysUser) initInnerCacheItems(ctx context.Context) {
	if s.mapInt64Items.Size() > 0 {
		return
	}

	items := daoctl.Scan[[]*sys_model.SysUser](
		sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).With(sys_model.SysUser{}.Detail).
			OrderDesc(sys_dao.SysUser.Columns().CreatedAt),
	)
	s.mapInt64Items.Clear()
	for _, sysPermission := range *items {
		s.mapInt64Items.Set(sysPermission.Id, sysPermission)
	}
}

// QueryUserList 获取用户列表
func (s *sSysUser) QueryUserList(ctx context.Context, info *sys_model.SearchParams, unionMainId int64, isExport bool) (response *sys_model.SysUserListRes, err error) {
	if info != nil {
		newFields := make([]sys_model.FilterInfo, 0)

		// 移除类型筛选条件
		for _, field := range info.Filter {
			if field.Field != sys_dao.SysUser.Columns().Type {
				newFields = append(newFields, field)
			}
		}
	}

	// 如果没有查询条件，则默认从缓存返回数据
	if info != nil && len(info.Filter) <= 0 {
		// 初始化内部缓存数据
		s.initInnerCacheItems(ctx)

		response = &sys_model.SysUserListRes{}
		if info.Pagination.PageSize <= 0 {
			info.PageSize = 20
		}
		if info.Pagination.PageNum <= 0 {
			info.PageSize = 1
		}
		// 如果缓存没有数据则直接返回
		if s.mapInt64Items.Size() <= 0 {
			response.PaginationRes = sys_model.PaginationRes{
				Pagination: info.Pagination,
				PageTotal:  0,
				Total:      0,
			}
			response.Records = []*sys_model.SysUser{}
			return
		}

		// 设置分页信息
		response.Pagination = info.Pagination
		// 初始化分页统计信息
		response.PaginationRes = sys_model.PaginationRes{
			Pagination: info.Pagination,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(s.mapInt64Items.Size()) / gconv.Float64(info.PageSize))),
			Total:      gconv.Int64(s.mapInt64Items.Size()),
		}
		beginRowIndex := info.PageNum*info.PageSize - info.PageSize
		s.mapInt64Items.Iterator(func(k int64, v *sys_model.SysUser) bool {
			if beginRowIndex > 0 {
				beginRowIndex--
			} else if len(response.Records) < info.PageSize {
				// 查询用户所拥有的角色
				roleIds, _ := sys_service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(v.Id), sys_consts.CasbinDomain)

				user := kconv.Struct(v, &sys_model.SysUser{})
				user.RoleNames = []string{}

				// 如果有角色信息则加载角色信息
				if len(roleIds) > 0 {
					roles, err := sys_service.SysRole().QueryRoleList(ctx, sys_model.SearchParams{
						Filter: append(make([]sys_model.FilterInfo, 0), sys_model.FilterInfo{
							Field:     sys_dao.SysRole.Columns().Id,
							Where:     "in",
							IsOrWhere: false,
							Value:     roleIds,
						}),
						Pagination: sys_model.Pagination{},
					}, unionMainId)
					if err == nil && len(roles.Records) > 0 {
						for _, role := range roles.Records {
							user.RoleNames = append(user.RoleNames, role.Name)
						}
					}
				}
				user = s.masker(user)
				response.Records = append(response.Records, user)
			}

			sort.Slice(response.Records, func(i, j int) bool {
				return response.Records[i].CreatedAt.After(response.Records[j].CreatedAt)
			})
			return true
		})
		return
	}

	result, err := daoctl.Query[*sys_model.SysUser](sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).With(sys_model.SysUser{}.Detail), info, isExport)

	newList := make([]*sys_model.SysUser, 0)
	if result != nil && result.Records != nil && len(result.Records) > 0 {
		for _, user := range result.Records {
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
				}, unionMainId)
				if err == nil && len(roles.Records) > 0 {
					for _, role := range roles.Records {
						user.RoleNames = append(user.RoleNames, role.Name)
					}
				}
			}
			user = s.masker(user)
			newList = append(newList, user)
		}
	}

	if newList != nil {
		result.Records = newList
	}

	return (*sys_model.SysUserListRes)(result), err
}

// SetUserRoleIds 设置用户角色
func (s *sSysUser) SetUserRoleIds(ctx context.Context, roleIds []int64, userId int64) (bool, error) {
	for _, roleId := range roleIds {
		roleInfo := &sys_entity.SysRole{}
		// 查找角色是否存在
		roleInfo, err := sys_service.SysRole().GetRoleById(ctx, roleId)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", sys_dao.SysUser.Table())
		}

		userInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "用户ID错误", sys_dao.SysUser.Table())
		}

		result, err := sys_service.Casbin().AddRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)

		if result == false || err != nil {
			return result, sys_service.SysLogs().ErrorSimple(ctx, err, "设置用户角色失败", sys_dao.SysUser.Table())
		}
	}

	return true, nil
}

// CreateUser 创建用户
func (s *sSysUser) CreateUser(ctx context.Context, info sys_model.UserInnerRegister, userState sys_enum.UserState, userType sys_enum.UserType, customId ...int64) (*sys_model.SysUser, error) {
	count, _ := sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).Unscoped().Count(sys_dao.SysUser.Columns().Username, info.Username)
	if count > 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名已经存在"), "", sys_dao.SysUser.Table())
	}

	data := sys_model.SysUser{
		SysUser: sys_entity.SysUser{
			Id:        idgen.NextId(),
			Username:  info.Username,
			Password:  info.Password,
			Mobile:    info.Mobile,
			State:     userState.Code(),
			Type:      userType.Code(),
			CreatedAt: gtime.Now(),
		},
	}

	if len(customId) > 0 && customId[0] > 0 {
		data.Id = customId[0]
	}

	pwdHash, err := en_crypto.PwdHash(info.Password, gconv.String(data.Id))

	// 密码赋值
	data.Password = pwdHash

	err = sys_dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建前
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.BeforeCreate.Code() == sys_enum.User.Event.BeforeCreate.Code() {
					res, _ := hook.Value.Value(ctx, sys_enum.User.Event.BeforeCreate, data)
					res.Detail.Id = data.Id
					data.Detail = res.Detail
				}
			}
		})

		{
			_, err = sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).OmitNilData().Data(data.SysUser).Insert()

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "账号注册失败", sys_dao.SysUser.Table())
			}
		}
		{
			if data.Detail.Id > 0 && (data.Detail.Realname != "" || data.Detail.UnionMainName != "") {
				_, err = sys_dao.SysUserDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).OmitNilData().Data(data.Detail).Insert()

				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "账号注册失败", sys_dao.SysUser.Table())
				}
			}

		}
		if len(info.RoleIds) > 0 {
			ret, err := s.SetUserRoleIds(ctx, info.RoleIds, data.Id)
			if ret != true || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "角色设置失败！"+err.Error(), sys_dao.SysUser.Table())
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
				res, _ := hook.Value.Value(ctx, sys_enum.User.Event.AfterCreate, data)
				res.Detail.Id = data.Id
				data.Detail = res.Detail
			}
		}
	})

	s.mapInt64Items.Set(data.Id, &data)

	return &data, nil
}

// SetUserPermissions 设置用户权限
func (s *sSysUser) SetUserPermissions(ctx context.Context, userId int64, permissionIds []int64) (bool, error) {
	_, err := s.GetSysUserById(ctx, userId)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败", sys_dao.SysRole.Table())
	}

	return sys_service.SysPermission().SetPermissionsByResource(ctx, gconv.String(userId), permissionIds)
}

// GetSysUserByUsername 根据用户名获取用户
func (s *sSysUser) GetSysUserByUsername(ctx context.Context, username string) (response *sys_model.SysUser, err error) {
	s.initInnerCacheItems(ctx)

	s.mapInt64Items.Iterator(func(k int64, v *sys_model.SysUser) bool {
		if v.Username == username {
			response = kconv.Struct(v, &sys_model.SysUser{})
			return false
		}
		return true
	})

	if response == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, sql.ErrNoRows, "用户信息不存在", sys_dao.SysUser.Table())
	}

	response = s.masker(response)
	return
}

// CheckPassword 检查密码是否正确
func (s *sSysUser) CheckPassword(ctx context.Context, userId int64, password string) (bool, error) {
	s.initInnerCacheItems(ctx)
	data, has := s.mapInt64Items.Search(userId)

	if !has {
		return false, sys_service.SysLogs().ErrorSimple(ctx, sql.ErrNoRows, "用户信息不存在", sys_dao.SysUser.Table())
	}

	// 取盐
	salt := gconv.String(userId)

	// 加密：用户输入的密码 + 他的id的后八位(盐)  --进行Hash--> 用户提供的密文
	pwdHash, err := en_crypto.PwdHash(password, salt)

	return data.Password == pwdHash, err
}

// HasSysUserByUsername 判断用户名是否存在
func (s *sSysUser) HasSysUserByUsername(ctx context.Context, username string) bool {
	data, _ := s.GetSysUserByUsername(ctx, username)
	return data != nil
}

// GetSysUserById 根据用户ID获取用户信息
func (s *sSysUser) GetSysUserById(ctx context.Context, userId int64) (*sys_model.SysUser, error) {
	s.initInnerCacheItems(ctx)

	data, has := s.mapInt64Items.Search(userId)

	if !has {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, sql.ErrNoRows, "用户信息不存在", sys_dao.SysUser.Table())
	}

	return s.masker(kconv.Struct(data, &sys_model.SysUser{})), nil
}

// SetUserPermissionIds 设置用户权限
func (s *sSysUser) SetUserPermissionIds(ctx context.Context, userId int64, permissionIds []int64) (bool, error) {
	err := sys_dao.SysCasbin.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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

// DeleteUser 删除用户信息，该方法一般由后端业务层内部调用
func (s *sSysUser) DeleteUser(ctx context.Context, id int64) (bool, error) {
	_, err := s.GetSysUserById(ctx, id)
	if err != nil {
		return false, err
	}

	err = sys_dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 移除员工权限
		_, err = sys_service.SysPermission().SetPermissionsByResource(ctx, gconv.String(id), []int64{0})
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		// 移除员工角色
		sys_service.SysUser().SetUserRoleIds(ctx, []int64{0}, id)
		if err != nil {
			return err
		}

		// 删除用户附加信息
		_, err = sys_dao.SysUserDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).Unscoped().Delete(sys_do.SysUserDetail{Id: id})
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		// 删除用户
		_, err = sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).Unscoped().Delete(sys_do.SysUser{Id: id})
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除员工信息失败", sys_dao.SysUser.Table())
	}

	s.mapInt64Items.Remove(id)
	return true, nil
}

// SetUsername 修改自己的账号登陆名称
func (s *sSysUser) SetUsername(ctx context.Context, newUsername string, userId int64) (bool, error) {
	result, err := sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Data(sys_do.SysUser{Username: newUsername}).
		Where(sys_do.SysUser{Id: userId}).
		Update()

	if err != nil || result == nil {
		return false, err
	}
	data := s.mapInt64Items.Get(userId)
	data.Username = newUsername
	return true, nil
}

// SetUserState 设置用户状态
func (s *sSysUser) SetUserState(ctx context.Context, userId int64, state sys_enum.UserType) (bool, error) {
	result, err := sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Data(sys_do.SysUser{State: state.Code()}).
		Where(sys_do.SysUser{Id: userId}).
		Update()

	if err != nil || result == nil {
		return false, err
	}

	data := s.mapInt64Items.Get(userId)
	data.State = state.Code()
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
					_, err = hook.Value.Value(ctx, sys_enum.User.Event.ChangePassword, *sysUserInfo)
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

	_, err = sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(sys_do.SysUser{Id: sysUserInfo.Id}).Update(sys_do.SysUser{Password: pwdHash})

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码修改失败")
	}

	data := s.mapInt64Items.Get(userId)
	data.Password = pwdHash
	return true, nil
}

// ResetUserPassword 重置用户密码 (超级管理员无需验证验证，XX商管理员重置员工密码无需验证)
func (s *sSysUser) ResetUserPassword(ctx context.Context, userId int64, password string, confirmPassword string) (bool, error) {
	// hook判断当前登录身份是否可以重置密码
	{
		s.initInnerCacheItems(ctx)

		user, err := s.GetSysUserById(ctx, userId)

		if err != nil {
			return false, err
		}

		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.ResetPassword.Code() == sys_enum.User.Event.ResetPassword.Code() {
					_, err = hook.Value.Value(ctx, sys_enum.User.Event.ResetPassword, *kconv.Struct(user, &sys_model.SysUser{}))
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

		result, err := sys_dao.SysUser.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(sys_do.SysUser{Id: userId}).Update(sys_do.SysUser{Password: pwdHash})

		// 受影响的行数
		count, _ := result.RowsAffected()

		if err != nil || count != 1 {
			return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "重置密码失败")
		}

		data := s.mapInt64Items.Get(userId)
		data.Password = pwdHash
	}

	return true, nil
}

// SetUserRoles 设置用户角色
func (s *sSysUser) SetUserRoles(ctx context.Context, userId int64, roleIds []int64, makeUserUnionMainId int64) (bool, error) {
	data, err := s.GetSysUserById(ctx, userId)
	if err != nil {
		return false, err
	}

	err = sys_dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, roleId := range roleIds {
			roleInfo, err := sys_service.SysRole().GetRoleById(ctx, roleId)

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "角色ID错误", sys_dao.SysRole.Table())
			}

			if roleInfo.UnionMainId != makeUserUnionMainId {
				return sys_service.SysLogs().ErrorSimple(ctx, err, roleInfo.Name+" 角色信息校验失败", sys_dao.SysRole.Table())
			}

			ret, _ := sys_service.Casbin().AddRoleForUserInDomain(gconv.String(userId), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)
			if ret == true {
				// 重置用户角色名称，并自动去重
				data.RoleNames = garray.NewSortedStrArrayFrom(append(data.RoleNames, roleInfo.Name)).Unique().Slice()
			}
		}
		return nil
	})
	return err == nil, err
}

// UpdateUserExDetail 更新用户扩展信息
func (s *sSysUser) UpdateUserExDetail(ctx context.Context, user *sys_model.SysUser) (*sys_model.SysUser, error) {
	s.initInnerCacheItems(ctx)

	data := sys_entity.SysUserDetail{}

	err := sys_dao.SysUserDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(sys_do.SysUserDetail{Id: user.Id}).Scan(&data)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		_, err = sys_dao.SysUserDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(user.Detail)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = sys_dao.SysUserDetail.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(sys_do.SysUserDetail{
			Realname:      user.Detail.Realname,
			UnionMainName: user.Detail.UnionMainName,
			LastLoginIp:   user.Detail.LastLoginIp,
			LastLoginArea: user.Detail.LastLoginArea,
			LastLoginAt:   user.Detail.LastLoginAt,
		}).Where(sys_do.SysUserDetail{Id: user.Id}).Update()
		if err != nil {
			return nil, err
		}
	}
	s.mapInt64Items.Set(user.Id, user)
	return user, nil
}

// GetUserDetail 查看用户详情，含完整手机号
func (s *sSysUser) GetUserDetail(ctx context.Context, userId int64) (*sys_entity.SysUser, error) {
	s.initInnerCacheItems(ctx)
	data, has := s.mapInt64Items.Search(userId)
	if !has {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户信息不存在")
	}
	result := kconv.Struct(data, &sys_entity.SysUser{})
	result.Password = masker.MaskString(result.Password, masker.Password)
	return result, nil
}

// SetUserMobile 设置用户手机号
func (s *sSysUser) SetUserMobile(ctx context.Context, newMobile int64, captcha string, userId int64) (bool, error) {
	_, err := sys_service.SysSms().Verify(ctx, newMobile, captcha)
	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysUser.Ctx(ctx).
		Data(sys_do.SysUser{Mobile: newMobile, UpdatedAt: gtime.Now()}).
		Where(sys_do.SysUser{Id: userId}).
		Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设置用户手机号失败", sys_dao.SysUser.Table())
	}

	return true, nil
}

func (s *sSysUser) masker(user *sys_model.SysUser) *sys_model.SysUser {
	user.Password = masker.MaskString(user.Password, masker.Password)
	user.Mobile = masker.MaskString(user.Mobile, masker.MaskPhone)
	return user
}
