package sys_model

import "github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"

type SysRole struct {
	Id          int64  `p:"id"         dc:"ID"`
	Name        string `p:"name"       v:"required|length:1,16#请输入角色名称|角色名称长度限定1~16字符" dc:"角色名称"`
	IsSystem    bool   `json:"isSystem"    description:"是否默认角色，true仅能修改名称"`
	IsSys       int    `json:"isSys"       description:"是否允许删除和修改: 0允许  1禁止,并拥有默认权限"`
	Description string `p:"description"       v:"max-length:128#角色描述长度限定128字符" dc:"角色描述"`
}

type RoleListRes CollectRes[sys_entity.SysRole]
