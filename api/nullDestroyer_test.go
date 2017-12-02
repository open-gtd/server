package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDestroy_ShouldReturnNil(t *testing.T) {
 	assert.Nil(t, NullDestroyer{}.Destroy())
}
