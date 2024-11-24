// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
)

type (
	IUEditor interface {
		UEditor(ctx context.Context, userId int64, unionMainId int64, fileUploadInput *sys_model.FileUploadInput) (*api_v1.MapRes, error)
		UploadScrawl(ctx context.Context) (*api_v1.MapRes, error)
	}
)

var (
	localUEditor IUEditor
)

func UEditor() IUEditor {
	if localUEditor == nil {
		panic("implement not found for interface IUEditor, forgot register?")
	}
	return localUEditor
}

func RegisterUEditor(i IUEditor) {
	localUEditor = i
}
