package sys_enum

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/audit"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/auth"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/business"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/casbin"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/file"
	sys_enum_invite "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/invite"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/license"
	sys_enum_menu "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/menu"
	sys_enum_organization "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/organization"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/oss"
	sys_enum_permissions "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/permissions"
	sys_enum_role "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/role"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/upload"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/user"
)

type (
	BusinessType sys_enum_business.TypeEnum

	OssType sys_enum_oss.OssTypeEnum

	AuthActionType sys_enum_auth.ActionTypeEnum

	UploadEventState sys_enum_upload.EventStateEnum

	UserEvent sys_enum_user.EventEnum
	UserType  sys_enum_user.TypeEnum
	UserState sys_enum_user.StateEnum

	CabinEvent sys_enum_casbin.EventEnum

	PermissionMatchMode sys_enum_permissions.MatchModeEnum

	AuditAction sys_enum_audit.ActionEnum
	AuditEvent  sys_enum_audit.EventEnum

	// InviteType 邀约类型
	InviteType sys_enum_invite.TypeEnum

	// InviteState 邀约状态
	InviteState sys_enum_invite.StateEnum
)

var (
	Business = sys_enum_business.Business

	Oss = sys_enum_oss.Oss

	// Captcha = sys_enum_captcha.Captcha 迁移到了base-library

	Auth   = sys_enum_auth.Auth
	Upload = sys_enum_upload.Upload
	User   = sys_enum_user.User
	Casbin = sys_enum_casbin.Casbin
	File   = sys_enum_file.File

	Organization = sys_enum_organization.Organization
	Role         = sys_enum_role.Role
	Permissions  = sys_enum_permissions.Permissions

	Menu = sys_enum_menu.Menu

	Audit   = sys_enum_audit.Audit
	License = sys_enum_license.License

	Invite = sys_enum_invite.Invite
)
