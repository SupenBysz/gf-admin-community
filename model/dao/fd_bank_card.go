// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/SupenBysz/gf-admin-community/model/dao/internal"
)

// internalFdBankCardDao is internal type for wrapping internal DAO implements.
type internalFdBankCardDao = *internal.FdBankCardDao

// fdBankCardDao is the data access object for table fd_bank_card.
// You can define custom methods on it to extend its functionality as you wish.
type fdBankCardDao struct {
	internalFdBankCardDao
}

var (
	// FdBankCard is globally public accessible object for table fd_bank_card operations.
	FdBankCard = fdBankCardDao{
		internal.NewFdBankCardDao(),
	}
)

// Fill with you ideas below.