package modules

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/sse"
)

type Module interface {
	Register(api.Registerer, sse.Registerer, eventBus.BusCollection, logging.Logger)
}

type ModuleManager interface {
	Register(Module)
}

func NewModuleManager(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	busCollection eventBus.BusCollection,
	logger logging.Logger) ModuleManager {
	return &moduleManager{
		apiRegisterer: apiRegisterer,
		sseRegisterer: sseRegisterer,
		busCollection: busCollection,
		logger:        logger,
	}
}

type moduleManager struct {
	apiRegisterer api.Registerer
	sseRegisterer sse.Registerer
	busCollection eventBus.BusCollection
	logger        logging.Logger
}

func (m *moduleManager) Register(module Module) {
	module.Register(
		m.apiRegisterer,
		m.sseRegisterer,
		m.busCollection,
		m.logger,
	)
}
