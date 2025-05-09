package internal

import (
	"github.com/SupenBysz/gf-admin-community/internal/boot"
	"github.com/SupenBysz/gf-admin-community/utility/i18n"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/kysion/base-library/utility/env"
)

func init() {
	env.LoadEnv()

	// 初始化国际化
	i18n.Init()

	idgen.InitIdGenerator()
	boot.InitPermissionFactory()
	boot.InitIp2region()
	boot.InitCustomRules()
	boot.InitGlobal()
	boot.InitRedisCache()
	boot.InitLogLevelToDatabase()
	boot.InitPermission()
	boot.InitEmail()

}
