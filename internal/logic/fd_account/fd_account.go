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

// CreateFdAccount 创建财务账号
func (s *sFdAccount) CreateFdAccount(ctx context.Context, info model.FdAccountRegister) (*entity.FdAccount, error) {
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		fmt.Println(err)
	}

	// 根据名称判断财务账号是否存在
	count, _ := dao.FdAccount.Ctx(ctx).Unscoped().Count(dao.FdAccount.Columns().Name, info.Name)
	if count > 0 {
		return nil, service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该公司已有财务账号"), "", dao.FdAccount.Table())
	}

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

// UpdateFdAccountIsEnable 修改财务账号状态（0禁用 1启用）
func (s *sFdAccount) UpdateFdAccountIsEnable(ctx context.Context, id int64, isEnabled int64) (bool, error) {
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
