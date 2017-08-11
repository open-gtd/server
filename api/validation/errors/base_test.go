package errors

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

type someErr struct {

}

func (someErr) Error() string { return "" }

//NewNotAllowedValue return Not Allowed Value Validation Error
func newSomeErr() ValidationError {
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