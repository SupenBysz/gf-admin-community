package model

// Context 请求上下文结构
type Context struct {
	ClaimsUser *JwtCustomClaims // 上下文用户信息
	Ipv4       string           // 客户端IP地址
}
