package auth

import (
	"github.com/open-gtd/server/api"
	apiAuth "github.com/open-gtd/server/auth/api"
	eventBusAuth "github.com/open-gtd/server/auth/eventBus"
	loggingAuth "github.com/open-gtd/server/auth/logging"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
	"github.com/open-gtd/server/modules"
	"github.com/open-gtd/server/sse"
)

func Module() modules.Module {
	return module{}
}

type module struct{}

func (m module) RegisterHandlers(r api.RestRegisterer) {
	apiAuth.RegisterHandlers(r)
}

func (m module) RegisterBus(busCollection eventBus.BusCollection) {
	eventBusAuth.RegisterBus(busCollection)
}

func (m module) RegisterBusHandlers(busCollection eventBus.BusCollection) {
}

func (m module) RegisterSse(sr sse.SseRegisterer) {
}

func (m module) RegisterLogger(l logging.Logger) {
	loggingAuth.RegisterLogger(l)
}
