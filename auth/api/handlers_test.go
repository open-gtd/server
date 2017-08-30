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
		mock.Anything,
	).Once()

	Initialize()

	r.AssertExpectations(t)
}

type registerer struct {
	mock.Mock
}

func (r *registerer) GET(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.Called(prefix, path, handlerFunc)
}
func (r *registerer) PATCH(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.Called(prefix, path, handlerFunc)
}
func (r *registerer) POST(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.Called(prefix, path, handlerFunc)
}
func (r *registerer) PUT(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.Called(prefix, path, handlerFunc)
}
func (r *registerer) DELETE(prefix string, path string, handlerFunc api.HandlerFunc) {
	r.Called(prefix, path, handlerFunc)
}
