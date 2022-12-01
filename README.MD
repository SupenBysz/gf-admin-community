

# 一、项目介绍

# ...
## 1、技术选型
### 开发工具
<img src="https://img.shields.io/badge/goland-2022.2.3+-blue"/>
<img src="https://img.shields.io/badge/vscode-1.71.2+-blue"/>

- 推荐 goland 2022.2.3+

#### 语言/框架/类库
<img src="https://img.shields.io/badge/golang-1.19+-green"/>
<img src="https://img.shields.io/badge/goframe-2.2.0+-green"/>
<img src="https://img.shields.io/badge/casbin-2.56.0+-green"/>
<img src="https://img.shields.io/badge/golang jwt-4.4.2+-green"/>
<img src="https://img.shields.io/badge/idgenerator go-1.3.2+-green"/>

- golang 1.19+
- goframe 2.2.4+
- goframe 2.56.0+
- golang-jwt 4.4.2+
- idgenerator-go 1.3.2+

#### 数据库
<img src="https://img.shields.io/badge/PostgreSQL-14.5+-red"/>

- PostgreSQL-14.5+

#### 脚本工具
<img src="https://img.shields.io/badge/goframe cli-2.2.0.plus-red"/>

- goframe cli 2.2.0.plus

# 二、项目结构
```
/
├── api_v1            对外接口：考虑到版本管理需要，往往以api/v1...存在。
│   └── sys_api                对外提供服务的输入/输出数据结构定义。
├── hack              工具脚本：存放项目开发工具、脚本等内容。例如，CLI工具的配置，各种shell/bat脚本等文件
├── internal          内部逻辑：业务逻辑存放目录。通过Golang internal特性对外部隐藏可见性。
│   ├── cmd           入口指令：命令行管理目录。可以管理维护多个命令行。
│   ├── consts        常量定义：内部逻辑专用常量。
│   ├── logic         业务封装：业务逻辑封装管理，特定的业务逻辑实现和封装。往往是项目中最复杂的部分。
├── manifest          交付清单：包含程序编译、部署、运行、配置的文件。
│   └── config        配置管理：配置文件存放目录。
├── model             结构模型：数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义。
|   ├── dao           数据访问：数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法
|   ├── do            领域对象：用于dao数据操作中业务模型与实例模型转换，由工具维护，用户不能修改。
│   └── entity        数据模型：数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改。
├── resource          静态资源：静态资源文件。这些文件往往可以通过 资源打包/镜像编译 的形式注入到发布文件中。
├── service           业务接口：用于业务模块解耦的接口定义层。具体的接口实现在logic中进行注入。
├── sys_controller    接口处理：接收/解析用户输入参数的入口/接口层。
├── utility           实用类库：一般以独立文件形式的通用类库
├── go.mod            依赖管理：使用Go Module包管理的依赖描述文件。
└── main.go           项目入口
```

### 以模块方式向业务项目服务，作为通用基础模块来独立维护
> 接口包含两部分：接口定义（api）+ 接口实现（sys_controller）<br>
> 负责提供标准接口服务，其它项目直接引用即可

> 服务包含两部分：模型结构（model）+ 服务接口（service）<br>
> 负责接收并响应调用的输入与输出，包括对输入参数的过滤、转换、校验，对输出数据结构的维护，并调用 service 实现业务逻辑处理。

### 主要功能
- [x] 权限管理：基于jwt和casbin实现的权限管理。
- [x] 分页封装
- [x] 用户管理
- [x] 角色管理
- [x] 组织管理
- [x] API管理
- [x] 条件搜索
- [x] 验证码生成
- [x] 文件上传
  - [x] 身份证照片识别
  - [x] 营业执照识别
  - [x] 银行卡照片识别
- [x] SDK配置管理
  - [x] 百度SDK
  - [x] 阿里云SDK
  - [x] 天翼云SDK
  - [x] 华为云SDK
  - [x] 腾讯云SDK
- [ ] 短信网关
  - [ ] 企业信使
  - [ ] 阿里云
  - [ ] 华为云
  - [ ] 腾讯云
  - [ ] 天翼云
- [ ] 数据存储网关
  - [ ] minio
  - [ ] 阿里云
  - [ ] 华为云
  - [ ] 腾讯云
  - [ ] 天翼云

# 三、项目编译/运行
### 创建并导入数据库
- 创建数据库<br>
- 导入db目录下的数据库脚本 latest.sql<br>
- 配置hack/config.yaml中的数据库链接
- 配置manifest/config/config.yaml中的数据库链接

### 编译
1、克隆项目
```Shell
git clone https://github.com/SupenBysz/gf-admin-community.git
```
2、初始化项目
```Shell
go mod tidy
```
3、编译
```Shell
gf build main.go
```
4、运行
> 由于是以组件方式提供服务，所以不能调试运行，需要先编译出来后可以执行。<br>
> 项目被引用运行时，是可以实现调试目的。

#### 项目脚手架命令: 
- `make cli`: 安装或升级脚手架命令。
- `make dao`: 根据“hack”文件夹中的配置文件为“Entity/DAO/DO”生成go文件。
- `make service`: 解析“逻辑”文件夹生成服务接口文件到“service”文件夹。
> 不推荐直接在本项目使用 gf gen 等命令，将会导致文件生成的位置不在预期位置，如果一定要用先了解相关参数后再用。

