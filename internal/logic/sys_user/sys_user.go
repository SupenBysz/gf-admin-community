package user

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/base_funs"
	"github.com/kysion/base-library/utility/base_verify"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/en_crypto"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/base-library/utility/masker"
)

type hookInfo sys_model.KeyValueT[int64, sys_hook.UserHookInfo]

type sSysUser struct {
	hookArr []hookInfo
	//redisCache *gcache.Cache
	Duration time.Duration

	heartbeatTimeout time.Duration
	//// 密码加密
	//CryptoPasswordFunc func(ctx context.Context, passwordStr string, user ...sys_entity.SysUser) (pwdEncode string)
}

func init() {
	sys_service.RegisterSysUser(New())

	// 初始化用户在线用户心跳超时设置
	_, _ = sys_service.SysUser().UpdateHeartbeatAt(context.Background(), 0)
}

func New() sys_service.ISysUser {

	return &sSysUser{
		//redisCache: gcache.New(),
		hookArr:          make([]hookInfo, 0),
		heartbeatTimeout: 60,
	}
}

//
//// 初始化缓存
//func (s *sSysUser) initInnerCacheItems(ctx context.Context) {
//	size, _ := s.redisCache.Size(ctx)
//	if size > 0 {
//		return
//	}
//
//	items := daoctl.Scan[[]*sys_model.SysUser](
//		sys_dao.SysUser.Ctx(ctx).
//			OrderDesc(sys_dao.SysUser.Columns().CreatedAt),
//	)
//
//	s.redisCache.Clear(ctx)
//	for _, sysUser := range *items {
//		s.redisCache.Set(ctx, sysUser.Id, sysUser, s.Duration)
//	}
//}

// InstallHook 安装Hook
func (s *sSysUser) InstallHook(event sys_enum.UserEvent, hookFunc sys_hook.UserHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_hook.UserHookInfo{Key: event, Value: hookFunc}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

//// SetCryptoPasswordFunc 用于业务端自定义密码规则
//func (s *sSysUser) SetCryptoPasswordFunc(f func(ctx context.Context, passwordStr string, user ...sys_entity.SysUser) (pwdEncode string)) {
//	s.CryptoPasswordFunc = f
//}
//
//// GetCryptoPasswordFunc 应用业务端自定义密码规则
//func (s *sSysUser) GetCryptoPasswordFunc() func(ctx context.Context, passwordStr string, user ...sys_entity.SysUser) (pwdEncode string) {
//	return s.CryptoPasswordFunc
//}

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

// UpdateHeartbeatAt 更新用户心跳监测时间
func (s *sSysUser) UpdateHeartbeatAt(ctx context.Context, heartbeatTimeout int) (api_v1.BoolRes, error) {

	// 加载心跳超时时间默认配置
	newHeartbeatTimeout := g.Cfg().MustGet(context.Background(), "service.heartbeatTimeout", 60).Duration() * time.Second

	// 如果配置了心跳超时时间,且大于10秒，则使用配置的时间，否则使用默认时间
	if heartbeatTimeout > 10 {
		newHeartbeatTimeout = time.Duration(heartbeatTimeout) * time.Second
	}

	heartbeatTimeoutKey := "heartbeatTimeout"
	// 从数据库中查询心跳超时时间设置
	settingInfo, _ := sys_service.SysSettings().GetByName(context.Background(), heartbeatTimeoutKey, nil)

	// 如果数据库中有设置，则覆盖默认值
	if settingInfo == nil {
		data := &sys_model.SysSettings{
			Name:   heartbeatTimeoutKey,
			Values: gjson.MustEncodeString(newHeartbeatTimeout.Seconds()),
			Desc:   g.I18n().T(ctx, "heartbeat_timeout_description"),
		}

		_, err := sys_service.SysSettings().Create(ctx, data)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_heartbeat_timeout_save_failed", sys_dao.SysUser.Table())
		}
	} else {
		_, err := sys_service.SysSettings().Update(ctx, &sys_model.SysSettings{
			Name:        settingInfo.Name,
			Values:      gjson.MustEncodeString(newHeartbeatTimeout.Seconds()),
			Desc:        settingInfo.Desc,
			UnionMainId: settingInfo.UnionMainId,
		})

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_heartbeat_timeout_save_failed", sys_dao.SysUser.Table())
		}
	}
	s.heartbeatTimeout = newHeartbeatTimeout
	return true, nil
}

