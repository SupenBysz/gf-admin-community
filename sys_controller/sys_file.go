package sysController

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// SysFile 文件
var SysFile = cSysFile{}

type cSysFile struct{}

// Upload 上传文件
func (c *cSysFile) Upload(ctx context.Context, req *sys_api.UploadReq) (res *sys_api.UploadRes, err error) {
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

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
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().UploadIDCard(ctx, req.OCRIDCardFileUploadInput, userId)

	return (*sys_api.IDCardWithOCRRes)(result), err
}

// UploadBusinessLicenseWithOCR 上传营业执照
func (c *cSysFile) UploadBusinessLicenseWithOCR(ctx context.Context, req *sys_api.UploadBusinessLicenseWithOCRReq) (*sys_api.UploadBusinessLicenseWithOCRRes, error) {
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().UploadBusinessLicense(ctx, req.OCRBusinessLicense, userId)

	return (*sys_api.UploadBusinessLicenseWithOCRRes)(result), err
}

// UploadBankCardWithOCR 上传银行卡
func (c *cSysFile) UploadBankCardWithOCR(ctx context.Context, req *sys_api.UploadBankCardWithOCRReq) (res *sys_api.BankCardWithOCRRes, err error) {
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

	if userId <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "请登陆后再操作"), "", "File")
	}

	result, err := sys_service.File().UploadBankCard(ctx, req.BankCardWithOCRInput, userId)

	return (*sys_api.BankCardWithOCRRes)(result), err
}

// GetFileById 通过id获取图片
func (c *cSysFile) GetFileById(ctx context.Context, _ *sys_api.GetFileByIdReq) (res *api_v1.MapRes, err error) {

	// 获取图片id 还有图片类型 临时还是永久
	id := g.RequestFromCtx(ctx).GetQuery("id").Int64()
	v := g.RequestFromCtx(ctx).GetQuery("v").Int() // v=1 永久 v=0 临时

	file, err := sys_service.File().GetFileById(ctx, id, v)

	return &file, err
}
