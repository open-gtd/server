package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRegisterer(t *testing.T) {
	checkSetRegisterer(t, &registererMock{})
	checkSetRegisterer(t, NullRegisterer{})
}

func TestGetRegisterer(t *testing.T) {
	checkGetRegisterer(t, &registererMock{})
	checkGetRegisterer(t, NullRegisterer{})
}

func checkSetRegisterer(t *testing.T, registerer Registerer) {
	SetRegisterer(registerer)
	assert.Equal(t, registerer, r)
}

func checkGetRegisterer(t *testing.T, registerer Registerer) {
	r = registerer
	assert.Equal(t, registerer, GetRegisterer())
}
