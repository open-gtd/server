package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/factories"
	"github.com/open-gtd/server/tags/presentation/controllers"
)

func RegisterHandlers(r api.Registerer) {
	r.GET("/tags", handler(factories.GetList))
	r.GET("/tags/:"+controllers.NameQueryParam, handler(factories.Get))
	r.POST("/tags", handler(factories.Create))
	r.PUT("/tags/:"+controllers.NameQueryParam, handler(factories.Update))
	r.DELETE("/tags/:"+controllers.NameQueryParam, handler(factories.Delete))
}
