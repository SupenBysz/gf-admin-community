package enum

type IEnumCodeInt interface {
	Code() int
	// Description returns the brief description for current code.
	Description() string
}
type IEnumCodeStr interface {
	Code() string
	// Description returns the brief description for current code.
	Description() string
}

type IEnumCode[TCode int | int64 | string] interface {
	Code() TCode
	// Description returns the brief description for current code.
	Description() string
}

type IEnumCodeWithData[TCode int | int64 | string, TData any] interface {
	Code() TCode
	Data() TData
	Description() string
}

// EnumType [T any] is an implementer for interface Code for internal usage only.
type enumType[TCode int | int64 | string, TData any] struct {
	code        TCode  // Error code, usually an integer.
	data        TData  // Brief data for this value.
	description string // Brief description for this code.
}

// Code returns the integer number of current code.
func (e *enumType[TCode, TData]) Code() TCode {
	return e.code
}

// Description returns the brief description for current code.
func (e *enumType[TCode, TData]) Description() string {
	return e.description
}

// Data returns the T data of current code.
func (e *enumType[TCode, TData]) Data() TData {
	return e.data
}

func New[R IEnumCode[TCode], TCode int | int64 | string](code TCode, description string) R {
	var result interface{}
	result = &enumType[TCode, interface{}]{
		code:        code,
		description: description,
	}
	return result.(R)
}

func NewWithData[TCode int | int64 | string, TData any](code TCode, data TData, description string) *IEnumCodeWithData[TCode, TData] {
	var result interface{}
	result = &enumType[TCode, TData]{
		code:        code,
		data:        data,
		description: description,
	}
	return result.(*IEnumCodeWithData[TCode, TData])
}
