package boot

import (
	"github.com/SupenBysz/gf-admin-community/internal/boot/internal"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/kysion/base-library/utility/env"
)

func init() {
	env.LoadEnv()

	internal.InitIp2region()
	internal.InitCustomRules()
	internal.InitIdGenerator()
	internal.InitGlobal()
	internal.InitRedisCache()
	internal.InitLogLevelToDatabase()
	internal.InitPermission()
	internal.InitEmail()
}
