package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
)

var Common = cCommon{}

type cCommon struct{}

// GetFile 获取图片文件，暴露此接口，
func (c *cCommon) GetFile(ctx context.Context, req *sys_api.GetFileReq) (res *sys_api.UploadFileRes, err error) {

	// 优先从缓存获取，缓存要是获取不到，那么从数据库加载文件信息，从而加载文件   缓存key == sign

	file, err := sys_service.File().GetFile(ctx, req.Sign, req.Path, req.Id, req.CId)

	if err != nil || file == nil {
		// 渲染默认的图片
		g.RequestFromCtx(ctx).Response.ServeFile("./resources/1x1_#00000000.png")
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(file.Src)

	//return (*sys_api.UploadFileRes)(&file.SysFile), err
	return nil, err
}
