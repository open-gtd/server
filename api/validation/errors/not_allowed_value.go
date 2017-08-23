package errors

import (
	"fmt"
	"strings"
)

const specSeparator = ", "

type notAllowedValueError struct {
	message string
}

func (notAllowedValueError) Type() ErrorType {
	return Validation
}

func (e notAllowedValueError) Error() string {
	return e.message
}

//IsNotAllowedValueError returns true, if error is of Not Allowed Value, otherwise false
func IsNotAllowedValueError(err error) bool {
	_, ok := err.(notAllowedValueError)
	return ok
}

//NewNotAllowedValue return Not Allowed Value Validation Error
func NewNotAllowedValue(name string, value string, allowedValues []string) TypedError {
	return notAllowedValueError{
		message: fmt.Sprintf(
			"Value '%v' is not allowed as '%v'. Allowed values are %v",
			value,
			name,
			strings.Join(allowedValues, specSeparator),
		),
	}
}