// QueryUserList 获取用户列表
func (s *sSysUser) QueryUserList(ctx context.Context, info *base_model.SearchParams, unionMainId int64, isExport bool) (response *sys_model.SysUserListRes, err error) {
	isFilterOnlineUser := false

	if info != nil {
		newFields := make([]base_model.FilterInfo, 0)

		for _, field := range info.Filter {
			if field.Field != sys_dao.SysUser.Columns().Type { // 移除类型筛选条件
				// 移除在线过滤标识
				if field.Field == "is_online" {
					isFilterOnlineUser = true
				} else {
					newFields = append(newFields, field)
				}
			}
		}

		// 是否过滤在线用户
		if isFilterOnlineUser {
			// 查询最后心跳小于30秒的用户作为在线用户，
			// 由于中线还有其他业务逻辑，可能影响在线判断逻辑，
			// 因此在 makeMore 方法附加数据中判断在线逻辑时，应大于30秒，避免返回的数据出现离线用户
			result, _ := sys_dao.SysUserDetail.Ctx(ctx).WhereGT(
				sys_dao.SysUserDetail.Columns().LastHeartbeatAt,
				time.Now().Add(-time.Second*s.heartbeatTimeout-time.Second*10),
			).Fields([]string{sys_dao.SysUserDetail.Columns().Id}).All()

			// 提取用户Ids
			userIds := make([]int64, 0)
			for _, value := range result.Array() {
				userIds = append(userIds, value.Int64())
			}

			// 附加查询条件
			newFields = append(newFields, base_model.FilterInfo{
				Field: sys_dao.SysUser.Columns().Id,
				Where: "in",
				Value: userIds,
			})
		}

		info.Filter = newFields
	}

	// 如果没有查询条件，则默认从缓存返回数据
	if info != nil && len(info.Filter) <= 0 {
		// 初始化内部缓存数据
		//s.initInnerCacheItems(ctx)

		response = &sys_model.SysUserListRes{}
		if info.Pagination.PageSize <= 0 {
			info.PageSize = 20
		}
		if info.Pagination.PageNum <= 0 {
			info.PageSize = 1
		}

		// 如果缓存没有数据则直接返回
		//size, _ := s.redisCache.Size(ctx)
		userList, err := daoctl.Query[sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), nil, true)
		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_list_query_failed", sys_dao.SysUser.Table())
		}
		size := len(userList.Records)

		if size <= 0 {
			response.PaginationRes = base_model.PaginationRes{
				Pagination: info.Pagination,
				PageTotal:  0,
				Total:      0,
			}
			response.Records = []*sys_model.SysUser{}
		}

		// 设置分页信息
		response.Pagination = info.Pagination
		// 初始化分页统计信息
		response.PaginationRes = base_model.PaginationRes{
			Pagination: info.Pagination,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(size) / gconv.Float64(info.PageSize))),
			Total:      gconv.Int64(size),
		}
		beginRowIndex := info.PageNum*info.PageSize - info.PageSize

		// 获得所有的key，遍历
		result, _ := daoctl.Query[*sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), info, isExport)

		//keys, _ := s.redisCache.Keys(ctx)

		for _, k := range result.Records {
			if beginRowIndex > 0 {
				beginRowIndex--
			} else if len(response.Records) < info.PageSize {
				// 查询用户所拥有的角色
				sysUser := &sys_model.SysUser{}
				sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: gconv.String(k.Id)}).Scan(&sysUser)

				_, err = s.getUserRole(ctx, sysUser, unionMainId)
				if err != nil {
					return nil, err
				}

				sysUser = s.masker(s.makeMore(ctx, sysUser))

				response.Records = append(response.Records, sysUser)
			}

			sort.Slice(response.Records, func(i, j int) bool {
				return response.Records[i].CreatedAt.After(response.Records[j].CreatedAt)
			})

		}
	}

	result, err := daoctl.Query[*sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), info, isExport)

	newList := make([]*sys_model.SysUser, 0)
	if result != nil && result.Records != nil && len(result.Records) > 0 {
		result.PaginationRes = base_model.PaginationRes{
			Pagination: info.Pagination,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(len(result.Records)) / gconv.Float64(info.PageSize))),
			Total:      gconv.Int64(len(result.Records)),
		}

		for _, user := range result.Records {
			user.RoleNames = make([]string, 0)
			roleIds, err := sys_service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(user.Id), sys_consts.CasbinDomain)

			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}

			if len(roleIds) > 0 {
				roles, err := sys_service.SysRole().QueryRoleList(ctx, base_model.SearchParams{
					Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
						Field:     sys_dao.SysRole.Columns().Id,
						Where:     "in",
						IsOrWhere: false,
						Value:     roleIds,
					}),
					Pagination: base_model.Pagination{},
				}, unionMainId)
				if err == nil && len(roles.Records) > 0 {
					for _, role := range roles.Records {
						user.RoleNames = append(user.RoleNames, role.Name)
					}
				}
			}
			user = s.masker(s.makeMore(ctx, user))
			newList = append(newList, user)
		}
	}

	if len(newList) > 0 {
		result.Records = newList
	}

	return (*sys_model.SysUserListRes)(result), err
}

