package daoctl

import (
	"github.com/gogf/gf/v2/database/gdb"
)

func Save[T any](model *gdb.Model, data ...interface{}) (lastInsertId int64, rowsAffected int64) {
	result, err := model.Save(data)

	if err != nil {
		return 0, 0
	}

	rowsAffected, err = result.RowsAffected()
	lastInsertId, err = result.LastInsertId()

	return lastInsertId, rowsAffected
}

func SaveWithError(model *gdb.Model, data ...interface{}) (lastInsertId int64, rowsAffected int64, err error) {
	result, err := model.Save(data)

	if err != nil {
		return 0, 0, err
	}

	rowsAffected, err1 := result.RowsAffected()
	lastInsertId, err2 := result.LastInsertId()

	if err == nil && err1 != nil {
		err = err1
	}
	if err == nil && err2 != nil {
		err = err2
	}

	return lastInsertId, rowsAffected, err
}
