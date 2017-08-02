package referenceLists

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/modules"
	apiReferenceLists "github.com/open-gtd/server/referenceLists/api"
	eventBusReferenceLists "github.com/open-gtd/server/referenceLists/eventBus"
	"github.com/open-gtd/server/sse"
	//sseReferenceLists "github.com/open-gtd/server/referenceLists/sse"
	"github.com/open-gtd/server/logging"
	loggingReferenceLists "github.com/open-gtd/server/referenceLists/logging"
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

	apiReferenceLists.RegisterHandlers(apiRegisterer)
	eventBusReferenceLists.RegisterBus(busCollection)
	//eventBusReferenceLists.RegisterBusHandlers(busCollection)
	//sseReferenceLists.RegisterSse(sseRegisterer)
	loggingReferenceLists.RegisterLogger(logger)
}