# 四、风格规范
&emsp;&emsp;代码风格并不影响程序的运行，也不会给你的程序带来潜在的危险。但一段好的代码风格，能让人阅读起来特别舒服，特别的阅读别人代码的时候。
## 代码风格规范
### 1、包名
&emsp;&emsp;统一使用小写，用下划线作为分隔符，以自然语义缩写或全称作为分隔
### 2、类名
&emsp;&emsp;由于golang属于面向过程编程，没有真正意义上的类。<br>&emsp;&emsp;本项目将逻辑封装成服务对象、控制器对象等，通过封装后的对象进行关业务调用，以实现仿类操作效果，这里称作“伪类”。那么伪类的定义规范如下：
> 伪类采用大驼峰的命名形式，所谓大驼峰就是首字母大写，例如UpperCameCase
### 3、常量
- 全局常量：遵行类名的命名一样的方式。
- 局部常量：局部常量则采用小驼峰的形式。所谓局部常量指的是方法/内/包内的常量。
### 4、大括号代码域
&emsp;&emsp;遵循JAVA风格，缩进直接紧接后面括号，例如：
```
int Foo(bool isBar) {
    if (isBar) {
        bar();
        return 1;
    }
    return 0
}
```
### 5、tab缩进
&emsp;&emsp;每个tab缩进，用4个空格代替。
## 接口风格规范
&emsp;&emsp;为了能更好的同步接口相关约束，规范接口文档的自动生成、并支持接口的相关参数自动校验项目接口做如下规范定义。
### 接口规范
1、所有逻辑接口需要封装成控制器对象，统一保存在controller文件夹下<br>
2、一个控制器接口方法首字母大写的驼峰命名方式<br>
3、每个接口方法只能包含2个参数，分别为接口上下文对象 Context 和 请求参数定义（以Req结尾的类型参数定义）。接口方法返回值只能有2个固定的返回值，第一个是要返回的数据，第二个是错误异常信息。<br>例如：
```
func (a *cAuth) Login(ctx context.Context, req *sysapi.LoginReq) (res *sysapi.LoginRes, err error) {
    return nil, nil
}
```
- 方法参数：ctx context.Context 为请求上下文，不需要用到上下文时，变量 ctx 可用下划线代替
- 方法参数：req *sysapi.LoginReq 为请求参数定义，必须以Req结尾，且必须是指针对象，不需要用到前端请求参数时，变量 req 可用下划线代替
- 方法返回：*sysapi.LoginRes 为方法第一个返回参数，且必须以Res结尾，返回值的变量名称为可选，如：(*sysapi.LoginRes,error)
- 方法返回：err error 为方法第二个返回参数，必填项，该参数类型固定且不可修改
### 请求参数规范
&emsp;&emsp;请求参数的定义统一在对象成员属性中实现约束、描述、验证规则等相关定义，例如：
```
type LoginReq struct {
    g.Meta `path:"/login" method:"post" summary:"登录" tags:"鉴权"`
	Username string `json:"username" v:"required#请输入用户名" dc:"登录账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"登录密码"`
    Captcha  string `json:"captcha" v:"required#请输入验证吗" dc:"验证码"`
}
```
- LoginReq 作为请求对象的类型名称，必须以首字母大写，且以Res结尾
- g.Meta path:定义请求路径，method:定义请求类型,summary:定义接口概要说明，tags:定义接口在文档生成上的分类
- 请求参数属性：Username，Password，Captcha，定义了请求参数的细则，包括但不限于json:说明参数格式，v:参数校验规则和校验失败后的提示信息,dc:参数描述

# 五、Git Commit 语义化规范
### 作用：
- 能加快 Code Review 的流程
- 根据 Git Commit 的元数据生成 Changelog
- 后续维护者可以知道 Feature 被修改的原因
- 统一团队 Git Commit 日志标准，便于后续代码 Review 和版本发布
- main、dev两个分支提交类型限定为：feat,fix,docs,style,refactor,perf,test,chore,revert等
- 提交格式参考：
```
<type>(<scope>):<subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```
- Type 代表了某次提交的类型，比如是修复一个 bug 还是增加了一个新的 feature。常用类型如下：
  - feat：新增 feature
  - fix：修复 bug
  - docs：仅仅修改了文档，比如 README、CHANGLELOG 等等
  - style：仅仅修改了空格、格式缩进、多行等等，不改变逻辑
  - refactor：代码重构，没有加新的功能或者修复 bug
  - perf：优化相关，比如提升性能、体验等
  - test：测试用例，包括单元测试、集成测试等
  - chore：改变构建流程，或者增加依赖库、工具等
  - revert：回滚到上一个版本
- Scope 字段用于说明本次 commit 所影响的范围，比如视图层、数据模型或者路由模块等，是一个可选参数。
- Subject 字段是本次 commit 的一个概要，需要用最简洁的语言来说明本次修改的内容。
- Body 可以使用多行文本详细地说明本次提交所改动的一些细节，从而帮助后续的使用者们更好地了解代码的内容。
- Footer 大部分只用于两种情况
  - 不兼容变动
    - 如果本次的 commit 与前一个版本的代码无法兼容，那么 Footer 部分需要以 BREAKING CHANGE 开头，后面描述本次变动的详细情况以及迁移到新版本代码的方法。
  - 关闭 Issue
    - 如果当前 commit 针对某个 issue，那么可以在 Footer 部分关闭这个 issue 。以 Close 开头，后面用 # 号标识对应的 Issue 号码。

### Commitizen
&emsp;&emsp;如果觉得自己写规范的 comment 信息太麻烦，那么可以使用Commitizen ，可以辅助撰写一个合格的 Commit message 工具。

