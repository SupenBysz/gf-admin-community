package component

import (
	"github.com/SupenBysz/gf-admin-community/utility/env"
	"github.com/SupenBysz/gf-admin-community/utility/validator"

	_ "github.com/SupenBysz/gf-admin-community/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func init() {
	env.LoadEnv()
	// 注册服务电话验证规则
	validator.RegisterServicePhone()
}
