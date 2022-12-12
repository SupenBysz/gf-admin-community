// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysSmsLogsDao is the data access object for table sys_sms_logs.
type SysSmsLogsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SysSmsLogsColumns // columns contains all the column names of Table for convenient usage.
}

// SysSmsLogsColumns defines and stores column names for table sys_sms_logs.
type SysSmsLogsColumns struct {
	Id        string //
	Type      string // 短信平台：qyxs：企业信使
	Context   string // 短信内容
	Mobile    string // 手机号
	State     string // 发送状态
	Result    string // 短信接口返回内容
	UserId    string // 用户ID
	LicenseId string // 主体ID
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// sysSmsLogsColumns holds the columns for table sys_sms_logs.
var sysSmsLogsColumns = SysSmsLogsColumns{
	Id:        "id",
	Type:      "type",
	Context:   "context",
	Mobile:    "mobile",
	State:     "state",
	Result:    "result",
	UserId:    "user_id",
	LicenseId: "license_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSysSmsLogsDao creates and returns a new DAO object for table data access.
func NewSysSmsLogsDao() *SysSmsLogsDao {
	return &SysSmsLogsDao{
		group:   "default",
		table:   "sys_sms_logs",
		columns: sysSmsLogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysSmsLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysSmsLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysSmsLogsDao) Columns() SysSmsLogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysSmsLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysSmsLogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysSmsLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}