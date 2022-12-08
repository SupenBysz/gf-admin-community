// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/SupenBysz/gf-admin-community/model/dao/internal"
)

// internalFdInvoiceDetailDao is internal type for wrapping internal DAO implements.
type internalFdInvoiceDetailDao = *internal.FdInvoiceDetailDao

// fdInvoiceDetailDao is the data access object for table fd_invoice_detail.
// You can define custom methods on it to extend its functionality as you wish.
type fdInvoiceDetailDao struct {
	internalFdInvoiceDetailDao
}

var (
	// FdInvoiceDetail is globally public accessible object for table fd_invoice_detail operations.
	FdInvoiceDetail = fdInvoiceDetailDao{
		internal.NewFdInvoiceDetailDao(),
	}
)

// Fill with you ideas below.