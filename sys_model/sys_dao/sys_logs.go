// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this sys_file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
)

// internalSysLogsDao is internal type for wrapping internal DAO implements.
type internalSysLogsDao = *internal.SysLogsDao

// sysLogsDao is the data access object for table sys_logs.
// You can define custom methods on it to extend its functionality as you wish.
type sysLogsDao struct {
	internalSysLogsDao
}

var (
	// SysLogs is globally public accessible object for table sys_logs operations.
	SysLogs = sysLogsDao{
		internal.NewSysLogsDao(),
	}
)
