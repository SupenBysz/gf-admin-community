package idgen

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
)

// InitIdGenerator 初始化ID生成器
func InitIdGenerator() {
	// 用户ID生成器，分布式终端标识，取值范围 1 ~ 63，解决分布式多服务横向扩展时保证生成的ID不重复
	serviceWorkerId := g.Cfg().MustGet(context.Background(), "service.idGeneratorWorkerId").Uint16()
	if serviceWorkerId < 1 || serviceWorkerId > 63 {
		g.Log().Fatal(context.Background(), "service.serviceWorkerId 取值范围只能是 1 ~ 63")
		return
	}

	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(serviceWorkerId) // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
	options.WorkerIdBitLength = 6                              // WorkerIdBitLength的值越小，那么生成的id就越小，默认值是6
	// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。
	// 保存参数（必须的操作，否则以上设置都不能生效）：
	idgen.SetIdGenerator(options)

	/*
			workerId 是分布式终端标识，多服务的模式下 ，如果workerID是一样的，那么就有概率出现重复的ID，虽然可能是小概率，但是如果多服务的workerID都是不一样的，那么就没有任何概率出现重复的ID。
		 	WorkerIdBitLength的值越小，那么生成的id就越小
	*/
}

var isInit = false

// NextId 构建新ID
func NextId() int64 {
	if isInit == false {
		isInit = true
		InitIdGenerator()
	}

	return idgen.NextId()
}
