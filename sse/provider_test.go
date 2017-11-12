package sse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRegisterer_ShouldReturnRegisterer(t *testing.T) {
	testGetRegisterer(t, &registererMock{})
	testGetRegisterer(t, &NullRegisterer{})
}

func testGetRegisterer(t *testing.T, r Registerer) {
	registerer = r
	result := GetRegisterer()
	assert.Equal(t, r, result)
}

func TestSetRegisterer_ShouldSetRegisterer(t *testing.T) {
	testSetRegisterer(t, &registererMock{})
	testSetRegisterer(t, &NullRegisterer{})
}
func testSetRegisterer(t *testing.T, r Registerer) {
	SetRegisterer(r)
	assert.Equal(t, r, registerer)
}
