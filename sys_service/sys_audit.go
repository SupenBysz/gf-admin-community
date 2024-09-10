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
		// GetAuditLatestByUnionMainId 获取最新的业务个人审核信息 (针对主体资质)
		GetAuditLatestByUnionMainId(ctx context.Context, unionMainId int64) *sys_entity.SysAudit
		// GetAuditLatestByUserId 获取最新的业务个人审核信息
		GetAuditLatestByUserId(ctx context.Context, userId int64) *sys_entity.SysAudit
		// CreateAudit 创建审核信息 // TODO 创建审核信息后，需要通过Hook将temp/upload 中的文件迁移到业务层的指定目录，例如 resource/upload
		CreateAudit(ctx context.Context, info sys_model.CreateAudit) (*sys_entity.SysAudit, error)
		// UpdateAudit 处理审核信息
		UpdateAudit(ctx context.Context, id int64, state int, reply string, auditUserId int64) (bool, error)
		// SetUnionMainId  设置审核关联的主体Id
		SetUnionMainId(ctx context.Context, id int64, unionMainId int64) (bool, error)
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
