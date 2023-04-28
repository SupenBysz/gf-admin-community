package sys_model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	SysUser
	UnionMainId  int64 `json:"unionMainId"    dc:"主体id"`
	ParentId     int64 `json:"parentId"    dc:"上级主体id"`
	IsAdmin      bool
	IsSuperAdmin bool
	jwt.RegisteredClaims
}
