package sys_controller

import (
	"github.com/SupenBysz/gf-admin-community/utility/i18n"
	"github.com/SupenBysz/gf-admin-community/utility/response"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SysCommon 系统公共控制器
var SysCommon = SysCommonController{}

// SysCommonController 系统公共控制器
type SysCommonController struct {
}

// SetLanguage 设置语言
// @Summary 设置语言
// @Description 设置用户界面语言，系统会将语言设置保存在Cookie中。请求中的lang参数优先级最高，其次是Accept-Language请求头，最后是Cookie。
// @Tags 系统公共
// @Accept json
// @Produce json
// @Param lang query string true "语言代码，如zh-CN, en-US"
// @Success 200 {object} response.JsonRes "设置成功"
// @Router /api/common/setLanguage [get]
func (c *SysCommonController) SetLanguage(r *ghttp.Request) {
	ctx := r.Context()
	lang := r.Get("lang").String()

	// 检查语言是否有效
	if lang != "zh-CN" && lang != "en-US" {
		response.JsonExit(r, 1, i18n.T(ctx, "error_parameter_error"))
		return
	}

	// 设置Cookie
	r.Cookie.Set("lang", lang)

	response.JsonExit(r, 0, i18n.T(ctx, "common_success"), lang)
}

// GetLanguages 获取支持的语言列表
// @Summary 获取支持的语言列表
// @Description 获取系统支持的语言列表
// @Tags 系统公共
// @Accept json
// @Produce json
// @Success 200 {object} response.JsonRes "语言列表"
// @Router /api/common/getLanguages [get]
func (c *SysCommonController) GetLanguages(r *ghttp.Request) {
	ctx := r.Context()

	// 支持的语言列表
	languages := []map[string]string{
		{"code": "zh-CN", "name": "简体中文"},
		{"code": "en-US", "name": "English"},
	}

	response.JsonExit(r, 0, i18n.T(ctx, "common_success"), languages)
}

// GetCurrentLanguage 获取当前语言
// @Summary 获取当前语言
// @Description 获取当前用户使用的语言
// @Tags 系统公共
// @Accept json
// @Produce json
// @Success 200 {object} response.JsonRes "当前语言"
// @Router /api/common/getCurrentLanguage [get]
func (c *SysCommonController) GetCurrentLanguage(r *ghttp.Request) {
	ctx := r.Context()

	lang := i18n.GetLanguage(ctx)

	response.JsonExit(r, 0, i18n.T(ctx, "common_success"), lang)
}
