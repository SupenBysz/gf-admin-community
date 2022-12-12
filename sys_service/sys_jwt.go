// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IJwt interface {
		GenerateToken(user *sys_entity.SysUser) (*sys_model.TokenInfo, error)
		CreateToken(claims *sys_model.JwtCustomClaims) (string, error)
		RefreshToken(oldToken string, claims *sys_model.JwtCustomClaims) (string, error)
		CustomMiddleware(r *ghttp.Request)
		Middleware(r *ghttp.Request)
		ParseToken(tokenString string) (*sys_model.JwtCustomClaims, error)
	}
)

var (
	localJwt IJwt
)

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}