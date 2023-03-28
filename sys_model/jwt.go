package sys_model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	SysUser
	UnionMainId  int64 `json:"unionMainId"    description:"主体id"`
	ParentId     int64 `json:"parentId"    description:"上级主体id"`
	IsAdmin      bool
	IsSuperAdmin bool
	jwt.RegisteredClaims
}
