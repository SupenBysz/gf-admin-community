package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyInvoice "github.com/SupenBysz/gf-admin-community/model/enum/invoice"
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
	service.RegisterFdInvoice(NewFdInvoice())
}

func NewFdInvoice() *sFdInvoice {
	return &sFdInvoice{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdInvoice.Table() + "_",
		// hookArr:       make([]hookInfo, 0),
	}
}

// CreateInvoice 创建发票
func (s *sFdInvoice) CreateInvoice(ctx context.Context, info model.FdInvoiceRegister) (*entity.FdInvoice, error) {
	// 判断审核状态
	if info.State == kyInvoice.AuditType.Reject.Code() && info.AuditReplayMsg == "" {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", dao.FdInvoice.Table())
	}

	// 创建发票
	data := entity.FdInvoice{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()

	data.State = kyInvoice.AuditType.WaitReview.Code()

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

	if result == nil {
		return nil, service.SysLogs().InfoSimple(ctx, nil, "当前没有发票抬头记录", dao.FdInvoice.Table())
	}

	return result, nil
}

// GetInvoiceList 获取发票抬头列表
func (s *sFdInvoice) GetInvoiceList(ctx context.Context, info *model.SearchParams, userId int64) (*model.FdInvoiceListRes, error) {
	newFields := make([]model.FilterInfo, 0)
	// 筛选条件强制指定所属用户
	newFields = append(newFields, model.FilterInfo{
		Field: dao.FdInvoice.Columns().UserId, // type
		Where: "=",
		Value: userId,
	})

	if info != nil {
		// 排除搜索参数中指定的所属用户参数
		for _, field := range info.Filter {
			if field.Field != dao.FdInvoice.Columns().UserId {
				newFields = append(newFields, field)
			}
		}
	}
	info.Filter = newFields

	result, err := daoctl.Query[entity.FdInvoice](dao.FdInvoice.Ctx(ctx), info, false)

	return (*model.FdInvoiceListRes)(result), err
}
