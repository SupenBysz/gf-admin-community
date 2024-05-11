package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

var Common = cCommon{}

type cCommon struct{}

// GetFile 获取图片文件，暴露此接口，
func (c *cCommon) GetFile(ctx context.Context, req *sys_api.GetFileReq) (res *sys_api.UploadFileRes, err error) {

	// 优先从缓存获取，缓存要是获取不到，那么从数据库加载文件信息，从而加载文件   缓存key == sign

	file, err := sys_service.File().GetFile(ctx, req.Sign, req.Path, req.Id, req.CId)

	if err != nil || file == nil {
		// 渲染默认的图片
		g.RequestFromCtx(ctx).Response.ServeFile("./resource/1x1_#00000000.png")
		return nil, err
	}

	if gfile.GetContents(file.Src) != "" { // 本地文件访问路径
		//  ServeFile 为响应提供文件 (通过给定文件路径path，ServeFile方法将会自动识别文件格式，如果是目录或者文本内容将会直接展示文件内容。如果path参数为目录，那么第二个参数allowIndex控制是否可以展示目录下的文件列表。)
		g.RequestFromCtx(ctx).Response.ServeFile(file.Src) // 图片、文本文件txt、pdf也是直接输出、docs文档

		// ServeFileDownload是相对使用频率比较高的方法，用于直接引导客户端下载指定路径的文件，并可以重新给定下载的文件名称。
		//g.RequestFromCtx(ctx).Response.ServeFileDownload(file.Src, "testDownload.jpg")

	} else { // oss 访问路径
		g.RequestFromCtx(ctx).Response.RedirectTo(file.Src) // 直接引导客户端下载指定路径的文件
	}

	//return (*sys_api.UploadFileRes)(&file.SysFile), err
	return nil, err
}
