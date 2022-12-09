package kyInvoice

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type StateEnum kyEnum.Code

type state struct {
	WaitAudit      StateEnum
	WaitForInvoice StateEnum
	Failure        StateEnum
	Success        StateEnum
	Cancel         StateEnum
}

var State = state{
	WaitAudit:      kyEnum.New(1, "待审核"),
	WaitForInvoice: kyEnum.New(2, "待开票"),
	Failure:        kyEnum.New(4, "开票失败"),
	Success:        kyEnum.New(8, "已开票"),
	Cancel:         kyEnum.New(16, "已撤销"),
}

func (e state) New(code int, description string) StateEnum {
	if (code&State.WaitAudit.Code()) == State.WaitAudit.Code() ||
		(code&State.WaitForInvoice.Code()) == State.WaitForInvoice.Code() ||
		(code&State.Failure.Code()) == State.Failure.Code() ||
		(code&State.Success.Code()) == State.Success.Code() ||
		(code&State.Cancel.Code()) == State.Cancel.Code() {
		return kyEnum.New(code, description)
	} else {
		panic("kyInvoice.State.New: error")
	}
}
