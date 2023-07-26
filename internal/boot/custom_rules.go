package boot

import (
	"github.com/SupenBysz/gf-admin-community/utility/rules"
	"github.com/kysion/base-library/utility/validator"
)

// InitCustomRules 注册自定义参数校验规则
func InitCustomRules() {
	// 注册电话验证规则
	validator.RegisterServicePhone()
	// 注册资质自定义规则
	rules.RequiredLicense()
}
