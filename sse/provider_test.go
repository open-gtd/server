package sse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRegisterer_ShouldReturnRegisterer(t *testing.T) {
	r := GetRegisterer()

	assert.Equal(t, registerer, r)
}

func TestSetRegisterer_ShouldSetRegisterer(t *testing.T) {
	r := &registererMock{}

	SetRegisterer(r)

	assert.Equal(t, r, registerer)
}
