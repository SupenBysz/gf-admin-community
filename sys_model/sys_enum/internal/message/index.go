package message

type messages struct {
	Type      messageType
	State     state
	SceneType sceneType
}

var Message = messages{
	Type:      MessageType,
	State:     State,
	SceneType: SceneType,
}
