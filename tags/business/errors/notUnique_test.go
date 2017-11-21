package errors

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNotUniqueError(t *testing.T) {

	assert.EqualError(
		t,
		NewNotFound(),
		"not unique",
	)
}