package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysRole struct {
	Id          int64  `p:"id"         dc:"ID"`
	Name        string `p:"name"       v:"required|length:1,16#请输入角色名称|角色名称长度限定1~16字符" dc:"角色名称"`
	IsSystem    bool   `json:"isSystem"    dc:"是否默认角色，true仅能修改名称，不允许删除和修改"`
	Description string `p:"description"       v:"max-length:128#角色描述长度限定128字符" dc:"角色描述"`
	UnionMainId int64  `json:"-"  dc:"主体id"`
}

type SysRoleRes sys_entity.SysRole

type RoleListRes base_model.CollectRes[*sys_entity.SysRole]

type RoleInfo sys_entity.SysRole
