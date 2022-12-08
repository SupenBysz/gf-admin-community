package fd_invoice_detail

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/internal/consts"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
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
	//hookArr       []hookInfo
}

func init() {
	service.RegisterFdInvoiceDetail(New())
}

func New() *sFdInvoiceDetail {
	return &sFdInvoiceDetail{
		CacheDuration: time.Hour,
		CachePrefix:   dao.FdInvoiceDetail.Table() + "_",
		//hookArr:       make([]hookInfo, 0),
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

// UpdateInvoiceDetail 修改发票详情 (审核发票后触发,把信息录进数据表)
func (s *sFdInvoiceDetail) UpdateInvoiceDetail(ctx context.Context, info entity.FdInvoiceDetail) (bool, error) {
	if info.State == 0 {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", dao.FdInvoiceDetail.Table())
	}

	if info.State == 4 && info.AuditReplyMsg == "" {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", dao.FdInvoiceDetail.Table())
	}

	invoiceDetailInfo, err := s.GetInvoiceDetailById(ctx, info.Id)
	if err != nil || invoiceDetailInfo == nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情ID参数错误", dao.FdInvoiceDetail.Table())
	}

	// 代表已审过的
	if invoiceDetailInfo.State == 4 || invoiceDetailInfo.State == 8 || invoiceDetailInfo.State == 16 { //4开票失败、8已开票、16已撤销
		return false, service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", dao.FdInvoiceDetail.Table())
	}

	// 事务里的Hook，还是事务外的Hook

	// 添加审核过后的数据
	_, err = dao.FdInvoiceDetail.Ctx(ctx).OmitNilData().Data(do.FdInvoiceDetail{
		// 修改非必选参数的数值
		AuditUserIds:  info.AuditUserIds,
		MakeType:      info.MakeType,
		MakeUserId:    info.MakeUserId,
		MakeAt:        gtime.Now(),
		CourierName:   info.CourierName,
		CourierNumber: info.CourierNumber,
		FdInvoiceId:   info.FdInvoiceId,
		AuditUserId:   info.AuditUserId,
		AuditReplyMsg: info.AuditReplyMsg,
		AuditAt:       gtime.Now(),

		// 修改状态
		State: info.State,
	}).Where(do.FdInvoiceDetail{
		Id: invoiceDetailInfo.Id,
	}).Update()

	if err != nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "发票详情数据修改失败", dao.FdInvoiceDetail.Table())
	}

	return true, nil
}

// GetInvoiceDetailList 获取发票详情列表
func (s *sFdInvoiceDetail) GetInvoiceDetailList(ctx context.Context, info *model.SearchParams, isExport bool) (*model.FdInvoiceDetailListRes, error) {
	if info != nil {
		newFields := make([]model.FilterInfo, 0)

		newFields = append(newFields, model.FilterInfo{
			Field: dao.SysUser.Columns().Type, //type
			Where: "=",
			Value: consts.Global.UserDefaultType,
		})

		for _, field := range info.Filter {
			if field.Field != dao.FdInvoiceDetail.Columns().Type {
				newFields = append(newFields, field)
			}
		}
	}

	result, err := daoctl.Query[entity.FdInvoiceDetail](dao.FdInvoiceDetail.Ctx(ctx), info, isExport)

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
		_, err = dao.FdInvoiceDetail.Ctx(ctx).Where(do.FdInvoiceDetail{Id: id}).Update(do.FdInvoiceDetail{State: 16})

		// 删除
		_, err = dao.FdInvoiceDetail.Ctx(ctx).Where(do.FdInvoiceDetail{Id: id}).Delete()

		if err != nil {
			return err
		}

		return nil
	})

	return err == nil, err
}
