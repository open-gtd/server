package auth

import (
	"github.com/open-gtd/server/api"
	apiAuth "github.com/open-gtd/server/auth/api"
	eventBusAuth "github.com/open-gtd/server/auth/eventBus"
	loggingAuth "github.com/open-gtd/server/auth/logging"
	"github.com/open-gtd/server/eventBus"
	"github.com/open-gtd/server/logging"
)

func Initialize(r api.Registerer, busCollection eventBus.BusCollection, l logging.Logger) {
	registerHandlers(r)
	registerBus(busCollection)
	registerLogger(l)
}

func registerHandlers(r api.Registerer) {
	apiAuth.RegisterHandlers(r)
}

func registerBus(busCollection eventBus.BusCollection) {
	eventBusAuth.RegisterBus(busCollection)
}

func registerLogger(l logging.Logger) {
	loggingAuth.RegisterLogger(l)
}
