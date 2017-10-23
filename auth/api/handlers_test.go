package api

import (
	"testing"

	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/mock"
)

func TestInitialize(t *testing.T) {
	r := &registererMock{}
	api.SetRegisterer(r)

	r.On(
		"POST",
		"",
		"/api/auth",
		mock.AnythingOfType("controllerHandler"),
	).Once()

	Initialize()

	r.AssertExpectations(t)
}
