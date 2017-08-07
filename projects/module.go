package projects

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/modules"
	apiProjects "github.com/open-gtd/server/projects/api"
	"github.com/open-gtd/server/sse"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) Register(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	reader config.Reader) {
	apiProjects.RegisterHandlers(apiRegisterer)
	//eventBusProjects.RegisterBus(busCollection)
	//sseProjects.RegisterSse(sseRegisterer)
}
