// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAnnouncement is the golang structure of table sys_announcement for DAO operations like Where/Data.
type SysAnnouncement struct {
	g.Meta          `orm:"table:sys_announcement, do:true"`
	Id              interface{} //
	Title           interface{} // 公告标题
	UnionMainId     interface{} // 发布主体，0值则代表平台发布的公告
	PublicAt        *gtime.Time // 公示时间，只有到了公示时间用户才可见
	Body            interface{} // 公告正文
	UserTypeScope   interface{} // 受众用户类型：0则所有，复合类型
	ExpireAt        *gtime.Time // 过期时间，过期后前端用户不可见
	State           interface{} // 状态：1草稿、2待发布、4已发布、8已过期、16已撤销
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
	CreatedBy       interface{} // 创建用户
	UpdatedBy       interface{} // 最后修改用户
	DeletedAt       *gtime.Time //
	DeletedBy       interface{} //
	ExtDataJson     interface{} // 扩展json数据
	CategoryId      interface{} // 公告分类ID
	Priority        interface{} // 优先级：1普通、2重要、3紧急
	IsPinned        interface{} // 是否置顶：0否、1是
	IsForceRead     interface{} // 是否强制阅读：0否、1是
	Tags            interface{} // 公告标签，多个用逗号分隔
	ReadCount       interface{} // 阅读次数
	ConfirmRequired interface{} // 是否需要确认：0否、1是
}
