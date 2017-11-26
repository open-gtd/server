package errors

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundError(t *testing.T) {

	assert.EqualError(
		t,
		NewNotFound(),
		"not found",
	)
}
