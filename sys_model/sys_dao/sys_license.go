// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao/internal"
)

// internalSysLicenseDao is internal type for wrapping internal DAO implements.
type internalSysLicenseDao = *internal.SysLicenseDao

// sysLicenseDao is the data access object for table sys_license.
// You can define custom methods on it to extend its functionality as you wish.
type sysLicenseDao struct {
	internalSysLicenseDao
}

var (
	// SysLicense is globally public accessible object for table sys_license operations.
	SysLicense = sysLicenseDao{
		internal.NewSysLicenseDao(),
	}
)