// SetUserRoleIds 设置用户角色
func (s *sSysUser) SetUserRoleIds(ctx context.Context, roleIds []int64, userId int64) (bool, error) {
	for _, roleId := range roleIds {
		roleInfo := &sys_model.SysRoleRes{}
		// 查找角色是否存在
		roleInfo, err := sys_service.SysRole().GetRoleById(ctx, roleId)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_role_id_invalid", sys_dao.SysUser.Table())
		}

		userInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_id_invalid", sys_dao.SysUser.Table())
		}

		result, err := sys_service.Casbin().AddRoleForUserInDomain(gconv.String(userInfo.Id), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)

		if !result || err != nil {
			return result, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_role_set_failed", sys_dao.SysUser.Table())
		}
	}

	return true, nil
}

// CreateUser 创建用户
func (s *sSysUser) CreateUser(ctx context.Context, info sys_model.UserInnerRegister, userState sys_enum.UserState, userType sys_enum.UserType, customId ...int64) (*sys_model.SysUser, error) {
	if info.Username == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_username_is_required"), "", sys_dao.SysUser.Table())
	}

	count, _ := sys_dao.SysUser.Ctx(ctx).WhereLike("LOWER(\""+sys_dao.SysUser.Columns().Username+"\")", strings.ToLower(info.Username)).Unscoped().Count()
	if count > 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_username_already_exists"), "", sys_dao.SysUser.Table())
	}

	data := sys_model.SysUser{
		SysUser: &sys_entity.SysUser{
			Id:        idgen.NextId(),
			Username:  info.Username,
			Password:  info.Password,
			Mobile:    info.Mobile,
			Email:     info.Email,
			State:     userState.Code(),
			Type:      userType.Code(),
			InviteCode: info.InviteCode,
			CreatedAt: gtime.Now(),
		},
	}

	if len(customId) > 0 && customId[0] > 0 {
		data.Id = customId[0]
	}
	pwdHash, err := en_crypto.PwdHash(info.Password, gconv.String(data.Id))

	// 业务层自定义密码加密规则
	if sys_consts.Global.CryptoPasswordFunc != nil {
		pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, info.Password, *data.SysUser)
	}

	// 密码赋值
	data.Password = pwdHash

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建前
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				// 枚举优化使用：直接调用Has
				//enumOb := sys_enum.User.Type.New(3, "")
				//if enumOb.Has(sys_enum.User.Event.BeforeCreate) { // 单个满足
				//
				//}
				//if hook.Value.Key.Has(sys_enum.User.Event.BeforeCreate, sys_enum.User.Event.AfterCreate) { // 多个满足
				//
				//}

				// 自增
				//enumOb.Add(sys_enum.User.Event.AfterCreate, sys_enum.User.Event.BeforeCreate)

				// 自减少
				//enumOb.Remove(sys_enum.User.Event.AfterCreate)

				if (hook.Value.Key.Code() & sys_enum.User.Event.BeforeCreate.Code()) == sys_enum.User.Event.BeforeCreate.Code() {
					res, _ := hook.Value.Value(ctx, sys_enum.User.Event.BeforeCreate, data)
					res.Detail = &sys_model.SysUserDetail{}
					res.Detail.Id = data.Id
					data.Detail = res.Detail
				}
			}
		})

		{
			_, err = sys_dao.SysUser.Ctx(ctx).OmitNilData().Data(data.SysUser).Insert()

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_account_registration_failed", sys_dao.SysUser.Table())
			}
		}

		{
			if data.Detail != nil && data.Detail.Id > 0 && (data.Detail.Realname != "" || data.Detail.UnionMainName != "") {
				_, err = sys_dao.SysUserDetail.Ctx(ctx).OmitNilData().Data(data.Detail).Insert()

				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "error_account_registration_failed", sys_dao.SysUser.Table())
				}
			}

			/*
				//	if data.Detail != nil && data.Detail.Id > 0 && (data.Detail.Realname != "" || data.Detail.UnionMainName != "") {

						_, err = sys_dao.SysUserDetail.Ctx(ctx).OmitNilData().Data(data.Detail).Insert()

						if err != nil {
							return sys_service.SysLogs().ErrorSimple(ctx, err, "账号注册失败", sys_dao.SysUser.Table())
						}
					//}
			*/

		}
		if len(info.RoleIds) > 0 {
			ret, err := s.SetUserRoleIds(ctx, info.RoleIds, data.Id)
			if !ret || err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_role_set_failed", sys_dao.SysUser.Table())
			}
		}

		// 建后
		err = g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.AfterCreate.Code() == sys_enum.User.Event.AfterCreate.Code() {
					res, err := hook.Value.Value(ctx, sys_enum.User.Event.AfterCreate, data)
					if err != nil {
						panic(err)
					}
					res.Detail.Id = data.Id
					data.Detail = res.Detail
				}
			}
		})

		return err
	})

	if err != nil {
		return nil, err
	}

	// 查询用户所拥有的角色 (指针传递)
	s.getUserRole(ctx, &data)

	// 增删改后不需要重新设置缓存，因为增伤缓存参数Duration为-1，就是删除缓存了
	// s.redisCache.Set(ctx, data.Id, data, s.Duration)

	return &data, nil
}

