package echo

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
)

type registerMock struct {
	mock.Mock
}

func (r registerMock) Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	args := r.Called(method, path, handler)
	return args.Get(0).(*echo.Route)
}
