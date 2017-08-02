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

func (m module) Register(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	busCollection eventBus.BusCollection,
	logger logging.Logger) {

	apiTasks.RegisterHandlers(apiRegisterer)
	eventBusTasks.RegisterBus(busCollection)
	//eventBusTasks.RegisterBusHandlers(busCollection)
	//sseTasks.RegisterSse(sseRegisterer)
	loggingTasks.RegisterLogger(logger)
}
