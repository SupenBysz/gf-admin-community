package sys_enum_upload

import (
	"github.com/SupenBysz/gf-admin-community/utility/enum"
)

type EventStateEnum enum.IEnumCode[int]

type eventState struct {
	AfterCache EventStateEnum
	BeforeSave EventStateEnum
	AfterSave  EventStateEnum
}

var EventState = eventState{
	AfterCache: enum.New[EventStateEnum](1, "已缓存"),
	BeforeSave: enum.New[EventStateEnum](2, "保存前"),
	AfterSave:  enum.New[EventStateEnum](4, "保存后"),
}

func (e eventState) New(code int, description string) EventStateEnum {
	if (code&EventState.AfterCache.Code()) == EventState.AfterCache.Code() ||
		(code&EventState.BeforeSave.Code()) == EventState.BeforeSave.Code() ||
		(code&EventState.AfterSave.Code()) == EventState.AfterSave.Code() {
		return enum.New[EventStateEnum](code, description)
	}
	panic("Upload.EventState.New: error")
}
