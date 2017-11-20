package api

import (
	"testing"

	"github.com/open-gtd/server/api"
)

func TestInitialize(t *testing.T) {

	loginHandler = &handlerMock{}

	r := &registererMock{}

	apiGetRegisterer = func() api.Registerer {
		return r
	}

	r.On("POST", "", "/api/auth", loginHandler)

	Initialize()

	r.AssertExpectations(t)
}
