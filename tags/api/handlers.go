package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/factories"
	"github.com/open-gtd/server/tags/presentation/controllers"
)

var r api.Registerer = api.GetRegisterer()

func Initialize() {
	r.GET("/api", "/tags", handler(factories.GetList))
	r.GET("/api", "/tags/:"+controllers.NameQueryParam, handler(factories.Get))
	r.POST("/api", "/tags", handler(factories.Create))
	r.PUT("/api", "/tags/:"+controllers.NameQueryParam, handler(factories.Update))
	r.DELETE("/api", "/tags/:"+controllers.NameQueryParam, handler(factories.Delete))
}
