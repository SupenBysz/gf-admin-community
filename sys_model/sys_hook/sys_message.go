package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type MessageTypeHookFunc func(ctx context.Context, messageType sys_enum.MessageType, info *sys_model.SysMessageRes) error
