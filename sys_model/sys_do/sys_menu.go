// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta               `orm:"table:sys_menu, do:true"`
	Id                   interface{} // ID
	Path                 interface{} // 路径
	Name                 interface{} // 名称
	SortTitle            interface{} // 简称
	I18NKey              interface{} // 国际化Key
	Redirect             interface{} // 跳转
	Title                interface{} // 标题
	Icon                 interface{} // 图标
	Layout               interface{} // 布局
	Component            interface{} // 组件
	ParentId             interface{} // 所属父级
	Sort                 interface{} // 排序
	State                interface{} // 状态：1启用，0禁用
	Hidden               interface{} // 是否隐藏：1是，0否
	Description          interface{} // 描述
	IconUrl              interface{} // 图标URL
	RedirectType         interface{} // 跳转类型：1当前页面打开、 2新的标签页打开
	Type                 interface{} // 类型：1菜单、2按钮
	DepPermissionIds     interface{} // 依赖权限Ids
	LimitHiddenRoleIds   interface{} // 限定不可见的角色
	LimitHiddenUserIds   interface{} // 限定不可见的用户
	LimitHiddenUserTypes interface{} // 限定不可见的用户类型
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
