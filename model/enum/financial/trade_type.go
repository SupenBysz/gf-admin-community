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
	Transfer:        kyEnum.New(1, "转账"),
	Consumption:     kyEnum.New(2, "消费"),
	Refund:          kyEnum.New(4, "退款"),
	Commission:      kyEnum.New(8, "佣金"),
	SecurityDeposit: kyEnum.New(16, "保证金"),
	EarnestMoney:    kyEnum.New(32, "诚意金"),
	ServiceCharge:   kyEnum.New(64, "手续费/服务费"),
	CashWithdrawal:  kyEnum.New(128, "提现"),
	Recharge:        kyEnum.New(256, "充值"),
	OperatingIncome: kyEnum.New(512, "营收"),
	Other:           kyEnum.New(8192, "其它"),
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
