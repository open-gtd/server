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

func (m module) RegisterHandlers(r api.RestRegisterer) {
	apiTags.RegisterHandlers(r)
}

func (m module) RegisterBus(busCollection eventBus.BusCollection) {
	eventBusTags.RegisterBus(busCollection)
}

func (m module) RegisterBusHandlers(busCollection eventBus.BusCollection) {
	//eventBusTags.RegisterBusHandlers(busCollection)
}

func (m module) RegisterSse(sr sse.SseRegisterer) {
	sseTags.RegisterSse(sr)
}

func (m module) RegisterLogger(l logging.Logger) {
	loggingTags.RegisterLogger(l)
}
