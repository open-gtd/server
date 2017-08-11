package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/open-gtd/server/api/validation/errors"
)

func TestAllowedValue_ShouldNotReturnError_IfValueIsOnAllowedValuesList(t *testing.T) {
	assert.Nil(
		t,
		AllowedValue(
			"someName",
			"allowedValue",
			[]string{"allowedValue", "otherValue"},
		),
	)
}

func TestAllowedValue_ShouldReturnError_IfValueIsNotOnAllowedValuesList(t *testing.T) {
	assert.IsType(
		t,
		errors.NewNotAllowedValue("", "", []string{}),
		AllowedValue(
			"someName",
			"notAllowedValue",
			[]string{"allowedValue", "otherValue"},
		),
	)
}
