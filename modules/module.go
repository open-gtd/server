package modules

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/sse"
)

type Module interface {
	Register(api.Registerer, sse.Registerer, config.Reader)
}

type ModuleManager interface {
	Register(Module)
}

func NewModuleManager(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	reader config.Reader) ModuleManager {
	return &moduleManager{
		apiRegisterer: apiRegisterer,
		sseRegisterer: sseRegisterer,
		reader:        reader,
	}
}

type moduleManager struct {
	apiRegisterer api.Registerer
	sseRegisterer sse.Registerer
	reader        config.Reader
}

func (m *moduleManager) Register(module Module) {
	module.Register(
		m.apiRegisterer,
		m.sseRegisterer,
		m.reader,
	)
}
