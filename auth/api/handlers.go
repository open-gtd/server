package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/factories"
)

func RegisterHandlers(r api.Registerer) {
	r.POST("/auth", handler(factories.Login))
}
