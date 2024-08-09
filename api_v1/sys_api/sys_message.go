package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetMessageByIdReq struct {
	g.Meta `path:"/getMessageById" method:"post" summary:"根据id查询消息｜信息" tags:"消息管理"`
	Id     int64 `json:"id"  v:"required#消息Id不能为空"`
}

type CreateMessageReq struct {
	g.Meta `path:"/createMessage" method:"post" summary:"添加消息｜信息" tags:"消息管理"`

	sys_model.SysMessage
}

type UpdateMessageReq struct {
	g.Meta `path:"/updateMessage" method:"post" summary:"编辑消息｜信息" tags:"消息管理"`
	Id     int64 `json:"id"  v:"required#消息Id不能为空"`

	sys_model.UpdateSysMessage
}

type QueryMessageReq struct {
	g.Meta `path:"/queryMessage" method:"post" summary:"查询消息｜列表" tags:"消息管理"`

	base_model.SearchParams
}

type GetUserMessageReq struct {
	g.Meta `path:"/getUserMessage" method:"post" summary:"查询登陆用户的消息｜列表" tags:"消息管理"`

	base_model.SearchParams
}

// TODO 这个可能不需要
//type QueryUnionMainMessageReq struct {
//	g.Meta `path:"/queryUnionMainMessage" method:"post" summary:"查询指定主体的消息｜列表" tags:"消息管理"`
//
//	base_model.SearchParams
//}

type HasUnReadMessageReq struct {
	g.Meta `path:"/hasUnReadMessage" method:"post" summary:"获取未读消息数量" tags:"消息管理"`
	Type   int `json:"type" dc:"消息类型"`
}

// TODO 在消息查看详情的时候，就给当前登录User进行追加进已读用户[]里面，不要直接暴露接口
//type SetMessageReadUserIdsReq struct {
//	g.Meta    `path:"/setMessageReadUserIds" method:"post" summary:"追加消息已读用户" tags:"消息管理"`
//	MessageId int64 `json:"messageId" dc:"消息Id" v:"required#请填写消息Id"`
//	UserId    int64 `json:"userId" dc:"已读用户UserId" v:"required#请填写已读用户Id"`
//}

type ReleaseMessageReq struct { // TODO  功能待补充
	g.Meta `path:"/releaseMessage" method:"post" summary:"发送消息" tags:"消息管理"`
}
