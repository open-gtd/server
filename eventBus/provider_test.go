package eventBus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBus(t *testing.T) {
	checkSetBus(t, &busMock{})
	checkSetBus(t, NullBus{})
}

func TestGetBus(t *testing.T) {
	checkGetBus(t, &busMock{})
	checkGetBus(t, NullBus{})
}

func checkSetBus(t *testing.T, bus Bus) {
	SetBus(bus)

	assert.Equal(t, bus, b)
}

func checkGetBus(t *testing.T, bus Bus) {
	b = bus

	assert.Equal(t, bus, GetBus())
}
