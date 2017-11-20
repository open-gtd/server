package api

import (
	"testing"

	"github.com/open-gtd/server/api"
)

func TestInitialize(t *testing.T) {

	getTagsHandler = &handlerMock{}
	getTagHandler = &handlerMock{}
	createTagsHandler = &handlerMock{}
	updateTagsHandler = &handlerMock{}
	deleteTagsHandler = &handlerMock{}

	r := &registererMock{}

	apiGetRegisterer = func() api.Registerer {
		return r
	}

	r.On("GET", "/api", "/tags", getTagsHandler)
	r.On("GET", "/api", "/tags/:name", getTagHandler)
	r.On("POST", "/api", "/tags", createTagsHandler)
	r.On("PUT", "/api", "/tags/:name", updateTagsHandler)
	r.On("DELETE", "/api", "/tags/:name", deleteTagsHandler)

	Initialize()

	r.AssertExpectations(t)
}
