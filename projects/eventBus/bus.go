package eventBus

import (
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
)

var getBus = eventBus.GetBus

func Initialize() {
	bus := getBus()

	bus.Subscribe("tagCreated", func(i interface{}) {
		logging.GetLogger().Info("")
	})
}
