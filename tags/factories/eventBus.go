package factories

import (
	"github.com/open-gtd/server/eventBus"
)

var bus = eventBus.GetBus()

func GetBus() eventBus.Bus {
	return bus
}
