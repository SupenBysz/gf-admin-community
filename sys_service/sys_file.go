// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
)

type (
	IFile interface {
		// InstallHook 安装Hook
		InstallHook(state sys_enum.UploadEventState, hookFunc sys_hook.FileHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除Hook
		CleanAllHook()
		// Upload 统一上传文件
		Upload(ctx context.Context, in sys_model.FileUploadInput) (*sys_entity.SysFile, error)
		// GetUploadFile 根据上传ID 获取上传文件信息
		GetUploadFile(ctx context.Context, uploadId int64, userId int64, message ...string) (*sys_model.FileInfo, error)
		// SaveFile 保存文件
		SaveFile(ctx context.Context, storageAddr string, info *sys_model.FileInfo) (*sys_model.FileInfo, error)
		// UploadIDCard 上传身份证照片
		UploadIDCard(ctx context.Context, in sys_model.OCRIDCardFileUploadInput) (*sys_model.IDCardWithOCR, error)
		// UploadBankCard 上传银行卡照片
		UploadBankCard(ctx context.Context, in sys_model.BankCardWithOCRInput) (*sys_model.BankCardWithOCR, error)
		// UploadBusinessLicense 上传营业执照照片
		UploadBusinessLicense(ctx context.Context, in sys_model.OCRBusinessLicense) (*sys_model.BusinessLicenseWithOCR, error)
		// DownLoadFile 下载文件
		DownLoadFile(ctx context.Context, savePath string, url string) (string, error)
		// GetFileById 根据id获取并返回文件信息
		GetFileById(ctx context.Context, id int64, errorMessage string) (*sys_model.FileInfo, error)
		// MakeFileUrl 图像id换取url: 拼接三个参数,缓存fileInfo、然后返回url + 三参
		MakeFileUrl(ctx context.Context, id int64, styleStr ...string) string
		// MakeFileUrlByPath 文件path换取url: 拼接三个参数,缓存签名数据、然后返回url + 三参
		MakeFileUrlByPath(ctx context.Context, path string) string
		// GetFile 获取图片 公开  (srcBase64 + srcMd5 + fileId) ==> md5加密
		GetFile(ctx context.Context, sign, srcBase64 string, id int64, cId int64, styleStr ...string) (*sys_model.FileInfo, error)
		// UploadPicture 上传图片并审核
		UploadPicture(ctx context.Context, input sys_model.PictureWithOCRInput) (*sys_model.PictureWithOCR, error)
		// GetOssFileSingUrl 获取文件的签名访问URL
		GetOssFileSingUrl(ctx context.Context, bucketName, objectKey string, styleStr ...string) (string, error)
		// GetOssFileWithURL 根据文件的签名访问URL获取文件
		GetOssFileWithURL(ctx context.Context, bucketName, filePath, singUrl string) (bool, error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
