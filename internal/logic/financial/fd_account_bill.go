package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyFinancial "github.com/SupenBysz/gf-admin-community/model/enum/financial"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type hookInfo model.HookEventType[model.AccountBillHookFilter, model.AccountBillHookFunc]

// 财务账单
type sFdAccountBill struct {
	CacheDuration time.Duration
	CachePrefix   string
	hookArr       []hookInfo
}

func init() {
	service.RegisterFdAccountBill(NewFdAccountBill())
}

func NewFdAccountBill() *sFdAccountBill {
	return &sFdAccountBill{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdAccountBill.Table() + "_",
		hookArr:       make([]hookInfo, 0),
	}
}

// CreateAccountBill 创建财务账单
func (s *sFdAccountBill) CreateAccountBill(ctx context.Context, info model.AccountBillRegister) (bool, error) {
	// 判断交易时间是否大于当前系统时间
	now := *gtime.Now()

	if !now.After(info.TradeAt) { // 系统时间是否在交易时间之后
		return false, gerror.New("非法操作！")
	}

	// 交易金额是否为负数
	if info.Amount < 0 {
		return false, gerror.New("交易金额不能是负数")
	}

	var success bool
	var err error

	// 判读收支类型  收入/支出
	if info.InOutType == 1 {
		success, err = s.income(ctx, info)
	} else if info.InOutType == 2 {
		success, err = s.spending(ctx, info)
	}

	if success == false {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), "", dao.FdAccountBill.Table())
	}

	return success == true, err
}

// 修改支付状态

// 接收微信支付宝的支付通知  （账单id，微信的交易id，支付结果）

// 接收微信支付宝的握手

// 待支付：

// HOOK (Hook)

// income 收入
func (s *sFdAccountBill) income(ctx context.Context, info model.AccountBillRegister) (bool, error) {
	// 判断接受者是否存在
	toUser, err := service.SysUser().GetSysUserById(ctx, info.ToUserId)
	if err != nil || toUser == nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收者不存在"), "", dao.FdAccountBill.Table())
	}

	// 先通过财务账号id查询账号出来
	account, err := service.FdAccount().GetAccountById(ctx, info.FdAccountId)

	// 判断需要收款的用户是否存在
	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收账户查询失败"), "", dao.FdAccountBill.Table())
	}

	bill := model.AccountBillInfo{}

	// 使用乐观锁校验余额，和更新余额
	err = dao.FdAccount.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 版本
		version := account.Version
		// 余额 = 之前的余额 + 本次交易的金额
		afterBalance := account.Balance + info.Amount

		// 1. 添加一条收入财务账单流水
		info.BeforeBalance = account.Balance
		info.AfterBalance = afterBalance
		gconv.Struct(info, &bill)
		bill.Id = idgen.NextId()

		result, err := dao.FdAccountBill.Ctx(ctx).Insert(bill)

		if result == nil || err != nil {
			return service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "财务账单创建失败"), "", dao.FdAccountBill.Table())
		}

		// 2.修改财务账号的余额
		// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
		affected, err := service.FdAccount().UpdateAccountBalance(ctx, account.Id, info.Amount, version, info.InOutType)

		if affected == 0 || err != nil {
			return service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新账户余额失败"), "", dao.FdAccountBill.Table())
		}

		for _, hook := range s.hookArr {
			if hook.Key.InTransaction && hook.Key.InOutType == kyFinancial.InOutType.In {
				if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
					hook.Value(ctx, hook.Key, bill)
				}
			}
		}
		return nil
	})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), err.Error(), dao.FdAccountBill.Table())
	}

	for _, hook := range s.hookArr {
		if !hook.Key.InTransaction && hook.Key.InOutType == kyFinancial.InOutType.In {
			if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
				hook.Value(ctx, hook.Key, bill)
			}
		}
	}

	return true, nil
}

// spending 支出
func (s *sFdAccountBill) spending(ctx context.Context, info model.AccountBillRegister) (bool, error) {
	// 先通过财务账号id查询账号出来
	account, err := service.FdAccount().GetAccountById(ctx, info.FdAccountId)
	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收账户查询失败"), "", dao.FdAccountBill.Table())
	}

	bill := model.AccountBillInfo{}

	// 使用乐观锁校验余额，和更新余额
	err = dao.FdAccount.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 版本
		version := account.Version
		// 余额 = 之前的余额 - 本次交易的金额
		afterBalance := account.Balance - info.Amount

		// 判断余额是否足够
		if account.Balance >= info.Amount { // 足够
			// 1. 添加一条财务账单流水
			info.BeforeBalance = account.Balance
			info.AfterBalance = afterBalance

			gconv.Struct(info, &bill)
			bill.Id = idgen.NextId()

			result, err := dao.FdAccountBill.Ctx(ctx).Insert(bill)

			if result == nil || err != nil {
				return service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "财务账单创建失败"), "", dao.FdAccountBill.Table())
			}

			// 2.修改财务账号的余额
			// 参数：上下文, 财务账号id, 需要修改的钱数目, 查询到的版本, 收支类型
			affected, err := service.FdAccount().UpdateAccountBalance(ctx, account.Id, info.Amount, version, info.InOutType)

			if affected == 0 || err != nil {
				return service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "更新账户余额失败"), "", dao.FdAccountBill.Table())
			}

			// Hook调用
			for _, hook := range s.hookArr {
				// 判断收支类型
				if hook.Key.InTransaction && hook.Key.InOutType == kyFinancial.InOutType.Out {
					// 判断交易类型
					if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
						hook.Value(ctx, hook.Key, bill)
					}
				}
			}
		} else {
			return gerror.New("交易发起者的账户余额不足")
		}

		return nil
	})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), err.Error(), dao.FdAccountBill.Table())
	}

	for _, hook := range s.hookArr {
		if !hook.Key.InTransaction && hook.Key.InOutType == kyFinancial.InOutType.Out {
			if hook.Key.TradeType.Code()&info.TradeType == info.TradeType {
				hook.Value(ctx, hook.Key, bill)
			}
		}
	}

	return true, nil
}

// GetAccountBillByAccountId  根据财务账号id获取账单
func (s *sFdAccountBill) GetAccountBillByAccountId(ctx context.Context, accountId int64, pagination *model.Pagination) (*model.AccountBillListRes, error) {
	if accountId == 0 {
		return nil, gerror.New("财务账号不能为0！")
	}

	if pagination == nil {
		pagination = &model.Pagination{
			Page:     1,
			PageSize: 20,
		}
	}

	result, err := daoctl.Query[entity.FdAccountBill](dao.FdAccountBill.Ctx(ctx), &model.SearchParams{
		Filter: append(make([]model.FilterInfo, 0), model.FilterInfo{
			Field: dao.FdAccountBill.Columns().FdAccountId,
			Where: "=",
			Value: accountId,
		}),
		OrderBy: append(make([]model.OrderBy, 0), model.OrderBy{
			Field: dao.FdAccountBill.Columns().CreatedAt,
			Sort:  "asc",
		}),
		Pagination: *pagination,
	}, false)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "账单查询失败", dao.FdAccountBill.Table())
	}

	return (*model.AccountBillListRes)(result), nil
}
