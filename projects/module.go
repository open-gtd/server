package projects

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/modules"
	apiProjects "github.com/open-gtd/server/projects/api"
	eventBusProjects "github.com/open-gtd/server/projects/eventBus"
	"github.com/open-gtd/server/sse"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) Register(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	busCollection eventBus.BusCollection,
	reader config.Reader) {
	apiProjects.RegisterHandlers(apiRegisterer)
	//eventBusProjects.RegisterBus(busCollection)
	eventBusProjects.RegisterBusHandlers(busCollection)
	//sseProjects.RegisterSse(sseRegisterer)
}
