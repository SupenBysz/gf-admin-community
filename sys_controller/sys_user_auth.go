package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// UserAuth 鉴权 - 需要登陆
var UserAuth = cUserAuth{}

type cUserAuth struct{}

// RefreshJwtToken 刷新用户jwtToken
func (c *cUserAuth) RefreshJwtToken(ctx context.Context, _ *sys_api.RefreshJwtTokenReq) (res *sys_model.LoginRes, err error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	res, err = sys_service.SysAuth().RefreshJwtToken(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
