package projects

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/modules"
	apiProjects "github.com/open-gtd/server/projects/api"
	eventBusProjects "github.com/open-gtd/server/projects/eventBus"
	loggingProjects "github.com/open-gtd/server/projects/logging"
	"github.com/open-gtd/server/sse"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) RegisterHandlers(r api.RestRegisterer) {
	apiProjects.RegisterHandlers(r)
}

func (m module) RegisterBus(busCollection eventBus.BusCollection) {
	//eventBusProjects.RegisterBus(busCollection)
}

func (m module) RegisterBusHandlers(busCollection eventBus.BusCollection) {
	eventBusProjects.RegisterBusHandlers(busCollection)
}

func (m module) RegisterSse(sr sse.SseRegisterer) {
	//sseProjects.RegisterSse(sr)
}

func (m module) RegisterLogger(l logging.Logger) {
	loggingProjects.RegisterLogger(l)
}
