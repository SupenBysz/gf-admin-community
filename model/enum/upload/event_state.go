package kyUpload

import (
	"github.com/SupenBysz/gf-admin-community/model/enum"
)

type EventStateEnum kyEnum.Code

type eventState struct {
	AfterCache EventStateEnum
	BeforeSave EventStateEnum
	AfterSave  EventStateEnum
}

var EventState = eventState{
	AfterCache: kyEnum.NewT[EventStateEnum](1, "已缓存"),
	BeforeSave: kyEnum.NewT[EventStateEnum](2, "保存前"),
	AfterSave:  kyEnum.NewT[EventStateEnum](4, "保存后"),
}

func (e eventState) New(code int, description string) EventStateEnum {
	if (code&EventState.AfterCache.Code()) == EventState.AfterCache.Code() ||
		(code&EventState.BeforeSave.Code()) == EventState.BeforeSave.Code() ||
		(code&EventState.AfterSave.Code()) == EventState.AfterSave.Code() {
		return kyEnum.NewT[EventStateEnum](code, description)
	}
	panic("uploadEventState: error")
}
