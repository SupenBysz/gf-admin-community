package sys_enum

import (
	sys_enum_auth "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/auth"
	sys_enum_casbin "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/casbin"
	sys_enum_oss "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/oss"
	sys_enum_upload "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/upload"
	sys_enum_user "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/user"
)

type (
	OssType sys_enum_oss.OssTypeEnum

	AuthActionType sys_enum_auth.ActionTypeEnum

	UploadEventState sys_enum_upload.EventStateEnum

	UserEvent sys_enum_user.EventEnum
	UserType  sys_enum_user.TypeEnum
	UserState sys_enum_user.StateEnum

	CabinEvent sys_enum_casbin.EventEnum
)

var (
	Oss    = sys_enum_oss.Oss
	Auth   = sys_enum_auth.Auth
	Upload = sys_enum_upload.Upload
	User   = sys_enum_user.User
	Casbin = sys_enum_casbin.Casbin
)
