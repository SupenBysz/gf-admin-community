package message

type messages struct {
	Type  messageType
	State state
}

var Message = messages{
	Type:  MessageType,
	State: State,
}
