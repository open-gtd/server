package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/factories"
	"github.com/open-gtd/server/tags/presentation/controllers"
)

func RegisterHandlers(r api.Registerer) {
	r.GET("/tags", handler(factories.NewGetList))
	r.GET("/tags/:"+controllers.NameQueryParam, handler(factories.NewGet))
	r.POST("/tags", handler(factories.NewCreate))
	r.PUT("/tags/:"+controllers.NameQueryParam, handler(factories.NewUpdate))
	r.DELETE("/tags/:"+controllers.NameQueryParam, handler(factories.NewDelete))
}
