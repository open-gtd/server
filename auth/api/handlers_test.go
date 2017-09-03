package api

import (
	"testing"

	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/mock"
)

func TestInitialize(t *testing.T) {
	r := &registerer{}
	api.SetRegisterer(r)

	r.On(
		"POST",
		"/api",
		"/auth",
		mock.AnythingOfType("controllerHandler"),
	).Once()

	Initialize()

	r.AssertExpectations(t)
}

type registerer struct {
	mock.Mock
}

func (r *registerer) GET(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}
func (r *registerer) PATCH(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}
func (r *registerer) POST(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}
func (r *registerer) PUT(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}
func (r *registerer) DELETE(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}