// SetUserPermissions 设置用户权限
func (s *sSysUser) SetUserPermissions(ctx context.Context, userId int64, permissionIds []int64) (bool, error) {
	_, err := s.GetSysUserById(ctx, userId)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_info_query_failed", sys_dao.SysRole.Table())
	}

	return sys_service.SysPermission().SetPermissionsByResource(ctx, gconv.String(userId), permissionIds)
}

// GetSysUserByUsername 根据用户名获取用户
func (s *sSysUser) GetSysUserByUsername(ctx context.Context, username string) (response *sys_model.SysUser, err error) {
	//s.initInnerCacheItems(ctx)

	// 获取所有keys
	// keys, err := s.redisCache.Keys(ctx)
	userList, err := daoctl.Query[sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), nil, true)

	if err != nil && err != sql.ErrNoRows {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_info_query_failed", sys_dao.SysUser.Table())
	}

	user := &sys_model.SysUser{}

	checkUsername := strings.ToLower(username)

	for _, k := range userList.Records {
		//sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: gconv.String(k.Id)}).Scan(&user)
		if strings.ToLower(k.Username) == checkUsername {
			user, err = daoctl.GetByIdWithError[sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), k.Id)
			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_info_query_failed", sys_dao.SysUser.Table())
			}
			response = s.masker(s.makeMore(ctx, user))
			// 查询用户所拥有的角色 (指针传递)
			s.getUserRole(ctx, response)
			return
		}
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, sql.ErrNoRows, "error_user_info_not_exist", sys_dao.SysUser.Table())
}

// CheckPassword 检查密码是否正确
func (s *sSysUser) CheckPassword(ctx context.Context, userId int64, password string) (bool, error) {
	//s.initInnerCacheItems(ctx)

	userInfo, err := daoctl.GetByIdWithError[sys_entity.SysUser](sys_dao.SysUser.Ctx(ctx), userId)

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, sql.ErrNoRows, "error_user_info_not_exist", sys_dao.SysUser.Table())
	}
	// if （）{hook()}
	// 取盐
	salt := gconv.String(userId)

	// 加密：用户输入的密码 + 他的id的后八位(盐)  --进行Hash--> 用户提供的密文
	pwdHash, err := en_crypto.PwdHash(password, salt)
	// 业务层自定义密码加密规则
	if sys_consts.Global.CryptoPasswordFunc != nil {
		pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, password, *userInfo)
	}

	return userInfo.Password == pwdHash, err
}

// HasSysUserByUsername 判断用户名是否存在
func (s *sSysUser) HasSysUserByUsername(ctx context.Context, username string) bool {
	data, _ := s.GetSysUserByUsername(ctx, username)
	return data != nil
}

