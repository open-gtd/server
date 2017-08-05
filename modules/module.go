package modules

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/sse"
)

type Module interface {
	Register(api.Registerer, sse.Registerer, eventBus.BusCollection, config.Reader)
}

type ModuleManager interface {
	Register(Module)
}

func NewModuleManager(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	busCollection eventBus.BusCollection,
	reader config.Reader) ModuleManager {
	return &moduleManager{
		apiRegisterer: apiRegisterer,
		sseRegisterer: sseRegisterer,
		busCollection: busCollection,
		reader:        reader,
	}
}

type moduleManager struct {
	apiRegisterer api.Registerer
	sseRegisterer sse.Registerer
	busCollection eventBus.BusCollection
	logger        logging.Logger
	reader        config.Reader
}

func (m *moduleManager) Register(module Module) {
	module.Register(
		m.apiRegisterer,
		m.sseRegisterer,
		m.busCollection,
		m.reader,
	)
}
