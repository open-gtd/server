package eventBus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetBus(t *testing.T) {
	testSetBus(t, &busMock{})
	testSetBus(t, &NullBus{})
}

func TestGetBus(t *testing.T) {
	testGetBus(&busMock{}, t)
	testGetBus(&NullBus{}, t)
}

func testGetBus(bus Bus, t *testing.T) {
	b = bus
	result := GetBus()
	assert.Equal(t, bus, result)
}

func testSetBus(t *testing.T, bus Bus) {
	SetBus(bus)
	assert.Equal(t, bus, b)
}
