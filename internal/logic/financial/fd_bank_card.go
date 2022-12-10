package financial

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 银行卡管理
type sFdBankCard struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	service.RegisterFdBankCard(NewFdBankCard())
}

func NewFdBankCard() *sFdBankCard {
	return &sFdBankCard{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdBankCard.Table() + "_",
	}
}

// CreateBankCard 添加银行卡账号
func (s *sFdBankCard) CreateBankCard(ctx context.Context, info model.BankCardRegister) (*entity.FdBankCard, error) {
	// 判断userid是否存在
	userInfo, err := service.SysUser().GetSysUserById(ctx, info.UserId)
	if err != nil || userInfo == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "用户信息查询失败,用户id错误", dao.SysUser.Table())
	}

	// 判断银行卡是否重复
	bankCard, err := s.GetBankCardByCardNumber(ctx, info.CardNumber)

	if bankCard != nil || err == nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "银行卡号已存在", dao.FdBankCard.Table())
	}

	bankCardInfo := do.FdBankCard{}
	gconv.Struct(info, &bankCardInfo)
	bankCardInfo.Id = idgen.NextId()

	// 添加银行卡
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

	if err != nil || result == nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "银行卡状态修改失败", dao.FdBankCard.Table())
	}

	return true, nil
}

// DeleteBankCardById 删除银行卡 (标记删除: 标记删除的银行卡号，将记录ID的后6位附加到卡号尾部，用下划线隔开,并修改状态)
func (s *sFdBankCard) DeleteBankCardById(ctx context.Context, bankCardId int64) (bool, error) {
	var result sql.Result
	var err error

	bankCard, err := s.GetBankCardById(ctx, bankCardId)
	if err != nil || bankCard == nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "银行卡不存在", dao.FdBankCard.Table())
	}

	err = dao.FdBankCard.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		bankCardIdLen := len(gconv.String(bankCard.Id))
		subId := gstr.SubStr(gconv.String(bankCard.Id), bankCardIdLen-6, 6)

		// 修改 1.修改状态为禁用0   2. 标记银行卡号：bankcardNum_bankCard.Id的后六位
		newBankCardNum := bankCard.CardNumber + "_" + subId

		result, err = dao.FdBankCard.Ctx(ctx).Where(do.FdBankCard{Id: bankCardId}).Data(do.FdBankCard{
			State:      0,
			CardNumber: newBankCardNum,
		}).Update()

		// 删除
		result, err = dao.FdBankCard.Ctx(ctx).Where(do.FdBankCard{Id: bankCardId}).Delete()

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil || result == nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "银行卡删除失败", dao.FdBankCard.Table())
	}

	return true, nil
}
