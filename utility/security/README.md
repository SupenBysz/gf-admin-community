# Token安全增强模块

本模块提供了多层次的Token安全增强功能，包括设备指纹验证、IP验证、频率限制、生命周期管理和安全监控等。

## 功能特性

### 1. 多层安全验证
- **设备指纹验证**: 基于浏览器特征生成设备指纹，防止Token被盗用
- **IP地址验证**: 检测IP地址变化，高安全级别用户需要重新验证
- **Token使用频率限制**: 防止Token被恶意频繁使用
- **Token生命周期管理**: 限制每用户活跃Token数量，支持强制登出

### 2. 安全事件监控
- **可疑活动记录**: 自动记录设备指纹不匹配、IP变化、频率超限等事件
- **自动告警**: 达到阈值时触发安全告警
- **账户保护**: 严重情况下自动锁定账户

### 3. Token轮换机制
- **自动刷新**: Token接近过期时自动刷新
- **旧Token撤销**: 生成新Token时立即撤销旧Token
- **并发安全**: 支持并发环境下的安全Token刷新

## 使用方法

### 1. 基础使用

```go
package main

import (
    "github.com/SupenBysz/gf-admin-community/utility/security"
    "github.com/gogf/gf/v2/net/ghttp"
)

func main() {
    s := g.Server()
    
    // 创建安全中间件
    securityMiddleware := security.NewSecurityMiddleware()
    
    // 使用Token安全处理中间件
    s.Use(securityMiddleware.TokenSecurityHandler)
    
    s.SetPort(8080)
    s.Run()
}
```

### 2. 增强Token使用

```go
// 生成增强Token
enhancedJwtService := security.NewEnhancedJwtService()
tokenInfo, err := enhancedJwtService.GenerateEnhancedToken(ctx, user, r)
if err != nil {
    // 处理错误
}

// 验证增强Token
claims, err := enhancedJwtService.ValidateEnhancedToken(ctx, tokenString, r)
if err != nil {
    // 处理验证失败
}
```

### 3. Token轮换

```go
// 刷新Token（自动轮换）
newTokenInfo, err := enhancedJwtService.RefreshTokenWithRotation(ctx, oldToken)
if err != nil {
    // 处理刷新失败
}
```

### 4. 用户Token管理

```go
// 获取用户当前活跃Token数量
count := enhancedJwtService.GetUserTokenCount(userID)

// 撤销用户所有Token（强制登出）
err := enhancedJwtService.RevokeAllUserTokens(ctx, userID)
```

## 配置说明

### 安全配置结构

```go
type SecurityConfig struct {
    EnableDeviceFingerprint bool     // 启用设备指纹验证
    EnableIPValidation      bool     // 启用IP验证
    EnableRateLimit         bool     // 启用频率限制
    MaxTokensPerUser        int      // 每用户最大Token数量
    RateLimitRequests       int      // 频率限制请求数
    RateLimitWindow         int      // 频率限制时间窗口（秒）
    AlertThreshold          int      // 告警阈值
    SkipAuthPaths           []string // 跳过认证的路径
}
```

### 默认配置

```go
config := security.DefaultSecurityConfig()
// 可以根据需要修改配置
config.EnableDeviceFingerprint = true
config.MaxTokensPerUser = 3
```

## 前端集成

### 1. 设备指纹生成

```javascript
// 生成设备指纹
function generateDeviceFingerprint() {
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');
    ctx.textBaseline = 'top';
    ctx.font = '14px Arial';
    ctx.fillText('Device fingerprint', 2, 2);
    
    const fingerprint = [
        navigator.userAgent,
        navigator.language,
        screen.width + 'x' + screen.height,
        new Date().getTimezoneOffset(),
        canvas.toDataURL()
    ].join('|');
    
    return btoa(fingerprint).substring(0, 32);
}

// 在请求头中添加设备指纹
axios.defaults.headers.common['X-Device-Fingerprint'] = generateDeviceFingerprint();
```

### 2. Token自动刷新

```javascript
// 设置Token自动刷新
axios.interceptors.response.use(
    response => response,
    async error => {
        if (error.response?.status === 401 && error.response?.data?.message === 'token_expired') {
            try {
                const refreshResponse = await axios.post('/api/v1/refresh-token', {
                    token: localStorage.getItem('token')
                });
                
                const newToken = refreshResponse.data.token;
                localStorage.setItem('token', newToken);
                
                // 重试原请求
                error.config.headers.Authorization = `Bearer ${newToken}`;
                return axios.request(error.config);
            } catch (refreshError) {
                // 刷新失败，跳转到登录页
                window.location.href = '/login';
            }
        }
        return Promise.reject(error);
    }
);
```

## 安全级别说明

| 级别 | 描述 | 验证策略 |
|------|------|----------|
| 1 | 基础用户 | 基础Token验证 |
| 2 | 普通用户 | Token验证 + 频率限制 |
| 3 | 重要用户 | Token验证 + 频率限制 + IP验证 |
| 4 | 管理员 | Token验证 + 频率限制 + IP验证 + 设备指纹 |
| 5 | 超级管理员 | 全部安全验证 + 严格监控 |

## 监控和告警

### 安全事件类型

- `device_fingerprint_mismatch`: 设备指纹不匹配
- `ip_address_changed`: IP地址变化
- `rate_limit_exceeded`: 频率限制超出
- `token_expired`: Token过期
- `invalid_token`: 无效Token

### 告警处理

系统会自动记录所有安全事件，并在达到告警阈值时：

1. 记录详细日志
2. 发送告警通知
3. 严重情况下自动锁定账户

## 性能优化

### 1. 缓存策略
- Token验证结果缓存
- 设备指纹缓存
- 频率限制使用内存存储

### 2. 并发处理
- 使用读写锁保护共享数据
- Token刷新支持并发安全
- 异步处理安全事件

### 3. 资源清理
- 定期清理过期的频率限制记录
- 自动清理过期的安全事件
- 撤销Token时清理相关缓存

## 注意事项

1. **设备指纹验证**: 需要前端支持，建议在生产环境中启用
2. **IP验证**: 对于移动用户可能频繁变化，需要合理设置安全级别
3. **频率限制**: 需要根据实际业务场景调整限制参数
4. **Token数量限制**: 避免设置过小导致用户体验问题
5. **安全事件存储**: 建议定期清理历史安全事件数据

## 扩展开发

如需扩展更多安全功能，可以：

1. 实现自定义安全验证器
2. 添加新的安全事件类型
3. 集成第三方安全服务
4. 实现更复杂的风险评估算法

## 故障排除

### 常见问题

1. **设备指纹验证失败**: 检查前端是否正确生成和发送设备指纹
2. **频率限制误报**: 调整频率限制参数或检查Token重复使用
3. **IP验证问题**: 检查代理服务器配置，确保获取真实客户端IP
4. **Token刷新失败**: 检查Token是否已过期或被撤销

### 调试模式

可以通过日志级别调整来获取更详细的调试信息：

```go
g.Log().SetLevel(glog.LEVEL_DEBUG)
```