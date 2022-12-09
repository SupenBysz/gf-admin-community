package fd_invoice_detail

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyInvoice "github.com/SupenBysz/gf-admin-community/model/enum/invoice"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 发票详情
type sFdInvoiceDetail struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	service.RegisterFdInvoiceDetail(New())
}

func New() *sFdInvoiceDetail {
	return &sFdInvoiceDetail{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdInvoiceDetail.Table() + "_",
	}
}

// CreateInvoiceDetail 创建发票详情，相当于创建审核列表，审核是人工审核
func (s *sFdInvoiceDetail) CreateInvoiceDetail(ctx context.Context, info model.FdInvoiceDetailRegister) (*entity.FdInvoiceDetail, error) {
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 获取发票详情

	// 创建发票审核详情记录
	data := entity.FdInvoiceDetail{}
	gconv.Struct(info, &data)

	data.Id = idgen.NextId()
	// 设置审核状态为待审核
	data.State = kyInvoice.AuditType.WaitReview.Code()

	result, err := dao.FdInvoiceDetail.Ctx(ctx).Data(data).Insert()

	if err != nil || result == nil {
		return nil, err
	}

	return s.GetInvoiceDetailById(ctx, data.Id)
}

// GetInvoiceDetailById 根据id获取发票详情
func (s *sFdInvoiceDetail) GetInvoiceDetailById(ctx context.Context, id int64) (*entity.FdInvoiceDetail, error) {
	if id == 0 {
		return nil, gerror.New("发票详情id不能为空")
	}

	result := daoctl.GetById[entity.FdInvoiceDetail](dao.FdInvoiceDetail.Ctx(ctx), id)

	invoiceDetail := entity.FdInvoiceDetail{}
	gconv.Struct(result, &invoiceDetail)

	return &invoiceDetail, nil
}

// MakeInvoiceDetail 开票
func (s *sFdInvoiceDetail) MakeInvoiceDetail(ctx context.Context, invoiceDetailId int64, makeInvoiceDetail model.FdMakeInvoiceDetail) (bool, error) {
	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, invoiceDetailId)
	if err != nil || invoiceDetailInfo == nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情ID参数错误", dao.FdInvoiceDetail.Table())
	}

	// 校验状态是否为待开票
	if invoiceDetailInfo.State != kyInvoice.State.WaitForInvoice.Code() {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "开票失败，状态类型不匹配", dao.FdInvoiceDetail.Table())
	}

	// 添加审核过后的数据
	_, err = dao.FdInvoiceDetail.Ctx(ctx).OmitNilData().Data(do.FdInvoiceDetail{
		MakeType:      makeInvoiceDetail.Type.Code(),
		MakeUserId:    makeInvoiceDetail.UserId,
		CourierName:   makeInvoiceDetail.CourierName,
		CourierNumber: makeInvoiceDetail.CourierNumber,
		State:         kyInvoice.State.Success,
		MakeAt:        gtime.Now(),
	}).Where(do.FdInvoiceDetail{
		Id: invoiceDetailInfo.Id,
	}).Update()

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情数据修改失败", dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// AuditInvoiceDetail 审核发票
func (s *sFdInvoiceDetail) AuditInvoiceDetail(ctx context.Context, invoiceDetailId int64, auditInfo model.FdInvoiceAuditInfo) (bool, error) {
	// 审核行仅允许 kyInvoice.State.WaitForInvoice 和 kyInvoice.State.Failure
	if auditInfo.State != kyInvoice.State.WaitForInvoice.Code() && auditInfo.State != kyInvoice.State.Failure.Code() {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", dao.FdInvoiceDetail.Table())
	}

	if auditInfo.State == kyInvoice.State.Failure.Code() && auditInfo.ReplyMsg == "" {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", dao.FdInvoiceDetail.Table())
	}

	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, invoiceDetailId)
	if err != nil || invoiceDetailInfo == nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情ID参数错误", dao.FdInvoiceDetail.Table())
	}

	// 代表已审过的
	if invoiceDetailInfo.State > kyInvoice.State.WaitAudit.Code() {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", dao.FdInvoiceDetail.Table())
	}

	// 添加审核过后的数据
	_, err = dao.FdInvoiceDetail.Ctx(ctx).OmitNilData().Data(do.FdInvoiceDetail{
		AuditUserId:   auditInfo.UserId,
		AuditReplyMsg: auditInfo.ReplyMsg,
		State:         auditInfo.State,
		AuditAt:       gtime.Now(),
	}).Where(do.FdInvoiceDetail{
		Id: invoiceDetailInfo.Id,
	}).Update()

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情数据修改失败", dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// GetInvoiceDetailList 根据财务账户，获取已开票的发票详情列表
func (s *sFdInvoiceDetail) GetInvoiceDetailList(ctx context.Context, fdAccountId int64) (*model.FdInvoiceDetailListRes, error) {
	account, err := service.FdAccount().GetAccountById(ctx, fdAccountId)
	if err != nil {
		return nil, err
	}

	result, err := daoctl.Query[entity.FdInvoiceDetail](dao.FdInvoiceDetail.Ctx(ctx), &model.SearchParams{
		Filter: append(make([]model.FilterInfo, 0), model.FilterInfo{
			Field: dao.FdInvoiceDetail.Columns().FdInvoiceId,
			Where: "=",
			Value: account.Id,
		}),
		Pagination: model.Pagination{
			Page:     1,
			PageSize: -1,
		},
	}, false)

	return (*model.FdInvoiceDetailListRes)(result), err
}

// DeleteInvoiceDetail 标记删除发票详情
func (s *sFdInvoiceDetail) DeleteInvoiceDetail(ctx context.Context, id int64) (bool, error) {
	// 判断是否存在该发票
	invoice, err := s.GetInvoiceDetailById(ctx, id)
	if err != nil || invoice == nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情id错误", dao.FdInvoiceDetail.Table())
	}

	err = dao.FdInvoiceDetail.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 状态修改为已撤消，
		_, err = dao.FdInvoiceDetail.Ctx(ctx).
			Where(do.FdInvoiceDetail{Id: id}).
			Update(do.FdInvoiceDetail{State: kyInvoice.State.Cancel.Code()})
		if err != nil {
			return err
		}

		// 删除
		_, err = dao.FdInvoiceDetail.Ctx(ctx).
			Where(do.FdInvoiceDetail{Id: id}).
			Delete()

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, err, "发票删除失败", dao.FdInvoiceDetail.Table())
	}

	return err == nil, err
}
