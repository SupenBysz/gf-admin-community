package kyFinancial

import kyEnum "github.com/SupenBysz/gf-admin-community/model/enum"

type TradeTypeEnum kyEnum.Code

type tradeType struct {
	Transfer        TradeTypeEnum
	Consumption     TradeTypeEnum
	Refund          TradeTypeEnum
	Commission      TradeTypeEnum
	SecurityDeposit TradeTypeEnum
	EarnestMoney    TradeTypeEnum
	ServiceCharge   TradeTypeEnum
	CashWithdrawal  TradeTypeEnum
	Recharge        TradeTypeEnum
	OperatingIncome TradeTypeEnum
	Other           TradeTypeEnum
}

var TradeType = tradeType{
	Transfer:        kyEnum.NewT[TradeTypeEnum](1, "转账"),
	Consumption:     kyEnum.NewT[TradeTypeEnum](2, "消费"),
	Refund:          kyEnum.NewT[TradeTypeEnum](4, "退款"),
	Commission:      kyEnum.NewT[TradeTypeEnum](8, "佣金"),
	SecurityDeposit: kyEnum.NewT[TradeTypeEnum](16, "保证金"),
	EarnestMoney:    kyEnum.NewT[TradeTypeEnum](32, "诚意金"),
	ServiceCharge:   kyEnum.NewT[TradeTypeEnum](64, "手续费/服务费"),
	CashWithdrawal:  kyEnum.NewT[TradeTypeEnum](128, "提现"),
	Recharge:        kyEnum.NewT[TradeTypeEnum](256, "充值"),
	OperatingIncome: kyEnum.NewT[TradeTypeEnum](512, "营收"),
	Other:           kyEnum.NewT[TradeTypeEnum](8192, "其它"),
}

func (e tradeType) New(code int, description string) TradeTypeEnum {
	if (code&TradeType.Transfer.Code()) == TradeType.Transfer.Code() ||
		(code&TradeType.Consumption.Code()) == TradeType.Consumption.Code() ||
		(code&TradeType.Refund.Code()) == TradeType.Refund.Code() ||
		(code&TradeType.Commission.Code()) == TradeType.Commission.Code() ||
		(code&TradeType.SecurityDeposit.Code()) == TradeType.SecurityDeposit.Code() ||
		(code&TradeType.EarnestMoney.Code()) == TradeType.EarnestMoney.Code() ||
		(code&TradeType.ServiceCharge.Code()) == TradeType.ServiceCharge.Code() ||
		(code&TradeType.CashWithdrawal.Code()) == TradeType.CashWithdrawal.Code() ||
		(code&TradeType.Recharge.Code()) == TradeType.Recharge.Code() ||
		(code&TradeType.OperatingIncome.Code()) == TradeType.OperatingIncome.Code() ||
		(code&TradeType.Other.Code()) == TradeType.Other.Code() {
		return kyEnum.NewT[TradeTypeEnum](code, description)
	}
	panic("kyFinancial.TradeType.New: error")
}
