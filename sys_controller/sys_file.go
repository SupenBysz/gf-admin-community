package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SysFile 文件
var SysFile = cSysFile{}

type cSysFile struct{}

// Upload 上传文件
func (c *cSysFile) Upload(ctx context.Context, req *sys_api.UploadReq) (res *sys_api.UploadRes, err error) {
	// sys_service.Jwt().CustomMiddleware(ghttp.RequestFromCtx(ctx))

	userId := sys_service.BizCtx().Get(ctx).ClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().Upload(ctx, req.FileUploadInput, userId)
	if err != nil {
		return nil, err
	}

	return (*sys_api.UploadRes)(result), nil
}

// UploadIDCardWithOCR 上传身份证
func (c *cSysFile) UploadIDCardWithOCR(ctx context.Context, req *sys_api.UploadIDCardWithOCRReq) (res *sys_api.IDCardWithOCRRes, err error) {
	// sys_service.Jwt().CustomMiddleware(ghttp.RequestFromCtx(ctx))

	userId := sys_service.BizCtx().Get(ctx).ClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().UploadIDCard(ctx, req.OCRIDCardFileUploadInput, userId)

	return (*sys_api.IDCardWithOCRRes)(result), err
}

// UploadBusinessLicenseWithOCR 上传营业执照
func (c *cSysFile) UploadBusinessLicenseWithOCR(ctx context.Context, req *sys_api.UploadBusinessLicenseWithOCRReq) (*sys_api.UploadBusinessLicenseWithOCRRes, error) {
	sys_service.Jwt().CustomMiddleware(ghttp.RequestFromCtx(ctx))

	userId := sys_service.BizCtx().Get(ctx).ClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().UploadBusinessLicense(ctx, req.OCRBusinessLicense, userId)

	return (*sys_api.UploadBusinessLicenseWithOCRRes)(result), err
}

// UploadBankCardWithOCR 上传银行卡
func (c *cSysFile) UploadBankCardWithOCR(ctx context.Context, req *sys_api.UploadBankCardWithOCRReq) (res *sys_api.BankCardWithOCRRes, err error) {
	userId := sys_service.BizCtx().Get(ctx).ClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().UploadBankCard(ctx, req.BankCardWithOCRInput, userId)

	return (*sys_api.BankCardWithOCRRes)(result), err
}

// GetFileById 通过id获取图片
func (c *cSysFile) GetFileById(ctx context.Context, req *sys_api.GetFileByIdReq) (res *sys_api.GetFileRes, err error) {

	file, err := sys_service.File().GetFileById(ctx, req.Id)

	return (*sys_api.GetFileRes)(file), err
}
