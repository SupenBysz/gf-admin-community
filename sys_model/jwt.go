package sys_model

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	sys_entity.SysUser
	UnionMainId int64 `json:"unionMainId"    description:"主体id"`
	IsAdmin     bool
	jwt.RegisteredClaims
}

type JwtHookFunc func(ctx context.Context, claims *JwtCustomClaims) (*JwtCustomClaims, error)

type JwtHookInfo HookEventType[sys_enum.UserType, JwtHookFunc]
