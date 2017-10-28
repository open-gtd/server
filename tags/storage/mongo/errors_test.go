package mongo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const notFoundErrorMessage = "not found"
const duplicateKeyErrorMessage = "duplicate key error index"
const otherMessage = "other error message"

func TestIsNotFoundShouldReturnTrueIfErrorHasNotFoundMessage(t *testing.T) {
	assert.True(t, isNotFound(errors.New(notFoundErrorMessage)))
}

func TestIsNotFoundShouldReturnFalseIfErrorHasOtherMessage(t *testing.T) {
	assert.False(t, isNotFound(errors.New(otherMessage)))
}

func TestIsDuplicateKeyShouldReturnTrueIfErrorHasDuplicateKeyMessage(t *testing.T) {
	assert.True(t, isDuplicateKey(errors.New(duplicateKeyErrorMessage)))
}

func TestIsDuplicateKeyShouldReturnFalseIfErrorHasOtherMessage(t *testing.T) {
	assert.False(t, isNotFound(errors.New(duplicateKeyErrorMessage)))
}
