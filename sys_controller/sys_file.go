package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
)

// SysFile 文件
var SysFile = cSysFile{}

type cSysFile struct{}

// Upload 上传文件
func (c *cSysFile) Upload(ctx context.Context, req *sys_api.UploadReq) (res *sys_api.UploadFileRes, err error) {
	result, err := sys_service.File().Upload(ctx, req.FileUploadInput)
	if err != nil {
		return nil, err
	}

	return (*sys_api.UploadFileRes)(result), nil
}

// UploadIDCardWithOCR 上传身份证
func (c *cSysFile) UploadIDCardWithOCR(ctx context.Context, req *sys_api.UploadIDCardWithOCRReq) (res *sys_api.IDCardWithOCRRes, err error) {

	result, err := sys_service.File().UploadIDCard(ctx, req.OCRIDCardFileUploadInput)

	return (*sys_api.IDCardWithOCRRes)(result), err
}

// UploadBusinessLicenseWithOCR 上传营业执照
func (c *cSysFile) UploadBusinessLicenseWithOCR(ctx context.Context, req *sys_api.UploadBusinessLicenseWithOCRReq) (*sys_api.UploadBusinessLicenseWithOCRRes, error) {
	result, err := sys_service.File().UploadBusinessLicense(ctx, req.OCRBusinessLicense)

	return (*sys_api.UploadBusinessLicenseWithOCRRes)(result), err
}

// UploadBankCardWithOCR 上传银行卡
func (c *cSysFile) UploadBankCardWithOCR(ctx context.Context, req *sys_api.UploadBankCardWithOCRReq) (res *sys_api.BankCardWithOCRRes, err error) {
	result, err := sys_service.File().UploadBankCard(ctx, req.BankCardWithOCRInput)

	return (*sys_api.BankCardWithOCRRes)(result), err
}

// GetFileById 通过图片ID获取图片文件
func (c *cSysFile) GetFileById(ctx context.Context, _ *sys_api.GetFileByIdReq) (res *sys_api.UploadFileRes, err error) {
	// 获取图片id 还有图片类型 临时还是永久
	id := g.RequestFromCtx(ctx).GetQuery("id").Int64()

	file, err := sys_service.File().GetFileById(ctx, id, "文件加载失败")

	if err != nil {
		return nil, err
	}

	// 加载显示图片
	g.RequestFromCtx(ctx).Response.ServeFile(file.Src)

	return (*sys_api.UploadFileRes)(&file.SysFile), err
}

// UploadPicture 上传图片并审核
func (c *cSysFile) UploadPicture(ctx context.Context, req *sys_api.UploadPictureReq) (res *sys_api.UploadPictureRes, err error) {
	result, err := sys_service.File().UploadPicture(ctx, req.PictureWithOCRInput)

	return (*sys_api.UploadPictureRes)(result), err
}

// SysFileAllowAnonymous 文件
var SysFileAllowAnonymous = cSysFileAllowAnonymous{}

type cSysFileAllowAnonymous struct{}

func (c *cSysFileAllowAnonymous) GetAnyFileById(ctx context.Context, _ *sys_api.GetFileAllowAnonymousByIdReq) (res *sys_api.UploadFileRes, err error) {
	// 获取图片id 还有图片类型 临时还是永久
	id := g.RequestFromCtx(ctx).GetQuery("id").Int64()

	file, err := sys_service.File().GetAnyFileById(ctx, id, "文件加载失败")

	if err != nil {
		return nil, err
	}

	// 加载显示图片
	g.RequestFromCtx(ctx).Response.ServeFile(file.Src)

	return (*sys_api.UploadFileRes)(&file.SysFile), err
}
