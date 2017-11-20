package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/auth/factories"
)

var loginHandler = NewControllerHandler(factories.Login)

var apiGetRegisterer = api.GetRegisterer

func Initialize() {
	r := apiGetRegisterer()

	r.POST("", "/api/auth", loginHandler)
}
