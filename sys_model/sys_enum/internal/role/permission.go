package organization

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type PermissionTypeEnum = *sys_model.SysPermissionTree

type permissionType struct {
	ViewDetail    PermissionTypeEnum
	ViewMember    PermissionTypeEnum
	List          PermissionTypeEnum
	Update        PermissionTypeEnum
	Delete        PermissionTypeEnum
	Create        PermissionTypeEnum
	SetMember     PermissionTypeEnum
	SetPermission PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail:    permission.New(5948684761759813, "ViewDetail", "查看角色", "查看某个角色"),
	ViewMember:    permission.New(5948684761759823, "ViewMember", "查看角色成员", "查看某个角色下的用户"),
	List:          permission.New(5948684761759814, "List", "角色列表", "查看所有角色"),
	Update:        permission.New(5948684761759815, "Update", "更新角色信息", "更新某个角色信息"),
	Delete:        permission.New(5948684761759816, "Delete", "删除角色", "删除某个角色"),
	Create:        permission.New(5948684761759817, "Create", "创建角色", "创建一个新角色"),
	SetMember:     permission.New(5950451522795973, "SetMember", "设置角色成员", "增加或移除角色成员"),
	SetPermission: permission.New(5950452043151813, "SetPermission", "设置角色权限", "设置某个角色的权限"),
}
