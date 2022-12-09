package kyFinancial

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type InOutTypeEnum kyEnum.Code

type inOutType struct {
	In  InOutTypeEnum
	Out InOutTypeEnum
}

var InOutType = inOutType{
	In:  kyEnum.New(1, "收入"),
	Out: kyEnum.New(2, "支出"),
}

func (e inOutType) New(code int, description string) InOutTypeEnum {
	if (code&InOutType.In.Code()) == InOutType.In.Code() ||
		(code&InOutType.Out.Code()) == InOutType.Out.Code() {
		return kyEnum.NewT[InOutTypeEnum](code, description)
	}
	panic("uploadEventState: error")
}
