// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package sys_dao

import (
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
	"{TplImportPrefix}/internal"
)

type {TplTableNameCamelCase}Dao = dao_interface.TIDao[internal.{TplTableNameCamelCase}Columns]

func New{TplTableNameCamelCase}(dao ...dao_interface.IDao) {TplTableNameCamelCase}Dao {
	return ({TplTableNameCamelCase}Dao)(internal.New{TplTableNameCamelCase}Dao(dao...))
}

var (
    {TplTableNameCamelCase} = New{TplTableNameCamelCase}()
)