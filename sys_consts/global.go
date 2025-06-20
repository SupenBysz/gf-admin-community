package sys_consts

import (
	"context"
	"strings"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type global struct {
	LogLevelToDatabaseArr *garray.SortedIntArray
	ApiPreFix             string
	ClientConfig          []sys_model.ClientConfig
	OrmCacheConf          []*sys_model.TableCacheConf
	PermissionTree        []base_permission.IPermission // PermissionTree 权限信息定义
	Searcher              *xdb.Searcher
	EmailConfig           sys_model.EmailConfig

	// 密码加密
	CryptoPasswordFunc func(ctx context.Context, passwordStr string, user ...sys_entity.SysUser) (pwdEncode string)
}

func (s global) GetClientConfig(ctx context.Context) (*sys_model.ClientConfig, error) {
	xClient := ghttp.RequestFromCtx(ctx).Header.Get("X-CLIENT-ID")

	for _, v := range s.ClientConfig {
		if strings.EqualFold(v.XClientToken, xClient) {
			return &v, nil
		}
	}

	return nil, gerror.NewCode(gcode.CodeNotAuthorized, "error_client_info_incorrect")
}

var (
	Global = global{
		ClientConfig:          []sys_model.ClientConfig{},
		LogLevelToDatabaseArr: garray.NewSortedIntArray(),
		ApiPreFix:             "",
		OrmCacheConf:          []*sys_model.TableCacheConf{},
		PermissionTree:        []base_permission.IPermission{},
		CryptoPasswordFunc:    nil,
		EmailConfig:           sys_model.EmailConfig{},
	}
)
