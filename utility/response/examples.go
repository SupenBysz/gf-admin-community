package response

import (
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
)

// 这个文件展示了新响应函数的使用示例
// 注意：这个文件仅用于示例，实际项目中可以删除

// ExampleBasicUsage 基本用法示例
func ExampleBasicUsage(r *ghttp.Request) {
	// 1. 简单成功响应
	userData := g.Map{"id": 1, "name": "张三"}
	Success(r, userData, "获取用户信息成功")
	
	// 2. 简单错误响应
	Error(r, 1001, "用户不存在")
	
	// 3. 带退出的响应
	SuccessExit(r, userData, "操作成功")
}

// ExampleAdvancedUsage 高级用法示例
func ExampleAdvancedUsage(r *ghttp.Request) {
	// 1. 使用构建器模式
	JsonBuilder(r,
		WithCode(0),
		WithMessage("操作成功"),
		WithData(g.Map{"result": "success"}),
	)
	
	// 2. 国际化响应
	I18nResponse(r, 0, "user.create.success", g.Map{"id": 123})
	
	// 3. 带错误处理的响应
	data, err := getUserData() // 假设的函数
	JsonWithError(r, err, data, "获取用户数据成功")
}

// ExampleValidation 验证示例
func ExampleValidation(r *ghttp.Request) {
	userData := g.Map{"age": 15}
	
	// 带验证的响应
	JsonWithValidation(r, userData, func(data interface{}) error {
		if userMap, ok := data.(g.Map); ok {
			if age, exists := userMap["age"]; exists {
				if ageInt, ok := age.(int); ok && ageInt < 18 {
					return errors.New("年龄必须大于18岁")
				}
			}
		}
		return nil
	}, WithCode(0), WithMessage("用户数据验证通过"))
}

// ExampleMigrationComparison 迁移对比示例
func ExampleMigrationComparison(r *ghttp.Request) {
	userData := g.Map{"id": 1, "name": "张三"}
	
	// 旧的用法（仍然支持）
	Json(r, 0, "操作成功", userData)
	
	// 新的推荐用法
	Success(r, userData, "操作成功")
	
	// 更灵活的构建器用法
	JsonBuilder(r,
		WithCode(0),
		WithMessage("操作成功"),
		WithData(userData),
	)
}

// 假设的辅助函数
func getUserData() (interface{}, error) {
	// 模拟可能出错的数据获取
	return g.Map{"id": 1, "name": "张三"}, nil
}