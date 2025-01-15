package boot

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

/*
	初始化全局的缓存服务：Redis｜内存缓存
*/

func InitRedisCache() {
	// 获取配置文件addr对象
	addr, _ := g.Cfg().Get(context.Background(), "redis.default.address")
	pass, _ := g.Cfg().Get(context.Background(), "redis.default.pass")
	db, _ := g.Cfg().Get(context.Background(), "redis.default.db")

	if addr == nil {
		return
	}

	conf, _ := gredis.GetConfig("default")

	// 不同的表分配不同的redis数据库
	conf.Db = 1

	// 设置服务端口和ip
	conf.Address = addr.String()
	if pass != nil {
		conf.Pass = pass.String()
	}
	if db != nil {
		conf.Db = db.Int()
	}

	// 没配置redis ip+端口,配置信息也为空，那么使用内存缓存
	if addr.String() == "" || conf.Address == "" {
		//g.DB().GetCache().SetAdapter(gcache.New())
		return
	}
	// 根据redis配置创建Redis客户端对象
	redis, _ := gredis.New(conf)
	// 全局设置Redis适配器
	g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(redis))
}
