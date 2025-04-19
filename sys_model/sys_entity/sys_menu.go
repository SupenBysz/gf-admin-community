// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id                   int64       `json:"id"                   orm:"id"                      description:"ID"`
	Path                 string      `json:"path"                 orm:"path"                    description:"路径"`
	Name                 string      `json:"name"                 orm:"name"                    description:"名称"`
	SortTitle            string      `json:"sortTitle"            orm:"sort_title"              description:"简称"`
	I18NKey              string      `json:"i18NKey"              orm:"i18n_key"                description:"国际化Key"`
	Redirect             string      `json:"redirect"             orm:"redirect"                description:"跳转"`
	Title                string      `json:"title"                orm:"title"                   description:"标题"`
	Icon                 string      `json:"icon"                 orm:"icon"                    description:"图标"`
	Layout               string      `json:"layout"               orm:"layout"                  description:"布局"`
	Component            string      `json:"component"            orm:"component"               description:"组件"`
	ParentId             int64       `json:"parentId"             orm:"parent_id"               description:"所属父级"`
	Sort                 int         `json:"sort"                 orm:"sort"                    description:"排序"`
	State                int         `json:"state"                orm:"state"                   description:"状态：1启用，0禁用"`
	Hidden               int         `json:"hidden"               orm:"hidden"                  description:"是否隐藏：1是，0否"`
	Description          string      `json:"description"          orm:"description"             description:"描述"`
	IconUrl              string      `json:"iconUrl"              orm:"icon_url"                description:"图标URL"`
	RedirectType         int         `json:"redirectType"         orm:"redirect_type"           description:"跳转类型：1当前页面打开、 2新的标签页打开"`
	Type                 int         `json:"type"                 orm:"type"                    description:"类型：1菜单、2按钮"`
	DepPermissionIds     string      `json:"depPermissionIds"     orm:"dep_permission_ids"      description:"依赖权限Ids"`
	LimitHiddenRoleIds   string      `json:"limitHiddenRoleIds"   orm:"limit_hidden_role_ids"   description:"限定不可见的角色"`
	LimitHiddenUserIds   string      `json:"limitHiddenUserIds"   orm:"limit_hidden_user_ids"   description:"限定不可见的用户"`
	LimitHiddenUserTypes string      `json:"limitHiddenUserTypes" orm:"limit_hidden_user_types" description:"限定不可见的用户类型"`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"              description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"              description:""`
}
