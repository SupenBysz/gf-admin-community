package fd_account_bill

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 财务账单
type sFdAccountBill struct {
	CacheDuration time.Duration
	CachePrefix   string
	//hookArr       []hookInfo
}

func init() {
	service.RegisterFdAccountBill(New())
}

func New() *sFdAccountBill {
	return &sFdAccountBill{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdAccountBill.Table() + "_",
		//hookArr:       make([]hookInfo, 0),
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

// income 收入
func (s *sFdAccountBill) income(ctx context.Context, info model.AccountBillRegister) (bool, error) {
	// 判断接受者是否存在
	toUser, err := service.SysUser().GetSysUserById(ctx, info.ToUserId)
	if err != nil || toUser == nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败,交易接收者不存在"), "", dao.SysUser.Table())
	}

	// 使用乐观锁校验余额，和更新余额
	err = dao.FdAccount.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 修改toUser财务账号余额，增加一条收入财务账单
		i := 0
		var affected int64
		for i == 0 {
			if affected == 0 { //未成功
				// 先通过财务账号id查询账号出来
				account, _ := service.FdAccount().GetAccountById(ctx, info.FdAccountId)
				// 版本
				version := account.Version
				// 余额 = 之前的余额 + 本次交易的余额
				afterBalance := account.Balance + info.Amount

				// 判断需要收款的用户是否存在
				if account != nil {

					// 1. 添加一条收入财务账单流水
					info.BeforeBalance = account.Balance
					info.AfterBalance = afterBalance
					bill := entity.FdAccountBill{}
					gconv.Struct(info, &bill)
					bill.Id = idgen.NextId()

					result, err := dao.FdAccountBill.Ctx(ctx).Insert(bill)

					if result == nil || err != nil {
						return gerror.New("财务账单添加失败！")
					}

					// 2.修改财务账号的余额
					affected, err = service.FdAccount().UpdateAccountBalance(ctx, account.Id, afterBalance, version)

				} else {
					return gerror.New("提示交易收款方不存在")
				}

			} else { // affected !=  0
				return nil
			}
		}

		return nil
	})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), err.Error(), dao.FdAccountBill.Table())
	}

	return true, nil
}

// spending 支出
func (s *sFdAccountBill) spending(ctx context.Context, info model.AccountBillRegister) (bool, error) {
	// 使用乐观锁校验余额，和更新余额

	err := dao.FdAccount.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 修改财务账号余额，增加一条财务账单
		i := 0
		var affected int64
		for i == 0 {
			if affected == 0 { //未成功
				// 先通过财务账号id查询账号出来
				account, _ := service.FdAccount().GetAccountById(ctx, info.FdAccountId)
				version := account.Version
				afterBalance := account.Balance - info.Amount

				// 判断余额是否足够
				if account.Balance >= info.Amount { // 足够

					// 1. 添加一条财务账单流水
					info.BeforeBalance = account.Balance
					info.AfterBalance = afterBalance

					bill := entity.FdAccountBill{}
					gconv.Struct(info, &bill)
					bill.Id = idgen.NextId()

					result, err := dao.FdAccountBill.Ctx(ctx).Insert(bill)

					if result == nil || err != nil {
						return gerror.New("财务账单添加失败！")
					}

					// 2.修改财务账号的余额
					affected, err = service.FdAccount().UpdateAccountBalance(ctx, account.Id, afterBalance, version)

				} else {
					return gerror.New("交易发起者的账户余额不足")
				}

			} else { // affected !=  0
				return nil
			}
		}

		return nil
	})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "交易失败"), err.Error(), dao.FdAccountBill.Table())
	}

	return true, nil
}
