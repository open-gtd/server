package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type someErr struct {
}

func (someErr) Error() string   { return "" }
func (someErr) Type() ErrorType { return Validation }

//NewNotAllowedValue return Not Allowed Value Validation Error
func newSomeErr() TypedError {
	return someErr{}
}

func TestIsValidationError_ShouldReturnTrue_IfErrorIsOfTypeValidationError(t *testing.T) {
	assert.True(
		t,
		IsValidationError(newSomeErr()),
		"someErr should be recognized as Validation Error!")
}

func TestIsValidationError_ShouldReturnFalse_IfErrorIsNotOfTypeValidationError(t *testing.T) {
	assert.False(
		t,
		IsValidationError(errors.New("x")),
		"someErr should not be recognized as Validation Error!")
}
