// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFile is the golang structure of table sys_file for DAO operations like Where/Data.
type SysFile struct {
	g.Meta         `orm:"table:sys_file, do:true"`
	Id             interface{} // 自增ID
	Name           interface{} // 文件名称
	Src            interface{} // 存储路径
	Url            interface{} // URL地址
	Ext            interface{} // 扩展名
	Size           interface{} // 文件大小
	Category       interface{} // 文件分类
	UserId         interface{} // 用户ID
	UnionMainId    interface{} // 关联主体ID
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	LocalPath      interface{} // 本地路径
	AllowAnonymous interface{} // 是否允许匿名访问
}
