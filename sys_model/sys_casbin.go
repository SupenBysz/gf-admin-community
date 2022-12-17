package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
)

type ReqCasbin struct {
	UserId    int64  `p:"userId"`
	Domain    string `p:"domain"`
	Interface string `p:"i"`
	Action    string `p:"a"`
}

type CasbinCheckObject struct {
	SysUser       sys_entity.SysUser
	SysPermission sys_entity.SysPermission
}
type CasbinHookFunc HookFunc[sys_enum.CabinEvent, CasbinCheckObject]
type CasbinHookInfo HookEventType[sys_enum.CabinEvent, CasbinHookFunc]
