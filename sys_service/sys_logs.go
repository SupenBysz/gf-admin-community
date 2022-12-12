// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type (
	ISysLogs interface {
		Write(ctx context.Context, err error, info sys_entity.SysLogs) error
		Error(ctx context.Context, err error, info sys_entity.SysLogs) error
		ErrorSimple(ctx context.Context, err error, context string, category string) error
		Info(ctx context.Context, err error, info sys_entity.SysLogs) error
		InfoSimple(ctx context.Context, err error, context string, category string) error
		Warn(ctx context.Context, err error, info sys_entity.SysLogs) error
		WarnSimple(ctx context.Context, err error, context string, category string) error
	}
)

var (
	localSysLogs ISysLogs
)

func SysLogs() ISysLogs {
	if localSysLogs == nil {
		panic("implement not found for interface ISysLogs, forgot register?")
	}
	return localSysLogs
}

func RegisterSysLogs(i ISysLogs) {
	localSysLogs = i
}