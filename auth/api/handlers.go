package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/factories"
)

func Initialize() {
	r := api.GetRegisterer()

	r.POST("/api", "/auth", handler(factories.Login))
}