// GetSysUserById 根据用户ID获取用户信息
func (s *sSysUser) GetSysUserById(ctx context.Context, userId int64) (*sys_model.SysUser, error) {
	//s.initInnerCacheItems(ctx)

	user := sys_model.SysUser{}
	err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{
		Id: userId,
	}).Scan(&user)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, sql.ErrNoRows, "error_user_info_not_exist", sys_dao.SysUser.Table())
	}

	// 查询用户所拥有的角色 (指针传递)
	_, err = s.getUserRole(ctx, &user)
	if err != nil {
		return nil, err
	}

	return s.masker(s.makeMore(ctx, &user)), nil
}

func (s *sSysUser) MakeSession(ctx context.Context, userId int64) {
	user, err := s.GetSysUserById(ctx, userId)

	if err != nil {
		return
	}

	token, _ := sys_service.Jwt().GenerateToken(ctx, user)
	if token != nil {
		sys_service.Jwt().MakeSession(ctx, token.Token)
	}
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
			if err != nil || !ret {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_permission_set_failed", sys_dao.SysUser.Table())
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
		_, err = sys_dao.SysUserDetail.Ctx(ctx).Unscoped().Delete(sys_do.SysUserDetail{Id: id})
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		// 删除用户
		_, err = sys_dao.SysUser.Ctx(ctx).Unscoped().Delete(sys_do.SysUser{Id: id})
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		return nil
	})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_delete_failed", sys_dao.SysUser.Table())
	}

	// daoctl.RemoveQueryCache(sys_dao.SysUser.DB(), sys_dao.SysUser.Table())

	return true, nil
}

// SetUsername 修改自己的账号登陆名称
func (s *sSysUser) SetUsername(ctx context.Context, newUsername string, userId int64) (bool, error) {
	result, err := sys_dao.SysUser.Ctx(ctx).Data(sys_do.SysUser{Username: newUsername}).Where(sys_do.SysUser{Id: userId}).Update()

	if err != nil || result == nil {
		return false, err
	}

	return true, nil
}

// SetUserState 设置用户状态
func (s *sSysUser) SetUserState(ctx context.Context, userId int64, state sys_enum.UserState) (bool, error) {
	result, err := sys_dao.SysUser.Ctx(ctx).Data(sys_do.SysUser{State: state.Code()}).Where(sys_do.SysUser{Id: userId}).Update()

	if err != nil || result == nil {
		return false, err
	}

	return true, nil
}

// SetUserType 设置用户类型
func (s *sSysUser) SetUserType(ctx context.Context, userId int64, state sys_enum.UserType) (bool, error) {
	result, err := sys_dao.SysUser.Ctx(ctx).Data(sys_do.SysUser{Type: state.Code()}).Where(sys_do.SysUser{Id: userId}).Update()

	if err != nil || result == nil {
		return false, err
	}

	return true, nil
}

// UpdateUserPassword 修改用户登录密码
func (s *sSysUser) UpdateUserPassword(ctx context.Context, info sys_model.UpdateUserPassword, userId int64) (bool, error) {
	// 查询到用户信息 不能使用这个操作去查询用户，因为masker操作会把用户密码变空
	// sysUserInfo, err := sys_service.SysUser().GetSysUserById(ctx, userId)

	sysUserInfo, err := daoctl.GetByIdWithError[sys_model.SysUser](sys_dao.SysUser.Ctx(ctx), userId)

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_user_not_exist")
	}

	// 判断输入的两次密码是否相同
	if info.Password != info.ConfirmPassword {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_password_mismatch")
	}

	{
		// 传入用户输入的原始密码，进行hash，看是否和数据库中原始密码一致
		hash1, _ := en_crypto.PwdHash(info.OldPassword, gconv.String(sysUserInfo.Id))
		// 业务层自定义密码加密规则
		if sys_consts.Global.CryptoPasswordFunc != nil {
			hash1 = sys_consts.Global.CryptoPasswordFunc(ctx, info.OldPassword, *sysUserInfo.SysUser)
		}
		if sysUserInfo.Password != hash1 {
			return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_invalid_old_password")
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
	// 业务层自定义密码加密规则
	if sys_consts.Global.CryptoPasswordFunc != nil {
		pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, info.Password, *sysUserInfo.SysUser)
	}

	_, err = sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: sysUserInfo.Id}).Update(sys_do.SysUser{Password: pwdHash})

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_password_update_failed")
	}

	return true, nil
}

