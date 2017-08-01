package eventBus

import "github.com/open-gtd/server/eventBus"

var bus eventBus.Bus

func RegisterBus(busCollection eventBus.BusCollection) {
	bus = busCollection.New("referenceLists")
}

//func RegisterBusHandlers(busCollection eventBus.BusCollection) {
//}

func Get() eventBus.Bus {
	return bus
}
