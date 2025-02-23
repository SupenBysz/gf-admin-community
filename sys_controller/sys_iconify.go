package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// SysIconify 图标
var SysIconify = cSysIconify{}

type cSysIconify struct{}

func (c *cSysIconify) Search(ctx context.Context, req *sys_api.SearchReq) (*api_v1.MapRes, error) {
	url := "https://api.iconify.design/search?query=" + req.Keyword + "&pretty=1&limit=" + gconv.String(req.Limit)

	body := g.Client().GetVar(ctx, url).Map()

	return (*api_v1.MapRes)(&body), nil
}

func (c *cSysIconify) GetIcons(ctx context.Context, req *sys_api.GetIconsReq) (*api_v1.MapRes, error) {
	url := "https://api.iconify.design/collections?prefix=" + req.Prefix + "&pretty=1"

	body := g.Client().GetVar(ctx, url).Map()

	return (*api_v1.MapRes)(&body), nil
}

func (c *cSysIconify) GetCollections(ctx context.Context, req *sys_api.GetCollectionsReq) (*api_v1.MapRes, error) {
	prefixes := ""
	if req.Prefixes != nil {
		prefixes = "?prefixes=" + *req.Prefixes + "&pretty=1"
	}
	url := "https://api.iconify.design/collections" + prefixes

	body := g.Client().GetVar(ctx, url).Map()

	return (*api_v1.MapRes)(&body), nil
}
