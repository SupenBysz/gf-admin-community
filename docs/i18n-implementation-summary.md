# 国际化(i18n)实现总结

本文档总结了在 gf-admin-community 项目中实现的国际化(i18n)功能。

## 已完成工作

### 1. 后端实现

1. **控制器实现**：
   - 创建了 `SysCommonController` 控制器，提供语言切换功能
   - 实现了三个API接口：
     - `GET /api/common/setLanguage` - 设置语言
     - `GET /api/common/getLanguages` - 获取支持的语言列表
     - `GET /api/common/getCurrentLanguage` - 获取当前使用的语言

2. **语言检测机制**：
   - 在 `i18n.GetLanguage()` 函数中实现了多级语言检测
   - 优先级顺序：查询参数 > Cookie > HTTP头部 Accept-Language > 默认值(zh-CN)

3. **国际化文件**：
   - 创建了 `i18n/zh-CN.yaml` 和 `i18n/en-US.yaml` 语言文件
   - 按模块组织了错误信息和常用词汇：
     - 系统错误信息
     - 用户模块
     - 邀约模块
     - 组织机构模块
     - 会员等级模块
     - 文件模块
     - 前端设置模块
     - 分类模块
     - 角色模块
     - 阿里云SDK模块
     - 消息模块
     - 通用词汇

4. **错误信息标准化**：
   - 采用 `error_模块名_错误描述` 的命名格式
   - 移除了错误处理中的直接翻译调用，统一使用错误键
   - 错误信息翻译延迟到响应处理阶段

5. **路由注册**：
   - 在 `/example/internal/boot/boot.go` 中注册了公共控制器路由
   - 添加了 `sys_controller.SysCommon` 到匿名访问路由组

### 2. 中间件处理

1. **响应处理中间件**：
   - 在 `sys_middleware.ResponseHandler` 中处理错误信息的国际化
   - 检测以 "error_" 开头的错误消息并使用 `i18n.T()` 进行翻译

2. **响应组件**：
   - 在 `response.Json` 和 `response.JsonExit` 中集成了翻译功能

### 3. 前端集成指南

1. **Vue组件示例**：
   - 创建了 `LanguageSwitcher.vue` 语言切换组件示例
   - 实现了获取语言列表、当前语言和语言切换功能

2. **集成文档**：
   - 提供了前端项目集成国际化的详细步骤
   - 包括依赖安装、组件集成、i18n设置和使用说明

## 架构设计

```
┌─────────────────┐      ┌─────────────────┐
│  前端应用       │      │  后端服务       │
│                 │◀────▶│                 │
└────────┬────────┘      └────────┬────────┘
         │                        │
         ▼                        ▼
┌─────────────────┐      ┌─────────────────┐
│  i18n实例       │      │ SysCommon控制器  │
│ (vue-i18n)      │      │ 语言设置API     │
└────────┬────────┘      └────────┬────────┘
         │                        │
         │                        ▼
┌────────▼────────┐      ┌─────────────────┐
│  语言文件       │      │  i18n工具        │
│  zh-CN.js      │      │  GetLanguage()   │
│  en-US.js      │      │  T()翻译函数     │
└─────────────────┘      └────────┬────────┘
                                  │
                                  ▼
                         ┌─────────────────┐
                         │  语言文件       │
                         │  zh-CN.yaml    │
                         │  en-US.yaml    │
                         └─────────────────┘
```

## 使用方式

### 后端使用

1. **错误信息**：

   ```go
   // 使用错误键（不直接翻译）
   return sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_status_update_failed", sys_dao.SysInvite.Table())
   ```

2. **直接翻译**：

   ```go
   // 在需要立即显示的地方使用翻译
   message := i18n.T(ctx, "common_success")
   ```

### 前端使用

1. **语言切换**：
   - 集成 `LanguageSwitcher` 组件到应用导航栏

2. **使用翻译**：

   ```vue
   <template>
     <!-- 在模板中使用 -->
     <div>{{ t('common.submit') }}</div>
   </template>

   <script>
   import { useI18n } from 'vue-i18n'

   export default {
     setup() {
       const { t } = useI18n()
       return { t }
     }
   }
   </script>
   ```

## 后续工作

1. **扩展语言支持**：
   - 添加更多语言选项，如繁体中文、日语、韩语等

2. **国际化覆盖率**：
   - 确保所有用户界面文本和错误信息都已国际化
   - 添加更多模块的错误信息和通用词汇

3. **用户偏好存储**：
   - 将用户语言偏好保存到用户配置中
   - 在用户登录时自动应用其语言偏好

4. **多语言内容管理**：
   - 为动态内容（如CMS内容）添加多语言支持
   - 实现内容的多语言版本管理
