// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package kconv

import (
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// Convert converts the variable `fromValue` to the type `toTypeName`, the type `toTypeName` is specified by string.
// The optional parameter `extraParams` is used for additional necessary parameter for this conversion.
// It supports common types conversion as its conversion based on type name string.
func Convert(fromValue interface{}, toTypeName string, extraParams ...interface{}) interface{} {
	return doConvert(doConvertInput{
		FromValue:  fromValue,
		ToTypeName: toTypeName,
		ReferValue: nil,
		Extra:      extraParams,
	})
}

type doConvertInput struct {
	FromValue  interface{}   // Value that is converted from.
	ToTypeName string        // Target value type name in string.
	ReferValue interface{}   // Referred value, a value in type `ToTypeName`.
	Extra      []interface{} // Extra values for implementing the converting.
	// Marks that the value is already converted and set to `ReferValue`. Caller can ignore the returned result.
	// It is an attribute for internal usage purpose.
	alreadySetToReferValue bool
}

// doConvert does commonly use types converting.
func doConvert(in doConvertInput) (convertedValue interface{}) {
	switch in.ToTypeName {
	case "int":
		return gconv.Int(in.FromValue)
	case "*int":
		if _, ok := in.FromValue.(*int); ok {
			return in.FromValue
		}
		v := gconv.Int(in.FromValue)
		return &v

	case "int8":
		return gconv.Int8(in.FromValue)
	case "*int8":
		if _, ok := in.FromValue.(*int8); ok {
			return in.FromValue
		}
		v := gconv.Int8(in.FromValue)
		return &v

	case "int16":
		return gconv.Int16(in.FromValue)
	case "*int16":
		if _, ok := in.FromValue.(*int16); ok {
			return in.FromValue
		}
		v := gconv.Int16(in.FromValue)
		return &v

	case "int32":
		return gconv.Int32(in.FromValue)
	case "*int32":
		if _, ok := in.FromValue.(*int32); ok {
			return in.FromValue
		}
		v := gconv.Int32(in.FromValue)
		return &v

	case "int64":
		return gconv.Int64(in.FromValue)
	case "*int64":
		if _, ok := in.FromValue.(*int64); ok {
			return in.FromValue
		}
		v := gconv.Int64(in.FromValue)
		return &v

	case "uint":
		return gconv.Uint(in.FromValue)
	case "*uint":
		if _, ok := in.FromValue.(*uint); ok {
			return in.FromValue
		}
		v := gconv.Uint(in.FromValue)
		return &v

	case "uint8":
		return gconv.Uint8(in.FromValue)
	case "*uint8":
		if _, ok := in.FromValue.(*uint8); ok {
			return in.FromValue
		}
		v := gconv.Uint8(in.FromValue)
		return &v

	case "uint16":
		return gconv.Uint16(in.FromValue)
	case "*uint16":
		if _, ok := in.FromValue.(*uint16); ok {
			return in.FromValue
		}
		v := gconv.Uint16(in.FromValue)
		return &v

	case "uint32":
		return gconv.Uint32(in.FromValue)
	case "*uint32":
		if _, ok := in.FromValue.(*uint32); ok {
			return in.FromValue
		}
		v := gconv.Uint32(in.FromValue)
		return &v

	case "uint64":
		return gconv.Uint64(in.FromValue)
	case "*uint64":
		if _, ok := in.FromValue.(*uint64); ok {
			return in.FromValue
		}
		v := gconv.Uint64(in.FromValue)
		return &v

	case "float32":
		return gconv.Float32(in.FromValue)
	case "*float32":
		if _, ok := in.FromValue.(*float32); ok {
			return in.FromValue
		}
		v := gconv.Float32(in.FromValue)
		return &v

	case "float64":
		return gconv.Float64(in.FromValue)
	case "*float64":
		if _, ok := in.FromValue.(*float64); ok {
			return in.FromValue
		}
		v := gconv.Float64(in.FromValue)
		return &v

	case "bool":
		return gconv.Bool(in.FromValue)
	case "*bool":
		if _, ok := in.FromValue.(*bool); ok {
			return in.FromValue
		}
		v := gconv.Bool(in.FromValue)
		return &v

	case "string":
		return gconv.String(in.FromValue)
	case "*string":
		if _, ok := in.FromValue.(*string); ok {
			return in.FromValue
		}
		v := gconv.String(in.FromValue)
		return &v

	case "[]byte":
		return gconv.Bytes(in.FromValue)
	case "[]int":
		return gconv.Ints(in.FromValue)
	case "[]int32":
		return gconv.Int32s(in.FromValue)
	case "[]int64":
		return gconv.Int64s(in.FromValue)
	case "[]uint":
		return gconv.Uints(in.FromValue)
	case "[]uint8":
		return gconv.Bytes(in.FromValue)
	case "[]uint32":
		return gconv.Uint32s(in.FromValue)
	case "[]uint64":
		return gconv.Uint64s(in.FromValue)
	case "[]float32":
		return gconv.Float32s(in.FromValue)
	case "[]float64":
		return gconv.Float64s(in.FromValue)
	case "[]string":
		return gconv.Strings(in.FromValue)

	case "Time", "time.Time":
		if len(in.Extra) > 0 {
			return gconv.Time(in.FromValue, gconv.String(in.Extra[0]))
		}
		return gconv.Time(in.FromValue)
	case "*time.Time":
		var v interface{}
		if len(in.Extra) > 0 {
			v = gconv.Time(in.FromValue, gconv.String(in.Extra[0]))
		} else {
			if _, ok := in.FromValue.(*time.Time); ok {
				return in.FromValue
			}
			v = gconv.Time(in.FromValue)
		}
		return &v

	case "GTime", "gtime.Time":
		if len(in.Extra) > 0 {
			if v := gconv.GTime(in.FromValue, gconv.String(in.Extra[0])); v != nil {
				return *v
			} else {
				return *gtime.New()
			}
		}
		if v := gconv.GTime(in.FromValue); v != nil {
			return *v
		} else {
			return *gtime.New()
		}
	case "*gtime.Time":
		if len(in.Extra) > 0 {
			if v := gconv.GTime(in.FromValue, gconv.String(in.Extra[0])); v != nil {
				return v
			} else {
				return gtime.New()
			}
		}
		if v := gconv.GTime(in.FromValue); v != nil {
			return v
		} else {
			return gtime.New()
		}

	case "Duration", "time.Duration":
		return gconv.Duration(in.FromValue)
	case "*time.Duration":
		if _, ok := in.FromValue.(*time.Duration); ok {
			return in.FromValue
		}
		v := gconv.Duration(in.FromValue)
		return &v

	case "map[string]string":
		return MapStrStr(in.FromValue)

	case "map[string]interface{}":
		return Map(in.FromValue)

	case "[]map[string]interface{}":
		return gconv.Maps(in.FromValue)

	case "json.RawMessage":
		return gconv.Bytes(in.FromValue)

	default:
		if in.ReferValue != nil {
			var referReflectValue reflect.Value
			if v, ok := in.ReferValue.(reflect.Value); ok {
				referReflectValue = v
			} else {
				referReflectValue = reflect.ValueOf(in.ReferValue)
			}
			defer func() {
				if recover() != nil {
					if err := bindVarToReflectValue(referReflectValue, in.FromValue, nil); err == nil {
						in.alreadySetToReferValue = true
						convertedValue = referReflectValue.Interface()
					}
				}
			}()
			in.ToTypeName = referReflectValue.Kind().String()
			in.ReferValue = nil
			return reflect.ValueOf(doConvert(in)).Convert(referReflectValue.Type()).Interface()
		}
		return in.FromValue
	}
}

func doConvertWithReflectValueSet(reflectValue reflect.Value, in doConvertInput) {
	convertedValue := doConvert(in)
	if !in.alreadySetToReferValue {
		reflectValue.Set(reflect.ValueOf(convertedValue))
	}
}
