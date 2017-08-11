package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotAllowedValue(t *testing.T) {
	e := NewNotAllowedValue("someName", "123", []string{"zxc", "qwe"})
	assert.Equal(t, e.Error(), "Value '123' is not allowed as 'someName'. Allowed values are zxc, qwe")

	e = NewNotAllowedValue("someName2", "153", []string{"281", "dsa"})
	assert.Equal(t, e.Error(), "Value '153' is not allowed as 'someName2'. Allowed values are 281, dsa")
}

func TestIsNotAllowedValueError_ShouldReturnTrue_IfErrorIsOfTypeNotAllowedValueError(t *testing.T) {
	assert.True(t, IsNotAllowedValueError(NewNotAllowedValue("", "", []string{})))
}

func TestIsNotAllowedValueError_ShouldReturnFalse_IfErrorIsNotOfTypeNotAllowedValueError(t *testing.T) {
	assert.False(t, IsNotAllowedValueError(errors.New("")))
}
