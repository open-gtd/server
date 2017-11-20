package api

import (
	"github.com/open-gtd/server/api"
	"github.com/open-gtd/server/tags/factories"
	"github.com/open-gtd/server/tags/presentation/controllers"
)

var getTagsHandler api.Handler = NewControllerHandler(factories.GetList)
var getTagHandler api.Handler = NewControllerHandler(factories.Get)
var createTagsHandler api.Handler = NewControllerHandler(factories.Create)
var updateTagsHandler api.Handler = NewControllerHandler(factories.Update)
var deleteTagsHandler api.Handler = NewControllerHandler(factories.Delete)

var apiGetRegisterer = api.GetRegisterer

func Initialize() {
	r := apiGetRegisterer()

	r.GET("/api", "/tags", getTagsHandler)
	r.GET("/api", "/tags/:"+controllers.NameQueryParam, getTagHandler)
	r.POST("/api", "/tags", createTagsHandler)
	r.PUT("/api", "/tags/:"+controllers.NameQueryParam, updateTagsHandler)
	r.DELETE("/api", "/tags/:"+controllers.NameQueryParam, deleteTagsHandler)
}
