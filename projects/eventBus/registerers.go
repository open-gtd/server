package eventBus

import (
	"fmt"
	"os"

	"github.com/open-gtd/server/eventBus"
)

var bus eventBus.Bus

func RegisterBus(busCollection eventBus.BusCollection) {
	bus = busCollection.New("projects")
}

func RegisterBusHandlers(busCollection eventBus.BusCollection) {
	busCollection.Get("tags").Subscribe("tagCreated", func(i interface{}) {
		fmt.Fprintln(os.Stdout, i)
	})
}

func Get() eventBus.Bus {
	return bus
}
