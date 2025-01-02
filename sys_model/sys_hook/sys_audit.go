package sys_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type AuditHookFunc func(ctx context.Context, state sys_enum.AuditEvent, info *sys_model.AuditRes) error
type AuditHookInfo struct {
	Key      sys_enum.AuditEvent
	Value    AuditHookFunc
	Category int `json:"category" dc:"业务类别"`
}
