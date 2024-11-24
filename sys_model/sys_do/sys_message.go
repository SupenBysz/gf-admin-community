// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMessage is the golang structure of table sys_message for DAO operations like Where/Data.
type SysMessage struct {
	g.Meta         `orm:"table:sys_message, do:true"`
	Id             interface{} // ID
	Title          interface{} // 标题
	Summary        interface{} // 摘要
	Content        interface{} // 内容
	Type           interface{} // 消息类型
	Link           interface{} // 跳转链接
	ToUserIds      interface{} // 接收者UserIds，允许有多个接收者
	ToUserType     interface{} // 接收者类型用户类型，和UserType保持一致
	FromUserId     interface{} // 发送者ID，为-1代表系统消息
	FromUserType   interface{} // 发送者类型
	SendAt         *gtime.Time // 发送时间
	ExtJson        interface{} // 拓展数据Json
	ReadUserIds    interface{} // 已读用户UserIds
	DataIdentifier interface{} // 关联的数据标识
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	SceneDesc      interface{} // 场景描述
	SceneType      interface{} // 场景类型【业务层自定义】例如：1活动即将开始提醒、2活动开始提醒、3活动即将结束提醒、4活动结束提醒、5活动获奖提醒、6券即将生效提醒、7券的生效提醒、8券的失效提醒、9券即将失效提醒、10券核销提醒、8192系统通知、
}
