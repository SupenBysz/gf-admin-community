package sys_consts

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type global struct {
	UserDefaultType          sys_enum.UserType
	UserDefaultState         sys_enum.UserState
	NotAllowLoginUserTypeArr *garray.SortedIntArray
	AllowLoginUserTypeArr    *garray.SortedIntArray
	LogLevelToDatabaseArr    *garray.SortedIntArray
	ApiPreFix                string
	OrmCacheConf             []*sys_model.TableCacheConf
	PermissionTree           []base_permission.IPermission // PermissionTree 权限信息定义
	Searcher                 *xdb.Searcher
	EmailConfig              sys_model.EmailConfig

	// 密码加密
	CryptoPasswordFunc func(ctx context.Context, passwordStr string, user ...sys_entity.SysUser) (pwdEncode string)

	// 注册是否需要邀约码
	RegisterIsNeedInviteCode bool
	// 邀约码默认时长天数
	InviteCodeExpireDay int
	// 邀约码次数上限
	InviteCodeMaxActivateNumber int
}

var (
	Global = global{
		UserDefaultType:             sys_enum.User.Type.SuperAdmin,
		UserDefaultState:            sys_enum.User.State.Normal,
		NotAllowLoginUserTypeArr:    garray.NewSortedIntArray(),
		AllowLoginUserTypeArr:       garray.NewSortedIntArray(),
		LogLevelToDatabaseArr:       garray.NewSortedIntArray(),
		ApiPreFix:                   "",
		OrmCacheConf:                []*sys_model.TableCacheConf{},
		PermissionTree:              []base_permission.IPermission{},
		CryptoPasswordFunc:          nil,
		EmailConfig:                 sys_model.EmailConfig{},
		RegisterIsNeedInviteCode:    false,
		InviteCodeExpireDay:         0,
		InviteCodeMaxActivateNumber: 0,
	}
)
