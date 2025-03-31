# GF Admin Community - 贡献者指南

欢迎来到 GF Admin Community 项目的 GitHub 仓库！本文档将帮助您了解如何参与项目开发以及我们的自动化流程。

## 目录

- [CI/CD 自动化流程](#cicd-自动化流程)
  - [持续集成检查](#持续集成检查)
  - [自动化测试](#自动化测试)
  - [安全扫描](#安全扫描)
  - [自动化发布](#自动化发布)
- [版本规则](#版本规则)
- [如何参与贡献](#如何参与贡献)
- [相关资源](#相关资源)

## CI/CD 自动化流程

我们使用 GitHub Actions 实现自动化的 CI/CD 流程，确保代码质量和简化发布过程。

### 持续集成检查

每次提交代码到 `main` 或 `develop` 分支，或创建 Pull Request 时，自动触发以下检查：

1. **代码风格检查** - 使用 `golangci-lint` 进行代码质量和风格检查
2. **依赖验证** - 确保所有依赖项正确声明
3. **编译检查** - 验证代码能够成功编译

查看 [ci.yml](../workflows/ci.yml) 了解详细配置。

### 自动化测试

CI 流程会自动运行单元测试和集成测试，确保新提交的代码不会破坏现有功能：

1. **单元测试** - 验证各个组件的功能
2. **覆盖率报告** - 生成测试覆盖率报告并上传到 Codecov
3. **性能基准测试** - 检测关键组件的性能变化

### 安全扫描

我们使用 `gosec` 工具进行自动化安全扫描，检测常见的安全问题：

1. **代码安全分析** - 检测潜在的安全漏洞
2. **依赖项安全检查** - 审查第三方依赖中的已知漏洞
3. **SARIF 报告** - 生成标准化的安全分析报告

### 自动化发布

我们的发布流程是自动化的，基于约定式提交消息：

1. **语义化版本** - 基于提交消息自动确定版本号
2. **自动化变更日志** - 根据提交信息生成变更日志
3. **GitHub Release** - 自动创建 GitHub Release 和标签

查看 [release.yml](../workflows/release.yml) 了解详细配置。

## 版本规则

项目遵循[语义化版本](https://semver.org/lang/zh-CN/)，版本号格式为：`MAJOR.MINOR.PATCH`

- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能性新增
- **PATCH**: 向后兼容的问题修正

您可以通过两种方式触发新版本发布：

1. **自动发布**: 当代码合并到 `main` 分支时，基于提交消息自动确定版本号并发布
2. **手动发布**: 通过 GitHub Actions 手动触发发布工作流，指定版本类型（patch/minor/major）

   步骤：
   - 转到 Actions 标签页
   - 选择 "Release Go Module" 工作流
   - 点击 "Run workflow"
   - 选择分支和版本类型
   - 点击 "Run workflow" 按钮

## 如何参与贡献

我们欢迎各种形式的贡献，包括但不限于：

1. **报告问题** - 使用 [Issue 模板](../ISSUE_TEMPLATE/)
2. **功能建议** - 通过 Issue 提出新功能想法
3. **提交代码** - 提交 Pull Request 实现新功能或修复问题
4. **完善文档** - 帮助改进文档和示例

详细的贡献指南请查看 [CONTRIBUTING.md](../CONTRIBUTING.md)。

## 相关资源

- [行为准则](../CODE_OF_CONDUCT.md)
- [安全策略](../SECURITY.md)
- [提交消息约定](../COMMIT_CONVENTION.md)
- [Pull Request 模板](../PULL_REQUEST_TEMPLATE.md)

---

感谢您对 GF Admin Community 的贡献！
