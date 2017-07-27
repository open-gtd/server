package eventBus

import "github.com/open-gtd/server/eventBus"

var bus eventBus.Bus

func RegisterBus(busCollection eventBus.BusCollection) {
	bus = busCollection.New("tasks")
}

func Get() eventBus.Bus {
	return bus
}
