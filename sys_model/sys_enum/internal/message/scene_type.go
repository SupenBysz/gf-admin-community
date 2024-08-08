package message

import "github.com/kysion/base-library/utility/enum"

type SceneTypeEnum enum.IEnumCode[int]

type sceneType struct {
	System SceneTypeEnum
}

var SceneType = sceneType{
	System: enum.New[SceneTypeEnum](8192, "系统通知"),
}

func (e sceneType) New(code int, description string) SceneTypeEnum {

	// 支持自定义场景
	return enum.New[SceneTypeEnum](code, description)
}
