package fd_bank_card

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 银行卡管理
type sFdBankCard struct {
	CacheDuration time.Duration
	CachePrefix   string
	//hookArr       []hookInfo
}

func init() {
	service.RegisterFdBankCard(New())
}

func New() *sFdBankCard {
	return &sFdBankCard{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdBankCard.Table() + "_",
		//hookArr:       make([]hookInfo, 0),
	}
}

// CreateBankCard 添加银行卡账号
func (s *sFdBankCard) CreateBankCard(ctx context.Context, info model.BankCardRegister) (*entity.FdBankCard, error) {
	// 判断userid是否存在
	userInfo, err := service.SysUser().GetSysUserById(ctx, info.UserId)
	if err != nil || userInfo == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败,用户id错误", dao.SysUser.Table())
	}

	bankCardInfo := do.FdBankCard{}
	gconv.Struct(info, &bankCardInfo)
	bankCardInfo.Id = idgen.NextId()

	_, err = dao.FdBankCard.Ctx(ctx).Data(bankCardInfo).Insert()

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "银行卡号添加失败", dao.SysUser.Table())
	}

	return s.GetBankCardById(ctx, gconv.Int64(bankCardInfo.Id))
}

// GetBankCardById 根据银行卡id获取银行卡信息
func (s *sFdBankCard) GetBankCardById(ctx context.Context, id int64) (*entity.FdBankCard, error) {
	if id == 0 {
		return nil, gerror.New("银行卡id不能为空")
	}
	result := daoctl.GetById[entity.FdBankCard](dao.FdBankCard.Ctx(ctx), id)

	return result, nil
}

// GetBankCardByCardNumber 根据银行卡号获取银行卡
func (s *sFdBankCard) GetBankCardByCardNumber(ctx context.Context, cardNumber string) (*entity.FdBankCard, error) {
	if cardNumber == "" {
		return nil, gerror.New("银行卡号码不能为空")
	}

	bankCard := entity.FdBankCard{}

	err := dao.FdBankCard.Ctx(ctx).Where(do.FdBankCard{CardNumber: cardNumber}).Scan(&bankCard)
	if err != nil {
		return nil, err
	}

	return &bankCard, nil
}

// UpdateBankCardState 修改银行卡状态 (0禁用 1正常)
func (s *sFdBankCard) UpdateBankCardState(ctx context.Context, bankCardId int64, state int) (bool, error) {
	bankCard, err := s.GetBankCardById(ctx, bankCardId)
	if err != nil || bankCard == nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "银行卡不存在", dao.FdBankCard.Table())
	}

	// 修改状态
	result, err := dao.FdBankCard.Ctx(ctx).Where(do.FdBankCard{Id: bankCardId}).Update(do.FdBankCard{State: state})

	//_, err = dao.SysUser.Ctx(ctx).Where(do.SysUser{Id: sysUserInfo.Id}).Update(do.SysUser{Password: pwdHash})

	if err != nil || result == nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "银行卡状态修改失败", dao.FdBankCard.Table())
	}

	return true, nil
}
