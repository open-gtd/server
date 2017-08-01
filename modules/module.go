package modules

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/sse"
)

type Module interface {
	RegisterHandlers(api.RestRegisterer)
	RegisterBus(busCollection eventBus.BusCollection)
	RegisterBusHandlers(busCollection eventBus.BusCollection)
	RegisterSse(sr sse.SseRegisterer)
	RegisterLogger(l logging.Logger)
}
