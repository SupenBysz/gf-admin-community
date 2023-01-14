package daoctl

import "github.com/gogf/gf/v2/database/gdb"

func Delete(model *gdb.Model, dataAndWhere ...interface{}) (rowsAffected int64) {
	result, err := model.Delete(dataAndWhere)

	if err != nil {
		return 0
	}

	rowsAffected, err = result.RowsAffected()

	return rowsAffected
}

func DeleteWithError(model *gdb.Model, dataAndWhere ...interface{}) (rowsAffected int64, err error) {
	result, err := model.Delete(dataAndWhere)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
