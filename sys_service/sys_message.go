// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/base_model"
)

type (
	IMessage interface {
		// GetMessageById 根据id查询消息
		GetMessageById(ctx context.Context, id int64) (*sys_model.SysMessageRes, error)
		// GetMessageDetailById 根据id查询消息详情
		GetMessageDetailById(ctx context.Context, messageId, userId int64) (*sys_model.SysMessageRes, error)
		// CreateMessage 添加消息
		CreateMessage(ctx context.Context, info *sys_model.SysMessage) (*sys_model.SysMessageRes, error)
		// UpdateMessage 编辑消息 （限制是还未发送的）
		UpdateMessage(ctx context.Context, id int64, info *sys_model.UpdateSysMessage) (*sys_model.SysMessageRes, error)
		// QueryMessage 查询消息列表
		QueryMessage(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysMessageListRes, error)
		// QueryUserMessage 查询指定用户的消息｜列表
		QueryUserMessage(ctx context.Context, userId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysMessageListRes, error)
		// QueryUnionMainMessage 查询指定主体发送的消息列表 （支持未发送消息列表，添加params参数）
		QueryUnionMainMessage(ctx context.Context, unionMainId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysMessageListRes, error)
		// HasUnReadMessage 是否存在未读消息
		HasUnReadMessage(ctx context.Context, userId int64) (int, error)
		// SetMessageReadUserIds 追加公告已读用户
		SetMessageReadUserIds(ctx context.Context, messageId int64, userId int64) (bool, error)
	}
)

var (
	localMessage IMessage
)

func Message() IMessage {
	if localMessage == nil {
		panic("implement not found for interface IMessage, forgot register?")
	}
	return localMessage
}

func RegisterMessage(i IMessage) {
	localMessage = i
}