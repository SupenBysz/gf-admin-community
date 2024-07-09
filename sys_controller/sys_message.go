package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	v1 "github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"

	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type cSysMessage struct{}

var SysMessage cSysMessage

// GetMessageById 根据id查询消息
func (c *cSysMessage) GetMessageById(ctx context.Context, req *v1.GetMessageByIdReq) (*sys_model.SysMessageRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Message().GetMessageDetailById(ctx, req.Id, user.UnionMainId)

	return ret, err
}

// CreateMessage 添加消息
func (c *cSysMessage) CreateMessage(ctx context.Context, req *v1.CreateMessageReq) (*sys_model.SysMessageRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 系统消息: 那么消息发起者是主体id，userType是用户Type
	if (req.Type & sys_enum.Message.Type.System.Code()) == sys_enum.Message.Type.System.Code() {
		req.FromUserId = user.UnionMainId
		req.FromUserType = user.Type
	}

	ret, err := sys_service.Message().CreateMessage(ctx, &req.SysMessage)

	return ret, err
}

// UpdateMessage 编辑消息 （限制是还未发送的）
func (c *cSysMessage) UpdateMessage(ctx context.Context, req *v1.UpdateMessageReq) (*sys_model.SysMessageRes, error) {
	ret, err := sys_service.Message().UpdateMessage(ctx, req.Id, &req.UpdateSysMessage)

	return ret, err
}

// QueryMessage 查询消息列表
func (c *cSysMessage) QueryMessage(ctx context.Context, req *v1.QueryMessageReq) (*sys_model.SysMessageListRes, error) {
	ret, err := sys_service.Message().QueryMessage(ctx, &req.SearchParams, false)

	return ret, err
}

// QueryUserMessage 查询指定用户消息列表
func (c *cSysMessage) QueryUserMessage(ctx context.Context, req *v1.GetUserMessageReq) (*sys_model.SysMessageListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Message().QueryUserMessage(ctx, user.Id, &req.SearchParams, false)

	return ret, err
}

//// QueryUnionMainMessage 查询指定主体消息列表 （支持未发送的消息列表，添加查询参数）
//func (c *cSysMessage) QueryUnionMainMessage(ctx context.Context, req *v1.QueryUnionMainMessageReq) (*sys_model.SysMessageListRes, error) {
//	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
//
//	ret, err := sys_service.SysMessage().QueryUnionMainMessage(ctx, user.UnionMainId, &req.SearchParams, false)
//
//	return ret, err
//}

// HasUnReadMessage 是否存在未读消息
func (c *cSysMessage) HasUnReadMessage(ctx context.Context, req *v1.HasUnReadMessageReq) (api_v1.IntRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.Message().HasUnReadMessage(ctx, user.Id, req.Type)

	return (api_v1.IntRes)(ret), err
}

// SetMessageReadUserIds 追加消息已读用户 // TODO 在消息查看详情的时候，就给当前登录User进行追加进已读用户[]里面，不要直接暴露接口
//func (c *cSysMessage) SetMessageReadUserIds(ctx context.Context, req *v1.SetMessageReadUserIdsReq) (api_v1.BoolRes, error) {
//	ret, err := sys_service.SysMessage().SetMessageReadUserIds(ctx, req.MessageId, req.UserId)
//
//	return ret == true, err
//}