// ResetUserPassword 重置用户密码 (超级管理员无需验证验证，XX商管理员重置员工密码无需验证)
func (s *sSysUser) ResetUserPassword(ctx context.Context, userId int64, password string, confirmPassword string) (bool, error) {
	// hook判断当前登录身份是否可以重置密码
	user, err := s.GetSysUserById(ctx, userId)
	{
		//s.initInnerCacheItems(ctx)

		if err != nil {
			return false, err
		}

		// 发布广播
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
			return false, gerror.NewCode(gcode.CodeValidationFailed, "error_password_mismatch")
		}
		// 取盐
		salt := gconv.String(userId)

		// 加密
		pwdHash, _ := en_crypto.PwdHash(password, salt)
		// 业务层自定义密码加密规则
		if sys_consts.Global.CryptoPasswordFunc != nil {
			pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, password, *user.SysUser)
		}

		result, err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: userId}).Update(sys_do.SysUser{Password: pwdHash})

		// 受影响的行数
		count, _ := result.RowsAffected()

		if err != nil || count != 1 {
			return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_password_reset_failed")
		}
	}

	return true, nil
}

// HasSysUserEmail 邮箱是否存在
func (s *sSysUser) HasSysUserEmail(ctx context.Context, email string) bool {
	response, _ := s.GetSysUserByEmail(ctx, email)

	return response != nil
}

// GetSysUserByEmail 根据邮箱获取用户信息
func (s *sSysUser) GetSysUserByEmail(ctx context.Context, email string) (response *sys_model.SysUser, err error) {

	err = sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Email: email}).Scan(response)

	return
}

// ResetUserEmail 重置用户邮箱
func (s *sSysUser) ResetUserEmail(ctx context.Context, userId int64, email string) (bool, error) {
	// hook判断当前登录身份是否可以重置密码
	user, err := s.GetSysUserById(ctx, userId)
	{
		//s.initInnerCacheItems(ctx)

		if err != nil {
			return false, err
		}

		// 发布广播
		err = g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if hook.Value.Key.Code()&sys_enum.User.Event.ResetEmail.Code() == sys_enum.User.Event.ResetEmail.Code() {
					_, err = hook.Value.Value(ctx, sys_enum.User.Event.ResetEmail, *kconv.Struct(user, &sys_model.SysUser{}))
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

	affected, err := daoctl.UpdateWithError(sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: userId}), sys_do.SysUser{Email: strings.ToLower(email)})

	return affected > 0, err
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
				return sys_service.SysLogs().ErrorSimple(ctx, err, "error_role_info_invalid", sys_dao.SysRole.Table())
			}

			if roleInfo.UnionMainId != makeUserUnionMainId {
				return sys_service.SysLogs().ErrorSimple(ctx, err, roleInfo.Name+" error_role_info_mismatch", sys_dao.SysRole.Table())
			}

			ret, _ := sys_service.Casbin().AddRoleForUserInDomain(gconv.String(userId), gconv.String(roleInfo.Id), sys_consts.CasbinDomain)
			if ret {
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
	//s.initInnerCacheItems(ctx)

	var data *sys_model.SysUserDetail

	err := sys_dao.SysUserDetail.Ctx(ctx).Where(sys_do.SysUserDetail{Id: user.Id}).Scan(&data)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if user.Detail == nil {
		if data != nil {
			user.Detail = data
		} else {
			user.Detail = &sys_model.SysUserDetail{
				SysUserDetail: sys_entity.SysUserDetail{
					Id:            user.Id,
					Realname:      "",
					UnionMainName: "",
					LastLoginIp:   "",
					LastLoginArea: "",
					LastLoginAt:   nil,
				},
			}
		}
	}

	if err == nil && data == nil || errors.Is(err, sql.ErrNoRows) {
		_, err = sys_dao.SysUserDetail.Ctx(ctx).Insert(user.Detail)
		if err != nil {
			return nil, err
		}
	} else {
		if data == nil {
			_, err = sys_dao.SysUserDetail.Ctx(ctx).Data(sys_do.SysUserDetail{
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
	}

	//s.redisCache.Set(ctx, user.Id, user, s.Duration)
	return user, nil
}

// GetUserDetail 查看用户详情，含完整手机号
func (s *sSysUser) GetUserDetail(ctx context.Context, userId int64) (*sys_model.SysUser, error) {
	//s.initInnerCacheItems(ctx)

	// Ctx()里面包含对所有Cache操作的赋值，查询不需要写Cache
	user := sys_model.SysUser{}
	err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{
		Id: userId,
	}).Scan(&user)

	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_user_info_not_exist")
	}

	user.Password = masker.MaskString(user.Password, masker.Password)

	// 查询用户所拥有的角色 (指针传递)
	_, err = s.getUserRole(ctx, &user)
	if err != nil {
		return nil, err
	}

	return s.makeMore(ctx, &user), nil
}

// GetUserListByMobileOrMail 根据手机号或者邮箱查询用户列表
func (s *sSysUser) GetUserListByMobileOrMail(ctx context.Context, info string) (*sys_model.SysUserListRes, error) {
	userModel := sys_dao.SysUser.Ctx(ctx)
	if base_verify.IsPhone(info) {
		userModel = userModel.Where(sys_do.SysUser{Mobile: strings.ToLower(info)})
	} else if base_verify.IsEmail(info) {
		userModel = userModel.Where(sys_do.SysUser{Email: strings.ToLower(info)})
	} else {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_invalid_phone_or_email_format")
	}

	userList, err := daoctl.Query[*sys_model.SysUser](userModel, nil, false)

	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_user_info_not_exist")
	}

	newList := make([]*sys_model.SysUser, 0)
	for _, user := range userList.Records {
		uData := kconv.Struct(user, &sys_model.SysUser{})
		uInfo := s.masker(s.makeMore(ctx, uData))
		newList = append(newList, uInfo)
	}

	if newList != nil {
		userList.Records = newList
	}

	return (*sys_model.SysUserListRes)(userList), nil
}

