package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMenuByIdReq struct {
	g.Meta `path:"/getMenuById" method:"post" summary:"根据菜单ID获取菜单|信息" tags:"菜单"`
	Id     int64 `json:"id" v:"required#菜单ID校验失败" dc:"菜单ID"`
}

type CreateMenuReq struct {
	g.Meta `path:"/createMenu" method:"post" summary:"创建菜单" tags:"菜单"`
	sys_model.SysMenu
}

type UpdateMenuReq struct {
	g.Meta `path:"/updateMenu" method:"post" summary:"更新菜单" tags:"菜单"`
	sys_model.UpdateSysMenu
}

type DeleteMenuReq struct {
	g.Meta `path:"/deleteMenu" method:"post" summary:"删除菜单" tags:"菜单"`
	Id     int64 `json:"id" v:"required#菜单ID校验失败" dc:"菜单ID"`
}

type GetMenuTreeReq struct {
	g.Meta `path:"/getMenuTree" method:"post" summary:"获取菜单树" tags:"菜单"`
	Id     int64 `json:"id" v:"required#菜单ID校验失败" dc:"菜单ID"`
}
