package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/factories"
)

var loginHandler = NewControllerHandler(factories.Login)

func Initialize() {
	r := api.GetRegisterer()

	r.POST("","/api/auth", loginHandler)
}