// SetUserMobile 设置用户手机号
func (s *sSysUser) SetUserMobile(ctx context.Context, newMobile, captcha, password string, userId int64) (bool, error) {
	//s.initInnerCacheItems(ctx)

	_, err := sys_service.SysSms().Verify(ctx, newMobile, captcha, base_enum.Captcha.Type.SetMobile)
	if err != nil {
		return false, err
	}

	userInfo := sys_model.SysUser{}
	err = sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{
		Id: userId,
	}).Scan(&userInfo)

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_user_info_not_exist")
	}
	if newMobile == userInfo.Mobile {
		return true, nil
	}

	// 检验密码
	user, _ := daoctl.GetByIdWithError[sys_entity.SysUser](sys_dao.SysUser.Ctx(ctx), userInfo.Id)

	pwdHash, _ := en_crypto.PwdHash(password, gconv.String(userId))

	// 业务层自定义密码加密规则
	if sys_consts.Global.CryptoPasswordFunc != nil {
		pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, password, *userInfo.SysUser)
	}

	if pwdHash != user.Password {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_invalid_login_password")
	}

	affected, err := daoctl.UpdateWithError(sys_dao.SysUser.Ctx(ctx).Data(sys_do.SysUser{Mobile: newMobile, UpdatedAt: gtime.Now()}).Where(sys_do.SysUser{
		Id: userId,
	}))

	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_mobile_set_failed", sys_dao.SysUser.Table())
	}

	// 清除redis验证码缓存
	//key := sys_enum.Sms.CaptchaType.SetMobile.Description() + "_" + newMobile
	//g.DB().GetCache().Remove(ctx, key)

	return true, nil
}

// SetUserMail 设置用户邮箱
func (s *sSysUser) SetUserMail(ctx context.Context, oldMail, newMail, captcha, password string, userId int64) (bool, error) {
	//s.initInnerCacheItems(ctx)

	_, err := sys_service.SysMails().Verify(ctx, newMail, captcha, base_enum.Captcha.Type.SetMail)
	if err != nil {
		return false, err
	}

	mailUser := sys_entity.SysUser{}
	err = sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Id: userId, Email: oldMail}).Scan(&mailUser)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invalid_original_email", sys_dao.SysUser.Table())
	}

	userInfo := sys_model.SysUser{}
	err = sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{
		Id: userId,
	}).Scan(&userInfo)

	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_user_info_not_exist")
	}
	if strings.EqualFold(newMail, userInfo.Email) {
		return true, nil
	}

	// 检验密码
	user, _ := daoctl.GetByIdWithError[sys_entity.SysUser](sys_dao.SysUser.Ctx(ctx), userInfo.Id)

	pwdHash, _ := en_crypto.PwdHash(password, gconv.String(userId))

	// 业务层自定义密码加密规则
	if sys_consts.Global.CryptoPasswordFunc != nil {
		pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, password, *userInfo.SysUser)
	}

	if pwdHash != user.Password {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "error_invalid_login_password")
	}

	affected, err := daoctl.UpdateWithError(sys_dao.SysUser.Ctx(ctx).Data(sys_do.SysUser{Email: newMail, UpdatedAt: gtime.Now()}).Where(sys_do.SysUser{
		Id: userId,
	}))

	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_email_set_failed", sys_dao.SysUser.Table())
	}

	return true, nil
}

