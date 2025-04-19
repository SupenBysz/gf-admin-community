// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCategory is the golang structure of table sys_category for DAO operations like Where/Data.
type SysCategory struct {
	g.Meta      `orm:"table:sys_category, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	ParentId    interface{} // 父级ID
	PicturePath interface{} // 分类图片
	Hidden      interface{} // 是否隐藏
	Sort        interface{} // 顺序
	UnionMainId interface{} // 关联主体ID（保留字段）
	Type        interface{} // 0全局；1商品分类；2资讯分类
	Description interface{} //
	UpdatedAt   *gtime.Time //
	CreatedAt   *gtime.Time //
}
