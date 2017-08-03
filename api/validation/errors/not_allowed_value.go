package errors

import (
	"fmt"
	"strings"
)

const specSeparator = ", "

type notAllowedValueError struct {
	message string
}

//IsNotAllowedValueError - test to check if error is Not Allowed Value
func IsNotAllowedValueError(err error) bool {
	_, ok := err.(notAllowedValueError)
	return ok
}

func (e notAllowedValueError) Error() string {
	return e.message
}

//NewNotAllowedValue - produces new NotAllowed Value validation error
func NewNotAllowedValue(name string, value string, allowedValues []string) ValidationError {
	return notAllowedValueError{
		message: fmt.Sprintf(
			"Value '%v' is not allowed as '%v'. Allowed values are %v",
			value,
			name,
			strings.Join(allowedValues, specSeparator),
		),
	}
}
