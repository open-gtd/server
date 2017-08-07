package auth

import (
	"github.com/open-gtd/server/api"
	apiAuth "github.com/open-gtd/server/auth/api"
)

func Initialize(r api.Registerer) {
	registerHandlers(r)
}

func registerHandlers(r api.Registerer) {
	apiAuth.RegisterHandlers(r)
}
