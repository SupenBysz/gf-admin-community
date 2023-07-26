package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_permission"
)

type SysPermission struct {
	Id          int64  `json:"id"             dc:"ID" v:"integer"`
	ParentId    int64  `json:"parentId"       dc:"父级ID" v:"min:0#必须是正整数，该属性创建后不支持修改"`
	Name        string `json:"name"           dc:"名称" v:"max-length:64#仅支持最大字符长度64"`
	Description string `json:"description"    dc:"描述" v:"max-length:128#仅支持最大字符长度128"`
	Identifier  string `json:"identifier"     dc:"标识符"`
	Type        int    `json:"type"           dc:"类型：1api、2menu"`
	MatchMode   int    `json:"matchMode"      dc:"匹配模式：ID：0，标识符：1"`
	IsShow      int    `json:"isShow"         dc:"是否显示：0不显示 1显示"`
	Sort        int    `json:"sort"           dc:"排序"`
}

type SysPermissionTree struct {
	*sys_entity.SysPermission
	Children []base_permission.IPermission `json:"children"       dc:"下级权限"`
}

type SysPermissionInfoRes sys_entity.SysPermission
type SysPermissionInfoListRes base_model.CollectRes[*sys_entity.SysPermission]
type SysPermissionInfoTreeRes []base_permission.IPermission
type MyPermissionListRes []*sys_entity.SysPermission

// 实现权限树接口

func (d *SysPermissionTree) GetIsEqual(father base_permission.IPermission, childId base_permission.IPermission) bool {
	return father.GetId() == childId.GetParentId()
}
func (d *SysPermissionTree) SetChild(father base_permission.IPermission, branchArr []base_permission.IPermission) {
	father.SetItems(branchArr)
}
func (d *SysPermissionTree) RetFather(father base_permission.IPermission) bool {
	// 顶级的ParentId这块可以看一下保存的时候ParentId 默认值是多少
	return father.GetParentId() == 0
}

// 实现权限接口

func (d *SysPermissionTree) GetId() int64 {
	return d.Id

}
func (d *SysPermissionTree) GetParentId() int64 {
	return d.ParentId

}
func (d *SysPermissionTree) GetName() string {
	return d.Name

}
func (d *SysPermissionTree) GetDescription() string {
	return d.Description

}
func (d *SysPermissionTree) GetIdentifier() string {
	return d.Identifier

}
func (d *SysPermissionTree) GetType() int {
	return d.Type

}
func (d *SysPermissionTree) GetMatchMode() int {
	return d.MatchMode

}
func (d *SysPermissionTree) GetIsShow() int {
	return d.IsShow

}
func (d *SysPermissionTree) GetSort() int {
	return d.Sort

}
func (d *SysPermissionTree) GetItems() []base_permission.IPermission {
	return d.Children

}
func (d *SysPermissionTree) GetData() interface{} {
	return d
}

func (d *SysPermissionTree) SetId(val int64) base_permission.IPermission {
	d.Id = val
	return d
}
func (d *SysPermissionTree) SetParentId(val int64) base_permission.IPermission {
	d.ParentId = val
	return d

}
func (d *SysPermissionTree) SetName(val string) base_permission.IPermission {
	d.Name = val
	return d

}
func (d *SysPermissionTree) SetDescription(val string) base_permission.IPermission {
	d.Description = val
	return d

}
func (d *SysPermissionTree) SetIdentifier(val string) base_permission.IPermission {
	d.Identifier = val
	return d

}
func (d *SysPermissionTree) SetType(val int) base_permission.IPermission {
	d.Type = val
	return d

}
func (d *SysPermissionTree) SetMatchMode(val int) base_permission.IPermission {
	d.MatchMode = val
	return d

}
func (d *SysPermissionTree) SetIsShow(val int) base_permission.IPermission {
	d.IsShow = val
	return d

}
func (d *SysPermissionTree) SetSort(val int) base_permission.IPermission {
	d.Sort = val
	return d

}
func (d *SysPermissionTree) SetItems(val []base_permission.IPermission) base_permission.IPermission {
	d.Children = val
	return d
}

//var PFactory func(data base_permission.IPermission) base_permission.IPermission
//var TFactory TPermission[base_permission.IPermission]
