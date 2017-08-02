package tags

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/modules"
	"github.com/open-gtd/server/sse"
	apiTags "github.com/open-gtd/server/tags/api"
	eventBusTags "github.com/open-gtd/server/tags/eventBus"
	loggingTags "github.com/open-gtd/server/tags/logging"
	sseTags "github.com/open-gtd/server/tags/sse"
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

	apiTags.RegisterHandlers(apiRegisterer)
	eventBusTags.RegisterBus(busCollection)
	//eventBusTags.RegisterBusHandlers(busCollection)
	sseTags.RegisterSse(sseRegisterer)
	loggingTags.RegisterLogger(logger)
}
