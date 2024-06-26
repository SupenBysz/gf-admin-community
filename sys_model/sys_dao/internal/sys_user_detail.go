// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDetailDao is the data access object for table sys_user_detail.
type SysUserDetailDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysUserDetailColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserDetailColumns defines and stores column names for table sys_user_detail.
type SysUserDetailColumns struct {
	Id            string // ID，保持与USERID一致
	Realname      string // 姓名
	UnionMainName string // 关联主体名称
	LastLoginIp   string // 最后登录IP
	LastLoginArea string // 最后登录地区
	LastLoginAt   string // 最后登录时间
}

// sysUserDetailColumns holds the columns for table sys_user_detail.
var sysUserDetailColumns = SysUserDetailColumns{
	Id:            "id",
	Realname:      "realname",
	UnionMainName: "union_main_name",
	LastLoginIp:   "last_login_ip",
	LastLoginArea: "last_login_area",
	LastLoginAt:   "last_login_at",
}

// NewSysUserDetailDao creates and returns a new DAO object for table data access.
func NewSysUserDetailDao(proxy ...dao_interface.IDao) *SysUserDetailDao {
	var dao *SysUserDetailDao
	if len(proxy) > 0 {
		dao = &SysUserDetailDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: sysUserDetailColumns,
		}
		return dao
	}

	return &SysUserDetailDao{
		group:   "default",
		table:   "sys_user_detail",
		columns: sysUserDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserDetailDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserDetailDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysUserDetailDao) Columns() SysUserDetailColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserDetailDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysUserDetailDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
