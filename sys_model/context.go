package sys_model

import "github.com/gogf/gf/v2/container/garray"

// Context 请求上下文结构
type Context struct {
	ClaimsUser        *JwtCustomClaims // 上下文用户信息
	Ipv4              string           // 客户端IP地址
	SessionErrorQueue *garray.Array    // 会话错误队列
}
