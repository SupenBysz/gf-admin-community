package sys_model

import (
	"github.com/gogf/gf/v2/container/garray"
)

// SessionContext 请求上下文结构
type SessionContext struct {
	JwtClaimsUser     *JwtCustomClaims // 上下文用户信息
	Ipv4              string           // 客户端IP地址
	SessionErrorQueue *garray.Array    // 会话错误队列
}
