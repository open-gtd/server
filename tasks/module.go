package tasks

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/modules"
	"github.com/open-gtd/server/sse"
	apiTasks "github.com/open-gtd/server/tasks/api"
	eventBusTasks "github.com/open-gtd/server/tasks/eventBus"
	//sseTasks "github.com/open-gtd/server/tasks/sse"
	"github.com/open-gtd/server/logging"
	loggingTasks "github.com/open-gtd/server/tasks/logging"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) RegisterHandlers(r api.RestRegisterer) {
	apiTasks.RegisterHandlers(r)
}

func (m module) RegisterBus(busCollection eventBus.BusCollection) {
	eventBusTasks.RegisterBus(busCollection)
}

func (m module) RegisterBusHandlers(busCollection eventBus.BusCollection) {
	//eventBusTasks.RegisterBusHandlers(busCollection)
}

func (m module) RegisterSse(sr sse.SseRegisterer) {
	//sseTasks.RegisterSse(sr)
}

func (m module) RegisterLogger(l logging.Logger) {
	loggingTasks.RegisterLogger(l)
}
