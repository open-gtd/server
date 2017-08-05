package tags

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/config"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/modules"
	"github.com/open-gtd/server/sse"
	apiTags "github.com/open-gtd/server/tags/api"
	configTags "github.com/open-gtd/server/tags/config"
	eventBusTags "github.com/open-gtd/server/tags/eventBus"
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
	config config.Reader) {

	apiTags.RegisterHandlers(apiRegisterer)
	eventBusTags.RegisterBus(busCollection)
	//eventBusTags.RegisterBusHandlers(busCollection)
	sseTags.RegisterSse(sseRegisterer)
	configTags.RegisterReader(config)
}
