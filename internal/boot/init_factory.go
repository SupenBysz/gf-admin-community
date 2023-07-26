package boot

import (
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/utility/base_permission"
)

//func init() {
//	InitIdGenerator()
//	InitFactory()
//}

func InitFactory() func() base_permission.IPermission {
	base_permission.Factory = func() base_permission.IPermission {
		return &sys_model.SysPermissionTree{
			SysPermission: &sys_entity.SysPermission{},
		}
	}

	fmt.Println(base_permission.Factory())
	return base_permission.Factory
}

//
//// InitIdGenerator 初始化ID生成器
//func InitIdGenerator() {
//	serviceWorkerId := g.Cfg().MustGet(context.Background(), "service.idGeneratorWorkerId").Uint16()
//	if serviceWorkerId < 1 || serviceWorkerId > 63 {
//		g.Log().Fatal(context.Background(), "service.serviceWorkerId 取值范围只能是 1 ~ 63")
//		return
//	}
//
//	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
//	var options = idgen.NewIdGeneratorOptions(serviceWorkerId)
//	options.WorkerIdBitLength = 10 // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
//	// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。
//	// 保存参数（必须的操作，否则以上设置都不能生效）：
//	idgen.SetIdGenerator(options)
//}
