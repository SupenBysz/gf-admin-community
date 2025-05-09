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

// SysPersonLicenseDao is the data access object for table sys_person_license.
type SysPersonLicenseDao struct {
	dao_interface.IDao
	table       string                  // table is the underlying table name of the DAO.
	group       string                  // group is the database configuration group name of current DAO.
	columns     SysPersonLicenseColumns // columns contains all the column names of Table for convenient usage.
	daoConfig   *dao_interface.DaoConfig
	ignoreCache bool
	exWhereArr  []string
}

// SysPersonLicenseColumns defines and stores column names for table sys_person_license.
type SysPersonLicenseColumns struct {
	Id               string // ID
	IdcardFrontPath  string // 身份证头像面照片
	IdcardBackPath   string // 身份证国徽面照片
	No               string // 身份证号
	Gender           string // 性别
	Nation           string // 名族
	Name             string // 姓名
	Birthday         string // 出生日期
	Address          string // 家庭住址
	IssuingAuthorit  string // 签发机关
	IssuingDate      string // 签发日期
	ExpriyDate       string //
	CreatedAt        string //
	UpdatedAt        string //
	DeletedAt        string //
	State            string // 状态：0失效、1正常
	AuthType         string // 认证类型:
	Remark           string // 备注信息
	LatestAuditLogId string // 最新的审核记录id
	UserId           string // 关联的用户ID
}

// sysPersonLicenseColumns holds the columns for table sys_person_license.
var sysPersonLicenseColumns = SysPersonLicenseColumns{
	Id:               "id",
	IdcardFrontPath:  "idcard_front_path",
	IdcardBackPath:   "idcard_back_path",
	No:               "no",
	Gender:           "gender",
	Nation:           "nation",
	Name:             "name",
	Birthday:         "birthday",
	Address:          "address",
	IssuingAuthorit:  "issuing_authorit",
	IssuingDate:      "issuing_date",
	ExpriyDate:       "expriy_date",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	State:            "state",
	AuthType:         "auth_type",
	Remark:           "remark",
	LatestAuditLogId: "latest_audit_logId",
	UserId:           "user_id",
}

// NewSysPersonLicenseDao creates and returns a new DAO object for table data access.
func NewSysPersonLicenseDao(proxy ...dao_interface.IDao) *SysPersonLicenseDao {
	var dao *SysPersonLicenseDao
	if len(proxy) > 0 {
		dao = &SysPersonLicenseDao{
			group:       proxy[0].Group(),
			table:       proxy[0].Table(),
			columns:     sysPersonLicenseColumns,
			daoConfig:   proxy[0].DaoConfig(context.Background()),
			IDao:        proxy[0].DaoConfig(context.Background()).Dao,
			ignoreCache: proxy[0].DaoConfig(context.Background()).IsIgnoreCache(),
			exWhereArr:  proxy[0].DaoConfig(context.Background()).Dao.GetExtWhereKeys(),
		}

		return dao
	}

	return &SysPersonLicenseDao{
		group:   "default",
		table:   "sys_person_license",
		columns: sysPersonLicenseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysPersonLicenseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysPersonLicenseDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *SysPersonLicenseDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *SysPersonLicenseDao) Columns() SysPersonLicenseColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysPersonLicenseDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *SysPersonLicenseDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) *dao_interface.DaoConfig {
	//if dao.daoConfig != nil && len(dao.exWhereArr) == 0 {
	//	return dao.daoConfig
	//}

	var daoConfig = daoctl.NewDaoConfig(ctx, dao, cacheOption...)
	dao.daoConfig = &daoConfig

	if len(dao.exWhereArr) > 0 {
		daoConfig.IgnoreExtModel(dao.exWhereArr...)
		dao.exWhereArr = []string{}

	}

	if dao.ignoreCache {
		daoConfig.IgnoreCache()
	}

	return dao.daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysPersonLicenseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

func (dao *SysPersonLicenseDao) GetExtWhereKeys() []string {
	return dao.exWhereArr
}

func (dao *SysPersonLicenseDao) IsIgnoreCache() bool {
	return dao.ignoreCache
}

func (dao *SysPersonLicenseDao) IgnoreCache() dao_interface.IDao {
	dao.ignoreCache = true
	return dao
}
func (dao *SysPersonLicenseDao) IgnoreExtModel(whereKey ...string) dao_interface.IDao {
	dao.exWhereArr = append(dao.exWhereArr, whereKey...)
	return dao
}
