package fd_account

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

//type hookInfo model.KeyValueT[int64, model.UserHookInfo]

type sFdAccount struct {
	CacheDuration time.Duration
	CachePrefix   string
	//hookArr       []hookInfo
}

func init() {
	service.RegisterFdAccount(New())
}

func New() *sFdAccount {
	return &sFdAccount{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdAccount.Table() + "_",
		//hookArr:       make([]hookInfo, 0),
	}
}

// CreateAccount 创建财务账号
func (s *sFdAccount) CreateAccount(ctx context.Context, info model.FdAccountRegister) (*entity.FdAccount, error) {
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		fmt.Println(err)
	}

	// 根据名称判断财务账号是否存在
	//count, _ := dao.FdAccount.Ctx(ctx).Unscoped().Count(dao.FdAccount.Columns().Name, info.Name)
	//if count > 0 {
	//	return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该公司已有财务账号"), "", dao.FdAccount.Table())
	//}

	// 关联用户id是否正确
	user := daoctl.GetById[entity.SysUser](dao.SysUser.Ctx(ctx), info.UnionUserId)
	if user == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "财务账号关联用户id错误"), "", dao.SysUser.Table())
	}

	// 判断货币代码是否符合标准
	currency, err := service.FdCurrenty().GetCurrentyByCurrencyCode(ctx, info.CurrencyCode)
	if err != nil || currency == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "货币代码错误", dao.FdCurrenty.Table())
	}
	if currency.IsLegalTender != 1 {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "请选择合法货币", dao.FdCurrenty.Table())

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

// UpdateAccountBalance 修改财务账户的余额
func (s *sFdAccount) UpdateAccountBalance(ctx context.Context, accountId int64, balance int64, version int) (int64, error) {
	// 根据id + 乐观锁version进行修改
	r, err := dao.FdAccount.Ctx(ctx).Where(do.FdAccount{
		Id:      accountId,
		Version: version, // 获取到的版本 = 数据库中的版本
	}).Update(do.FdAccount{
		Balance: balance,
		Version: version + 1,
	})

	affected, err := r.RowsAffected()

	return affected, err
}

// GetAccountByUnionUserId 根据用户union_user_id获取财务账号
func (s *sFdAccount) GetAccountByUnionUserId(ctx context.Context, unionUserId int64) (*entity.FdAccount, error) {
	if unionUserId == 0 {
		return nil, gerror.New("财务账号用户id不能为空")
	}

	result := entity.FdAccount{}

	dao.FdAccount.Ctx(ctx).Where(do.FdAccount{UnionUserId: unionUserId}).Scan(&result)

	return &result, nil
}
