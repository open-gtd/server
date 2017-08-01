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

func (m module) RegisterHandlers(r api.RestRegisterer) {
	apiReferenceLists.RegisterHandlers(r)
}

func (m module) RegisterBus(busCollection eventBus.BusCollection) {
	eventBusReferenceLists.RegisterBus(busCollection)
}

func (m module) RegisterBusHandlers(busCollection eventBus.BusCollection) {
	//eventBusReferenceLists.RegisterBusHandlers(busCollection)
}

func (m module) RegisterSse(sr sse.SseRegisterer) {
	//sseReferenceLists.RegisterSse(sr)
}

func (m module) RegisterLogger(l logging.Logger) {
	loggingReferenceLists.RegisterLogger(l)
}
