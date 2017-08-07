package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/factories"
)

var r api.Registerer = api.GetRegisterer()

func Initialize() {
	r.POST("/api", "/auth", handler(factories.Login))
}
