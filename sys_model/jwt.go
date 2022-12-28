package sys_model

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Id          int64       `json:"id"               description:"ID"`
	Username    string      `json:"username"         description:"用户名"`
	State       int         `json:"state"            description:"状态：0未激活、1正常、-1封号、-2异常、-3已注销"`
	Type        int         `json:"type"             description:"用户类型，0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心"`
	UnionMainId int64       `json:"union_main_id" description:"主体id"`
	CreatedAt   *gtime.Time `json:"createdAt"        description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"        description:""`
	jwt.RegisteredClaims
}

type JwtHookFunc func(ctx context.Context, userInfo sys_entity.SysUser) (int64, error)

type JwtHookInfo HookEventType[sys_enum.UserType, JwtHookFunc]
