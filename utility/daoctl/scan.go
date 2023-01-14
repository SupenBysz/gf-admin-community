package daoctl

import "github.com/gogf/gf/v2/database/gdb"

func Scan[T any](model *gdb.Model) *T {
	result := new(T)

	if err := model.Scan(result); err != nil {
		return nil
	}
	return result
}

func ScanWithError[T any](model *gdb.Model) (*T, error) {
	result := new(T)

	if err := model.Scan(result); err != nil {
		return nil, err
	}
	return result, nil
}

func ScanList[T any](model *gdb.Model, bindToAttrName string, relationAttrNameAndFields ...string) *T {
	result := new(T)
	if err := model.ScanList(result, bindToAttrName, relationAttrNameAndFields...); err != nil {
		return nil
	}
	return result
}

func ScanListWithError[T any](model *gdb.Model, bindToAttrName string, relationAttrNameAndFields ...string) (*T, error) {
	result := new(T)

	if err := model.ScanList(result, bindToAttrName, relationAttrNameAndFields...); err != nil {
		return nil, err
	}
	return result, nil
}
