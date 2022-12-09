package fd_account

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// type hookInfo model.KeyValueT[int64, model.UserHookInfo]

type sFdAccount struct {
	CacheDuration time.Duration
	CachePrefix   string
	// hookArr       []hookInfo
}

func init() {
	service.RegisterFdAccount(New())
}

func New() *sFdAccount {
	return &sFdAccount{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdAccount.Table() + "_",
		// hookArr:       make([]hookInfo, 0),
	}
}

// CreateAccount 创建财务账号
func (s *sFdAccount) CreateAccount(ctx context.Context, info model.FdAccountRegister) (*entity.FdAccount, error) {
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 根据名称判断财务账号是否存在
	// count, _ := dao.FdAccount.Ctx(ctx).Unscoped().Count(dao.FdAccount.Columns().Name, info.Name)
	// if count > 0 {
	//	return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该公司已有财务账号"), "", dao.FdAccount.Table())
	// }

	// 关联用户id是否正确
	user := daoctl.GetById[entity.SysUser](dao.SysUser.Ctx(ctx), info.UnionUserId)
	if user == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "财务账号关联用户id错误"), "", dao.SysUser.Table())
	}

	// 判断货币代码是否符合标准
	currency, err := service.FdCurrency().GetCurrencyByCurrencyCode(ctx, info.CurrencyCode)
	if err != nil || currency == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "货币代码错误", dao.FdCurrency.Table())
	}
	if currency.IsLegalTender != 1 {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "请选择合法货币", dao.FdCurrency.Table())

	}
	// 生产随机id
	data := do.FdAccount{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()

	// 插入财务账号
	_, err = dao.FdAccount.Ctx(ctx).Insert(data)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "财务账号添加失败", dao.FdAccount.Table())
	}

	return s.GetAccountById(ctx, gconv.Int64(data.Id))
}

// GetAccountById 根据ID获取财务账号
func (s *sFdAccount) GetAccountById(ctx context.Context, id int64) (*entity.FdAccount, error) {
	if id == 0 {
		return nil, gerror.New("财务账号id不能为空")
	}
	result := daoctl.GetById[entity.FdAccount](dao.FdAccount.Ctx(ctx), id)

	return result, nil
}

// UpdateAccountIsEnable 修改财务账号状态（是否启用：0禁用 1启用）
func (s *sFdAccount) UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int64) (bool, error) {
	account := daoctl.GetById[entity.FdAccount](dao.FdAccount.Ctx(ctx), id)
	if account == nil {
		return false, gerror.New("财务账号不存在")
	}

	_, err := dao.FdAccount.Ctx(ctx).Where(do.FdAccount{Id: id}).Update(do.FdAccount{IsEnabled: isEnabled})
	if err != nil {
		return false, err
	}

	return true, nil
}

// HasAccountByName 根据账户名查询财务账户
func (s *sFdAccount) HasAccountByName(ctx context.Context, name string) (*entity.FdAccount, error) {
	data := &entity.FdAccount{}
	err := dao.FdAccount.Ctx(ctx).Where(do.FdAccount{Name: name}).Scan(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateAccountLimitState 修改财务账号的限制状态 （0不限制，1限制支出、2限制收入）
func (s *sFdAccount) UpdateAccountLimitState(ctx context.Context, id int64, limitState int64) (bool, error) {
	_, err := dao.FdAccount.Ctx(ctx).Where(do.FdAccount{Id: id}).Update(do.FdAccount{LimitState: limitState})
	if err != nil {
		return false, err
	}

	return true, nil
}

// QueryAccountListByUserId 获取指定用户的所有财务账号
func (s *sFdAccount) QueryAccountListByUserId(ctx context.Context, userId int64) (*model.AccountList, error) {
	accountList := model.AccountList{}

	if userId == 0 {
		return nil, gerror.New("用户id不能为空")
	}

	err := dao.FdAccount.Ctx(ctx).Where(do.FdAccount{UnionUserId: userId}).Scan(&accountList)

	if err != nil || len(accountList) <= 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "该账户没有财务账号", dao.FdAccount.Table())
	}

	return &accountList, nil
}

// UpdateAccountBalance 修改财务账户余额(上下文, 财务账号id, 需要修改的钱数目, 版本, 收支类型)
func (s *sFdAccount) UpdateAccountBalance(ctx context.Context, accountId int64, amount int64, version int, inOutType int) (int64, error) {
	db := dao.FdAccount.Ctx(ctx)

	data := do.FdAccount{
		Version: gdb.Raw(dao.FdAccount.Columns().Version + "+1"),
	}

	if inOutType == 1 { // 收入
		// 余额 = 之前的余额 + 本次交易的余额
		data.Balance = gdb.Raw(dao.FdAccount.Columns().Balance + "+" + gconv.String(amount))
	} else if inOutType == 2 { // 支出
		// 余额 = 之前的余额 - 本次交易的余额
		data.Balance = gdb.Raw(dao.FdAccount.Columns().Balance + "-" + gconv.String(amount))
	}

	result, err := db.Data(data).Where(do.FdAccount{
		Id:      accountId,
		Version: version,
	}).Update()

	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()

	return affected, err
}

// GetAccountByUnionUserIdAndCurrencyCode 根据用户union_user_id和货币代码currency_code获取财务账号
func (s *sFdAccount) GetAccountByUnionUserIdAndCurrencyCode(ctx context.Context, unionUserId int64, currencyCode string) (*entity.FdAccount, error) {
	if unionUserId == 0 {
		return nil, gerror.New("财务账号用户id不能为空")
	}

	result := entity.FdAccount{}

	// 查找指定用户名下指定货币类型的财务账号
	err := dao.FdAccount.Ctx(ctx).Where(do.FdAccount{
		UnionUserId:  unionUserId,
		CurrencyCode: currencyCode,
	}).Scan(&result)

	return &result, err
}
