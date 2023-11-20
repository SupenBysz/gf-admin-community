package internal

import (
	"github.com/SupenBysz/gf-admin-community/internal/boot"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/kysion/base-library/utility/env"
)

func init() {
	env.LoadEnv()

	boot.InitIdGenerator()
	//init_factory.InitFactory()
	boot.InitIp2region()
	boot.InitCustomRules()
	boot.InitGlobal()
	boot.InitRedisCache()
	boot.InitLogLevelToDatabase()
	boot.InitPermission()
	boot.InitEmail()

}