func (s *sSysUser) getUserRole(ctx context.Context, sysUser *sys_model.SysUser, unionMainId ...int64) (*sys_model.SysUser, error) {

	if len(unionMainId) <= 0 || unionMainId[0] == 0 {
		getS := sys_service.SysSession().Get(ctx)
		if getS != nil {
			sessionUser := getS.JwtClaimsUser
			if sessionUser != nil {
				unionMainId = make([]int64, 1)
				unionMainId[0] = sessionUser.UnionMainId
			}
		}

	}

	roleIds, _ := sys_service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(sysUser.Id), sys_consts.CasbinDomain)

	sysUser.RoleNames = []string{}

	// 如果有角色信息则加载角色信息
	if len(roleIds) > 0 {
		var unionMainIdInfo int64 = 0
		if len(unionMainId) > 0 {
			unionMainIdInfo = unionMainId[0]
		}
		roles, err := sys_service.SysRole().QueryRoleList(ctx, base_model.SearchParams{
			Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
				Field:     sys_dao.SysRole.Columns().Id,
				Where:     "in",
				IsOrWhere: false,
				Value:     roleIds,
			}),
			Pagination: base_model.Pagination{},
		}, unionMainIdInfo)
		if err == nil && len(roles.Records) > 0 {
			for _, role := range roles.Records {
				sysUser.RoleNames = append(sysUser.RoleNames, role.Name)
			}
		}
	}

	return sysUser, nil
}

// Heartbeat 用户在线心跳
func (s *sSysUser) Heartbeat(ctx context.Context, userId int64) (bool, error) {
	affected, err := daoctl.UpdateWithError(
		sys_dao.SysUserDetail.Ctx(ctx).Where(sys_do.SysUserDetail{Id: userId}),
		sys_do.SysUserDetail{LastHeartbeatAt: gtime.Now()},
	)

	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_user_heartbeat_failed", sys_dao.SysUser.Table())
	}

	return true, nil
}

func (s *sSysUser) masker(user *sys_model.SysUser) *sys_model.SysUser {
	user.Password = masker.MaskString(user.Password, masker.Password)
	user.Mobile = masker.MaskString(user.Mobile, masker.MaskPhone)
	return user
}

// makeMore 处理订阅请求，获取订阅数据回调返回
func (s *sSysUser) makeMore(ctx context.Context, data *sys_model.SysUser) *sys_model.SysUser {
	base_funs.AttrMake[sys_model.SysUser](ctx,
		sys_dao.SysUser.Columns().Id,
		func() *sys_model.SysUserDetail {

			data.Detail, _ = daoctl.GetByIdWithError[sys_model.SysUserDetail](sys_dao.SysUserDetail.Ctx(ctx), data.Id)

			// 最后心跳小于 60秒内的认为在线
			if data.Detail != nil && data.Detail.LastHeartbeatAt != nil && data.Detail.LastHeartbeatAt.After(gtime.Now().Add(-time.Second*s.heartbeatTimeout)) {
				data.Detail.IsOnline = true
			}

			return data.Detail

			//resultArr, _ := daoctl.Query[sys_model.SysUserDetail](sys_dao.SysUserDetail.Ctx(ctx), nil, true)
			//result, _ := daoctl.ScanWithError[sys_entity.SysUserDetail](sys_dao.SysUserDetail.Ctx(ctx).Where(sys_do.SysUserDetail{Id: data.Id}))
			//var result *sys_model.SysUserDetail
			//for _, record := range resultArr.Records {
			//	// 最后心跳小于 60秒内的认为在线
			//	if record.LastHeartbeatAt != nil && record.LastHeartbeatAt.Before(gtime.Now().Add(-time.Second*60)) {
			//		record.IsOnline = true
			//	}
			//
			//	if record.Id == data.Id {
			//		result = &record
			//		break
			//	}
			//}
			//if result == nil {
			//	return nil
			//}
			//res := kconv.Struct[sys_model.SysUserDetail](ctx, *result)
			//if res.LastLoginIp == "" {
			//	return nil
			//}

		},
	)
	return data
}
