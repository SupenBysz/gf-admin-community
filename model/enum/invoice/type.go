package kyInvoice

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type TypeEnum kyEnum.Code

type makeType struct {
	Normal       TypeEnum
	Special      TypeEnum
	professional TypeEnum
}

var Type = makeType{
	Normal:       kyEnum.New(1, "普通发票"),
	Special:      kyEnum.New(2, "增值税专用发票"),
	professional: kyEnum.New(3, "专业发票"),
}

func (e makeType) New(code int, description string) TypeEnum {
	if (code&Type.Normal.Code()) == Type.Normal.Code() ||
		(code&Type.Special.Code()) == Type.Special.Code() ||
		(code&Type.professional.Code()) == Type.professional.Code() {
		return kyEnum.New(code, description)
	} else {
		panic("kyInvoice.Type.New: error")
	}
}
