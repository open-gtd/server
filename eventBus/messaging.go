package eventBus

import (
	"github.com/asaskevich/EventBus"
)

type Bus interface {
	Subscribe(topic string, fn interface{}) error
	Publish(topic string, args ...interface{})
}
type Name string

type BusCollection interface {
	New(name Name) Bus
	Get(name Name) Bus
}

type busCollection struct {
	m map[Name]Bus
}

func New() BusCollection {
	return &busCollection{
		m: map[Name]Bus{},
	}
}

func (bc *busCollection) New(name Name) Bus {
	bus := EventBus.New()
	bc.m[name] = bus

	return bus
}

func (bc busCollection) Get(name Name) Bus {
	bus, _ := bc.m[name]
	return bus
}
