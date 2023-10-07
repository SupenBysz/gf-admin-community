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
	"github.com/kysion/base-library/base_model"
)

type (
	ISysAudit interface {
		// InstallHook 安装Hook
		InstallHook(state sys_enum.AuditEvent, category int, hookFunc sys_hook.AuditHookFunc) int64
		// UnInstallHook 卸载Hook
		UnInstallHook(savedHookId int64)
		// CleanAllHook 清除所有Hook
		CleanAllHook()
		// QueryAuditList 获取审核信息列表
		QueryAuditList(ctx context.Context, filter *base_model.SearchParams) (*sys_model.AuditListRes, error)
		// GetAuditById 根据ID获取审核信息
		GetAuditById(ctx context.Context, id int64) *sys_entity.SysAudit
		// GetAuditByLatestUnionMainId 获取最新的业务个人审核信息
		GetAuditByLatestUnionMainId(ctx context.Context, unionMainId int64) *sys_entity.SysAudit
		// GetAuditByLatestUserId 获取最新的业务个人审核信息
		GetAuditByLatestUserId(ctx context.Context, userId int64) *sys_entity.SysAudit
		// CreateAudit 创建审核信息
		CreateAudit(ctx context.Context, info sys_model.CreateAudit) (*sys_entity.SysAudit, error)
		// UpdateAudit 处理审核信息
		UpdateAudit(ctx context.Context, id int64, state int, reply string, auditUserId int64) (bool, error)
	}
)

var (
	localSysAudit ISysAudit
)

func SysAudit() ISysAudit {
	if localSysAudit == nil {
		panic("implement not found for interface ISysAudit, forgot register?")
	}
	return localSysAudit
}

func RegisterSysAudit(i ISysAudit) {
	localSysAudit = i
}
