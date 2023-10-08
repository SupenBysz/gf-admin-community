package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysMenu struct {
	Id          int64   `json:"id"        dc:"ID，ID值为0时则新增菜单" v:"min:0#ID不能小于0"`
	Path        *string `json:"path"      dc:"路径" v:"max-length:128#路径最大长度请不要超过128字符"`
	Name        *string `json:"name"      dc:"名称" v:"required|max-length:128#请输入名称|名称最大长度请不要超过128字符"`
	Redirect    *string `json:"redirect"  dc:"跳转" v:"max-length:128#跳转URL最大长度请不要超过128字符"`
	Title       *string `json:"title"     dc:"标题" v:"required|max-length:64#请输入标题|标题最大长度请不要超过128字符"`
	Icon        *string `json:"icon"      dc:"图标" v:"max-length:128#图标名称最大长度请不要超过128字符"`
	IconUrl     *string `json:"iconUrl"     description:"图标URL"`
	Component   *string `json:"component" dc:"组件" v:"max-length:128#组件地址最大长度请不要超过128字符"`
	ParentId    *int64  `json:"parentId"  dc:"所属父级" v:"integer|min:0#父级ID参数错误|父级ID不能小于0"`
	Sort        *int    `json:"sort"      dc:"排序" v:"integer#排序参数错误"`
	State       *int    `json:"state"    dc:"状态：0隐藏，1显示" v:"in:0,1#请选择状态类型"`
	Description *string `json:"description" description:"描述"`
}

type UpdateSysMenu struct {
	Id        int64  `json:"id"        dc:"ID，ID值为0时则新增菜单" v:"min:0#ID不能小于0"`
	Path      string `json:"path"      dc:"路径" v:"max-length:128#路径最大长度请不要超过128字符"`
	Redirect  string `json:"redirect"  dc:"跳转" v:"max-length:128#跳转URL最大长度请不要超过128字符"`
	Title     string `json:"title"     dc:"标题" v:"required|max-length:64#请输入标题|标题最大长度请不要超过128字符"`
	Icon      string `json:"icon"      dc:"图标" v:"max-length:128#图标名称最大长度请不要超过128字符"`
	IconUrl   string `json:"iconUrl"     description:"图标URL"`
	Component string `json:"component" dc:"组件" v:"max-length:128#组件地址最大长度请不要超过128字符"`
	Sort      int    `json:"sort"      dc:"排序" v:"integer#排序参数错误"`
	State     int    `json:"state"    dc:"状态：0隐藏，1显示" v:"in:0,1#请选择状态类型"`
}

type SysMenuRes sys_entity.SysMenu
type SysMenuListRes base_model.CollectRes[*sys_entity.SysMenu]

type SysMenuTreeRes struct {
	*sys_entity.SysMenu
	Children []*SysMenuTreeRes `json:"children" dc:"菜单子级"`
}

type SysMenuTreeListRes []*SysMenuTreeRes
