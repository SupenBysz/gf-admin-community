// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPermissionDao is the data access object for table sys_permission.
type SysPermissionDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysPermissionColumns // columns contains all the column names of Table for convenient usage.
}

// SysPermissionColumns defines and stores column names for table sys_permission.
type SysPermissionColumns struct {
	Id          string // ID
	ParentId    string // 父级ID
	Name        string // 名称
	Description string // 描述
	Identifier  string // 标识符
	Type        string // 类型：1api，2menu
	CreatedAt   string //
	UpdatedAt   string //
}

// sysPermissionColumns holds the columns for table sys_permission.
var sysPermissionColumns = SysPermissionColumns{
	Id:          "id",
	ParentId:    "parent_id",
	Name:        "name",
	Description: "description",
	Identifier:  "identifier",
	Type:        "type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSysPermissionDao creates and returns a new DAO object for table data access.
func NewSysPermissionDao() *SysPermissionDao {
	return &SysPermissionDao{
		group:   "default",
		table:   "sys_permission",
		columns: sysPermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysPermissionDao) Columns() SysPermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}