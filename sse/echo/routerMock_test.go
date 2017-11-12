package echo

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
)

type routerMock struct {
	mock.Mock
}

func (r *routerMock) Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	args := r.Called(method, path, handler, middleware)
	return args.Get(0).(*echo.Route)
}
