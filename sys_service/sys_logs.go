// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type (
	ISysLogs interface {
		// Write 写日志
		Write(ctx context.Context, err error, info sys_entity.SysLogs) error
		// Write 写错误日志
		Error(ctx context.Context, err error, info sys_entity.SysLogs) error
		// ErrorSimple 写错误日志
		ErrorSimple(ctx context.Context, err error, context string, category string) error
		// Info 写日志信息
		Info(ctx context.Context, err error, info sys_entity.SysLogs) error
		// InfoSimple 写日志信息
		InfoSimple(ctx context.Context, err error, context string, category string) error
		// Warn 写警示日志
		Warn(ctx context.Context, err error, info sys_entity.SysLogs) error
		// WarnSimple 写警示日志
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
