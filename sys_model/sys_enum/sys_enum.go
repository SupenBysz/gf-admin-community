package sys_enum

import (
	sys_enum_announcement "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/announcement"
	sys_enum_audit "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/audit"
	sys_enum_auth "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/auth"
	sys_enum_business "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/business"
	sys_enum_casbin "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/casbin"
	sys_enum_category "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/category"
	sys_enum_comment "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/comment"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/delivery"
	sys_enum_file "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/file"
	sys_enum_industry "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/industry"
	sys_enum_invite "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/invite"
	sys_enum_license "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/license"
	sys_member_level "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/member_level"
	sys_enum_menu "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/menu"
	sys_enum_messages "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/message"
	sys_enum_organization "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/organization"
	sys_enum_oss "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/oss"
	sys_enum_permissions "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/permissions"
	sys_enum_role "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/role"
	sys_enum_upload "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/upload"
	sys_enum_user "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/user"
)

type (
	BusinessType = sys_enum_business.TypeEnum

	OssType = sys_enum_oss.OssTypeEnum

	AuthActionType = sys_enum_auth.ActionTypeEnum

	UploadEventState = sys_enum_upload.EventStateEnum

	UserEvent = sys_enum_user.EventEnum
	UserType  = sys_enum_user.TypeEnum
	UserState = sys_enum_user.StateEnum

	CabinEvent = sys_enum_casbin.EventEnum

	PermissionMatchMode = sys_enum_permissions.MatchModeEnum

	AuditAction = sys_enum_audit.AuditStateEnum
	AuditEvent  = sys_enum_audit.EventEnum

	// InviteType 邀约类型
	InviteType = sys_enum_invite.TypeEnum

	// InviteState 邀约状态
	InviteState = sys_enum_invite.StateEnum

	// RoleMemberChange 角色成员事件
	RoleMemberChange = sys_enum_role.ChangeEnum

	// LicenseState 个人资质状态
	LicenseState = sys_enum_license.StateEnum
	// LicenseAuthType 个人资质认证类型
	LicenseAuthType = sys_enum_license.AuthTypeEnum

	// MessageType 消息类型
	MessageType = sys_enum_messages.TypeEnum
	// MessageSceneType 消息场景
	MessageSceneType = sys_enum_messages.SceneTypeEnum
	// MessageState 消息状态
	MessageState = sys_enum_messages.StateEnum

	// AnnouncementState 公告状态
	AnnouncementState = sys_enum_announcement.StateEnum

	MemberLevelEvent = sys_member_level.EventEnum
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

	MemberLevel = sys_member_level.MemberLevel

	Organization = sys_enum_organization.Organization
	Role         = sys_enum_role.Role
	Permissions  = sys_enum_permissions.Permissions

	Menu = sys_enum_menu.Menu

	Audit = sys_enum_audit.Audit

	Invite = sys_enum_invite.Invite

	// Industry 行业类别
	Industry = sys_enum_industry.Industry

	// License 个人资质
	License = sys_enum_license.License

	// Message 消息
	Message = sys_enum_messages.Message

	// Announcement 公告
	Announcement = sys_enum_announcement.Announcement

	// Category 类目
	Category = sys_enum_category.Category

	// Delivery 物流公司
	Delivery = delivery.Delivery

	// Comment 评论
	Comment = sys_enum_comment.Comment
)
