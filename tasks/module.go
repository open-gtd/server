package tasks

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/modules"
	"github.com/open-gtd/server/sse"
	apiTasks "github.com/open-gtd/server/tasks/api"
	//sseTasks "github.com/open-gtd/server/tasks/sse"
	"github.com/open-gtd/server/config"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) Register(
	apiRegisterer api.Registerer,
	sseRegisterer sse.Registerer,
	reader config.Reader) {

	apiTasks.RegisterHandlers(apiRegisterer)
	//eventBusTasks.RegisterBusHandlers(busCollection)
	//sseTasks.RegisterSse(sseRegisterer)
}
