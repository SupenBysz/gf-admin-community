package i18n

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

var instance *gi18n.Manager

// Init 初始化国际化
func Init(path ...string) {
	var err error
	i18nPath := "i18n"
	if len(path) > 0 && path[0] != "" {
		i18nPath = path[0]
	}

	// 检查目录是否存在
	if !gfile.IsDir(i18nPath) {
		g.Log().Warning(context.Background(), "i18n directory not found:", i18nPath)
		return
	}

	// 加载国际化文件
	instance = gi18n.New()
	err = instance.SetPath(i18nPath)
	if err != nil {
		g.Log().Error(context.Background(), "failed to set i18n path:", err)
		return
	}

	g.Log().Debug(context.Background(), "i18n initialized with path:", i18nPath)
}

// GetLanguage 获取当前语言
func GetLanguage(ctx context.Context) string {
	language := "zh-CN" // 默认语言

	if ctx != nil {
		r := ghttp.RequestFromCtx(ctx)
		if r != nil {
			// 尝试从Header获取语言
			acceptLanguage := r.Header.Get("Accept-Language")
			if acceptLanguage != "" {
				language = parseAcceptLanguage(acceptLanguage)
			}

			// 尝试从Cookie获取语言
			langCookie := r.Cookie.Get("lang")
			if langCookie != nil && langCookie.String() != "" {
				language = langCookie.String()
			}

			// 尝试从请求参数获取语言（最高优先级）
			if r.Get("lang").String() != "" {
				language = r.Get("lang").String()
			}
		}
	}

	return language
}

// parseAcceptLanguage 解析Accept-Language头
func parseAcceptLanguage(acceptLanguage string) string {
	if acceptLanguage == "" {
		return "zh-CN"
	}

	parts := gstr.Split(acceptLanguage, ",")
	if len(parts) > 0 {
		langParts := gstr.Split(parts[0], ";")
		if len(langParts) > 0 {
			lang := langParts[0]

			// 标准化语言代码
			if strings.Contains(lang, "zh") {
				if strings.Contains(lang, "TW") || strings.Contains(lang, "HK") {
					return "zh-TW"
				}
				return "zh-CN"
			} else if strings.Contains(lang, "en") {
				return "en-US"
			}

			return lang
		}
	}

	return "zh-CN"
}

// T 翻译字符串
func T(ctx context.Context, key string, args ...interface{}) string {
	if instance == nil {
		Init()
	}

	if instance == nil {
		glog.Warning(ctx, "i18n instance is nil, returning key:", key)
		return key
	}

	language := GetLanguage(ctx)

	// 设置语言
	//instance.SetLanguage(language)
	n18nCtx := gi18n.WithLanguage(ctx, language)

	// 翻译
	translated := ""
	// langKey := `{#` + key + `}`
	if len(args) > 0 {
		translated = instance.Tf(n18nCtx, key, args...)
	} else {
		translated = instance.T(n18nCtx, key)
		// translated = g.I18n().T(gi18n.WithLanguage(context.TODO(), language), langKey)
	}

	if translated == key {
		// 可能是包含参数的情况，尝试在参数下标前添加%
		for _, pattern := range []string{"%d", "%s", "%v", "%f"} {
			if gstr.Contains(translated, pattern) {
				return translated
			}
		}

		// 记录未翻译的字符串，方便后续补充
		glog.Debug(ctx, "untranslated i18n key:", key, "language:", language)

		// 如果没有翻译，尝试将key转换为可读的格式
		if gstr.Contains(key, "_") {
			parts := gstr.Split(key, "_")
			for i, part := range parts {
				if len(part) > 0 {
					parts[i] = gstr.UcFirst(part)
				}
			}
			return gstr.Join(parts, " ")
		}
	}

	return translated
}

// GetJson 以JSON格式获取特定路径下的翻译
func GetJson(ctx context.Context, pattern string) *gjson.Json {
	if instance == nil {
		Init()
	}

	if instance == nil {
		return nil
	}

	language := GetLanguage(ctx)
	instance.SetLanguage(language)

	data := map[string]interface{}{}
	prefix := pattern + "_"

	// 手动从文件夹中获取翻译
	path := "i18n/" + language
	if !gfile.IsDir(path) {
		return gjson.New(data)
	}

	files, err := gfile.ScanDir(path, "*.yaml", false)
	if err != nil {
		return gjson.New(data)
	}

	for _, file := range files {
		content := gfile.GetBytes(file)
		j, err := gjson.LoadContent(content)
		if err != nil {
			continue
		}

		m := j.Map()
		for k, v := range m {
			if gstr.HasPrefix(k, prefix) {
				key := gstr.TrimLeftStr(k, prefix)
				data[key] = v
			}
		}
	}

	return gjson.New(data)
}
