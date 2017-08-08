package eventBus

import (
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
)

func Initialize() {
	bus := eventBus.GetBus()

	bus.Subscribe("tagCreated", func(i interface{}) {
		logging.GetLogger().Info("")
	})
}
