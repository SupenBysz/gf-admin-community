package sys_enum

import (
	"fmt"
)

// Code is universal code interface definition.
type Code interface {
	// Code returns the integer number of current code.
	Code() int

	// Description returns the brief description for current code.
	Description() string

	// String returns current code as a string.
	String() string
}

// EnumCode is an implementer for interface Code for internal usage only.
type EnumCode struct {
	code        int    // Error code, usually an integer.
	description string // Brief description for this code.
}

// Code returns the integer number of current code.
func (c EnumCode) Code() int {
	return c.code
}

// Description returns the brief description for current code.
func (c EnumCode) Description() string {
	return c.description
}

// String returns current code as a string.
func (c EnumCode) String() string {
	if c.description != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.description)
	}
	return fmt.Sprintf(`%d`, c.code)
}

func New(code int, description string) Code {
	result := EnumCode{
		code:        code,
		description: description,
	}
	return (Code)(&result)
}
