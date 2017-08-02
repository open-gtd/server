package eBus

import (
	"github.com/open-gtd/server/eventBus"
)

func NewCollection() eventBus.BusCollection {
	return &busCollection{
		m: map[eventBus.Name]eventBus.Bus{},
	}
}

type busCollection struct {
	m map[eventBus.Name]eventBus.Bus
}

func (bc *busCollection) New(name eventBus.Name) eventBus.Bus {
	bus := newBus()
	bc.m[name] = bus

	return bus
}

func (bc busCollection) Get(name eventBus.Name) eventBus.Bus {
	bus, _ := bc.m[name]
	return bus
}
