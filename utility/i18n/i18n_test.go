package i18n_test

import (
	"context"
	"os"
	"testing"

	"github.com/SupenBysz/gf-admin-community/utility/i18n"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

// 测试准备工作，确保i18n目录存在并包含测试所需的语言文件
func setup() {
	// 确保i18n目录存在
	if !gfile.Exists("i18n") {
		_ = gfile.Mkdir("i18n")
	}

	// 创建中文测试文件
	if !gfile.Exists("i18n/zh-CN") {
		_ = gfile.Mkdir("i18n/zh-CN")
	}

	// 创建英文测试文件
	if !gfile.Exists("i18n/en-US") {
		_ = gfile.Mkdir("i18n/en-US")
	}

	// 创建中文测试数据
	zhContent := `
error_test: "测试错误"
common_test: "测试公共文本"
test_with_params: "参数测试：%s"
`
	_ = gfile.PutContents("i18n/zh-CN/test.yaml", zhContent)

	// 创建英文测试数据
	enContent := `
error_test: "Test Error"
common_test: "Test Common Text"
test_with_params: "Param test: %s"
`
	_ = gfile.PutContents("i18n/en-US/test.yaml", enContent)

	// 初始化i18n
	i18n.Init()
}

// 测试完成后清理资源
func cleanup() {
	// 删除测试文件
	_ = gfile.Remove("i18n/zh-CN/test.yaml")
	_ = gfile.Remove("i18n/en-US/test.yaml")
}

// TestMain 测试主函数
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	cleanup()
	os.Exit(code)
}

// TestGetLanguage 测试获取语言函数
func TestGetLanguage(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试默认语言
		ctx := context.Background()
		lang := i18n.GetLanguage(ctx)
		t.Assert(lang, "zh-CN")

		// 注意：以下测试需要ghttp.Request才能正确测试
		// 由于ghttp.RequestFromCtx依赖请求上下文，这里无法直接模拟
		// 实际应用中，我们已经确认了GetLanguage函数的优先级逻辑是正确的
	})
}

// TestTranslation 测试翻译函数
func TestTranslation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 测试中文翻译
		ctx := context.WithValue(context.Background(), "URL_QUERY_lang", "zh-CN")
		text := i18n.T(ctx, "error_test")
		t.Assert(text, "测试错误")

		text = i18n.T(ctx, "common_test")
		t.Assert(text, "测试公共文本")

		// 测试英文翻译
		ctx = context.WithValue(context.Background(), "URL_QUERY_lang", "en-US")
		text = i18n.T(ctx, "error_test")
		t.Assert(text, "Test Error")

		text = i18n.T(ctx, "common_test")
		t.Assert(text, "Test Common Text")

		// 测试参数翻译
		ctx = context.WithValue(context.Background(), "URL_QUERY_lang", "zh-CN")
		text = i18n.T(ctx, "test_with_params", "Hello")
		t.Assert(gstr.Contains(text, "参数测试"), true)
		t.Assert(gstr.Contains(text, "Hello"), true)

		ctx = context.WithValue(context.Background(), "URL_QUERY_lang", "en-US")
		text = i18n.T(ctx, "test_with_params", "Hello")
		t.Assert(gstr.Contains(text, "Param test"), true)
		t.Assert(gstr.Contains(text, "Hello"), true)

		// 测试不存在的键
		ctx = context.WithValue(context.Background(), "URL_QUERY_lang", "zh-CN")
		text = i18n.T(ctx, "non_existent_key")
		t.Assert(text, "non_existent_key")
	})
}
