package factories

import (
	"testing"

	"github.com/open-gtd/server/eventBus"
	"github.com/stretchr/testify/assert"
)

func TestGetBus(t *testing.T) {

	b := GetBus()

	var bus eventBus.Bus
	assert.Implements(t, &bus, b)
}
