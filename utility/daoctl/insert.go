package daoctl

import "github.com/gogf/gf/v2/database/gdb"

func Insert(model *gdb.Model, data ...interface{}) (rowsAffected int64) {
	result, err := model.Insert(data...)

	if err != nil {
		return 0
	}

	rowsAffected, err = result.RowsAffected()

	return rowsAffected
}

func InsertWithError(model *gdb.Model, data ...interface{}) (rowsAffected int64, err error) {
	result, err := model.Insert(data...)

	if err != nil {
		return 0, err
	}

	rowsAffected, err1 := result.RowsAffected()

	if err == nil && err1 != nil {
		err = err1
	}

	return rowsAffected, err
}

func InsertIgnore(model *gdb.Model, data ...interface{}) (rowsAffected int64) {
	result, err := model.InsertIgnore(data...)

	if err != nil {
		return 0
	}

	rowsAffected, err = result.RowsAffected()

	return rowsAffected
}

func InsertIgnoreWithError(model *gdb.Model, data ...interface{}) (rowsAffected int64, err error) {
	result, err := model.Insert(data...)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
