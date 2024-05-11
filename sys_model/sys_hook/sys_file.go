package sys_hook

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type FileHookFunc sys_model.HookFunc[sys_enum.UploadEventState, *sys_entity.SysFile]
type FileHookInfo sys_model.HookEventType[sys_enum.UploadEventState, FileHookFunc]
