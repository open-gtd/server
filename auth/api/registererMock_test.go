package api

import (
	"github.com/open-gtd/server/api"
	"github.com/stretchr/testify/mock"
)

type registererMock struct {
	mock.Mock
}

func (r *registererMock) GET(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}

func (r *registererMock) PATCH(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}

func (r *registererMock) POST(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}

func (r *registererMock) PUT(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}

func (r *registererMock) DELETE(prefix string, path string, handler api.Handler) {
	r.Called(prefix, path, handler)
}
