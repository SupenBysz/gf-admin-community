package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type cSysUEditor struct{}

var SysUEditor = cSysUEditor{}

func (c *cSysUEditor) UEditor(ctx context.Context, req *sys_api.UEditorReq) (*api_v1.MapRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	return sys_service.UEditor().UEditor(ctx, user.Id, user.UnionMainId, req.FileUploadInput)
}
