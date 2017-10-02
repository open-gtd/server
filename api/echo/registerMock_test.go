package echo

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
)

type registerMock struct {
	mock.Mock
}

func (r registerMock) Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	args := r.Called(method, path, handler)
	return route(args.Get(0))
}

func route(obj interface{}) *echo.Route {
	var r *echo.Route
	var ok bool
	if obj == nil {
		return nil
	}
	if r, ok = obj.(*echo.Route); !ok {
		panic(fmt.Sprintf("assert: arguments: Controller failed because object wasn't correct type: %v", obj))
	}
	return r
}
