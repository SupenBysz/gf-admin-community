package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
)

type UEditorReq struct {
	g.Meta `path:"/plus" method:"GET,POST" summary:"UEditor富文本支持" tags:"工具"`
	*sys_model.FileUploadInput
	File       *ghttp.UploadFile `json:"file"  dc:"请选择文件，以form-data方式提交"`    // 上传文件对象
}
