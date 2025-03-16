package sdk_tencent

import (
	"context"
	"fmt"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// GetAccessToken 获取 AccessToken
func (s *sSdkTencent) GetAccessToken(ctx context.Context, wBAppId, wBAppSecret, version string) (accessToken string, err error) {
	{
		// 注意：Access Token 必须缓存在磁盘，并定时刷新，且不能并发刷新，建议每20分钟请求新的 Access Token，获取之后立即使用最新的 Access Token。旧的只有一分钟的并存期。
		// GET https://kyc1.qcloud.com/api/oauth2/access_token?app_id=xxx&secret=xxx&grant_type=client_credential&version=1.0.0

		url := fmt.Sprintf("https://kyc1.qcloud.com/api/oauth2/access_token?app_id=%s&secret=%s&grant_type=client_credential&version=%s",
			wBAppId,
			wBAppSecret,
			version,
		)

		response := g.Client().GetContent(ctx, url)

		// 接受返回数据，json解析
		//{"code":"0","msg":"请求成功","bizSeqNo":"24072620001184415214363812482752","transactionTime":"20240726143638","success":true,"access_token":"WAA0f-****","expire_in":7200,"expire_time":"20240726163638"}
		newTokenInfo := sys_model.AccessTokenByTencent{}
		_ = gjson.DecodeTo(response, &newTokenInfo)
		if &newTokenInfo == nil || newTokenInfo.AccessToken == "" {
			return "", sys_service.SysLogs().ErrorSimple(ctx, err, "error_tencent_api_access_token_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}
		if newTokenInfo.AccessToken != "" {
			accessToken = newTokenInfo.AccessToken
		}

		// TODO 将API访问认证的Token存入缓存
	}

	return accessToken, err
}

// GetApiTicket 获取API Ticket
func (s *sSdkTencent) GetApiTicket(ctx context.Context, wBAppId, accessToken, version string) (ticket string, err error) {
	{
		// GET https://kyc1.qcloud.com/api/oauth2/api_ticket?app_id=xxx&access_token=xxx&type=SIGN&version=1.0.0
		url := fmt.Sprintf("https://kyc1.qcloud.com/api/oauth2/api_ticket?app_id=%s&access_token=%s&type=SIGN&version=%s",
			wBAppId,
			accessToken,
			version,
		)
		response := g.Client().GetContent(ctx, url)
		// 接受返回数据，json解析
		// {"code":"0","msg":"请求成功","bizSeqNo":"24072620001184433114383908047162","transactionTime":"20240726143840","success":true,"tickets":[{"value":"*****","expire_in":3600,"expire_time":"20240726153839"}]}
		signTicketRes := sys_model.SignTicketRes{}
		_ = gjson.DecodeTo(response, &signTicketRes)
		if &signTicketRes == nil || signTicketRes.Tickets == nil || len(signTicketRes.Tickets) == 0 {
			return "", sys_service.SysLogs().ErrorSimple(ctx, err, "error_tencent_api_ticket_failed", sys_dao.SysConfig.Table()+":"+s.sysConfigName)
		}

		if len(signTicketRes.Tickets) > 0 {
			ticket = signTicketRes.Tickets[0].Value
		}
	}

	return ticket, nil
}
