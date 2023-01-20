package daoctl

import (
	"github.com/gogf/gf/v2/database/gdb"
)

func Save(model *gdb.Model, data ...interface{}) (rowsAffected int64) {
	result, err := model.Save(data...)

	if err != nil {
		return 0
	}

	rowsAffected, err = result.RowsAffected()

	return rowsAffected
}

func SaveWithError(model *gdb.Model, data ...interface{}) (rowsAffected int64, err error) {
	result, err := model.Save(data...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
