package fd_invoice

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/internal/consts"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyInvoice "github.com/SupenBysz/gf-admin-community/model/enum/audit"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sFdInvoice struct {
	CacheDuration time.Duration
	CachePrefix   string
	// hookArr       []hookInfo
}

func init() {
	service.RegisterFdInvoice(New())
}

func New() *sFdInvoice {
	return &sFdInvoice{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdInvoice.Table() + "_",
		// hookArr:       make([]hookInfo, 0),
	}
}

// CreateInvoice 创建发票
func (s *sFdInvoice) CreateInvoice(ctx context.Context, info model.FdInvoiceRegister) (*entity.FdInvoice, error) {
	// 判断审核状态
	if info.State == kyInvoice.Action.Reject.Code() && info.AuditReplayMsg == "" {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", dao.FdInvoice.Table())
	}

	// 创建发票
	data := entity.FdInvoice{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()

	_, err := dao.FdInvoice.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "创建发票记录失败", dao.FdInvoice.Table())
	}

	return s.GetInvoiceById(ctx, data.Id)
}

// GetInvoiceById 根据id获取发票
func (s *sFdInvoice) GetInvoiceById(ctx context.Context, id int64) (*entity.FdInvoice, error) {
	if id == 0 {
		return nil, gerror.New("id不能为空")
	}

	result := daoctl.GetById[entity.FdInvoice](dao.FdInvoice.Ctx(ctx), id)

	// result := dao.FdInvoice.Ctx(ctx).Where(do.FdAccount{Id: id})

	return result, nil
}

// GetInvoiceList 获取发票抬头列表
func (s *sFdInvoice) GetInvoiceList(ctx context.Context, info *model.SearchParams, isExport bool) (*model.FdInvoiceListRes, error) {
	if info != nil {
		newFields := make([]model.FilterInfo, 0)

		newFields = append(newFields, model.FilterInfo{
			Field: dao.SysUser.Columns().Type, // type
			Where: "=",
			Value: consts.Global.UserDefaultType,
		})

		// 这是干嘛的 ？？？？？？
		for _, field := range info.Filter {
			if field.Field != dao.FdInvoice.Columns().State {
				newFields = append(newFields, field)
			}
		}
	}

	result, err := daoctl.Query[entity.FdInvoice](dao.FdInvoice.Ctx(ctx), info, isExport)

	return (*model.FdInvoiceListRes)(result), err
}
