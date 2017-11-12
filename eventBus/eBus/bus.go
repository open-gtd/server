package eBus

import (
	"github.com/asaskevich/EventBus"
	"github.com/open-gtd/server/eventBus"
)

var innerBus = EventBus.New()

func NewBus() eventBus.Bus {
	return &bus{
		eBus: innerBus,
	}
}

type bus struct {
	eBus EventBus.Bus
}

func (bus *bus) Subscribe(topic eventBus.Topic, fn eventBus.BusHandler) error {
	return bus.eBus.Subscribe(string(topic), fn)
}

func (bus *bus) Unsubscribe(topic eventBus.Topic, handler eventBus.BusHandler) error {
	return bus.eBus.Unsubscribe(string(topic), handler)
}

func (bus *bus) Publish(topic eventBus.Topic, arg interface{}) {
	bus.eBus.Publish(string(topic), arg)
}
